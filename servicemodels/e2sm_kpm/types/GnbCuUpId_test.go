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

func TestGnbCuUpID_NewBGnbCuUpID(t *testing.T) {

	variable := NewGnbCuUpID()
	assert.Equal(t, reflect.TypeOf(GnbCuUpID{}), reflect.TypeOf(*variable), "GnbCuUpID{} types are mismatched")
}

func TestGnbCuUpID_SetValue(t *testing.T) {

	var value int64 = 12
	gnbCuUpID := NewGnbCuUpID().SetValue(value)

	assert.Equal(t, gnbCuUpID.Value, value, "Mismatch of GnbCuUpID values")
}

func TestGnbCuUpID_GetValue(t *testing.T) {

	var value int64 = 12
	gnbCuUpID := NewGnbCuUpID().SetValue(value)

	assert.Equal(t, gnbCuUpID.GetValue(), value, "Mismatch of GnbCuUpID values")
}

func TestGnbCuUpID_GetGnbCuUpID(t *testing.T) {

	gnbCuUpID1 := NewGnbCuUpID().SetValue(13)
	gnbCuUpID2 := gnbCuUpID1.GetGnbCuUpID()

	assert.Equal(t, gnbCuUpID1.GetValue(), gnbCuUpID2.GetValue(), "Mismatch of GnbCuUpID values")
}
