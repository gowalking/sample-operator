/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

// Code generated by informer-gen. DO NOT EDIT.

package daas

import (
	v1 "sample-operator/pkg/generated/informers/externalversions/daas/v1"
	internalinterfaces "sample-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.factory, g.namespace, g.tweakListOptions)
}
