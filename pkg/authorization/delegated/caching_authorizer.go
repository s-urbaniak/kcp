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

package delegated

import (
	"context"
	"fmt"
	"time"

	kcpkubernetesclientset "github.com/kcp-dev/client-go/kubernetes"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/util/cache"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/klog/v2"

	corev1alpha1 "github.com/kcp-dev/kcp/pkg/apis/core/v1alpha1"
	"github.com/kcp-dev/kcp/pkg/apis/tenancy/initialization"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	dynamiccontext "github.com/kcp-dev/kcp/pkg/virtual/framework/dynamic/context"
)

type CachingOptions struct {
	Options

	// Verb is the verb being authorized, passed in authorizer.AttributesRecord.
	Verb string

	// Resource is the resource being authorized, passed in authorizer.AttributesRecord.
	Resource string

	// TTL is the default time-to-live when a delegated authorizer
	// is stored in the internal cache.
	TTL time.Duration
}

func (c *CachingOptions) defaults() {
	c.Options.defaults()
	if c.TTL == 0 {
		c.TTL = 12 * time.Hour
	}
}

// NewCaching creates a new Authorizer that holds an internal cache of
// Delegated Authorizer(s).
// A Delegated Authorizer, in turn, caches Authorize requests for the
func NewCaching(client kcpkubernetesclientset.ClusterInterface, opts CachingOptions) *cachingAuthorizer {
	opts.defaults()
	return &cachingAuthorizer{
		opts:   &opts,
		cache:  cache.NewExpiring(),
		client: client,
	}
}

// cachingAuthorizer is a wrapper around authorizer.Authorize that uses
// an internal expiring cache.
type cachingAuthorizer struct {
	opts  *CachingOptions
	cache *cache.Expiring

	client kcpkubernetesclientset.ClusterInterface
}

// load loads the authorizer from the cache, if any.
func (c *cachingAuthorizer) load(clusterName logicalcluster.Name) authorizer.Authorizer {
	value, ok := c.cache.Get(clusterName)
	if !ok && value == nil {
		return nil
	}
	authz, ok := value.(authorizer.Authorizer)
	if !ok {
		return nil
	}
	return authz
}

func (c *cachingAuthorizer) loadOrStore(clusterName logicalcluster.Name) (authorizer.Authorizer, error) {
	if authz := c.load(clusterName); authz != nil {
		return authz, nil
	}

	// Create the delegated authorizer.
	authz, err := NewDelegatedAuthorizer(clusterName, c.client, c.opts.Options)
	if err != nil {
		return nil, err
	}

	// Store the cache and return.
	c.cache.Set(clusterName, authz, c.opts.TTL)
	return authz, nil
}

func (c *cachingAuthorizer) Authorize(ctx context.Context, attr authorizer.Attributes) (authorized authorizer.Decision, reason string, err error) {
	logger := klog.FromContext(ctx)
	clusterName, name, err := initialization.TypeFrom(
		corev1alpha1.LogicalClusterInitializer(dynamiccontext.APIDomainKeyFrom(ctx)),
	)
	if err != nil {
		logger.V(2).Info(err.Error())
		return authorizer.DecisionNoOpinion, "unable to determine initializer", fmt.Errorf("access not permitted")
	}

	authz, err := c.loadOrStore(clusterName)
	if err != nil {
		return authorizer.DecisionNoOpinion, "", err
	}

	SARAttributes := authorizer.AttributesRecord{
		APIGroup:        tenancyv1alpha1.SchemeGroupVersion.Group,
		APIVersion:      tenancyv1alpha1.SchemeGroupVersion.Version,
		User:            attr.GetUser(),
		Name:            name,
		ResourceRequest: c.opts.Resource != "",
		Verb:            c.opts.Verb,
		Resource:        c.opts.Resource,
	}

	return authz.Authorize(ctx, SARAttributes)
}
