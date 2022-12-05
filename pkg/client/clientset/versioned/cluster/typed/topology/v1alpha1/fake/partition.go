//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	topologyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/topology/v1alpha1"
	topologyv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/topology/v1alpha1"
)

var partitionsResource = schema.GroupVersionResource{Group: "topology.kcp.dev", Version: "v1alpha1", Resource: "partitions"}
var partitionsKind = schema.GroupVersionKind{Group: "topology.kcp.dev", Version: "v1alpha1", Kind: "Partition"}

type partitionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *partitionsClusterClient) Cluster(clusterPath logicalcluster.Path) topologyv1alpha1client.PartitionInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &partitionsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Partitions that match those selectors across all clusters.
func (c *partitionsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*topologyv1alpha1.PartitionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(partitionsResource, partitionsKind, logicalcluster.Wildcard, opts), &topologyv1alpha1.PartitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &topologyv1alpha1.PartitionList{ListMeta: obj.(*topologyv1alpha1.PartitionList).ListMeta}
	for _, item := range obj.(*topologyv1alpha1.PartitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Partitions across all clusters.
func (c *partitionsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(partitionsResource, logicalcluster.Wildcard, opts))
}

type partitionsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *partitionsClient) Create(ctx context.Context, partition *topologyv1alpha1.Partition, opts metav1.CreateOptions) (*topologyv1alpha1.Partition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(partitionsResource, c.ClusterPath, partition), &topologyv1alpha1.Partition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*topologyv1alpha1.Partition), err
}

func (c *partitionsClient) Update(ctx context.Context, partition *topologyv1alpha1.Partition, opts metav1.UpdateOptions) (*topologyv1alpha1.Partition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(partitionsResource, c.ClusterPath, partition), &topologyv1alpha1.Partition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*topologyv1alpha1.Partition), err
}

func (c *partitionsClient) UpdateStatus(ctx context.Context, partition *topologyv1alpha1.Partition, opts metav1.UpdateOptions) (*topologyv1alpha1.Partition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(partitionsResource, c.ClusterPath, "status", partition), &topologyv1alpha1.Partition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*topologyv1alpha1.Partition), err
}

func (c *partitionsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(partitionsResource, c.ClusterPath, name, opts), &topologyv1alpha1.Partition{})
	return err
}

func (c *partitionsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(partitionsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &topologyv1alpha1.PartitionList{})
	return err
}

func (c *partitionsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*topologyv1alpha1.Partition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(partitionsResource, c.ClusterPath, name), &topologyv1alpha1.Partition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*topologyv1alpha1.Partition), err
}

// List takes label and field selectors, and returns the list of Partitions that match those selectors.
func (c *partitionsClient) List(ctx context.Context, opts metav1.ListOptions) (*topologyv1alpha1.PartitionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(partitionsResource, partitionsKind, c.ClusterPath, opts), &topologyv1alpha1.PartitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &topologyv1alpha1.PartitionList{ListMeta: obj.(*topologyv1alpha1.PartitionList).ListMeta}
	for _, item := range obj.(*topologyv1alpha1.PartitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *partitionsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(partitionsResource, c.ClusterPath, opts))
}

func (c *partitionsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*topologyv1alpha1.Partition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(partitionsResource, c.ClusterPath, name, pt, data, subresources...), &topologyv1alpha1.Partition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*topologyv1alpha1.Partition), err
}
