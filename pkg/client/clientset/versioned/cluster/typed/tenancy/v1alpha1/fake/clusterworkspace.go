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

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	tenancyv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1alpha1"
)

var clusterWorkspacesResource = schema.GroupVersionResource{Group: "tenancy.kcp.dev", Version: "v1alpha1", Resource: "clusterworkspaces"}
var clusterWorkspacesKind = schema.GroupVersionKind{Group: "tenancy.kcp.dev", Version: "v1alpha1", Kind: "ClusterWorkspace"}

type clusterWorkspacesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterWorkspacesClusterClient) Cluster(clusterPath logicalcluster.Path) tenancyv1alpha1client.ClusterWorkspaceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterWorkspacesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ClusterWorkspaces that match those selectors across all clusters.
func (c *clusterWorkspacesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.ClusterWorkspaceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterWorkspacesResource, clusterWorkspacesKind, logicalcluster.Wildcard, opts), &tenancyv1alpha1.ClusterWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.ClusterWorkspaceList{ListMeta: obj.(*tenancyv1alpha1.ClusterWorkspaceList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.ClusterWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterWorkspaces across all clusters.
func (c *clusterWorkspacesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterWorkspacesResource, logicalcluster.Wildcard, opts))
}

type clusterWorkspacesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *clusterWorkspacesClient) Create(ctx context.Context, clusterWorkspace *tenancyv1alpha1.ClusterWorkspace, opts metav1.CreateOptions) (*tenancyv1alpha1.ClusterWorkspace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterWorkspacesResource, c.ClusterPath, clusterWorkspace), &tenancyv1alpha1.ClusterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), err
}

func (c *clusterWorkspacesClient) Update(ctx context.Context, clusterWorkspace *tenancyv1alpha1.ClusterWorkspace, opts metav1.UpdateOptions) (*tenancyv1alpha1.ClusterWorkspace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterWorkspacesResource, c.ClusterPath, clusterWorkspace), &tenancyv1alpha1.ClusterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), err
}

func (c *clusterWorkspacesClient) UpdateStatus(ctx context.Context, clusterWorkspace *tenancyv1alpha1.ClusterWorkspace, opts metav1.UpdateOptions) (*tenancyv1alpha1.ClusterWorkspace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterWorkspacesResource, c.ClusterPath, "status", clusterWorkspace), &tenancyv1alpha1.ClusterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), err
}

func (c *clusterWorkspacesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterWorkspacesResource, c.ClusterPath, name, opts), &tenancyv1alpha1.ClusterWorkspace{})
	return err
}

func (c *clusterWorkspacesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterWorkspacesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &tenancyv1alpha1.ClusterWorkspaceList{})
	return err
}

func (c *clusterWorkspacesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*tenancyv1alpha1.ClusterWorkspace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterWorkspacesResource, c.ClusterPath, name), &tenancyv1alpha1.ClusterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), err
}

// List takes label and field selectors, and returns the list of ClusterWorkspaces that match those selectors.
func (c *clusterWorkspacesClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.ClusterWorkspaceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterWorkspacesResource, clusterWorkspacesKind, c.ClusterPath, opts), &tenancyv1alpha1.ClusterWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.ClusterWorkspaceList{ListMeta: obj.(*tenancyv1alpha1.ClusterWorkspaceList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.ClusterWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterWorkspacesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterWorkspacesResource, c.ClusterPath, opts))
}

func (c *clusterWorkspacesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*tenancyv1alpha1.ClusterWorkspace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterWorkspacesResource, c.ClusterPath, name, pt, data, subresources...), &tenancyv1alpha1.ClusterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), err
}
