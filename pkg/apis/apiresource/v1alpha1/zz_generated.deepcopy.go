//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The KCP Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceImport) DeepCopyInto(out *APIResourceImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceImport.
func (in *APIResourceImport) DeepCopy() *APIResourceImport {
	if in == nil {
		return nil
	}
	out := new(APIResourceImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIResourceImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceImportCondition) DeepCopyInto(out *APIResourceImportCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceImportCondition.
func (in *APIResourceImportCondition) DeepCopy() *APIResourceImportCondition {
	if in == nil {
		return nil
	}
	out := new(APIResourceImportCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceImportList) DeepCopyInto(out *APIResourceImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIResourceImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceImportList.
func (in *APIResourceImportList) DeepCopy() *APIResourceImportList {
	if in == nil {
		return nil
	}
	out := new(APIResourceImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIResourceImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceImportSpec) DeepCopyInto(out *APIResourceImportSpec) {
	*out = *in
	in.CommonAPIResourceSpec.DeepCopyInto(&out.CommonAPIResourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceImportSpec.
func (in *APIResourceImportSpec) DeepCopy() *APIResourceImportSpec {
	if in == nil {
		return nil
	}
	out := new(APIResourceImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceImportStatus) DeepCopyInto(out *APIResourceImportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]APIResourceImportCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceImportStatus.
func (in *APIResourceImportStatus) DeepCopy() *APIResourceImportStatus {
	if in == nil {
		return nil
	}
	out := new(APIResourceImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnDefinition) DeepCopyInto(out *ColumnDefinition) {
	*out = *in
	out.TableColumnDefinition = in.TableColumnDefinition
	if in.JSONPath != nil {
		in, out := &in.JSONPath, &out.JSONPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnDefinition.
func (in *ColumnDefinition) DeepCopy() *ColumnDefinition {
	if in == nil {
		return nil
	}
	out := new(ColumnDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ColumnDefinitions) DeepCopyInto(out *ColumnDefinitions) {
	{
		in := &in
		*out = make(ColumnDefinitions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnDefinitions.
func (in ColumnDefinitions) DeepCopy() ColumnDefinitions {
	if in == nil {
		return nil
	}
	out := new(ColumnDefinitions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonAPIResourceSpec) DeepCopyInto(out *CommonAPIResourceSpec) {
	*out = *in
	out.GroupVersion = in.GroupVersion
	in.CustomResourceDefinitionNames.DeepCopyInto(&out.CustomResourceDefinitionNames)
	in.OpenAPIV3Schema.DeepCopyInto(&out.OpenAPIV3Schema)
	if in.SubResources != nil {
		in, out := &in.SubResources, &out.SubResources
		*out = make(SubResources, len(*in))
		copy(*out, *in)
	}
	if in.ColumnDefinitions != nil {
		in, out := &in.ColumnDefinitions, &out.ColumnDefinitions
		*out = make(ColumnDefinitions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonAPIResourceSpec.
func (in *CommonAPIResourceSpec) DeepCopy() *CommonAPIResourceSpec {
	if in == nil {
		return nil
	}
	out := new(CommonAPIResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersion) DeepCopyInto(out *GroupVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersion.
func (in *GroupVersion) DeepCopy() *GroupVersion {
	if in == nil {
		return nil
	}
	out := new(GroupVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegotiatedAPIResource) DeepCopyInto(out *NegotiatedAPIResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegotiatedAPIResource.
func (in *NegotiatedAPIResource) DeepCopy() *NegotiatedAPIResource {
	if in == nil {
		return nil
	}
	out := new(NegotiatedAPIResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NegotiatedAPIResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegotiatedAPIResourceCondition) DeepCopyInto(out *NegotiatedAPIResourceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegotiatedAPIResourceCondition.
func (in *NegotiatedAPIResourceCondition) DeepCopy() *NegotiatedAPIResourceCondition {
	if in == nil {
		return nil
	}
	out := new(NegotiatedAPIResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegotiatedAPIResourceList) DeepCopyInto(out *NegotiatedAPIResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NegotiatedAPIResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegotiatedAPIResourceList.
func (in *NegotiatedAPIResourceList) DeepCopy() *NegotiatedAPIResourceList {
	if in == nil {
		return nil
	}
	out := new(NegotiatedAPIResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NegotiatedAPIResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegotiatedAPIResourceSpec) DeepCopyInto(out *NegotiatedAPIResourceSpec) {
	*out = *in
	in.CommonAPIResourceSpec.DeepCopyInto(&out.CommonAPIResourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegotiatedAPIResourceSpec.
func (in *NegotiatedAPIResourceSpec) DeepCopy() *NegotiatedAPIResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NegotiatedAPIResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegotiatedAPIResourceStatus) DeepCopyInto(out *NegotiatedAPIResourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NegotiatedAPIResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegotiatedAPIResourceStatus.
func (in *NegotiatedAPIResourceStatus) DeepCopy() *NegotiatedAPIResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NegotiatedAPIResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource) DeepCopyInto(out *SubResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource.
func (in *SubResource) DeepCopy() *SubResource {
	if in == nil {
		return nil
	}
	out := new(SubResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SubResources) DeepCopyInto(out *SubResources) {
	{
		in := &in
		*out = make(SubResources, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResources.
func (in SubResources) DeepCopy() SubResources {
	if in == nil {
		return nil
	}
	out := new(SubResources)
	in.DeepCopyInto(out)
	return *out
}
