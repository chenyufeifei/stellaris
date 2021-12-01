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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MultiClusterResourceAggregateRuleListLister helps list MultiClusterResourceAggregateRuleLists.
// All objects returned here must be treated as read-only.
type MultiClusterResourceAggregateRuleListLister interface {
	// List lists all MultiClusterResourceAggregateRuleLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceAggregateRuleList, err error)
	// MultiClusterResourceAggregateRuleLists returns an object that can list and get MultiClusterResourceAggregateRuleLists.
	MultiClusterResourceAggregateRuleLists(namespace string) MultiClusterResourceAggregateRuleListNamespaceLister
	MultiClusterResourceAggregateRuleListListerExpansion
}

// multiClusterResourceAggregateRuleListLister implements the MultiClusterResourceAggregateRuleListLister interface.
type multiClusterResourceAggregateRuleListLister struct {
	indexer cache.Indexer
}

// NewMultiClusterResourceAggregateRuleListLister returns a new MultiClusterResourceAggregateRuleListLister.
func NewMultiClusterResourceAggregateRuleListLister(indexer cache.Indexer) MultiClusterResourceAggregateRuleListLister {
	return &multiClusterResourceAggregateRuleListLister{indexer: indexer}
}

// List lists all MultiClusterResourceAggregateRuleLists in the indexer.
func (s *multiClusterResourceAggregateRuleListLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceAggregateRuleList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResourceAggregateRuleList))
	})
	return ret, err
}

// MultiClusterResourceAggregateRuleLists returns an object that can list and get MultiClusterResourceAggregateRuleLists.
func (s *multiClusterResourceAggregateRuleListLister) MultiClusterResourceAggregateRuleLists(namespace string) MultiClusterResourceAggregateRuleListNamespaceLister {
	return multiClusterResourceAggregateRuleListNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MultiClusterResourceAggregateRuleListNamespaceLister helps list and get MultiClusterResourceAggregateRuleLists.
// All objects returned here must be treated as read-only.
type MultiClusterResourceAggregateRuleListNamespaceLister interface {
	// List lists all MultiClusterResourceAggregateRuleLists in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceAggregateRuleList, err error)
	// Get retrieves the MultiClusterResourceAggregateRuleList from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MultiClusterResourceAggregateRuleList, error)
	MultiClusterResourceAggregateRuleListNamespaceListerExpansion
}

// multiClusterResourceAggregateRuleListNamespaceLister implements the MultiClusterResourceAggregateRuleListNamespaceLister
// interface.
type multiClusterResourceAggregateRuleListNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MultiClusterResourceAggregateRuleLists in the indexer for a given namespace.
func (s multiClusterResourceAggregateRuleListNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceAggregateRuleList, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResourceAggregateRuleList))
	})
	return ret, err
}

// Get retrieves the MultiClusterResourceAggregateRuleList from the indexer for a given namespace and name.
func (s multiClusterResourceAggregateRuleListNamespaceLister) Get(name string) (*v1alpha1.MultiClusterResourceAggregateRuleList, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("multiclusterresourceaggregaterulelist"), name)
	}
	return obj.(*v1alpha1.MultiClusterResourceAggregateRuleList), nil
}
