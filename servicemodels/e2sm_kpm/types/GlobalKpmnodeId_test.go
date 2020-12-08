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

func TestGlobalKpmnodeID_NewGlobalKpmnodeID(t *testing.T) {

	var value int64 = 12
	gnbCuUpID := NewGnbCuUpID(value)

	assert.Equal(t, reflect.TypeOf(GnbCuUpID{}), reflect.TypeOf(*gnbCuUpID), "GnbCuUpID{} types are mismatched")
	assert.Equal(t, gnbCuUpID.Value, value, "GnbCuUpID{} values are mismatched")
}

func TestGlobalKpmnodeID_SetGlobalKpmnodeGnbID(t *testing.T) {

	bytes := []byte{0x22, 0x21, 0x20}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	gnbID := NewGnbIDChoice(NewBitString(value, length))

	var gnbCuID int64 = 12
	var gnbDuID int64 = 13
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID)).
		SetGnbCuUpID(NewGnbCuUpID(gnbCuID)).SetGnbDuID(NewGnbDuID(gnbDuID))
	globalKpmnodeID := NewGlobalKpmnodeID().SetGlobalKpmnodeGnbID(globalKpmnodeGnbID)

	assert.Equal(t, globalKpmnodeID.Gnb.GetGnbCuUpID().GetValue(), gnbCuID, "TestGlobalKpmnodeID gnbCuUpID values mismatch")
	assert.Equal(t, globalKpmnodeID.Gnb.GetGnbDuID().GetValue(), gnbDuID, "TestGlobalKpmnodeID gnbDuID values mismatch")
	assert.DeepEqual(t, globalKpmnodeID.Gnb.GetGlobalGnbID().GetPlmnID(), plmnID)
	assert.DeepEqual(t, globalKpmnodeID.Gnb.GetGlobalGnbID().GetGnbID(), gnbID)
}

func TestGlobalKpmnodeID_GetGlobalKpmnodeGnbID(t *testing.T) {

	bytes := []byte{0x22, 0x21, 0x20}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	gnbID := NewGnbIDChoice(NewBitString(value, length))

	var gnbCuID int64 = 12
	var gnbDuID int64 = 13
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID)).
		SetGnbCuUpID(NewGnbCuUpID(gnbCuID)).SetGnbDuID(NewGnbDuID(gnbDuID))
	globalKpmnodeID := NewGlobalKpmnodeID().SetGlobalKpmnodeGnbID(globalKpmnodeGnbID)
	globalKpmnodeGnbID2 := globalKpmnodeID.GetGlobalKpmnodeGnbID()

	assert.DeepEqual(t, globalKpmnodeGnbID, globalKpmnodeGnbID2)
}

func TestGlobalKpmnodeID_GetGlobalKpmnodeID(t *testing.T) {

	bytes := []byte{0x22, 0x21, 0x20}
	var value uint64 = 0x9bcd4
	var length uint32 = 22
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	gnbID := NewGnbIDChoice(NewBitString(value, length))

	var gnbCuID int64 = 12
	var gnbDuID int64 = 13
	globalKpmnodeGnbID := NewGlobalKpmnodeGnbID().SetGlobalGnbID(NewGlobalGnbID().SetPlmnID(plmnID).SetGnbID(gnbID)).
		SetGnbCuUpID(NewGnbCuUpID(gnbCuID)).SetGnbDuID(NewGnbDuID(gnbDuID))
	globalKpmnodeID1 := NewGlobalKpmnodeID().SetGlobalKpmnodeGnbID(globalKpmnodeGnbID)
	globalKpmnodeID2 := globalKpmnodeID1.GetGlobalKpmnodeID()

	assert.DeepEqual(t, globalKpmnodeID1, globalKpmnodeID2)
}
