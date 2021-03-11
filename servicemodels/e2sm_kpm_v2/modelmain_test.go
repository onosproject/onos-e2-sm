// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var kpmv2TestSm servicemodel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	plmnID := []byte{0x37, 0x34, 0x37}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion string = "txt"
	var senderName string = "ONF"
	var senderType string = "someType"
	var vendorName string = "onf"

	globalKpmNodeID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp, fileFormatVersion, senderName, senderType, vendorName, *globalKpmNodeID)

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	assert.Equal(t, 68, len(protoBytes))

	asn1Bytes, err := kpmv2TestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 42, len(asn1Bytes))
	t.Logf("E2SM-KPM-Indicationheader gNB asn1Bytes are \n%x", asn1Bytes)
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-Indication-Header-gNB.xml
	indicationHeaderAsn1Bytes := []byte{0x1f, 0x21, 0x22, 0x23, 0x24, 0x18, 0x74, 0x78, 0x74, 0x00, 0x00, 0x03, 0x4f, 0x4e,
		0x46, 0x00, 0x06, 0x6f, 0x6e, 0x66, 0x0c, 0x37, 0x34, 0x37, 0x00, 0xd4, 0xbc, 0x08, 0x80, 0x30, 0x39, 0x20, 0x1a, 0x85}

	protoBytes, err := kpmv2TestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 42, len(protoBytes))
	testIH := &e2sm_kpm_v2.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	//PER decoding is wrong - always sticks to eNB-ID
	//t.Logf("Whole message is \n%v", testIH)
	//t.Logf("Timestamp is \n%v", testIH.GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	//t.Logf("GnbCuUp is \n%v", testIH.GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbCuUpId())
	//t.Logf("PLMN ID value is \n%v", testIH.GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	//assert.DeepEqual(t, []byte{0x37, 0x34, 0x37}, testIH.GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
}

//func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
//	var plmnID = "ONF"
//	var cellIdentityValue uint64 = 0
//	var cellIdentityLen uint32 = 0
//	var ulTotalAvlblProbs int32 = 0
//	var dlTotalAvlblProbs int32 = 0
//	var fiveQi int32 = 0
//	var dlPrbusage int32 = 0
//	var ulPrbusage int32 = 0
//	var qCi int32 = 0
//	var qciDlPrbusage int32 = 0
//	var qciUlPrbusage int32 = 0
//	var sst = "1"
//	var sd = "SD1"
//	var gNbCuName = "OpenNetworking"
//	var cuCpNumberActvts int32 = 0
//	var ranContainer = "ranContainer"
//	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMsg(plmnID, cellIdentityValue, cellIdentityLen, dlTotalAvlblProbs,
//		ulTotalAvlblProbs, fiveQi, dlPrbusage, ulPrbusage, qCi, qciDlPrbusage, qciUlPrbusage, sst,
//		sd, gNbCuName, cuCpNumberActvts, ranContainer)
//	assert.NilError(t, err, "error creating E2SmPDU")
//
//	err = newE2SmKpmPdu.Validate()
//	assert.NilError(t, err, "error validating E2SmPDU")
//
//	assert.NilError(t, err)
//	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
//	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationMessage to bytes")
//	assert.Equal(t, 44, len(protoBytes)) //with ODU it will raise on 110
//
//	asn1Bytes, err := kpmTestSm.IndicationMessageProtoToASN1(protoBytes)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
//	assert.Assert(t, asn1Bytes != nil)
//}

//func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
//	// This message is taken from Shad
//	indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}
//
//	protoBytes, err := kpmTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
//	assert.Assert(t, protoBytes != nil)
//	assert.Equal(t, 12, len(protoBytes))
//	testIM := e2sm_kpm_v2.E2SmKpmIndicationMessage{}
//	err = proto.Unmarshal(protoBytes, &testIM)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(testIM.GetIndicationMessageFormat1().GetPmContainers()))
//	pm1 := testIM.GetIndicationMessageFormat1().GetPmContainers()[0]
//	assert.Equal(t, 0, int(pm1.GetPerformanceContainer().GetOCuCp().GetCuCpResourceStatus().GetNumberOfActiveUes()))
//}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var rfSn string = "onf"
	var rfE2SMoid string = "oid123"
	var rfd string = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	cellGlobalID, err := pdubuilder.CreateCellGlobalID_NRCGI(plmnID, 0xabcdef012 << 28) // 36 bit
	assert.NilError(t, err)
	var cellObjID string = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, *cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

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

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd, rfi, rknl, retsl, rrsl)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmKpmPdu != nil)

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmRanfunctionDescription to bytes")
	assert.Equal(t, 174, len(protoBytes))

	asn1Bytes, err := kpmv2TestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 77, len(asn1Bytes))
	t.Logf("E2SM-KPM-RANfunctionDescription asn1Bytes are \n%x", asn1Bytes)
}

func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	// This message is taken as an output from the function above
	//ToDo: fill it with correct ASN1 bytes - ask Shad
	ranFuncDescriptionAsn1 := []byte{0x64, 0x04, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x05, 0x6f, 0x69, 0x64, 0x31, 0x32, 0x33,
		0x07, 0x00, 0x73, 0x6f, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x01, 0x15,
		0x00, 0x00, 0x43, 0x00, 0x21, 0x22, 0x23, 0x00, 0xd4, 0xbc, 0x08, 0x80, 0x30, 0x39, 0x20, 0x1a, 0x85, 0x00, 0x00,
		0x00, 0x00, 0x03, 0x4f, 0x4e, 0x46, 0x00, 0x21, 0x22, 0x23, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x01, 0x0b, 0x01,
		0x00, 0x6f, 0x6e, 0x66, 0x01, 0x0f}

	protoBytes, err := kpmv2TestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 118, len(protoBytes))
	testIM := &e2sm_kpm_v2.E2SmKpmRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	assert.Equal(t, "oid123", testIM.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(21), testIM.GetRanFunctionName().GetRanFunctionInstance())
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int32 = 12
	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmKpmEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmKpmEventTriggerDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmEventTriggerDefinition to bytes")
	assert.Equal(t, 4, len(protoBytes))

	asn1Bytes, err := kpmv2TestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 3, len(asn1Bytes))
	t.Logf("E2SM-KPM-EventTriggerDefinition asn1Bytes are \n%x", asn1Bytes)
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-EventTriggerDefinition.xml
	eventTriggerDefinitionAsn1 := []byte{0x00, 0x01, 0x0c}

	protoBytes, err := kpmv2TestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 4, len(protoBytes))
	testIM := &e2sm_kpm_v2.E2SmKpmEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Reporting period is \n%v", testIM.GetEventDefinitionFormat1().GetReportingPeriod())
	assert.Equal(t, int32(12), testIM.GetEventDefinitionFormat1().GetReportingPeriod())
}

//func TestServicemodel_ActionDefinitionProtoToASN1(t *testing.T) {
//	var ricStyleType int32 = 12
//	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(ricStyleType)
//	assert.NilError(t, err, "error creating E2SmPDU")
//	assert.Assert(t, e2SmKpmActionDefinition != nil, "Created E2SmPDU is nil")
//
//	err = e2SmKpmActionDefinition.Validate()
//	assert.NilError(t, err, "error validating E2SmPDU")
//
//	assert.NilError(t, err)
//	protoBytes, err := proto.Marshal(e2SmKpmActionDefinition)
//	assert.NilError(t, err, "unexpected error marshalling E2SmKpmActionDefinition to bytes")
//	assert.Equal(t, 4, len(protoBytes))
//
//	asn1Bytes, err := kpmTestSm.ActionDefinitionProtoToASN1(protoBytes)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
//	assert.Assert(t, asn1Bytes != nil)
//	assert.Equal(t, 3, len(asn1Bytes))
//}
//
//func TestServicemodel_ActionDefinitionASN1toProto(t *testing.T) {
//	// This value is taken from XxX and passed as a byte array directly to the function
//	// It's the encoding of what's in the file ../test/E2SM-KPM-EventTriggerDefinition.xml
//	// TODO: Take real values
//	actionDefinitionAsn1 := []byte{0x20, 0x38, 0x37, 0xDB, 0xFD, 0x7F, 0x00,
//		0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00}
//
//	protoBytes, err := kpmTestSm.ActionDefinitionASN1toProto(actionDefinitionAsn1)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
//	assert.Assert(t, protoBytes != nil)
//	assert.Equal(t, 4, len(protoBytes))
//	testIM := e2smkpmies.E2SmKpmActionDefinition{}
//	err = proto.Unmarshal(protoBytes, &testIM)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, testIM.GetRicStyleType().GetValue())
//}
