// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	//"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	//"io/ioutil"
	"reflect"
	"testing"
)

func TestOcucpPfContainerCuCpResourceStatus001_NewOcucpPfContainerCuCpResourceStatus001(t *testing.T) {

	var nUEs int32 = 12
	nues := NewOcucpPfContainerCuCpResourceStatus001(nUEs)

	assert.Equal(t, reflect.TypeOf(OcucpPfContainerCuCpResourceStatus001{}), reflect.TypeOf(*nues), "OcucpPfContainerCuCpResourceStatus001{} types are mismatched")
	assert.Equal(t, nues.NumberOfActiveUes, nUEs, "OcucpPfContainerCuCpResourceStatus001{} NumberOfActiveUes are mismatched")
}

func TestOcucpPfContainerCuCpResourceStatus001_GetNumberOfActiveUes(t *testing.T) {

	var nUEs int32 = 12
	nues := NewOcucpPfContainerCuCpResourceStatus001(nUEs)

	assert.Equal(t, nues.GetNumberOfActiveUes(), nUEs, "Test_OcucpPfContainer_CuCpResourceStatus001 GetNumberOfActiveUes values mismatch")
}

func TestOcucpPfContainer_CuCpResourceStatus001_GetOcucpPfContainer_CuCpResourceStatus001(t *testing.T) {

	var nUEs int32 = 12
	nues1 := NewOcucpPfContainerCuCpResourceStatus001(nUEs)
	nues2 := nues1.GetOcucpPfContainerCuCpResourceStatus001()

	assert.Equal(t, nues1, nues2, "Test_OcucpPfContainer_CuCpResourceStatus001 GetNumberOfActiveUes values mismatch")
}
