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

func TestGnbCuCpName_NewGnbCuCpName(t *testing.T) {

	var name string = "ONF"
	gnbCuCpName := NewGnbCuCpName(name)

	assert.Equal(t, reflect.TypeOf(GnbCuCpName{}), reflect.TypeOf(*gnbCuCpName), "GnbCuCpName{} types are mismatched")
	assert.Equal(t, gnbCuCpName.Value, name, "GnbCuCpName{} names are mismatched")
}

func TestGnbCuCpName_GetValue(t *testing.T) {

	var name string = "ONF"
	gnbCuCpName1 := NewGnbCuCpName(name)

	assert.Equal(t, gnbCuCpName1.GetValue(), name, "TestGnbCuCpName_GetValue GetValue values mismatch")
}

func TestGnbCuCpName_GetGnbCuCpName(t *testing.T) {

	var name string = "ONF"
	gnbCuCpName1 := NewGnbCuCpName(name)
	gnbCuCpName2 := gnbCuCpName1.GetGnbCuCpName()

	assert.Equal(t, gnbCuCpName1, gnbCuCpName2, "TestGnbCuCpName_GetGnbCuCpName GetGnbCuCpName values mismatch")
}
