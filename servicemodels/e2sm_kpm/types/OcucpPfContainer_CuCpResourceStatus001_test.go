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

func TestOcucpPfContainer_CuCpResourceStatus001_NewOcucpPfContainer_CuCpResourceStatus001(t *testing.T) {

	var nUEs int32 = 12
	nues := NewOcucpPfContainer_CuCpResourceStatus001(nUEs)

	assert.Equal(t, reflect.TypeOf(OcucpPfContainer_CuCpResourceStatus001{}), reflect.TypeOf(*nues), "OcucpPfContainer_CuCpResourceStatus001{} types are mismatched")
	assert.Equal(t, nues.NumberOfActiveUes, nUEs, "OcucpPfContainer_CuCpResourceStatus001{} NumberOfActiveUes are mismatched")
}

func TestOcucpPfContainer_CuCpResourceStatus001_GetNumberOfActiveUes(t *testing.T) {

	var nUEs int32 = 12
	nues := NewOcucpPfContainer_CuCpResourceStatus001(nUEs)

	assert.Equal(t, nues.GetNumberOfActiveUes(), nUEs, "Test_OcucpPfContainer_CuCpResourceStatus001 GetNumberOfActiveUes values mismatch")
}

func TestOcucpPfContainer_CuCpResourceStatus001_GetOcucpPfContainer_CuCpResourceStatus001(t *testing.T) {

	var nUEs int32 = 12
	nues1 := NewOcucpPfContainer_CuCpResourceStatus001(nUEs)
	nues2 := nues1.GetOcucpPfContainer_CuCpResourceStatus001()

	assert.Equal(t, nues1, nues2, "Test_OcucpPfContainer_CuCpResourceStatus001 GetNumberOfActiveUes values mismatch")
}
