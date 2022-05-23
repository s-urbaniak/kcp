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

package internalapis

import (
	"k8s.io/apimachinery/pkg/runtime"
	common "k8s.io/kube-openapi/pkg/common"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	generatedopenapi "k8s.io/kubernetes/pkg/generated/openapi"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
)

// Schemas contains a list of internal APIs that should be exposed for the syncer of any WorkloadCluster.
var Schemas []*apisv1alpha1.APIResourceSchema

func init() {
	schemes := []*runtime.Scheme{legacyscheme.Scheme}
	openAPIDefinitionsGetters := []common.GetOpenAPIDefinitions{generatedopenapi.GetOpenAPIDefinitions}

	if apis, err := createAPIResourceSchemas(schemes, openAPIDefinitionsGetters, KCPInternalAPIs...); err != nil {
		panic(err)
	} else {
		Schemas = apis
	}
}
