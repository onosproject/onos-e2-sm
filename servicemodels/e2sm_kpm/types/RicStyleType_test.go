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

func TestRicStyleType_NewRicStyleType(t *testing.T) {

	var value int32 = 14
	ricStyleType := NewRicStyleType(value)

	assert.Equal(t, reflect.TypeOf(RicStyleType{}), reflect.TypeOf(*ricStyleType), "RicStyleType{} types are mismatched")
	assert.Equal(t, ricStyleType.Value, value, "RicStyleType{} values are mismatched")
}

func TestRicStyleType_GetValue(t *testing.T) {

	var value int32 = 14
	ricStyleType := NewRicStyleType(value)

	assert.Equal(t, ricStyleType.GetValue(), value, "Test_RicStyleType RicStyleType values mismatch")
}

func TestRicStyleType_GetRicStyleType(t *testing.T) {

	var value int32 = 14
	ricStyleType1 := NewRicStyleType(value)
	ricStyleType2 := ricStyleType1.GetRicStyleType()

	assert.Equal(t, ricStyleType1.GetValue(), ricStyleType2.GetValue(), "Test_RicStyleType RicStyleType values mismatch")
	assert.DeepEqual(t, ricStyleType1, ricStyleType2)
}
