// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestGlobalGnbID_NewGlobalGnbID(t *testing.T) {

	globalGnb := NewGlobalGnbID()

	assert.Equal(t, reflect.TypeOf(GlobalGnbID{}), reflect.TypeOf(*globalGnb), "GlobalGnbID{} types are mismatched")
}

func TestGlobalGnbID_SetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21}
	globalGnb := NewGlobalGnbID().SetPlmnID(NewPlmnID(value))

	assert.DeepEqual(t, globalGnb.PlmnID.GetValue(), value)
}

func TestGlobalGnbID_SetGnbID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	choiceID := NewGnbIDChoice(NewBitString(value, length))
	globalGnb := NewGlobalGnbID().SetGnbID(NewGnbIDChoice(NewBitString(value, length)))

	assert.Equal(t, globalGnb.GnbID.ID.GetValue(), choiceID.GetID().GetValue(), "Test_GlobalGnbID() SetGnbID values mismatch")
	assert.Equal(t, globalGnb.GnbID.ID.GetLen(), choiceID.GetID().GetLen(), "Test_GlobalGnbID() SetGnbID lengths mismatch")
	assert.DeepEqual(t, globalGnb.GnbID.GetID(), choiceID.GetID())
}

func TestGlobalGnbID_GetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21}
	plmnID := NewPlmnID(value)
	globalGnb := NewGlobalGnbID().SetPlmnID(NewPlmnID(value))

	assert.DeepEqual(t, globalGnb.GetPlmnID().GetValue(), plmnID.GetValue())
}

func TestGlobalGnbID_GetGnbID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	choiceID := NewGnbIDChoice(NewBitString(value, length))
	globalGnb := NewGlobalGnbID().SetGnbID(NewGnbIDChoice(NewBitString(value, length)))

	assert.Equal(t, globalGnb.GetGnbID().GetID().GetValue(), choiceID.GetID().GetValue(), "Test_GlobalGnbID() GetGnbID values mismatch")
	assert.Equal(t, globalGnb.GetGnbID().GetID().GetLen(), choiceID.GetID().GetLen(), "Test_GlobalGnbID() GetGnbID lengths mismatch")
}

func TestGlobalGnbID_GetGlobalGnbID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bytes := []byte{0x22, 0x21}
	globalGnb1 := NewGlobalGnbID().SetGnbID(NewGnbIDChoice(NewBitString(value, length))).SetPlmnID(NewPlmnID(bytes))
	globalGnb2 := globalGnb1.GetGlobalGnbID()

	assert.Equal(t, globalGnb1.GetGnbID().GetID().GetValue(), globalGnb2.GetGnbID().GetID().GetValue(), "Test_GlobalGnbID() GetGlobalGnbID values mismatch")
	assert.Equal(t, globalGnb1.GetGnbID().GetID().GetLen(), globalGnb2.GetGnbID().GetID().GetLen(), "Test_GlobalGnbID() GetGlobalGnbID lengths mismatch")
	assert.DeepEqual(t, globalGnb1.GetPlmnID().GetValue(), globalGnb2.GetPlmnID().GetValue())
}
