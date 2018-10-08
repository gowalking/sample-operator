/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Dreamx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DreamxSpec   `json:"spec"`
	Status            DreamxStatus `json:"status"`
}

type DreamxSpec struct {
	Gv1 []Gv `json:"gvl"`
}

type Gv struct {
	Groupversion string     `json:"groupversion"`
	Consumers    []Consumer `json:"consumers"`
	Providers    []Provider `json:"providers"`
}

type Consumer struct {
	Host   string `json:"host"`
	Params string `json:"params"`
}

type Provider struct {
	HostPort string `json:"hostport"`
	Params   string `json:"params"`
	Pod      string `json:"pod"`
}

type DreamxStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DreamxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Dreamx `json:"items"`
}
