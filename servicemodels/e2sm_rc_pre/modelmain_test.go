// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var rcPreTestSm servicemodel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(plmnIDBytes, &cellID)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreIndicationHeader to bytes")
	assert.Equal(t, 24, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.IndicationHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)

	assert.Equal(t, 8, len(asn1Bytes))
	t.Logf("E2SM-KPM-IndicationHeader asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	var dlEarfcn int32 = 253
	var pci int32 = 11

	pciPool := &e2sm_rc_pre_ies.PciRange{
		LowerPci: &e2sm_rc_pre_ies.Pci{
			Value: 10,
		},
		UpperPci: &e2sm_rc_pre_ies.Pci{
			Value: 20,
		},
	}

	neighbors := &e2sm_rc_pre_ies.Nrt{
		NrIndex: 1,
		Cgi: &e2sm_rc_pre_ies.CellGlobalId{
			CellGlobalId: &e2sm_rc_pre_ies.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_rc_pre_ies.Eutracgi{
					PLmnIdentity: &e2sm_rc_pre_ies.PlmnIdentity{
						Value: plmnIDBytes,
					},
					EUtracellIdentity: &e2sm_rc_pre_ies.EutracellIdentity{
						Value: &e2sm_rc_pre_ies.BitString{
							Value: 0x9bcd4ac, //uint64
							Len:   28,        //uint32
						},
					},
				},
			},
		},
		Pci: &e2sm_rc_pre_ies.Pci{
			Value: 11,
		},
		CellSize: e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO,
		DlEarfcn: &e2sm_rc_pre_ies.Earfcn{
			Value: 253,
		},
	}

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationMsg(plmnIDBytes, &cellID, dlEarfcn, e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO, pci, pciPool, neighbors)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreIndicationMessage to bytes")

	assert.Equal(t, 82, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.IndicationMessageProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-IndicationMessage asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-RC-PRE-Indication-Header.xml

	indicationHeaderAsn1Bytes := []byte{0x28, 0x4f, 0x4e, 0x46, 0xab, 0xd4, 0xbc, 0x00}

	protoBytes, err := rcPreTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 24, len(protoBytes))
	testIH := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(protoBytes, &testIH)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testIH.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	//indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}

	indicationMessageAsn1 := []byte{0x48, 0x4f, 0x4e, 0x46, 0xab, 0xd4, 0xbc, 0x00, 0xfd, 0x60,
		0x00, 0x00, 0x0a, 0x00, 0x14, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x4f,
		0x4e, 0x46, 0xab, 0xd4, 0xbc, 0x00, 0xfd, 0x60, 0x00, 0x0b}
	protoBytes, err := rcPreTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 82, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreIndicationMessage{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(testIM.GetIndicationMessageFormat1().GetNeighbors()))
	nbr := testIM.GetIndicationMessageFormat1().GetNeighbors()[0]
	assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, nbr.GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmRcPrePdu != nil)

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreRanfunctionDescription to bytes")
	assert.Equal(t, 81, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 63, len(asn1Bytes))
	t.Logf("E2SM-KPM-RanFunctionDescription asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int32 = 12
	e2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmRcPreEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmRcPreEventTriggerDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreEventTriggerDefinition to bytes")
	assert.Equal(t, 6, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Equal(t, 3, len(asn1Bytes))
	t.Logf("E2SM-KPM-EventTriggerDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x14, 0x01, 0x0c}
	protoBytes, err := rcPreTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 6, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 12, int(testIM.GetEventDefinitionFormat1().GetReportingPeriodMs()))
}

func TestServicemodel_ActionDefinitionProtoToASN1(t *testing.T) {
	var ricStyleType int32 = 12
	e2SmRcPreActionDefinition, err := pdubuilder.CreateE2SmRcPreActionDefinition(ricStyleType)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmRcPreActionDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmRcPreActionDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmRcPreActionDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreActionDefinition to bytes")
	assert.Equal(t, 4, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.ActionDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 3, len(asn1Bytes))
	t.Logf("E2SM-KPM-ActionDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderProtoToASN1(t *testing.T) {

	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader(controlMessagePriority, plmnIDBytes, &cellID)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlHeader to bytes")
	assert.Equal(t, 28, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.ControlHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 10, len(asn1Bytes))
	t.Logf("E2SM-KPM-ControlHeader asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x34, 0x4f, 0x4e, 0x46, 0xab, 0xd4, 0xbc, 0x00, 0x01, 0x01}

	protoBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 28, len(protoBytes))
	testCH := e2sm_rc_pre_ies.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(protoBytes, &testCH)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testCH.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())

}

func TestServicemodel_ControlMessageProtoToASN1(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue int32 = 20
	var ranParameterID int32 = 1

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameterValue)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlMessage to bytes")

	assert.Equal(t, 19, len(protoBytes))
	asn1Bytes, err := rcPreTestSm.ControlMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-ControlMessage asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlMessageASN1toProto(t *testing.T) {
	ControlMessageAsn1 := []byte{0x00, 0x00, 0x14, 0x01, 0x00, 0x50, 0x43, 0x49, 0x00, 0x06,
		0x00, 0xc0, 0x00, 0x0c, 0xbb, 0x70}

	protoBytes, err := rcPreTestSm.ControlMessageASN1toProto(ControlMessageAsn1)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 21, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreControlMessage{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 834416, int(testIM.GetControlMessage().GetParameterVal().GetValueInt()))
}

func TestServicemodel_ControlOutcomeProtoToASN1(t *testing.T) {
	var ranParameterID int32 = 20
	var ranParameterValue int32 = 10

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID, ranParameterValue)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlOutcome to bytes")

	assert.Equal(t, 12, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.ControlOutcomeProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-KPM-ControlOutcome asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlOutcomeASN1toProto(t *testing.T) {

	ControlOutcomeAsn1 := []byte{0x20, 0x00, 0x00, 0x00, 0x00, 0x14, 0x00, 0x06,
		0x00, 0xc0, 0x00, 0x0c, 0xbb, 0x40}
	protoBytes, err := rcPreTestSm.ControlOutcomeASN1toProto(ControlOutcomeAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 14, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(testIM.GetControlOutcomeFormat1().GetOutcomeElementList()))
	outcome := testIM.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, 834368, int(outcome.GetRanParameterValue().GetValueInt()))
}

//func TestServicemodel_ActionDefinitionASN1toProto(t *testing.T) {
//	// This value is taken from XxX and passed as a byte array directly to the function
//	// It's the encoding of what's in the file ../test/E2SM-RC-PRE-EventTriggerDefinition.xml
//	// TODO: Take real values
//	actionDefinitionAsn1 := []byte{0x20, 0x38, 0x37, 0xDB, 0xFD, 0x7F, 0x00,
//		0x00, 0x28, 0x00, 0x00, 0x00a0126, 0x00, 0x00}
//
//	protoBytes, err := rcPreTestSm.ActionDefinitionASN1toProto(actionDefinitionAsn1)
//	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
//	assert.Assert(t, protoBytes != nil)
//	assert.Equal(t, 4, len(protoBytes))
//	testIM := e2sm_rc_pre_ies.E2SmRcPreActionDefinition{}
//	err = proto.Unmarshal(protoBytes, &testIM)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, testIM.GetRicStyleType().GetValue())
//}

func TestServicemodel_Asn1BytesToByte(t *testing.T) {
	indHdrHex, err := rcPreTestSm.Asn1BytesToByte("2812f410abd4bc00")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indHdrHex)

	indHdrBytes, err := rcPreTestSm.IndicationHeaderASN1toProto(indHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, indHdrBytes != nil)
	assert.Equal(t, 24, len(indHdrBytes))
	testIH := &e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(indHdrBytes, testIH)
	assert.NilError(t, err)
	//assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testIH.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
	t.Logf("Decoded message (IndicationHeader) is \n%v\n", testIH)

	indMsgHex, err := rcPreTestSm.Asn1BytesToByte("4812f410abd4bc00fd6000000a0014000b00000000004012f410acd4bc00fd60000b")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indMsgHex)

	indMsgBytes, err := rcPreTestSm.IndicationMessageASN1toProto(indMsgHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, indMsgBytes != nil)
	assert.Equal(t, 82, len(indMsgBytes))
	testIM := &e2sm_rc_pre_ies.E2SmRcPreIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationMessage) is \n%v", testIM)

	actDefHex, err := rcPreTestSm.Asn1BytesToByte("00010c")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", actDefHex)

	actDefBytes, err := rcPreTestSm.ActionDefinitionASN1toProto(actDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, actDefBytes != nil)
	assert.Equal(t, 4, len(actDefBytes))
	testAD := &e2sm_rc_pre_ies.E2SmRcPreActionDefinition{}
	err = proto.Unmarshal(actDefBytes, testAD)
	assert.NilError(t, err)
	assert.Equal(t, int32(12), testAD.GetRicStyleType().GetValue())
	t.Logf("Decoded message (ActionDefinition) is \n%v\n", testAD)

	evntTrigDefHex, err := rcPreTestSm.Asn1BytesToByte("14010c")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", evntTrigDefHex)

	evntTrigDefBytes, err := rcPreTestSm.EventTriggerDefinitionASN1toProto(evntTrigDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, evntTrigDefBytes != nil)
	assert.Equal(t, 6, len(evntTrigDefBytes))
	testETD := &e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition{}
	err = proto.Unmarshal(evntTrigDefBytes, testETD)
	assert.NilError(t, err)
	assert.Equal(t, 12, int(testETD.GetEventDefinitionFormat1().GetReportingPeriodMs()))
	t.Logf("Decoded message (EventTriggerDefinition) is \n%v\n", testETD)

	ranFuncDescHex, err := rcPreTestSm.Asn1BytesToByte("20204f4e460000024f696406804f70656e4e6574776f726b696e67010360" +
		"00010d03804f4e466576656e74012a00010c04004f4e467265706f727401150138")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ranFuncDescHex)

	ranFuncDescBytes, err := rcPreTestSm.RanFuncDescriptionASN1toProto(ranFuncDescHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ranFuncDescBytes != nil)
	assert.Equal(t, 81, len(ranFuncDescBytes))
	testRFD := &e2sm_rc_pre_ies.E2SmRcPreRanfunctionDescription{}
	err = proto.Unmarshal(ranFuncDescBytes, testRFD)
	t.Logf("Decoded message (RanFunctionDescription) is \n%v\n", testRFD)
	assert.NilError(t, err)

	ctrlHdrHex, err := rcPreTestSm.Asn1BytesToByte("3412f410abd4bc000101")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ctrlHdrHex)

	ctrlHdrBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ctrlHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, ctrlHdrBytes != nil)
	assert.Equal(t, 28, len(ctrlHdrBytes))
	testCH := &e2sm_rc_pre_ies.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(ctrlHdrBytes, testCH)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	//assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testCH.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
	t.Logf("Decoded message (ControlHeader) is \n%v\n", testCH)

	ctrlMsgHex, err := rcPreTestSm.Asn1BytesToByte("0000010100504349000114")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ctrlMsgHex)

	ctrlMsgBytes, err := rcPreTestSm.ControlMessageASN1toProto(ctrlMsgHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, ctrlMsgBytes != nil)
	assert.Equal(t, 19, len(ctrlMsgBytes))
	testCM := &e2sm_rc_pre_ies.E2SmRcPreControlMessage{}
	err = proto.Unmarshal(ctrlMsgBytes, testCM)
	assert.NilError(t, err)
	assert.Equal(t, 20, int(testCM.GetControlMessage().GetParameterVal().GetValueInt()))
	t.Logf("Decoded message (ControlMessage) is \n%v\n", testCM)

	ctrlOutHex, err := rcPreTestSm.Asn1BytesToByte("20000000001400010a")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ctrlOutHex)

	ctrlOutBytes, err := rcPreTestSm.ControlOutcomeASN1toProto(ctrlOutHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ctrlOutBytes != nil)
	assert.Equal(t, 12, len(ctrlOutBytes))
	testCO := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(ctrlOutBytes, testCO)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(testCO.GetControlOutcomeFormat1().GetOutcomeElementList()))
	outcome := testCO.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, 10, int(outcome.GetRanParameterValue().GetValueInt()))
}

func TestServicemodel_HexDumpToByte(t *testing.T) {
	indHdrHex, err := rcPreTestSm.HexDumpToByte("00000000  28 12 f4 10 ab d4 bc 00                           |(.......|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", indHdrHex)

	indHdrBytes, err := rcPreTestSm.IndicationHeaderASN1toProto(indHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, indHdrBytes != nil)
	assert.Equal(t, 24, len(indHdrBytes))
	testIH := &e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(indHdrBytes, testIH)
	assert.NilError(t, err)
	//assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testIH.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
	t.Logf("Decoded message (IndicationHeader) is \n%v", testIH)

	indMsgHex, err := rcPreTestSm.HexDumpToByte("00000000  48 12 f4 10 ab d4 bc 00  fd 60 00 00 0a 00 14 00  |H........`......|" +
		"	00000010  0b 00 00 00 00 00 40 12  f4 10 ac d4 bc 00 fd 60  |......@........`|" +
		"        00000020  00 0b                                             |..|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", indMsgHex)

	indMsgBytes, err := rcPreTestSm.IndicationMessageASN1toProto(indMsgHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, indMsgBytes != nil)
	assert.Equal(t, 82, len(indMsgBytes))
	testIM := &e2sm_rc_pre_ies.E2SmRcPreIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message (IndicationMessage) is \n%v", testIM)

	actDefHex, err := rcPreTestSm.HexDumpToByte("00000000  00 01 0c                                          |...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", actDefHex)

	actDefBytes, err := rcPreTestSm.ActionDefinitionASN1toProto(actDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, actDefBytes != nil)
	assert.Equal(t, 4, len(actDefBytes))
	testAD := &e2sm_rc_pre_ies.E2SmRcPreActionDefinition{}
	err = proto.Unmarshal(actDefBytes, testAD)
	assert.NilError(t, err)
	assert.Equal(t, int32(0), testIM.GetRicStyleType().GetValue())
	t.Logf("Decoded message (ActionDefinition) is \n%v\n", testAD)

	evntTrigDefHex, err := rcPreTestSm.HexDumpToByte("00000000  14 01 0c                                          |...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", evntTrigDefHex)

	evntTrigDefBytes, err := rcPreTestSm.EventTriggerDefinitionASN1toProto(evntTrigDefHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, evntTrigDefBytes != nil)
	assert.Equal(t, 6, len(evntTrigDefBytes))
	testETD := &e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition{}
	err = proto.Unmarshal(evntTrigDefBytes, testETD)
	assert.NilError(t, err)
	assert.Equal(t, 12, int(testETD.GetEventDefinitionFormat1().GetReportingPeriodMs()))
	t.Logf("Decoded message (EventTriggerDefinition) is \n%v\n", testETD)

	ranFuncDescHex, err := rcPreTestSm.HexDumpToByte("00000000  20 20 4f 4e 46 00 00 02  4f 69 64 06 80 4f 70 65  |  ONF...Oid..Ope|" +
		"	00000010  6e 4e 65 74 77 6f 72 6b  69 6e 67 01 03 60 00 01  |nNetworking..`..|" +
		"        00000020  0d 03 80 4f 4e 46 65 76  65 6e 74 01 2a 00 01 0c  |...ONFevent.*...|" +
		"        00000030  04 00 4f 4e 46 72 65 70  6f 72 74 01 15 01 38     |..ONFreport...8|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ranFuncDescHex)

	ranFuncDescBytes, err := rcPreTestSm.RanFuncDescriptionASN1toProto(ranFuncDescHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ranFuncDescBytes != nil)
	assert.Equal(t, 81, len(ranFuncDescBytes))
	testRFD := &e2sm_rc_pre_ies.E2SmRcPreRanfunctionDescription{}
	err = proto.Unmarshal(ranFuncDescBytes, testRFD)
	t.Logf("Decoded message (RanFunctionDescription) is \n%v\n", testRFD)
	assert.NilError(t, err)

	ctrlHdrHex, err := rcPreTestSm.HexDumpToByte("00000000  34 12 f4 10 ab d4 bc 00  01 01                    |4.........|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ctrlHdrHex)

	ctrlHdrBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ctrlHdrHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, ctrlHdrBytes != nil)
	assert.Equal(t, 28, len(ctrlHdrBytes))
	testCH := &e2sm_rc_pre_ies.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(ctrlHdrBytes, testCH)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	//assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, testCH.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
	t.Logf("Decoded message (ControlHeader) is \n%v\n", testRFD)

	ctrlMsgHex, err := rcPreTestSm.HexDumpToByte("00000000  00 00 01 01 00 50 43 49  00 01 14                 |.....PCI...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ctrlMsgHex)

	ctrlMsgBytes, err := rcPreTestSm.ControlMessageASN1toProto(ctrlMsgHex)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, ctrlMsgBytes != nil)
	assert.Equal(t, 19, len(ctrlMsgBytes))
	testCM := &e2sm_rc_pre_ies.E2SmRcPreControlMessage{}
	err = proto.Unmarshal(ctrlMsgBytes, testCM)
	assert.NilError(t, err)
	assert.Equal(t, 20, int(testCM.GetControlMessage().GetParameterVal().GetValueInt()))
	t.Logf("Decoded message (ControlMessage) is \n%v\n", testCM)

	ctrlOutHex, err := rcPreTestSm.HexDumpToByte("00000000  20 00 00 00 00 14 00 01  0a                       | ........|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ctrlMsgHex)

	ctrlOutBytes, err := rcPreTestSm.ControlOutcomeASN1toProto(ctrlOutHex)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, ctrlOutBytes != nil)
	assert.Equal(t, 12, len(ctrlOutBytes))
	testCO := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(ctrlOutBytes, testCO)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(testCO.GetControlOutcomeFormat1().GetOutcomeElementList()))
	outcome := testCO.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, 10, int(outcome.GetRanParameterValue().GetValueInt()))
	t.Logf("Decoded message (ControlOutcome) is \n%v\n", testCO)
}
