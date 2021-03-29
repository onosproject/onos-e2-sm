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
		Value: 0x9bcd4,
		Len:   22,
	}
	cellGlobalID, _ := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, 0xABCDEF012<<28) // 36 bits

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
	assert.Equal(t, 2704, len(xer))
	//assert.Equal(t, 372, len(xer)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 2704, len(xer))
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
	assert.Equal(t, 112, len(per))
	//assert.Equal(t, 33, len(per)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription PER\n%s", hex.Dump(per))
}

func Test_perDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	rfd, err := createE2SMKPMRanfunctionDescription()
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	assert.Equal(t, 112, len(per))
	//assert.Equal(t, 33, len(per)) // without lists
	t.Logf("E2SmKpmRanfunctionDescription PER\n%s", hex.Dump(per))

	result, err := PerDecodeE2SmKpmRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmRanfunctionDescription PER - decoded\n%s", result)
}
