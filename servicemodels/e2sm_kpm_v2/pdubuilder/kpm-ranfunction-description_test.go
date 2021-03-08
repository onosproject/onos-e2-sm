// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmRanfunctionDescription(t *testing.T) {
	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	cellGlobalID, err := CreateCellGlobalID_NRCGI(plmnID, bs)
	assert.NilError(t, err)
	var cellObjID string = "ONF"
	cellMeasObjItem := CreateCellMeasurementObjectItem(cellObjID, *cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	cmol := make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := CreateRicKpmnodeItem(*globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName string = "onf"
	var ricFormatType int32 = 15
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	var measTypeName string = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := CreateMeasurementInfoActionItem(measTypeName, measTypeID)
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rfi, rknl, retsl, rrsl)
	assert.NilError(t, err)

	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
