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
	assert.Assert(t, reflect.TypeOf(GlobalGnbID{}) == reflect.TypeOf(*globalGnb), "plmnID{} types are mismatched")

}

func TestGlobalGnbID_SetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21}
	plmnID := NewPlmnID()
	plmnID.SetValue(value)

	globalGnb := NewGlobalGnbID()
	globalGnb.SetPlmnID(*plmnID)

	assert.DeepEqual(t, globalGnb.PlmnID.GetValue(), value)

}

func TestGlobalGnbID_SetGnbID(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID := NewGnbIDChoice()
	choiceID.SetID(id)

	globalGnb := NewGlobalGnbID()
	globalGnb.SetGnbID(*choiceID)

	assert.Equal(t, globalGnb.GetGnbID().GetID().GetValue(), choiceID.GetID().GetValue())
	assert.Equal(t, globalGnb.GetGnbID().GetID().GetLen(), choiceID.GetID().GetLen())
	assert.Equal(t, globalGnb.GetGnbID().GetID(), choiceID.GetID())

}

func TestGlobalGnbID_GetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21}
	plmnID := NewPlmnID()
	plmnID.SetValue(value)

	globalGnb := NewGlobalGnbID()
	globalGnb.SetPlmnID(*plmnID)

	assert.DeepEqual(t, globalGnb.GetPlmnID().GetValue(), plmnID.GetValue())

}

func TestGlobalGnbID_GetGnbID(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID := NewGnbIDChoice()
	choiceID.SetID(id)

	globalGnb := NewGlobalGnbID()
	globalGnb.SetGnbID(*choiceID)

	assert.DeepEqual(t, globalGnb.GetGnbID().GetID().GetValue(), id.GetValue())
	assert.Equal(t, globalGnb.GetGnbID().GetID().GetLen(), id.GetLen())

}

func TestGlobalGnbID_GetGlobalGnbID(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID := NewGnbIDChoice()
	choiceID.SetID(id)

	value := []byte{0x22, 0x21}
	plmnID := NewPlmnID()
	plmnID.SetValue(value)

	globalGnb1 := NewGlobalGnbID()
	globalGnb1.SetGnbID(*choiceID)
	globalGnb1.SetPlmnID(*plmnID)

	globalGnb2 := globalGnb1.GetGlobalGnbID()

	assert.DeepEqual(t, globalGnb1.GetGnbID().GetID().GetValue(), globalGnb2.GetGnbID().GetID().GetValue())
	assert.Equal(t, globalGnb1.GetGnbID().GetID().GetLen(), globalGnb2.GetGnbID().GetID().GetLen())
	assert.DeepEqual(t, globalGnb1.GetPlmnID().GetValue(), globalGnb2.GetPlmnID().GetValue())

}
