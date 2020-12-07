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

func TestGnbCuUpID_NewGnbCuUpID(t *testing.T) {

	var value int64 = 12
	gnbCuUpID := NewGnbCuUpID(value)

	assert.Equal(t, reflect.TypeOf(GnbCuUpID{}), reflect.TypeOf(*gnbCuUpID), "GnbCuUpID{} types are mismatched")
	assert.Equal(t, gnbCuUpID.Value, value, "GnbCuUpID{} values are mismatched")
}

func TestGnbCuUpID_GetValue(t *testing.T) {

	var value int64 = 12
	gnbCuUpID := NewGnbCuUpID(value)

	assert.Equal(t, gnbCuUpID.GetValue(), value, "Test_GnbCuUpID GetValue values mismatch")
}

func TestGnbCuUpID_GetGnbCuUpID(t *testing.T) {

	var value int64 = 12
	gnbCuUpID1 := NewGnbCuUpID(value)
	gnbCuUpID2 := gnbCuUpID1.GetGnbCuUpID()

	assert.Equal(t, gnbCuUpID1.GetValue(), gnbCuUpID2.GetValue(), "Test_GnbCuUpID GetGnbCuUpID values mismatch")
}
