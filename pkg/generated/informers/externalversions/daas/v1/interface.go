/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "sample-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Dreamxes returns a DreamxInformer.
	Dreamxes() DreamxInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Dreamxes returns a DreamxInformer.
func (v *version) Dreamxes() DreamxInformer {
	return &dreamxInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
