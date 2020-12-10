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

func TestRicStyleName_NewRicStyleName(t *testing.T) {

	var value string = "ONF"
	ricStyleName := NewRicStyleName(value)

	assert.Equal(t, reflect.TypeOf(RicStyleName{}), reflect.TypeOf(*ricStyleName), "RicStyleName{} types are mismatched")
	assert.Equal(t, ricStyleName.Value, value, "RicStyleName{} values are mismatched")
}

func TestRicStyleName_GetValue(t *testing.T) {

	var value string = "ONF"
	ricStyleName := NewRicStyleName(value)

	assert.Equal(t, ricStyleName.GetValue(), value, "Test_RicStyleName RicStyleName values mismatch")
}

func TestRicStyleName_GetRicStyleName(t *testing.T) {

	var value string = "ONF"
	ricStyleName1 := NewRicStyleName(value)
	ricStyleName2 := ricStyleName1.GetRicStyleName()

	assert.Equal(t, ricStyleName1.GetValue(), ricStyleName2.GetValue(), "Test_RicStyleName RicStyleName values mismatch")
	assert.DeepEqual(t, ricStyleName1, ricStyleName2)
}
