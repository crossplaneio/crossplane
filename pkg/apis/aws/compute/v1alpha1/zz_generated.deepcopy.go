// +build !ignore_autogenerated

/*
Copyright 2018 The Crossplane Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCluster) DeepCopyInto(out *EKSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCluster.
func (in *EKSCluster) DeepCopy() *EKSCluster {
	if in == nil {
		return nil
	}
	out := new(EKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterList) DeepCopyInto(out *EKSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterList.
func (in *EKSClusterList) DeepCopy() *EKSClusterList {
	if in == nil {
		return nil
	}
	out := new(EKSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterSpec) DeepCopyInto(out *EKSClusterSpec) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.WorkerNodes.DeepCopyInto(&out.WorkerNodes)
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ClassRef != nil {
		in, out := &in.ClassRef, &out.ClassRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterSpec.
func (in *EKSClusterSpec) DeepCopy() *EKSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EKSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterStatus) DeepCopyInto(out *EKSClusterStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.BindingStatusPhase = in.BindingStatusPhase
	out.ConnectionSecretRef = in.ConnectionSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterStatus.
func (in *EKSClusterStatus) DeepCopy() *EKSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EKSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpPermission) DeepCopyInto(out *IpPermission) {
	*out = *in
	if in.IpRanges != nil {
		in, out := &in.IpRanges, &out.IpRanges
		*out = make([]IpRange, len(*in))
		copy(*out, *in)
	}
	if in.Ipv6Ranges != nil {
		in, out := &in.Ipv6Ranges, &out.Ipv6Ranges
		*out = make([]Ipv6Range, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupsIDs != nil {
		in, out := &in.SecurityGroupsIDs, &out.SecurityGroupsIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpPermission.
func (in *IpPermission) DeepCopy() *IpPermission {
	if in == nil {
		return nil
	}
	out := new(IpPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRange) DeepCopyInto(out *IpRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRange.
func (in *IpRange) DeepCopy() *IpRange {
	if in == nil {
		return nil
	}
	out := new(IpRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ipv6Range) DeepCopyInto(out *Ipv6Range) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ipv6Range.
func (in *Ipv6Range) DeepCopy() *Ipv6Range {
	if in == nil {
		return nil
	}
	out := new(Ipv6Range)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupList) DeepCopyInto(out *SecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupList.
func (in *SecurityGroupList) DeepCopy() *SecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSpec) DeepCopyInto(out *SecurityGroupSpec) {
	*out = *in
	if in.IpPermissions != nil {
		in, out := &in.IpPermissions, &out.IpPermissions
		*out = make([]IpPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IpPermissionsEgress != nil {
		in, out := &in.IpPermissionsEgress, &out.IpPermissionsEgress
		*out = make([]IpPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ClassRef != nil {
		in, out := &in.ClassRef, &out.ClassRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSpec.
func (in *SecurityGroupSpec) DeepCopy() *SecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupStatus) DeepCopyInto(out *SecurityGroupStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.BindingStatusPhase = in.BindingStatusPhase
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupStatus.
func (in *SecurityGroupStatus) DeepCopy() *SecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerNodesSpec) DeepCopyInto(out *WorkerNodesSpec) {
	*out = *in
	if in.NodeAutoScalingGroupMinSize != nil {
		in, out := &in.NodeAutoScalingGroupMinSize, &out.NodeAutoScalingGroupMinSize
		*out = new(int)
		**out = **in
	}
	if in.NodeAutoScalingGroupMaxSize != nil {
		in, out := &in.NodeAutoScalingGroupMaxSize, &out.NodeAutoScalingGroupMaxSize
		*out = new(int)
		**out = **in
	}
	if in.NodeVolumeSize != nil {
		in, out := &in.NodeVolumeSize, &out.NodeVolumeSize
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerNodesSpec.
func (in *WorkerNodesSpec) DeepCopy() *WorkerNodesSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerNodesSpec)
	in.DeepCopyInto(out)
	return out
}
