// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestGlobalKpmnodeGnbID_NewGlobalKpmnodeGnbID(t *testing.T) {

	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID()

	assert.Equal(t, reflect.TypeOf(GlobalKpmnodeGnbID{}), reflect.TypeOf(*globalKpmnodeGnbID), "GlobalKpmnodeGnbID{} types are mismatched")
}

func TestGlobalKpmnodeGnbID_SetGlobalGnbID(t *testing.T) {

	bytes := []byte{0x22, 0x21}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID := NewPlmnID(bytes)
	gnbID := NewGnbIDChoice(NewBitString(value, length))
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID))

	assert.DeepEqual(t, globalKpmnodeGnbID.GlobalGnbID.PlmnID.GetValue(), bytes)
	assert.Equal(t, globalKpmnodeGnbID.GlobalGnbID.GnbID.ID.GetValue(), value, "Test_GlobalKpmnodeGnbID() SetGlobalGnbID BitString values mismatch")
	assert.Equal(t, globalKpmnodeGnbID.GlobalGnbID.GnbID.ID.GetLen(), length, "Test_GlobalKpmnodeGnbID() SetGlobalGnbID BitString lengths mismatch")
}

func TestGlobalKpmnodeGnbID_SetGnbCuUpID(t *testing.T) {

	var value int64 = 12
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGnbCuUpID(NewGnbCuUpID(value))

	assert.Equal(t, globalKpmnodeGnbID.GnbCuUpID.GetValue(), value, "Test_GlobalKpmnodeGnbID() SetGnbCuUpID GnbCuUpID values mismatch")
}

func TestGlobalKpmnodeGnbID_SetGnbDuID(t *testing.T) {

	var value int64 = 13
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGnbDuID(NewGnbDuID(value))

	assert.Equal(t, globalKpmnodeGnbID.GnbDuID.GetValue(), value, "Test_GlobalKpmnodeGnbID() SetGnbDuID GnbDuID values mismatch")
}

func TestGlobalKpmnodeGnbID_GetGlobalGnbID(t *testing.T) {

	bytes := []byte{0x22, 0x21}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID := NewPlmnID(bytes)
	gnbID := NewGnbIDChoice(NewBitString(value, length))
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID))

	assert.DeepEqual(t, globalKpmnodeGnbID.GetGlobalGnbID().GetPlmnID().GetValue(), bytes)
	assert.Equal(t, globalKpmnodeGnbID.GetGlobalGnbID().GetGnbID().GetGnbIDChoice().GetID().GetValue(), value, "Test_GlobalKpmnodeGnbID() GetGlobalGnbID BitString values mismatch")
	assert.Equal(t, globalKpmnodeGnbID.GetGlobalGnbID().GetGnbID().GetGnbIDChoice().GetID().GetLen(), length, "Test_GlobalKpmnodeGnbID() GetGlobalGnbID BitString lengths mismatch")
}

func TestGlobalKpmnodeGnbID_GetGnbCuUpID(t *testing.T) {

	var value int64 = 12
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGnbCuUpID(NewGnbCuUpID(value))

	assert.Equal(t, globalKpmnodeGnbID.GetGnbCuUpID().GetValue(), value, "Test_GlobalKpmnodeGnbID() GetGnbCuUpID GnbCuUpID values mismatch")
}

func TestGlobalKpmnodeGnbID_GetGnbDuID(t *testing.T) {

	var value int64 = 13
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGnbDuID(NewGnbDuID(value))

	assert.Equal(t, globalKpmnodeGnbID.GetGnbDuID().GetValue(), value, "Test_GlobalKpmnodeGnbID() GetGnbDuID GnbDuID values mismatch")
}

func TestGlobalKpmnodeGnbID_GetGlobalKpmnodeGnbID(t *testing.T) {

	bytes := []byte{0x22, 0x21}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID := NewPlmnID(bytes)
	gnbID := NewGnbIDChoice(NewBitString(value, length))

	var gnbCuID int64 = 12
	var gnbDuID int64 = 13
	globalKpmnodeGnbID1 := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID)).
		SetGnbCuUpID(NewGnbCuUpID(gnbCuID)).SetGnbDuID(NewGnbDuID(gnbDuID))

	globalKpmnodeGnbID2 := globalKpmnodeGnbID1.GetGlobalKpmnodeGnbID()

	assert.DeepEqual(t, globalKpmnodeGnbID1, globalKpmnodeGnbID2)
}
