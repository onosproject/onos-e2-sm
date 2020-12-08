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

func TestGnbDuID_NewGnbDuID(t *testing.T) {

	var value int64 = 13
	gnbDuID := NewGnbDuID(value)

	assert.Equal(t, reflect.TypeOf(GnbDuID{}), reflect.TypeOf(*gnbDuID), "GnbDuID{} types are mismatched")
	assert.Equal(t, gnbDuID.Value, value, "GnbDuID{} values are mismatched")
}

func TestGGnbDuID_GetValue(t *testing.T) {

	var value int64 = 13
	gnbDuID := NewGnbDuID(value)

	assert.Equal(t, gnbDuID.GetValue(), value, "Test_GnbDuID GetValue values mismatch")
}

func TestGnbDuID_GetGnbCuUpID(t *testing.T) {

	var value int64 = 13
	gnbDuID1 := NewGnbDuID(value)
	gnbDuID2 := gnbDuID1.GetGnbDuID()

	assert.Equal(t, gnbDuID1.GetValue(), gnbDuID2.GetValue(), "Test_GnbDuID GetGnbCuUpID values mismatch")
}
