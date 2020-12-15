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

func TestE2SmKpmIndicationHeaderFormat1_NewE2SmKpmIndicationHeaderFormat1(t *testing.T) {

	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1()

	assert.Equal(t, reflect.TypeOf(E2SmKpmIndicationHeaderFormat1{}), reflect.TypeOf(*e2SmKpmIndicationHeaderFormat1), "E2SmKpmIndicationHeaderFormat1{} types are mismatched")
}

func TestE2SmKpmIndicationHeaderFormat1_SetIDglobalKpmnodeID(t *testing.T) {

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
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID)

	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.IDglobalKpmnodeID.Gnb.GetGnbCuUpID().GetValue(), gnbCuID, "TestE2SmKpmIndicationHeaderFormat1() SetIDglobalKpmnodeID gnbCuUpID values mismatch")
	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.IDglobalKpmnodeID.Gnb.GetGnbDuID().GetValue(), gnbDuID, "TestE2SmKpmIndicationHeaderFormat1() SetIDglobalKpmnodeID gnbDuID values mismatch")
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.IDglobalKpmnodeID.Gnb.GetGlobalGnbID().GetPlmnID(), plmnID)
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.IDglobalKpmnodeID.Gnb.GetGlobalGnbID().GetGnbID(), gnbID)
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.IDglobalKpmnodeID, globalKpmnodeID)
}

func TestE2SmKpmIndicationHeaderFormat1_SetNRcgi(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bytes := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	nrcgi := NewNRcgi().SetPlmnID(plmnID).SetNRcellID(NewNrcellIdentity(NewBitString(value, length)))
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetNRcgi(nrcgi)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.NRcgi.GetPlmnID().GetValue(), bytes)
	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.NRcgi.GetNRcellID().GetValue().GetValue(), value, "TestE2SmKpmIndicationHeaderFormat1() SetNRcgi values mismatch")
	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.NRcgi.GetNRcellID().GetValue().GetLen(), length, "TestE2SmKpmIndicationHeaderFormat1() SetNRcgi lengths mismatch")
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.NRcgi, nrcgi)
}

func TestE2SmKpmIndicationHeaderFormat1_SetPlmnID(t *testing.T) {

	bytes := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetPlmnID(plmnID)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.PlmnID.GetValue(), bytes)
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.PlmnID, plmnID)
}

func TestE2SmKpmIndicationHeaderFormat1_SetSliceID(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetSliceID(snssai)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.SliceID.GetSD(), sd)
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.SliceID.GetSSt(), sst)
	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.SliceID, snssai)
}

func TestE2SmKpmIndicationHeaderFormat1_SetFiveQi(t *testing.T) {

	var fiveqi int32 = 12
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetFiveQi(fiveqi)

	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.FiveQi, fiveqi, "TestE2SmKpmIndicationHeaderFormat1() SetFiveQi values mismatch")
}

func TestE2SmKpmIndicationHeaderFormat1_SetQci(t *testing.T) {

	var qci int32 = 13
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetQci(qci)

	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.Qci, qci, "TestE2SmKpmIndicationHeaderFormat1() SetQci values mismatch")
}

func TestE2SmKpmIndicationHeaderFormat1_GetIDglobalKpmnodeID(t *testing.T) {

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
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.GetIDlobalKpmnodeID(), globalKpmnodeID)
}

func TestE2SmKpmIndicationHeaderFormat1_GetNRcgi(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bytes := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	nrcgi := NewNRcgi().SetPlmnID(plmnID).SetNRcellID(NewNrcellIdentity(NewBitString(value, length)))
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetNRcgi(nrcgi)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.GetNRcgi(), nrcgi)
}

func TestE2SmKpmIndicationHeaderFormat1_GetPlmnID(t *testing.T) {

	bytes := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetPlmnID(plmnID)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.GetPlmnID(), plmnID)
}

func TestE2SmKpmIndicationHeaderFormat1_GetSliceID(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetSliceID(snssai)

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1.GetSliceID(), snssai)
}

func TestE2SmKpmIndicationHeaderFormat1_GetFiveQi(t *testing.T) {

	var fiveqi int32 = 12
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetFiveQi(fiveqi)

	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.GetFiveQi(), fiveqi, "TestE2SmKpmIndicationHeaderFormat1() GetFiveQi values mismatch")
}

func TestE2SmKpmIndicationHeaderFormat1_GetQci(t *testing.T) {

	var qci int32 = 12
	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetQci(qci)

	assert.Equal(t, e2SmKpmIndicationHeaderFormat1.GetQci(), qci, "TestE2SmKpmIndicationHeaderFormat1() Qci values mismatch")
}

func TestE2SmKpmIndicationHeaderFormat1_GetE2SmKpmIndicationHeaderFormat1(t *testing.T) {

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

	var nrcgiValue uint64 = 0x9bcd4
	var nrcgiLength uint32 = 22
	nrcgiBytes := []byte{0x22, 0x21, 0x20}
	nrcgiPlmn, err := NewPlmnID(nrcgiBytes)
	assert.NilError(t, err)
	nrcgi := NewNRcgi().SetPlmnID(nrcgiPlmn).SetNRcellID(NewNrcellIdentity(NewBitString(nrcgiValue, nrcgiLength)))

	indHdrBytes := []byte{0x22, 0x21, 0x20}
	indHdrPlmnID, err := NewPlmnID(indHdrBytes)
	assert.NilError(t, err)

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)

	var fiveqi int32 = 12
	var qci int32 = 12

	e2SmKpmIndicationHeaderFormat1_1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID).
		SetNRcgi(nrcgi).SetPlmnID(indHdrPlmnID).SetSliceID(snssai).SetFiveQi(fiveqi).SetQci(qci)
	e2SmKpmIndicationHeaderFormat1_2 := e2SmKpmIndicationHeaderFormat1_1.GetE2SmKpmIndicationHeaderFormat1()

	assert.DeepEqual(t, e2SmKpmIndicationHeaderFormat1_1, e2SmKpmIndicationHeaderFormat1_2)
}
