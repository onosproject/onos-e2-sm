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

	variable := NewGnbDuID()
	assert.Equal(t, reflect.TypeOf(GnbDuID{}), reflect.TypeOf(*variable), "GnbDuID{} types are mismatched")
}

func TestGnbDuID_SetValue(t *testing.T) {

	var value int64 = 13
	gnbDuID := NewGnbDuID().SetValue(value)

	assert.Equal(t, gnbDuID.Value, value, "Mismatch of GnbDuID values")
}

func TestGGnbDuID_GetValue(t *testing.T) {

	var value int64 = 13
	gnbDuID := NewGnbDuID().SetValue(value)

	assert.Equal(t, gnbDuID.GetValue(), value, "Mismatch of GnbDuID values")
}

func TestGnbDuID_GetGnbCuUpID(t *testing.T) {

	gnbDuID1 := NewGnbDuID().SetValue(13)
	gnbDuID2 := gnbDuID1.GetGnbDuID()

	assert.Equal(t, gnbDuID1.GetValue(), gnbDuID2.GetValue(), "Mismatch of GnbDuID values")
}
