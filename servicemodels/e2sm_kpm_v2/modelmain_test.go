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
	var vendorName string = "onf"

	globalKpmNodeID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp, fileFormatVersion, senderName, vendorName, *globalKpmNodeID)

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	assert.Equal(t, 58, len(protoBytes))

	asn1Bytes, err := kpmv2TestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 34, len(asn1Bytes))
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
	//ToDo - something is wrong here...
	t.Logf("Whole message is \n%v", testIH)
	t.Logf("Timestamp is \n%v", testIH.GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	t.Logf("GnbCuUp is \n%v", testIH.GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbCuUpId())
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

//func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
//	var ranFunctionShortName = "ONF"
//	var ranFunctionE2SmOid = "Oid"
//	var ranFunctionDescription = "OpenNetworking"
//	var ranFunctionInstance int32 = 1
//	var ricEventStyleType int32 = 13
//	var ricEventStyleName = "ONFevent"
//	var ricEventFormatType int32 = 42
//	var ricReportStyleType int32 = 12
//	var ricReportStyleName = "ONFreport"
//	var ricIndicationHeaderFormatType int32 = 21
//	var ricIndicationMessageFormatType int32 = 56
//	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
//		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
//		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
//	assert.NilError(t, err, "error creating E2SmPDU")
//	assert.Assert(t, newE2SmKpmPdu != nil)
//
//	err = newE2SmKpmPdu.Validate()
//	assert.NilError(t, err, "error validating E2SmPDU")
//
//	assert.NilError(t, err)
//	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
//	assert.NilError(t, err, "unexpected error marshalling E2SmKpmRanfunctionDescription to bytes")
//	assert.Equal(t, 81, len(protoBytes))
//
//	asn1Bytes, err := kpmTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
//	assert.Assert(t, asn1Bytes != nil)
//	assert.Equal(t, 63, len(asn1Bytes))
//}

//func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
//	// This message is taken as an output from the function above
//	//ToDo: fill it with correct ASN1 bytes - ask Shad
//	ranFuncDescriptionAsn1 := []byte{0x00, 0xC0, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D, 0x00,
//		0x00, 0x12, 0x31, 0x2E, 0x33, 0x2E, 0x36, 0x2E, 0x31, 0x2E, 0x34, 0x2E, 0x31, 0x2E, 0x31, 0x2E,
//		0x31, 0x2E, 0x32, 0x2E, 0x32, 0x05, 0x00, 0x4B, 0x50, 0x4D, 0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74,
//		0x6F, 0x72, 0x60, 0x00, 0x01, 0x01, 0x03, 0x80, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x31,
//		0x01, 0x01, 0x00, 0x01, 0x06, 0x1E, 0x80, 0x4F, 0x2D, 0x43, 0x55, 0x2D, 0x55, 0x50, 0x20, 0x4D,
//		0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61,
//		0x69, 0x6E, 0x65, 0x72, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x50, 0x43,
//		0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F,
//		0x79, 0x6D, 0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01}
//
//	protoBytes, err := kpmTestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
//	assert.Assert(t, protoBytes != nil)
//	assert.Equal(t, 10, len(protoBytes))
//	testIM := e2smkpmies.E2SmKpmRanfunctionDescription{}
//	err = proto.Unmarshal(protoBytes, &testIM)
//	assert.NilError(t, err)
//	//assert.Equal(t, 1, len(testIM.GetEventDefinitionFormat1().GetPolicyTestList()))
//}

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
