// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

type servicemodel string

const smname = "e2sm_ni"
const smversion = "v1beta1"
const modulename = "e2sm_ni.so.1.0.1"

func (sm servicemodel) ServiceModelData() (string, string, string) {
	return smname, smversion, modulename
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
