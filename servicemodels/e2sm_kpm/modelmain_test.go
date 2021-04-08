// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	e2smkpmies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var kpmTestSm servicemodel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var gNbCuUpID int64 = 0
	var gNbDuID int64 = 0
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpID, gNbDuID, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	assert.Equal(t, 68, len(protoBytes))

	asn1Bytes, err := kpmTestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 31, len(asn1Bytes))
	t.Logf("E2SM-KPM-IndicationHeader asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var cellIdentityValue uint64 = 0
	var cellIdentityLen uint32 = 0
	var ulTotalAvlblProbs int32 = 0
	var dlTotalAvlblProbs int32 = 0
	var fiveQi int32 = 0
	var dlPrbusage int32 = 0
	var ulPrbusage int32 = 0
	var qCi int32 = 0
	var qciDlPrbusage int32 = 0
	var qciUlPrbusage int32 = 0
	var sst = "1"
	var sd = "SD1"
	var gNbCuName = "OpenNetworking"
	var cuCpNumberActvts int32 = 0
	var ranContainer = "ranContainer"
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMsg(plmnID, cellIdentityValue, cellIdentityLen, dlTotalAvlblProbs,
		ulTotalAvlblProbs, fiveQi, dlPrbusage, ulPrbusage, qCi, qciDlPrbusage, qciUlPrbusage, sst,
		sd, gNbCuName, cuCpNumberActvts, ranContainer)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationMessage to bytes")
	assert.Equal(t, 44, len(protoBytes)) //with ODU it will raise on 110

	asn1Bytes, err := kpmTestSm.IndicationMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-IndicationMessage asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-Indication-Header.xml
	indicationHeaderAsn1Bytes := []byte{0x3F, 0x08, 0x37, 0x34, 0x37, 0x38, 0xB5, 0xC6, 0x77, 0x88, 0x02, 0x37, 0x34, 0x37, 0x22, 0x5B,
		0xD6, 0x00, 0x70, 0x37, 0x34, 0x37, 0x98, 0x80, 0x31, 0x30, 0x30, 0x09, 0x09}

	protoBytes, err := kpmTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 74, len(protoBytes))
	testIH := e2smkpmies.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(protoBytes, &testIH)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x37, 0x34, 0x37}, testIH.GetIndicationHeaderFormat1().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	// This message is taken from Shad
	indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}

	protoBytes, err := kpmTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 12, len(protoBytes))
	testIM := e2smkpmies.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(testIM.GetIndicationMessageFormat1().GetPmContainers()))
	pm1 := testIM.GetIndicationMessageFormat1().GetPmContainers()[0]
	assert.Equal(t, 0, int(pm1.GetPerformanceContainer().GetOCuCp().GetCuCpResourceStatus().GetNumberOfActiveUes()))
}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 1
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmKpmPdu != nil)

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmRanfunctionDescription to bytes")
	assert.Equal(t, 81, len(protoBytes))

	asn1Bytes, err := kpmTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 63, len(asn1Bytes))
	t.Logf("E2SM-KPM-RanFunctionDescription asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

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
	assert.Equal(t, 6, len(protoBytes))

	asn1Bytes, err := kpmTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 2, len(asn1Bytes))
	t.Logf("E2SM-KPM-EventTriggerDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-EventTriggerDefinition.xml
	eventTriggerDefinitionAsn1 := []byte{0x20, 0x38, 0x37, 0xDB, 0xFD, 0x7F, 0x00,
		0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00}

	protoBytes, err := kpmTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 6, len(protoBytes))
	testIM := e2smkpmies.E2SmKpmEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(testIM.GetEventDefinitionFormat1().GetPolicyTestList()))
}

func TestServicemodel_ActionDefinitionProtoToASN1(t *testing.T) {
	var ricStyleType int32 = 12
	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(ricStyleType)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmKpmActionDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmKpmActionDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmKpmActionDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmActionDefinition to bytes")
	assert.Equal(t, 4, len(protoBytes))

	asn1Bytes, err := kpmTestSm.ActionDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 3, len(asn1Bytes))
	t.Logf("E2SM-KPM-ActionDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

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

func TestServicemodel_Asn1BytesToByte(t *testing.T) {
	indHdrHex, err := kpmTestSm.Asn1BytesToByte("3f0c4f4e4600d4bc08000000006f6e66efabd4bc004f4e4698805344310000")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indHdrHex)

	inHdrBytes, err := kpmTestSm.IndicationHeaderASN1toProto(indHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, inHdrBytes != nil)
	assert.Equal(t, 67, len(inHdrBytes))
	testIH := &e2smkpmies.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(inHdrBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationHeader) is \n%v\n", testIH)

	indMsgHex, err := kpmTestSm.Asn1BytesToByte("4000006c1a4f70656e4e6574776f726b696e67c001000c72616e436f6e7461696e6572")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indMsgHex)

	indMsgBytes, err := kpmTestSm.IndicationMessageASN1toProto(indMsgHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, indMsgBytes != nil)
	assert.Equal(t, 44, len(indMsgBytes))
	testIM := &e2smkpmies.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationMessage) is \n%v", testIM)

	actDefHex, err := kpmTestSm.Asn1BytesToByte("00010c")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", actDefHex)

	actDefBytes, err := kpmTestSm.ActionDefinitionASN1toProto(actDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, actDefBytes != nil)
	assert.Equal(t, 4, len(actDefBytes))
	testAD := &e2smkpmies.E2SmKpmActionDefinition{}
	err = proto.Unmarshal(actDefBytes, testAD)
	assert.NilError(t, err)
	t.Logf("Decoded message (ActionDefinition) is \n%v\n", testAD)

	evntTrigDefHex, err := kpmTestSm.Asn1BytesToByte("2030")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", evntTrigDefHex)

	evntTrigDefBytes, err := kpmTestSm.EventTriggerDefinitionASN1toProto(evntTrigDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, evntTrigDefBytes != nil)
	assert.Equal(t, 6, len(evntTrigDefBytes))
	testETD := &e2smkpmies.E2SmKpmEventTriggerDefinition{}
	err = proto.Unmarshal(evntTrigDefBytes, testETD)
	assert.NilError(t, err)
	t.Logf("Decoded message (EventTriggerDefinition) is \n%v\n", testETD)

	ranFuncDescHex, err := kpmTestSm.Asn1BytesToByte("20204f4e460000024f696406804f70656e4e6574776f726b696e6701016000" +
		"010d03804f4e466576656e74012a00010c04004f4e467265706f727401150138")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ranFuncDescHex)

	ranFuncDescBytes, err := kpmTestSm.RanFuncDescriptionASN1toProto(ranFuncDescHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ranFuncDescBytes != nil)
	assert.Equal(t, 81, len(ranFuncDescBytes))
	testRFD := &e2smkpmies.E2SmKpmRanfunctionDescription{}
	err = proto.Unmarshal(ranFuncDescBytes, testRFD)
	t.Logf("Decoded message (RanFunctionDescription) is \n%v\n", testRFD)
	assert.NilError(t, err)
}

func TestServicemodel_HexDumpToByte(t *testing.T) {
	indHdrHex, err := kpmTestSm.HexDumpToByte("00000000  3f 0c 4f 4e 46 00 d4 bc  08 00 00 00 00 6f 6e 66  |?.ONF........onf|" +
		"	00000010  ef ab d4 bc 00 4f 4e 46  98 80 53 44 31 00 00     |.....ONF..SD1..|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", indHdrHex)

	indHdrBytes, err := kpmTestSm.IndicationHeaderASN1toProto(indHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, indHdrBytes != nil)
	assert.Equal(t, 67, len(indHdrBytes))
	testIH := &e2smkpmies.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(indHdrBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationHeader) is \n%v", testIH)

	indMsgHex, err := kpmTestSm.HexDumpToByte("00000000  40 00 00 6c 1a 4f 70 65  6e 4e 65 74 77 6f 72 6b  |@..l.OpenNetwork|" +
		"	00000010  69 6e 67 c0 01 00 0c 72  61 6e 43 6f 6e 74 61 69  |ing....ranContai|" +
		"		00000020  6e 65 72                                          |ner|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", indMsgHex)

	indMsgBytes, err := kpmTestSm.IndicationMessageASN1toProto(indMsgHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, indMsgBytes != nil)
	assert.Equal(t, 44, len(indMsgBytes))
	testIM := &e2smkpmies.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationMessage) is \n%v", testIM)

	actDefHex, err := kpmTestSm.HexDumpToByte("00000000  00 01 0c                                          |...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", actDefHex)

	actDefBytes, err := kpmTestSm.ActionDefinitionASN1toProto(actDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, actDefBytes != nil)
	assert.Equal(t, 4, len(actDefBytes))
	testAD := &e2smkpmies.E2SmKpmActionDefinition{}
	err = proto.Unmarshal(actDefBytes, testAD)
	assert.NilError(t, err)
	t.Logf("Decoded message (ActionDefinition) is \n%v\n", testAD)

	evntTrigDefHex, err := kpmTestSm.HexDumpToByte("00000000  20 30                                             | 0|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", evntTrigDefHex)

	evntTrigDefBytes, err := kpmTestSm.EventTriggerDefinitionASN1toProto(evntTrigDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, evntTrigDefBytes != nil)
	assert.Equal(t, 6, len(evntTrigDefBytes))
	testETD := &e2smkpmies.E2SmKpmEventTriggerDefinition{}
	err = proto.Unmarshal(evntTrigDefBytes, testETD)
	assert.NilError(t, err)
	t.Logf("Decoded message (EventTriggerDefinition) is \n%v\n", testETD)

	ranFuncDescHex, err := kpmTestSm.HexDumpToByte("00000000  20 20 4f 4e 46 00 00 02  4f 69 64 06 80 4f 70 65  |  ONF...Oid..Ope|" +
		"	00000010  6e 4e 65 74 77 6f 72 6b  69 6e 67 01 01 60 00 01  |nNetworking..`..|" +
		"        00000020  0d 03 80 4f 4e 46 65 76  65 6e 74 01 2a 00 01 0c  |...ONFevent.*...|" +
		"        00000030  04 00 4f 4e 46 72 65 70  6f 72 74 01 15 01 38     |..ONFreport...8|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ranFuncDescHex)

	ranFuncDescBytes, err := kpmTestSm.RanFuncDescriptionASN1toProto(ranFuncDescHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ranFuncDescBytes != nil)
	assert.Equal(t, 81, len(ranFuncDescBytes))
	testRFD := &e2smkpmies.E2SmKpmRanfunctionDescription{}
	err = proto.Unmarshal(ranFuncDescBytes, testRFD)
	t.Logf("Decoded message (RanFunctionDescription) is \n%v\n", testRFD)
	assert.NilError(t, err)
}
