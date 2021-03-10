// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMRanfunctionDescription() *e2sm_kpm_v2.E2SmKpmRanfunctionDescription {

	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	cellGlobalID, _ := pdubuilder.CreateCellGlobalID_NRCGI(plmnID, 0xABCDEF012 << 28) // 36 bits

	var cellObjID string = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, *cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, _ := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)

	cmol := make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(*globalKpmnodeID, cmol)

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
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName, measTypeID)
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rfi, rknl, retsl, rrsl)

	return newE2SmKpmPdu
}

func Test_xerEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd := createE2SMKPMRanfunctionDescription()

	xer, err := XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 2702, len(xer))
	t.Logf("E2SmKpmRanfunctionDescription XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd := createE2SMKPMRanfunctionDescription()

	xer, err := XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 2702, len(xer))
	t.Logf("E2SmKpmRanfunctionDescription XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmRanfunctionDescription XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd := createE2SMKPMRanfunctionDescription()

	per, err := PerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 111, len(per))
	t.Logf("E2SmKpmRanfunctionDescription PER\n%s", hex.Dump(per))
}

//func Test_perDecodeE2SmKpmRanfunctionDescription(t *testing.T) {
//
//	rfd := createE2SMKPMRanfunctionDescription()
//
//	per, err := PerEncodeE2SmKpmRanfunctionDescription(rfd)
//	assert.NilError(t, err)
//	assert.Equal(t, 111, len(per))
//	t.Logf("E2SmKpmRanfunctionDescription PER\n%s", hex.Dump(per))
//
//	result, err := PerDecodeE2SmKpmRanfunctionDescription(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2SmKpmRanfunctionDescription PER - decoded\n%s", result)
//}
