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
	"time"

	kcpkubernetesclientset "github.com/kcp-dev/client-go/kubernetes"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/util/cache"
	"k8s.io/apiserver/pkg/authorization/authorizer"
)

type CachingOptions struct {
	Options

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

// NewDelegatedAuthorizerCache creates a new cache that holds  Delegated Authorizer(s) per cluster name.
func NewDelegatedAuthorizerCache(client kcpkubernetesclientset.ClusterInterface, opts CachingOptions) *DelegatedAuthorizerCache {
	opts.defaults()
	return &DelegatedAuthorizerCache{
		opts:   &opts,
		cache:  cache.NewExpiring(),
		client: client,
	}
}

type DelegatedAuthorizerCache struct {
	opts  *CachingOptions
	cache *cache.Expiring

	client kcpkubernetesclientset.ClusterInterface
}

// load loads the authorizer from the cache, if any.
func (c *DelegatedAuthorizerCache) load(clusterName logicalcluster.Name) authorizer.Authorizer {
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

// GetDelegatedAuthorizer returns or creates a new DelegatedAuthorizer for the given cluster name if it doesn't exist.
func (c *DelegatedAuthorizerCache) GetDelegatedAuthorizer(clusterName logicalcluster.Name) (authorizer.Authorizer, error) {
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
