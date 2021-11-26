// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	kpmctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmRanfunctionDescription(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellGlobalID, err := CreateCellGlobalIDNRCGI(plmnID, []byte{0xd5, 0xbc, 0x09, 0x00, 0x00}) // 36 bit
	assert.NilError(t, err)
	var cellObjID = "ONF"
	cellMeasObjItem := CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	assert.NilError(t, err)
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := CreateRicKpmnodeItem(globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rknl, retsl, rrsl)
	assert.NilError(t, err)
	if newE2SmKpmPdu != nil {
		newE2SmKpmPdu.RanFunctionName.RanFunctionInstance = rfi
	}

	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

var sercommBytesFromE2 = []byte{
	0x07, 0x01, 0x80, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D,
	0x00, 0x00, 0x18, 0x31, 0x2E, 0x33, 0x2E, 0x36, 0x2E, 0x31, 0x2E, 0x34, 0x2E, 0x31, 0x2E, 0x35,
	0x33, 0x31, 0x34, 0x38, 0x2E, 0x31, 0x2E, 0x32, 0x2E, 0x32, 0x2E, 0x32, 0x05, 0x00, 0x4B, 0x50,
	0x4D, 0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74, 0x6F, 0x72, 0x00, 0x00, 0x40, 0x00, 0x13, 0xF1, 0x84,
	0x50, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xF1, 0x84, 0x00, 0x00,
	0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6F, 0x64, 0x69, 0x63, 0x20,
	0x72, 0x65, 0x70, 0x6F, 0x72, 0x74, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4E, 0x6F,
	0x64, 0x65, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x00, 0x01,
	0x08, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74, 0x61, 0x62,
	0x41, 0x74, 0x74, 0x2E, 0x54, 0x6F, 0x74, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52, 0x43, 0x2E,
	0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2E, 0x54, 0x6F,
	0x74, 0x00, 0x00, 0x01, 0x42, 0xA0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65,
	0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x54, 0x6F, 0x74, 0x00, 0x00, 0x02, 0x43,
	0xC0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62,
	0x41, 0x74, 0x74, 0x2E, 0x72, 0x65, 0x63, 0x6F, 0x6E, 0x66, 0x69, 0x67, 0x46, 0x61, 0x69, 0x6C,
	0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x48, 0x4F, 0x46, 0x61, 0x69, 0x6C, 0x00, 0x00,
	0x04, 0x42, 0xE0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45, 0x73, 0x74,
	0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x4F, 0x74, 0x68, 0x65, 0x72, 0x00, 0x00, 0x05, 0x41, 0x60,
	0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x2E, 0x41, 0x76, 0x67, 0x00, 0x00, 0x06, 0x41,
	0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x2E, 0x4D, 0x61, 0x78, 0x00, 0x00, 0x07,
	0x00, 0x01, 0x00, 0x01}

func TestSercomm(t *testing.T) {
	t.Logf("KPMv2 RFD PER bytes are \n%v", hex.Dump(sercommBytesFromE2))

	e2apPdu, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(sercommBytesFromE2)
	assert.NilError(t, err)
	t.Logf("Decoded KPMv2 RFD message is \n%v", e2apPdu)

	xer, err := kpmctypes.XerEncodeE2SmKpmRanfunctionDescription(e2apPdu)
	assert.NilError(t, err)
	t.Logf("KPMv2 RFD  XER\n%s", string(xer))
}

func TestSercommSampleMsg(t *testing.T) {
	var rfSn = "ORAN-E2SM-KPM"
	var rfE2SMoid = "1.3.6.1.4.1.53148.1.2.2.2"
	var rfd = "KPM monitor"
	var rfi int32 = 0

	plmnID := []byte{0x03, 0x11, 0x48}
	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	}
	cellGlobalID, err := CreateCellGlobalIDNRCGI(plmnID, []byte{0x00, 0x00, 0x00, 0x00, 0x10}) // 36 bit
	assert.NilError(t, err)
	var cellObjID = "0"
	cellMeasObjItem := CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	cmol := make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	globalKpmnodeID, err := CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	assert.NilError(t, err)

	kpmNodeItem := CreateRicKpmnodeItem(globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 1
	var ricStyleName = "Periodic report"
	var ricFormatType int32 = 1
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	var measTypeName1 = "RRC.ConnEstabAtt.Tot"
	var measTypeID1 int32 = 1
	measInfoActionItem1 := CreateMeasurementInfoActionItem(measTypeName1)
	measInfoActionItem1.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID1,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem1)
	var measTypeName2 = "RRC.ConnEstabSucc.Tot"
	var measTypeID2 int32 = 2
	measInfoActionItem2 := CreateMeasurementInfoActionItem(measTypeName2)
	measInfoActionItem2.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID2,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem2)
	var measTypeName3 = "RRC.ConnReEstabAtt.Tot"
	var measTypeID3 int32 = 3
	measInfoActionItem3 := CreateMeasurementInfoActionItem(measTypeName3)
	measInfoActionItem3.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID3,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem3)
	var measTypeName4 = "RRC.ConnReEstabAtt.reconfigFail"
	var measTypeID4 int32 = 4
	measInfoActionItem4 := CreateMeasurementInfoActionItem(measTypeName4)
	measInfoActionItem4.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID4,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem4)
	var measTypeName5 = "RRC.ConnEstabAtt.HOFail"
	var measTypeID5 int32 = 5
	measInfoActionItem5 := CreateMeasurementInfoActionItem(measTypeName5)
	measInfoActionItem5.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID5,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem5)
	var measTypeName6 = "RRC.ConnEstabAtt.Other"
	var measTypeID6 int32 = 6
	measInfoActionItem6 := CreateMeasurementInfoActionItem(measTypeName6)
	measInfoActionItem6.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID6,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem6)
	var measTypeName7 = "RRC.Conn.Avg"
	var measTypeID7 int32 = 7
	measInfoActionItem7 := CreateMeasurementInfoActionItem(measTypeName7)
	measInfoActionItem7.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID7,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem7)
	var measTypeName8 = "RRC.Conn.Max"
	var measTypeID8 int32 = 8
	measInfoActionItem8 := CreateMeasurementInfoActionItem(measTypeName8)
	measInfoActionItem8.MeasId = &e2sm_kpm_v2.MeasurementTypeId{
		Value: measTypeID8,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem8)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 1
	rrsi := CreateRicReportStyleItem(3, "E2 Node Measurement", ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rknl, retsl, rrsl)
	assert.NilError(t, err)
	if newE2SmKpmPdu != nil {
		newE2SmKpmPdu.RanFunctionName.RanFunctionInstance = rfi
	}

	xer, err := kpmctypes.XerEncodeE2SmKpmRanfunctionDescription(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("KPMv2 RanFunction-Description by Sercomm in XER is\n%s", string(xer))

	per, err := kpmctypes.PerEncodeE2SmKpmRanfunctionDescription(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("KPMv2 RanFunction-Description by Sercomm in PER is\n%s", hex.Dump(per))

	e2apPdu, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(per)
	assert.NilError(t, err)
	t.Logf("Decoded KPMv2 RFD message is \n%v", e2apPdu)
}
