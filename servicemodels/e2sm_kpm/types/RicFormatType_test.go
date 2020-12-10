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

func TestRicFormatType_NewRicFormatType(t *testing.T) {

	var value int32 = 15
	ricFormatType := NewRicFormatType(value)

	assert.Equal(t, reflect.TypeOf(RicFormatType{}), reflect.TypeOf(*ricFormatType), "RicFormatType{} types are mismatched")
	assert.Equal(t, ricFormatType.Value, value, "RicFormatType{} values are mismatched")
}

func TestRicFormatType_GetValue(t *testing.T) {

	var value int32 = 15
	ricFormatType := NewRicFormatType(value)

	assert.Equal(t, ricFormatType.GetValue(), value, "Test_RicFormatType RicFormatType values mismatch")
}

func TestRicFormatType_GetRicFormatType(t *testing.T) {

	var value int32 = 15
	ricFormatType1 := NewRicFormatType(value)
	ricFormatType2 := ricFormatType1.GetRicFormatType()

	assert.Equal(t, ricFormatType1.GetValue(), ricFormatType2.GetValue(), "Test_RicFormatType RicFormatType values mismatch")
	assert.DeepEqual(t, ricFormatType1, ricFormatType2)
}
