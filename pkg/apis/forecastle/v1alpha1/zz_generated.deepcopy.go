// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForecastleApp) DeepCopyInto(out *ForecastleApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForecastleApp.
func (in *ForecastleApp) DeepCopy() *ForecastleApp {
	if in == nil {
		return nil
	}
	out := new(ForecastleApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForecastleApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForecastleAppList) DeepCopyInto(out *ForecastleAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ForecastleApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForecastleAppList.
func (in *ForecastleAppList) DeepCopy() *ForecastleAppList {
	if in == nil {
		return nil
	}
	out := new(ForecastleAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForecastleAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForecastleAppSpec) DeepCopyInto(out *ForecastleAppSpec) {
	*out = *in
	if in.URLFrom != nil {
		in, out := &in.URLFrom, &out.URLFrom
		*out = new(URLSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForecastleAppSpec.
func (in *ForecastleAppSpec) DeepCopy() *ForecastleAppSpec {
	if in == nil {
		return nil
	}
	out := new(ForecastleAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForecastleAppStatus) DeepCopyInto(out *ForecastleAppStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForecastleAppStatus.
func (in *ForecastleAppStatus) DeepCopy() *ForecastleAppStatus {
	if in == nil {
		return nil
	}
	out := new(ForecastleAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressURLSource) DeepCopyInto(out *IngressURLSource) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressURLSource.
func (in *IngressURLSource) DeepCopy() *IngressURLSource {
	if in == nil {
		return nil
	}
	out := new(IngressURLSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URLSource) DeepCopyInto(out *URLSource) {
	*out = *in
	if in.IngressRef != nil {
		in, out := &in.IngressRef, &out.IngressRef
		*out = new(IngressURLSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URLSource.
func (in *URLSource) DeepCopy() *URLSource {
	if in == nil {
		return nil
	}
	out := new(URLSource)
	in.DeepCopyInto(out)
	return out
}
