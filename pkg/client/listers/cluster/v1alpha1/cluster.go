/*
Copyright 2022 The KCP Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/cluster/v1alpha1"
)

// ClusterLister helps list Clusters.
// All objects returned here must be treated as read-only.
type ClusterLister interface {
	Scoped(scope rest.Scope) ClusterLister
	// List lists all Clusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Cluster, err error)
	// Get retrieves the Cluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Cluster, error)
	ClusterListerExpansion
}

// clusterLister implements the ClusterLister interface.
type clusterLister struct {
	indexer cache.Indexer
	scope   rest.Scope
}

// NewClusterLister returns a new ClusterLister.
func NewClusterLister(indexer cache.Indexer) ClusterLister {
	return &clusterLister{indexer: indexer}
}

func (s *clusterLister) Scoped(scope rest.Scope) ClusterLister {
	return &clusterLister{
		indexer: s.indexer,
		scope:   scope,
	}
}

// List lists all Clusters in the indexer.
func (s *clusterLister) List(selector labels.Selector) (ret []*v1alpha1.Cluster, err error) {
	appendFunc := func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cluster))
	}

	if s.scope == nil {
		err = cache.ListAll(s.indexer, selector, appendFunc)
		return ret, err
	}

	indexValue := s.scope.Name()
	err = cache.ListAllByIndexAndValue(s.indexer, cache.ListAllIndex, indexValue, selector, appendFunc)
	return ret, err
}

// Get retrieves the Cluster from the index for a given name.
func (s *clusterLister) Get(name string) (*v1alpha1.Cluster, error) {
	key := name
	if s.scope != nil {
		key = s.scope.CacheKey(key)
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cluster"), name)
	}
	return obj.(*v1alpha1.Cluster), nil
}
