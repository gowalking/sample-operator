/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package controller

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (self *DaasController) endpointsUpdate(oldObj, newObj interface{}) {
	oldEndpoint := oldObj.(*corev1.Endpoints)
	newEndpoint := newObj.(*corev1.Endpoints)
	if newEndpoint.ResourceVersion == oldEndpoint.ResourceVersion {
		return
	}
	fmt.Println("EndUpdate --------> ", newEndpoint.Name)
}

func (self *DaasController) endpointsDelete(obj interface{}) {
	endpoint := obj.(*corev1.Endpoints)
	fmt.Println("EndDelete --------> ", endpoint.Name)
}
