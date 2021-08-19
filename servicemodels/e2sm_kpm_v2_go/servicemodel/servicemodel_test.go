// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package servicemodel

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var kpmv2TestSm Kpm2ServiceModel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x37, 0x34, 0x37}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion = "txt"
	var senderName = "ONF"
	var senderType = "someType"
	var vendorName = "onf"

	globalKpmNodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp)
	assert.NilError(t, err)
	newE2SmKpmPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetSenderType(senderType).SetVendorName(vendorName).SetGlobalKPMnodeID(globalKpmNodeID)
	assert.NilError(t, err)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	//assert.Equal(t, 69, len(protoBytes))

	asn1Bytes, err := kpmv2TestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-IndicationHeader (gNB) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-Indication-Header-gNB.xml
	indicationHeaderAsn1Bytes := []byte{0x1f, 0x21, 0x22, 0x23, 0x24, 0x18, 0x74, 0x78, 0x74, 0x00, 0x00, 0x03, 0x4f, 0x4e,
		0x46, 0x40, 0x73, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x06, 0x6f, 0x6e, 0x66, 0x0c, 0x37, 0x34, 0x37, 0x00,
		0xd4, 0xbc, 0x08, 0x80, 0x30, 0x39, 0x20, 0x1a, 0x85}

	protoBytes, err := kpmv2TestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testIH := &e2sm_kpm_v2_go.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testIH)
	assert.DeepEqual(t, []byte{0x21, 0x22, 0x23, 0x24}, testIH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	assert.DeepEqual(t, []byte{0x37, 0x34, 0x37}, testIH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	assert.Equal(t, int64(12345), testIH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbCuUpId().GetValue())
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var integer int64 = 12345
	//var rl float64 = 6789
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START
	var measurementName = "trial"
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem := pdubuilder.CreateMeasurementInfoItem(measName).SetLabelInfoList(&labelInfoList)

	measInfoList := e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	//measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, err := pdubuilder.CreateMeasurementDataItem(&measRecord)
	assert.NilError(t, err)
	measDataItem.SetIncompleteFlag()

	measData := e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, &measData).SetGranularityPeriod(granularity).SetCellObjectID(cellObjID).SetMeasInfoList(&measInfoList)
	assert.Assert(t, newE2SmKpmPdu != nil)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationMessage (Format1) to bytes")

	asn1Bytes, err := kpmv2TestSm.IndicationMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-IndicationMessage (Format1)) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	indicationMessageAsn1 := []byte{0x0e, 0x80, 0x30, 0x38, 0x00, 0x00, 0x03, 0x6f, 0x6e, 0x66, 0x00, 0x14, 0x00, 0x00, 0x40, 0x20,
		0x74, 0x72, 0x69, 0x61, 0x6c, 0x01, 0x3f, 0xff, 0xe0, 0x21, 0x22, 0x23, 0x40, 0x40, 0x01, 0x02,
		0x03, 0x00, 0x0a, 0x7c, 0x0f, 0x00, 0x0f, 0x00, 0x01, 0x72, 0x40, 0x00, 0xfa, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x7a, 0x00, 0x01, 0xc7, 0x00, 0x03, 0x14, 0x00, 0x00, 0x00, 0x40, 0x02, 0x08, 0x30,
		0x39, 0x40}

	protoBytes, err := kpmv2TestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 124, len(protoBytes))
	testIM := &e2sm_kpm_v2_go.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testIM)
	assert.Equal(t, int64(12345), testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetSubscriptId().GetValue())
	assert.Equal(t, "onf", testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetCellObjId().GetValue())
	md := testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0]
	assert.Equal(t, 2, len(md.GetMeasRecord().GetValue()))
	mil := testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0]
	assert.Equal(t, "trial", mil.GetMeasType().GetMeasName().GetValue())
	lil := mil.GetLabelInfoList().GetValue()[0]
	assert.Equal(t, int32(62), lil.GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, int32(15), lil.GetMeasLabel().GetARpmax().GetValue())
}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellGlobalID, err := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, []byte{0xd5, 0xbc, 0x09, 0x00, 0x00}) // 36 bit
	assert.NilError(t, err)
	var cellObjID = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID).SetCellMeasurementObjectList(cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicKpmNodeList(rknl).SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmKpmPdu != nil)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmRanfunctionDescription to bytes")

	asn1Bytes, err := kpmv2TestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	// This message is taken as an output from the function above
	ranFuncDescriptionAsn1 := []byte{0x74, 0x04, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x05, 0x6f, 0x69, 0x64, 0x31, 0x32, 0x33, 0x07, 0x00,
		0x73, 0x6f, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x00,
		0x15, 0x00, 0x00, 0x43, 0x00, 0x21, 0x22, 0x23, 0x00, 0xd4, 0xbc, 0x08, 0x80, 0x30, 0x39, 0x20,
		0x1a, 0x85, 0x00, 0x00, 0x00, 0x00, 0x03, 0x4f, 0x4e, 0x46, 0x00, 0x21, 0x22, 0x23, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0x0b, 0x01, 0x00, 0x6f, 0x6e, 0x66, 0x00, 0x0f, 0x00, 0x0b, 0x01, 0x00,
		0x6f, 0x6e, 0x66, 0x00, 0x0f, 0x00, 0x00, 0x41, 0xa0, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x65, 0x74,
		0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x00, 0x00, 0x17, 0x00, 0x2f, 0x00, 0x18}

	protoBytes, err := kpmv2TestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 171, len(protoBytes))
	testRFD := &e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, testRFD)
	t.Logf("Decoded message is \n%v", testRFD)
	assert.NilError(t, err)
	assert.Equal(t, "oid123", testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(21), testRFD.GetRanFunctionName().GetRanFunctionInstance())
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int64 = 12
	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmKpmEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	//err = e2SmKpmEventTriggerDefinition.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmEventTriggerDefinition to bytes")

	asn1Bytes, err := kpmv2TestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-EventTriggerDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-EventTriggerDefinition.xml
	eventTriggerDefinitionAsn1 := []byte{0x00, 0x0b}

	protoBytes, err := kpmv2TestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 6, len(protoBytes))
	testETD := &e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, testETD)
	t.Logf("Decoded message is \n%v", testETD)
	assert.NilError(t, err)
	t.Logf("Reporting period is \n%v", testETD.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriod())
	assert.Equal(t, int64(12), testETD.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriod())
}

func TestServicemodel_ActionDefinitionProtoToASN1(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))
	assert.NilError(t, err)

	mci, err := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)
	assert.NilError(t, err)

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measCondItem, err := pdubuilder.CreateMeasurementCondItem(measName, mcl)
	assert.NilError(t, err)

	measCondList := e2sm_kpm_v2_go.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, err := pdubuilder.CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmActionDefinitionFormat3(ricStyleType, actionDefinitionFormat3)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmKpmPdu != nil, "Created E2SmPDU is nil")

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmActionDefinition to bytes")

	asn1Bytes, err := kpmv2TestSm.ActionDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ActionDefinitionASN1toProto(t *testing.T) {
	actionDefinitionAsn1 := []byte{0x00, 0x0c, 0x40, 0x00, 0x03, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x00, 0x40, 0x74, 0x72, 0x69, 0x61,
		0x6c, 0x00, 0x00, 0x48, 0x21, 0x02, 0x00, 0xc9, 0x00, 0x14, 0x40, 0x30, 0x38}

	protoBytes, err := kpmv2TestSm.ActionDefinitionASN1toProto(actionDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testAD := &e2sm_kpm_v2_go.E2SmKpmActionDefinition{}
	err = proto.Unmarshal(protoBytes, testAD)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testAD)
	assert.Equal(t, int32(12), testAD.GetRicStyleType().GetValue())
	adf3 := testAD.GetActionDefinitionFormats().GetActionDefinitionFormat3()
	assert.Equal(t, "onf", adf3.GetCellObjId().GetValue())
	assert.Equal(t, int64(12345), adf3.GetSubscriptId().GetValue())
	assert.Equal(t, int64(21), adf3.GetGranulPeriod().GetValue())
	mcl := adf3.GetMeasCondList().GetValue()[0]
	assert.Equal(t, "trial", mcl.GetMeasType().GetMeasName().GetValue())
	mc := mcl.GetMatchingCond().GetValue()[0]
	assert.Equal(t, int64(201), mc.GetTestCondInfo().GetTestValue().GetValueEnum())
}
