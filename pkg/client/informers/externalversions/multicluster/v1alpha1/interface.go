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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "harmonycloud.cn/multi-cluster-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AggregatedResources returns a AggregatedResourceInformer.
	AggregatedResources() AggregatedResourceInformer
	// AggregatedResourceLists returns a AggregatedResourceListInformer.
	AggregatedResourceLists() AggregatedResourceListInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterResources returns a ClusterResourceInformer.
	ClusterResources() ClusterResourceInformer
	// ClusterSets returns a ClusterSetInformer.
	ClusterSets() ClusterSetInformer
	// MultiClusterResources returns a MultiClusterResourceInformer.
	MultiClusterResources() MultiClusterResourceInformer
	// MultiClusterResourceAggregatePolicies returns a MultiClusterResourceAggregatePolicyInformer.
	MultiClusterResourceAggregatePolicies() MultiClusterResourceAggregatePolicyInformer
	// MultiClusterResourceAggregatePolicyLists returns a MultiClusterResourceAggregatePolicyListInformer.
	MultiClusterResourceAggregatePolicyLists() MultiClusterResourceAggregatePolicyListInformer
	// MultiClusterResourceAggregateRules returns a MultiClusterResourceAggregateRuleInformer.
	MultiClusterResourceAggregateRules() MultiClusterResourceAggregateRuleInformer
	// MultiClusterResourceAggregateRuleLists returns a MultiClusterResourceAggregateRuleListInformer.
	MultiClusterResourceAggregateRuleLists() MultiClusterResourceAggregateRuleListInformer
	// NamespaceMappings returns a NamespaceMappingInformer.
	NamespaceMappings() NamespaceMappingInformer
	// ResourceAggregatePolicies returns a ResourceAggregatePolicyInformer.
	ResourceAggregatePolicies() ResourceAggregatePolicyInformer
	// ResourceAggregatePolicyLists returns a ResourceAggregatePolicyListInformer.
	ResourceAggregatePolicyLists() ResourceAggregatePolicyListInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AggregatedResources returns a AggregatedResourceInformer.
func (v *version) AggregatedResources() AggregatedResourceInformer {
	return &aggregatedResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AggregatedResourceLists returns a AggregatedResourceListInformer.
func (v *version) AggregatedResourceLists() AggregatedResourceListInformer {
	return &aggregatedResourceListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterResources returns a ClusterResourceInformer.
func (v *version) ClusterResources() ClusterResourceInformer {
	return &clusterResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterSets returns a ClusterSetInformer.
func (v *version) ClusterSets() ClusterSetInformer {
	return &clusterSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MultiClusterResources returns a MultiClusterResourceInformer.
func (v *version) MultiClusterResources() MultiClusterResourceInformer {
	return &multiClusterResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterResourceAggregatePolicies returns a MultiClusterResourceAggregatePolicyInformer.
func (v *version) MultiClusterResourceAggregatePolicies() MultiClusterResourceAggregatePolicyInformer {
	return &multiClusterResourceAggregatePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterResourceAggregatePolicyLists returns a MultiClusterResourceAggregatePolicyListInformer.
func (v *version) MultiClusterResourceAggregatePolicyLists() MultiClusterResourceAggregatePolicyListInformer {
	return &multiClusterResourceAggregatePolicyListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterResourceAggregateRules returns a MultiClusterResourceAggregateRuleInformer.
func (v *version) MultiClusterResourceAggregateRules() MultiClusterResourceAggregateRuleInformer {
	return &multiClusterResourceAggregateRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterResourceAggregateRuleLists returns a MultiClusterResourceAggregateRuleListInformer.
func (v *version) MultiClusterResourceAggregateRuleLists() MultiClusterResourceAggregateRuleListInformer {
	return &multiClusterResourceAggregateRuleListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NamespaceMappings returns a NamespaceMappingInformer.
func (v *version) NamespaceMappings() NamespaceMappingInformer {
	return &namespaceMappingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceAggregatePolicies returns a ResourceAggregatePolicyInformer.
func (v *version) ResourceAggregatePolicies() ResourceAggregatePolicyInformer {
	return &resourceAggregatePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceAggregatePolicyLists returns a ResourceAggregatePolicyListInformer.
func (v *version) ResourceAggregatePolicyLists() ResourceAggregatePolicyListInformer {
	return &resourceAggregatePolicyListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
