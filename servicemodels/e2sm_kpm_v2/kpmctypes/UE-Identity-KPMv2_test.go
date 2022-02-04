// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeUeIdentity(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2.UeIdentity{
		Value: "SomeUE",
	}
	xer, err := xerEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	//assert.Equal(t, 45, len(xer))
	t.Logf("UeIdentity XER\n%s", string(xer))
}

func Test_xerDecodeUeIdentity(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2.UeIdentity{
		Value: "SomeUE",
	}
	xer, err := xerEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	//assert.Equal(t, 45, len(xer))
	t.Logf("UeIdentity XER\n%s", string(xer))

	result, err := xerDecodeUeIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("UeIdentity XER - decoded\n%s", result)
}

func Test_perEncodeUeIdentity(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2.UeIdentity{
		Value: "SomeUE",
	}
	per, err := perEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	assert.Equal(t, 7, len(per))
	t.Logf("UeIdentity PER\n%v", hex.Dump(per))
}

func Test_perDecodeUeIdentity(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2.UeIdentity{
		Value: "SomeUE",
	}
	per, err := perEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	assert.Equal(t, 7, len(per))
	t.Logf("UeIdentity PER\n%v", hex.Dump(per))

	result, err := perDecodeUeIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("UeIdentity PER - decoded\n%v", result)
}
