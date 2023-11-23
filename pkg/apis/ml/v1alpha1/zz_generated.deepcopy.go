//go:build !ignore_autogenerated
// +build !ignore_autogenerated

//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJob) DeepCopyInto(out *ArangoMLBatchJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJob.
func (in *ArangoMLBatchJob) DeepCopy() *ArangoMLBatchJob {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLBatchJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobList) DeepCopyInto(out *ArangoMLBatchJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLBatchJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobList.
func (in *ArangoMLBatchJobList) DeepCopy() *ArangoMLBatchJobList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLBatchJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobSpec) DeepCopyInto(out *ArangoMLBatchJobSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobSpec.
func (in *ArangoMLBatchJobSpec) DeepCopy() *ArangoMLBatchJobSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobStatus) DeepCopyInto(out *ArangoMLBatchJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobStatus.
func (in *ArangoMLBatchJobStatus) DeepCopy() *ArangoMLBatchJobStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJob) DeepCopyInto(out *ArangoMLCronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJob.
func (in *ArangoMLCronJob) DeepCopy() *ArangoMLCronJob {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLCronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobList) DeepCopyInto(out *ArangoMLCronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLCronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobList.
func (in *ArangoMLCronJobList) DeepCopy() *ArangoMLCronJobList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLCronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobSpec) DeepCopyInto(out *ArangoMLCronJobSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobSpec.
func (in *ArangoMLCronJobSpec) DeepCopy() *ArangoMLCronJobSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobStatus) DeepCopyInto(out *ArangoMLCronJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobStatus.
func (in *ArangoMLCronJobStatus) DeepCopy() *ArangoMLCronJobStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtension) DeepCopyInto(out *ArangoMLExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtension.
func (in *ArangoMLExtension) DeepCopy() *ArangoMLExtension {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionList) DeepCopyInto(out *ArangoMLExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionList.
func (in *ArangoMLExtensionList) DeepCopy() *ArangoMLExtensionList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpec) DeepCopyInto(out *ArangoMLExtensionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpec.
func (in *ArangoMLExtensionSpec) DeepCopy() *ArangoMLExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionStatus) DeepCopyInto(out *ArangoMLExtensionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionStatus.
func (in *ArangoMLExtensionStatus) DeepCopy() *ArangoMLExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorage) DeepCopyInto(out *ArangoMLStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorage.
func (in *ArangoMLStorage) DeepCopy() *ArangoMLStorage {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageList) DeepCopyInto(out *ArangoMLStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageList.
func (in *ArangoMLStorageList) DeepCopy() *ArangoMLStorageList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageS3Spec) DeepCopyInto(out *ArangoMLStorageS3Spec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageS3Spec.
func (in *ArangoMLStorageS3Spec) DeepCopy() *ArangoMLStorageS3Spec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpec) DeepCopyInto(out *ArangoMLStorageSpec) {
	*out = *in
	if in.ListenPort != nil {
		in, out := &in.ListenPort, &out.ListenPort
		*out = new(uint16)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ArangoMLStorageS3Spec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpec.
func (in *ArangoMLStorageSpec) DeepCopy() *ArangoMLStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageStatus) DeepCopyInto(out *ArangoMLStorageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageStatus.
func (in *ArangoMLStorageStatus) DeepCopy() *ArangoMLStorageStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageStatus)
	in.DeepCopyInto(out)
	return out
}
