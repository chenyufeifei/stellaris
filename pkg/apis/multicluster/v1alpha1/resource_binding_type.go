package v1alpha1

import (
	"harmonycloud.cn/stellaris/pkg/apis/multicluster/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

type MultiClusterResourceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultiClusterResourceBindingSpec   `json:"spec,omitempty"`
	Status MultiClusterResourceBindingStatus `json:"status,omitempty"`
}

type MultiClusterResourceBindingSpec struct {
	Resources []MultiClusterResourceBindingResource `json:"resources,omitempty"`
}

type MultiClusterResourceBindingResource struct {
	Name      string                               `json:"name,omitempty"`
	Namespace string                               `json:"namespace,omitempty"`
	Clusters  []MultiClusterResourceBindingCluster `json:"clusters,omitempty"`
}

type MultiClusterResourceBindingCluster struct {
	Name     string             `json:"name,omitempty"`
}

type MultiClusterResourceBindingStatus struct {
	ClusterStatus []common.MultiClusterResourceClusterStatus `json:"clusters,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterResourceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MultiClusterResourceBinding `json:"items"`
}
