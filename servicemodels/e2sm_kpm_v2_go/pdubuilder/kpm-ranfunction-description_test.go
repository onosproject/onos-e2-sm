// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmRanfunctionDescription(t *testing.T) {
	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x09, 0x00},
		Len:   22,
	}
	cellIDbits := []byte{0x12, 0xF0, 0xDE, 0xBC, 0x0A, 0x00}
	cellGlobalID, err := CreateCellGlobalIDNRCGI(plmnID, cellIDbits) // 36 bit
	assert.NilError(t, err)
	var cellObjID string = "ONF"
	cellMeasObjItem := CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	assert.NilError(t, err)
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := CreateRicKpmnodeItem(globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName string = "onf"
	var ricFormatType int32 = 15
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName string = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rknl, retsl, rrsl)
	assert.NilError(t, err)
	if newE2SmKpmPdu != nil {
		newE2SmKpmPdu.RanFunctionName.RanFunctionInstance = &rfi
	}
	assert.Assert(t, newE2SmKpmPdu != nil)
}
