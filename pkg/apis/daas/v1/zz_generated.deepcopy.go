// +build !ignore_autogenerated

/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consumer) DeepCopyInto(out *Consumer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consumer.
func (in *Consumer) DeepCopy() *Consumer {
	if in == nil {
		return nil
	}
	out := new(Consumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dreamx) DeepCopyInto(out *Dreamx) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dreamx.
func (in *Dreamx) DeepCopy() *Dreamx {
	if in == nil {
		return nil
	}
	out := new(Dreamx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dreamx) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DreamxList) DeepCopyInto(out *DreamxList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dreamx, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DreamxList.
func (in *DreamxList) DeepCopy() *DreamxList {
	if in == nil {
		return nil
	}
	out := new(DreamxList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DreamxList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DreamxSpec) DeepCopyInto(out *DreamxSpec) {
	*out = *in
	if in.Gv1 != nil {
		in, out := &in.Gv1, &out.Gv1
		*out = make([]Gv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DreamxSpec.
func (in *DreamxSpec) DeepCopy() *DreamxSpec {
	if in == nil {
		return nil
	}
	out := new(DreamxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DreamxStatus) DeepCopyInto(out *DreamxStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DreamxStatus.
func (in *DreamxStatus) DeepCopy() *DreamxStatus {
	if in == nil {
		return nil
	}
	out := new(DreamxStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gv) DeepCopyInto(out *Gv) {
	*out = *in
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]Consumer, len(*in))
		copy(*out, *in)
	}
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]Provider, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gv.
func (in *Gv) DeepCopy() *Gv {
	if in == nil {
		return nil
	}
	out := new(Gv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}