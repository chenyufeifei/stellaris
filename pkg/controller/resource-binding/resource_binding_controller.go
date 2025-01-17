package resource_binding

import (
	"context"
	"fmt"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"strings"

	"harmonycloud.cn/stellaris/pkg/apis/multicluster/common"

	"github.com/go-logr/logr"
	"harmonycloud.cn/stellaris/pkg/apis/multicluster/v1alpha1"
	managerCommon "harmonycloud.cn/stellaris/pkg/common"
	controllerCommon "harmonycloud.cn/stellaris/pkg/controller/common"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type Reconciler struct {
	client.Client
	log    logr.Logger
	Scheme *runtime.Scheme
}

func (r *Reconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	r.log.Info("Reconciling MultiClusterResourceBinding")
	// get resource
	instance := &v1alpha1.MultiClusterResourceBinding{}
	err := r.Client.Get(ctx, request.NamespacedName, instance)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// add Finalizers
	if controllerCommon.ShouldAddFinalizer(instance) {
		return r.addFinalizer(ctx, instance)
	}

	// the object is being deleted
	// TODO (chenkun) delete clusterResource
	if !instance.GetDeletionTimestamp().IsZero() {
		return r.removeFinalizer(ctx, instance)
	}

	// add labels
	if shouldChangeBindingLabels(instance) {
		return r.addBindingLabels(ctx, instance)
	}

	return r.updateStatusAndSyncClusterResource(ctx, instance)
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.MultiClusterResourceBinding{}).
		Watches(&source.Kind{Type: &v1alpha1.MultiClusterResourceOverride{}},
		handler.EnqueueRequestsFromMapFunc(NewOverridePolicyFunc(r.Client))).
		Complete(r)
}

func Setup(mgr ctrl.Manager, controllerCommon controllerCommon.Args) error {
	reconciler := Reconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		log:    logf.Log.WithName("resource_binding_controller"),
	}
	return reconciler.SetupWithManager(mgr)
}

func (r *Reconciler) addFinalizer(ctx context.Context, instance *v1alpha1.MultiClusterResourceBinding) (ctrl.Result, error) {
	err := controllerCommon.AddFinalizer(ctx, r.Client, instance)
	if err != nil {
		r.log.Error(err, fmt.Sprintf("append finalizer filed, resource(%s)", instance.Name))
		return controllerCommon.ReQueueResult(err)
	}
	return ctrl.Result{}, nil
}

func (r *Reconciler) removeFinalizer(ctx context.Context, instance *v1alpha1.MultiClusterResourceBinding) (ctrl.Result, error) {
	err := controllerCommon.RemoveFinalizer(ctx, r.Client, instance)
	if err != nil {
		r.log.Error(err, fmt.Sprintf("delete finalizer filed, resource(%s)", instance.Name))
		return controllerCommon.ReQueueResult(err)
	}
	return ctrl.Result{}, nil
}

func (r *Reconciler) updateStatusAndSyncClusterResource(ctx context.Context, instance *v1alpha1.MultiClusterResourceBinding) (ctrl.Result, error) {
	// get ClusterResourceList
	clusterResourceList, err := getClusterResourceListForBinding(ctx, r.Client, instance)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			r.log.Error(err, fmt.Sprintf("get clusterResource for resource failed, resource(%s)", instance.Name))
			return controllerCommon.ReQueueResult(err)
		}
	}

	// sync ClusterResource
	err = syncClusterResource(ctx, r.Client, clusterResourceList, instance)
	if err != nil {
		r.log.Error(err, fmt.Sprintf("sync ClusterResource failed, resource(%s)", instance.Name))
		return controllerCommon.ReQueueResult(err)
	}

	// update status
	err = updateBindingStatus(ctx, r.Client, instance, clusterResourceList)
	if err != nil {
		r.log.Error(err, fmt.Sprintf("update binding status failed, resource(%s)", instance.Name))
		return controllerCommon.ReQueueResult(err)
	}

	return ctrl.Result{}, nil
}

func getMultiClusterResourceForName(ctx context.Context, clientSet client.Client, multiClusterResourceName, multiClusterResourceNamespace string) (*v1alpha1.MultiClusterResource, error) {
	object := &v1alpha1.MultiClusterResource{}
	if len(multiClusterResourceNamespace) == 0 {
		multiClusterResourceNamespace = managerCommon.ManagerNamespace
	}
	namespacedName := types.NamespacedName{
		Namespace: multiClusterResourceNamespace,
		Name:      multiClusterResourceName,
	}
	err := clientSet.Get(ctx, namespacedName, object)
	return object, err
}

func updateBindingStatus(ctx context.Context, clientSet client.Client, binding *v1alpha1.MultiClusterResourceBinding, clusterResourceList *v1alpha1.ClusterResourceList) error {
	updateStatus := false
	for _, clusterResource := range clusterResourceList.Items {
		// no status
		if len(clusterResource.Status.Phase) <= 0 {
			continue
		}

		bindingStatusMap := bindingClusterStatusMap(binding)
		// find target multiClusterResourceClusterStatus
		key := bindingClusterStatusMapKey(managerCommon.ClusterName(clusterResource.GetNamespace()), clusterResource.GetName())
		bindingStatus, ok := bindingStatusMap[key]
		if ok {
			if statusEqual(clusterResource.Status, *bindingStatus) {
				continue
			}
			binding.Status.ClusterStatus = removeItemForClusterStatusList(binding.Status.ClusterStatus, *bindingStatus)
		}
		// should update binding status
		updateStatus = true
		// new resourceStatus
		multiClusterResourceClusterStatus := common.MultiClusterResourceClusterStatus{
			Name:                      managerCommon.ClusterName(clusterResource.GetNamespace()),
			Resource:                  clusterResource.Name,
			ObservedReceiveGeneration: clusterResource.Status.ObservedReceiveGeneration,
			Phase:                     clusterResource.Status.Phase,
			Message:                   clusterResource.Status.Message,
			Binding:                   binding.Name,
		}
		binding.Status.ClusterStatus = append(binding.Status.ClusterStatus, multiClusterResourceClusterStatus)
	}
	if updateStatus {
		// update binding status
		return clientSet.Status().Update(ctx, binding)
	}
	return nil
}

func bindingClusterStatusMap(binding *v1alpha1.MultiClusterResourceBinding) map[string]*common.MultiClusterResourceClusterStatus {
	statusMap := map[string]*common.MultiClusterResourceClusterStatus{}
	for _, item := range binding.Status.ClusterStatus {
		statusKey := bindingClusterStatusMapKey(item.Name, item.Resource)
		statusMap[statusKey] = &item
	}
	return statusMap
}

func bindingClusterStatusMapKey(clusterName, resourceName string) string {
	return clusterName + "." + resourceName
}

func statusEqual(clusterResourceStatus v1alpha1.ClusterResourceStatus, bindingStatus common.MultiClusterResourceClusterStatus) bool {
	if clusterResourceStatus.Phase != bindingStatus.Phase || clusterResourceStatus.Message != bindingStatus.Message || clusterResourceStatus.ObservedReceiveGeneration != bindingStatus.ObservedReceiveGeneration {
		return false
	}
	return true
}

// add labels
func shouldChangeBindingLabels(binding *v1alpha1.MultiClusterResourceBinding) bool {
	if len(binding.Spec.Resources) <= 0 {
		return false
	}
	currentLabels := getMultiClusterResourceLabels(binding)
	if len(currentLabels) <= 0 {
		return true
	}
	existLabels := shouldExistLabels(binding)
	if reflect.DeepEqual(existLabels, currentLabels) {
		return false
	}
	return true
}
func (r *Reconciler) addBindingLabels(ctx context.Context, binding *v1alpha1.MultiClusterResourceBinding) (ctrl.Result, error) {
	currentLabels := getMultiClusterResourceLabels(binding)
	existLabels := shouldExistLabels(binding)

	binding.SetLabels(replaceLabels(binding.GetLabels(), currentLabels, existLabels))
	err := r.Client.Update(ctx, binding)
	if err != nil {
		return controllerCommon.ReQueueResult(err)
	}
	return ctrl.Result{}, nil
}

func replaceLabels(bindingLabels, removeLabels, addLabels map[string]string) map[string]string {
	if len(bindingLabels) <= 0 || len(removeLabels) <= 0 {
		return addLabels
	}
	if reflect.DeepEqual(bindingLabels, removeLabels) {
		return addLabels
	}
	for removeKey, _ := range removeLabels {
		delete(bindingLabels, removeKey)
	}
	for addKey, addValue := range addLabels {
		bindingLabels[addKey] = addValue
	}
	return bindingLabels
}

func shouldExistLabels(binding *v1alpha1.MultiClusterResourceBinding) map[string]string {
	existLabels := map[string]string{}
	for _, resource := range binding.Spec.Resources {
		existLabels[managerCommon.MultiClusterResourceLabelName+"."+resource.Name] = "1"
	}
	return existLabels
}

func getMultiClusterResourceLabels(binding *v1alpha1.MultiClusterResourceBinding) map[string]string {
	labels := map[string]string{}
	for k, v := range binding.GetLabels() {
		if strings.HasPrefix(k, managerCommon.MultiClusterResourceLabelName) {
			labels[k] = v
		}
	}
	return labels
}
