/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "sample-operator/pkg/generated/clientset/versioned/typed/daas/v1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDaasV1 struct {
	*testing.Fake
}

func (c *FakeDaasV1) Dreamxes(namespace string) v1.DreamxInterface {
	return &FakeDreamxes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDaasV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}