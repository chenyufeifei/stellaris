/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAggregatedResources implements AggregatedResourceInterface
type FakeAggregatedResources struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var aggregatedresourcesResource = schema.GroupVersionResource{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Resource: "aggregatedresources"}

var aggregatedresourcesKind = schema.GroupVersionKind{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Kind: "AggregatedResource"}

// Get takes name of the aggregatedResource, and returns the corresponding aggregatedResource object, and an error if there is any.
func (c *FakeAggregatedResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AggregatedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aggregatedresourcesResource, c.ns, name), &v1alpha1.AggregatedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregatedResource), err
}

// List takes label and field selectors, and returns the list of AggregatedResources that match those selectors.
func (c *FakeAggregatedResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AggregatedResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aggregatedresourcesResource, aggregatedresourcesKind, c.ns, opts), &v1alpha1.AggregatedResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AggregatedResourceList{ListMeta: obj.(*v1alpha1.AggregatedResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AggregatedResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aggregatedResources.
func (c *FakeAggregatedResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aggregatedresourcesResource, c.ns, opts))

}

// Create takes the representation of a aggregatedResource and creates it.  Returns the server's representation of the aggregatedResource, and an error, if there is any.
func (c *FakeAggregatedResources) Create(ctx context.Context, aggregatedResource *v1alpha1.AggregatedResource, opts v1.CreateOptions) (result *v1alpha1.AggregatedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aggregatedresourcesResource, c.ns, aggregatedResource), &v1alpha1.AggregatedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregatedResource), err
}

// Update takes the representation of a aggregatedResource and updates it. Returns the server's representation of the aggregatedResource, and an error, if there is any.
func (c *FakeAggregatedResources) Update(ctx context.Context, aggregatedResource *v1alpha1.AggregatedResource, opts v1.UpdateOptions) (result *v1alpha1.AggregatedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aggregatedresourcesResource, c.ns, aggregatedResource), &v1alpha1.AggregatedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregatedResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAggregatedResources) UpdateStatus(ctx context.Context, aggregatedResource *v1alpha1.AggregatedResource, opts v1.UpdateOptions) (*v1alpha1.AggregatedResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(aggregatedresourcesResource, "status", c.ns, aggregatedResource), &v1alpha1.AggregatedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregatedResource), err
}

// Delete takes name of the aggregatedResource and deletes it. Returns an error if one occurs.
func (c *FakeAggregatedResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(aggregatedresourcesResource, c.ns, name), &v1alpha1.AggregatedResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAggregatedResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aggregatedresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AggregatedResourceList{})
	return err
}

// Patch applies the patch and returns the patched aggregatedResource.
func (c *FakeAggregatedResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AggregatedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aggregatedresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AggregatedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregatedResource), err
}