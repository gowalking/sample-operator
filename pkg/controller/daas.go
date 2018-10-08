/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package controller

import (
	"fmt"
	daasv1 "sample-operator/pkg/apis/daas/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
)

func (self *DaasController) daasAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	self.daasWorkQueue.AddRateLimited(key)
}

func (self *DaasController) daasUpdate(oldObj, newObj interface{}) {
	oldDaas := oldObj.(*daasv1.Dreamx)
	newDaas := newObj.(*daasv1.Dreamx)
	if newDaas.ResourceVersion == oldDaas.ResourceVersion {
		return
	}
	fmt.Println("DaasUpdate ------> ", newDaas.Name)
}

func (self *DaasController) daasSync(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(err)
		return err
	}
	dreamx, err := self.dClientSet.Daas().Dreamxes(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		runtime.HandleError(err)
		return err
	}
	fmt.Println("DaasAdd ------> ", dreamx.Name)
	return nil
}
