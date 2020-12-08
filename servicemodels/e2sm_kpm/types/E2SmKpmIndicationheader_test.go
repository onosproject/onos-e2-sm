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

func TestE2SmKpmIndicationHeader_NewE2SmKpmIndicationHeader(t *testing.T) {

	e2SmKpmIndicationHeader := NewE2SmKpmIndicationHeader()

	assert.Equal(t, reflect.TypeOf(E2SmKpmIndicationHeader{}), reflect.TypeOf(*e2SmKpmIndicationHeader), "E2SmKpmIndicationHeader{} types are mismatched")
}

func TestE2SmKpmIndicationHeader_SetE2SmKpmIndicationHeaderFormat1(t *testing.T) {

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

	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID).
		SetNRcgi(nrcgi).SetPlmnID(indHdrPlmnID).SetSliceID(snssai).SetFiveQi(fiveqi).SetQci(qci)
	e2SmKpmIndicationHeader := NewE2SmKpmIndicationHeader().SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1)

	assert.DeepEqual(t, e2SmKpmIndicationHeader.E2SmKpmIndicationHeaderFormat1, e2SmKpmIndicationHeaderFormat1)
}

func TestE2SmKpmIndicationHeader_GetE2SmKpmIndicationHeaderFormat1(t *testing.T) {

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

	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID).
		SetNRcgi(nrcgi).SetPlmnID(indHdrPlmnID).SetSliceID(snssai).SetFiveQi(fiveqi).SetQci(qci)
	e2SmKpmIndicationHeader := NewE2SmKpmIndicationHeader().SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1)

	assert.DeepEqual(t, e2SmKpmIndicationHeader.GetE2SmKpmIndicationHeaderFormat1(), e2SmKpmIndicationHeaderFormat1)
}

func TestE2SmKpmIndicationHeader_GetE2SmKpmIndicationHeader(t *testing.T) {

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

	e2SmKpmIndicationHeaderFormat1 := NewE2SmKpmIndicationHeaderFormat1().SetIDglobalKpmnodeID(globalKpmnodeID).
		SetNRcgi(nrcgi).SetPlmnID(indHdrPlmnID).SetSliceID(snssai).SetFiveQi(fiveqi).SetQci(qci)
	e2SmKpmIndicationHeader1 := NewE2SmKpmIndicationHeader().SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1)
	e2SmKpmIndicationHeader2 := e2SmKpmIndicationHeader1.GetE2SmKpmIndicationHeader()

	assert.DeepEqual(t, e2SmKpmIndicationHeader1, e2SmKpmIndicationHeader2)
}
