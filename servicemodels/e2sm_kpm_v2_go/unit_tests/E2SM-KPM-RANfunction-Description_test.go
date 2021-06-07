// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerE2SmKpmRanFunctionDescription = "00000000  74 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |t.onf...oid123..|" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 00  |someDescription.|" +
	"00000020  15 00 00 43 00 21 22 23  00 d4 bc 08 80 30 39 20  |...C.!\"#.....09 |" +
	"00000030  1a 85 00 00 00 00 03 4f  4e 46 00 21 22 23 00 00  |.......ONF.!\"#..|" +
	"00000040  00 20 00 00 0b 01 00 6f  6e 66 00 0f 00 0b 01 00  |. .....onf......|" +
	"00000050  6f 6e 66 00 0f 00 00 41  a0 4f 70 65 6e 4e 65 74  |onf....A.OpenNet|" +
	"00000060  77 6f 72 6b 69 6e 67 00  00 17 00 02 00 01        |working.......|\""

func createE2SmKpmRanFunctionDescription() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	cellGlobalID, _ := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, 0xABCDEF012<<28) // 36 bits

	var cellObjID string = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, _ := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID, cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName string = "onf"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName string = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 2
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rknl, retsl, rrsl)
	newE2SmKpmPdu.RanFunctionName.RanFunctionInstance = &rfi
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmRanFunctionDescription(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescription()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*rfd, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescription)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
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

	result := e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription{}
	err := aper.UnmarshalWithParams(radisysBytesRanFunctionDefinition, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription (Radisys) PER - decoded\n%v", result)
}
