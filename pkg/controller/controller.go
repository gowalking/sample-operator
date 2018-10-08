/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package controller

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeInformerv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	kubeListerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"sample-operator/pkg/generated/clientset/versioned"
	daasScheme "sample-operator/pkg/generated/clientset/versioned/scheme"
	daasInformerv1 "sample-operator/pkg/generated/informers/externalversions/daas/v1"
	daasListerv1 "sample-operator/pkg/generated/listers/daas/v1"
)

type DaasController struct {
	endpointSynced cache.InformerSynced
	endpointLister kubeListerv1.EndpointsLister

	daasSynced    cache.InformerSynced
	daasLister    daasListerv1.DreamxLister
	daasWorkQueue workqueue.RateLimitingInterface

	kClientSet kubernetes.Interface
	dClientSet versioned.Interface

	recorder record.EventRecorder
	cronSets map[string]interface{}
	cronLock *sync.RWMutex
}

func NewController(
	kClientSet kubernetes.Interface,
	dClientSet versioned.Interface,
	endPointsInfomer kubeInformerv1.EndpointsInformer,
	daasInformer daasInformerv1.DreamxInformer) *DaasController {

	daasScheme.AddToScheme(scheme.Scheme)
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kClientSet.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "sample-controller"})

	DaasController := &DaasController{
		endpointSynced: endPointsInfomer.Informer().HasSynced,
		endpointLister: endPointsInfomer.Lister(),

		daasSynced:    daasInformer.Informer().HasSynced,
		daasLister:    daasInformer.Lister(),
		daasWorkQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Dreamx"),

		kClientSet: kClientSet,
		dClientSet: dClientSet,

		recorder: recorder,
		cronLock: new(sync.RWMutex),
		cronSets: make(map[string]interface{}),
	}

	daasInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    DaasController.daasAdd,
		UpdateFunc: DaasController.daasUpdate,
	})

	endPointsInfomer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: DaasController.endpointsUpdate,
		DeleteFunc: DaasController.endpointsDelete,
	})
	return DaasController
}

func (self *DaasController) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer self.daasWorkQueue.ShutDown()

	if ok := cache.WaitForCacheSync(stopCh, self.daasSynced); !ok {
		glog.Errorln("failed to wait for caches to sync. ")
		return fmt.Errorf("failed to wait for caches to sync. ")
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(self.runWorker(self.daasWorkQueue, self.daasSync), time.Second, stopCh)
	}

	<-stopCh
	return nil
}

func (self *DaasController) runWorker(workQueue workqueue.RateLimitingInterface, syncHandler func(key string) error) func() {
	return func() {
		for self.processNextWorkItem(workQueue, syncHandler) {
		}
	}
}

func (self *DaasController) processNextWorkItem(workQueue workqueue.RateLimitingInterface, syncHandler func(key string) error) bool {
	obj, shutdown := workQueue.Get()
	if shutdown {
		return false
	}
	err := func(obj interface{}) error {
		defer workQueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			workQueue.Forget(obj)
			glog.Errorf("expected string in workqueue but got %#v.", obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v. ", obj))
			return nil
		}
		if err := syncHandler(key); err != nil {
			glog.Errorf("error syncing '%s': %s. ", key, err.Error())
			return fmt.Errorf("error syncing '%s': %s. ", key, err.Error())
		}
		workQueue.Forget(obj)
		glog.Infof("Successfully synced '%s'.", key)
		return nil
	}(obj)
	if err != nil {
		glog.Errorln("processNextWorkItem anonymous function occur error.")
		runtime.HandleError(err)
		return true
	}
	return true
}
