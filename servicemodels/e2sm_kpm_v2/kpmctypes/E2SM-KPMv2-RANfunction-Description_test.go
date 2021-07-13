// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMRanfunctionDescription() (*e2sm_kpm_v2.E2SmKpmRanfunctionDescription, error) {

	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellGlobalID, _ := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, []byte{0xd5, 0xbc, 0x09, 0x00, 0x00}) // 36 bits

	var cellObjID string = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, _ := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName string = "onf"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	var measTypeName string = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 2
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rknl, retsl, rrsl)
	newE2SmKpmPdu.RanFunctionName.RanFunctionInstance = rfi
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func Test_xerEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	//assert.Equal(t, 2704, len(xer))
	//assert.Equal(t, 372, len(xer)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	//assert.Equal(t, 2704, len(xer))
	//assert.Equal(t, 372, len(xer)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmRanfunctionDescription XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 110, len(per))
	//assert.Equal(t, 33, len(per)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription PER\n%v", hex.Dump(per))
}

func Test_perDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 110, len(per))
	//assert.Equal(t, 33, len(per)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmKpmRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmRanfunctionDescription PER - decoded\n%s", result)
}

func Test_perDecodeRadysisBytes(t *testing.T) {

	radisysBytesRanFunctionDefinition := []byte{0x70, 0x18, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D, 0x00,
		0x00, 0x18, 0x31, 0x2E, 0x33, 0x2E, 0x36, 0x2E, 0x31, 0x2E, 0x34, 0x2E, 0x31, 0x2E, 0x35, 0x33,
		0x31, 0x34, 0x38, 0x2E, 0x31, 0x2E, 0x32, 0x2E, 0x32, 0x2E, 0x32, 0x05, 0x00, 0x4B, 0x50, 0x4D,
		0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74, 0x6F, 0x72, 0x00, 0x00, 0x40, 0x00, 0x13, 0xF1, 0x84, 0x50,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xF1, 0x84, 0x00, 0x00,
		0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6F, 0x64, 0x69, 0x63, 0x20,
		0x72, 0x65, 0x70, 0x6F, 0x72, 0x74, 0x00, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4E,
		0x6F, 0x64, 0x65, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x00,
		0x01, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74,
		0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52,
		0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2E,
		0x53, 0x75, 0x6D, 0x00, 0x00, 0x01, 0x42, 0xA0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E,
		0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00,
		0x02, 0x43, 0xC0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45, 0x73, 0x74,
		0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x72, 0x65, 0x63, 0x6F, 0x6E, 0x66, 0x69, 0x67, 0x46, 0x61,
		0x69, 0x6C, 0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52,
		0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x48, 0x4F, 0x46, 0x61, 0x69, 0x6C,
		0x00, 0x00, 0x04, 0x42, 0xE0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45,
		0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x4F, 0x74, 0x68, 0x65, 0x72, 0x00, 0x00, 0x05,
		0x41, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x65, 0x61, 0x6E, 0x00, 0x00,
		0x06, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x61, 0x78, 0x00, 0x00,
		0x07, 0x00, 0x01, 0x00, 0x01}
	t.Logf("Radysis E2SmKpmRanfunctionDescription PER\n%v", hex.Dump(radisysBytesRanFunctionDefinition))

	result, err := PerDecodeE2SmKpmRanfunctionDescription(radisysBytesRanFunctionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Radysis E2SmKpmRanfunctionDescription PER - decoded\n%s", result)
}
