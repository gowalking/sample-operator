/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"sample-operator/pkg/controller"
	"sample-operator/pkg/generated/clientset/versioned"
	"sample-operator/pkg/generated/informers/externalversions"
	"sample-operator/pkg/signals"
)

var (
	master string
	config string
)

func init() {
	flag.StringVar(&config, "config", "", "Path to a Kubernetes config. Only required if out-of-cluster.")
	flag.StringVar(&master, "master", "", "The address of the Kubernetes API server. Overrides any value in Kubernetes config. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(master, config)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	daasClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := informers.NewSharedInformerFactory(kubeClient, time.Second*30)
	daasInformerFactory := externalversions.NewSharedInformerFactory(daasClient, time.Second*30)

	daasController := controller.NewController(
		kubeClient,
		daasClient,
		kubeInformerFactory.Core().V1().Endpoints(),
		daasInformerFactory.Daas().V1().Dreamxes(),
	)

	go kubeInformerFactory.Start(stopCh)
	go daasInformerFactory.Start(stopCh)
	if err = daasController.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}
