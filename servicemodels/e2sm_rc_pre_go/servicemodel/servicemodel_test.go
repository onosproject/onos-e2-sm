// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/pdubuilder"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var rcPreTestSm RcPreServiceModel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreIndicationHeader to bytes")

	asn1Bytes, err := rcPreTestSm.IndicationHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-IndicationHeader are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-RC-PRE-Indication-Header.xml

	indicationHeaderAsn1Bytes := []byte{0x28, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x00}

	protoBytes, err := rcPreTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)

	testIH := &e2sm_rc_pre_go.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-IndicationHeader is \n%v", testIH)
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, testIH.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationHeaderNrCGIProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreIndicationHeader (NrCGI) to bytes")

	asn1Bytes, err := rcPreTestSm.IndicationHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-IndicationHeader (NrCGI) are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderNrCGIASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-RC-PRE-Indication-Header.xml

	indicationHeaderAsn1Bytes := []byte{0x20, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x09, 0x00}

	protoBytes, err := rcPreTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testIH := &e2sm_rc_pre_go.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-IndicationHeader (NrCGI) is \n%v", testIH)
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, testIH.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	dlArfcn, err := pdubuilder.CreateEArfcn(253)
	assert.NilError(t, err)
	var pci int32 = 11
	cellSize := e2sm_rc_pre_go.CellSize_CELL_SIZE_MACRO

	cellID := asn1.BitString{
		Value: []byte{0xac, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	neighbor, err := pdubuilder.CreateNrt(cgi, dlArfcn, cellSize, &e2sm_rc_pre_go.Pci{
		Value: pci,
	})
	assert.NilError(t, err)

	neighbors := make([]*e2sm_rc_pre_go.Nrt, 0)
	neighbors = append(neighbors, neighbor)

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes, dlArfcn, cellSize, pci, neighbors)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmRcPrePdu != nil)

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreIndicationMessage to bytes")

	asn1Bytes, err := rcPreTestSm.IndicationMessageProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-IndicationMessage (Format1) are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	//indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}

	indicationMessageAsn1 := []byte{0x40, 0xfd, 0x60, 0x00, 0x0b, 0x00, 0x01, 0x20, 0x12, 0xf4,
		0x10, 0xac, 0xd4, 0xbc, 0x00, 0xfd, 0x60, 0x00, 0x0b}
	protoBytes, err := rcPreTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testIM := &e2sm_rc_pre_go.E2SmRcPreIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-IndicationMessage is \n%v", testIM)

	assert.Equal(t, 1, len(testIM.GetIndicationMessageFormat1().GetNeighbors()))
	nbr := testIM.GetIndicationMessageFormat1().GetNeighbors()[0]
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, nbr.GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
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

	retsl := make([]*e2sm_rc_pre_go.RicEventTriggerStyleList, 0)
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_rc_pre_go.RicReportStyleList, 0)
	rrsi := pdubuilder.CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl).GetRanFunctionName().SetRanFunctionInstance(ranFunctionInstance)

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreRanfunctionDescription to bytes")

	asn1Bytes, err := rcPreTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-RanFunctionDescription are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	// This message is taken as an output from the function above
	ranFuncDescriptionAsn1 := []byte{0x20, 0x20, 0x4f, 0x4e, 0x46, 0x00, 0x00, 0x02, 0x4f, 0x69, 0x64, 0x06, 0x80, 0x4f,
		0x70, 0x65, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x00, 0x03, 0x60, 0x00, 0x0d, 0x03,
		0x80, 0x4f, 0x4e, 0x46, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x00, 0x2a, 0x00, 0x0c, 0x04, 0x00, 0x4f, 0x4e, 0x46, 0x72,
		0x65, 0x70, 0x6f, 0x72, 0x74, 0x00, 0x15, 0x00, 0x38}

	protoBytes, err := rcPreTestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testRFD := &e2sm_rc_pre_go.E2SmRcPreRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, testRFD)
	t.Logf("Decoded RC-PRE-RanFunctionDescription is \n%v", testRFD)
	assert.NilError(t, err)
	assert.Equal(t, "Oid", testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(3), testRFD.GetRanFunctionName().GetRanFunctionInstance())
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int64 = 12
	e2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmRcPreEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreEventTriggerDefinition to bytes")

	asn1Bytes, err := rcPreTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	t.Logf("ASN1 bytes for RC-PRE-EventTriggerDefinition are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x14, 0x0b}
	protoBytes, err := rcPreTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testIM := &e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-EventTriggerDefinition is \n%v", testIM)
	assert.Equal(t, int64(12), testIM.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriodMs())
}

func TestServicemodel_ControlHeaderProtoToASN1(t *testing.T) {

	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}

	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetCGI(cgi).SetRicControlMessagePriority(controlMessagePriority)
	assert.NilError(t, err, "error creating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlHeader to bytes")

	asn1Bytes, err := rcPreTestSm.ControlHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-ControlHeader are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x34, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x00, 0x01}

	protoBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testCH := &e2sm_rc_pre_go.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(protoBytes, testCH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-ControlHeader is \n%v", testCH)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, testCH.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.Equal(t, uint32(28), testCH.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func TestServicemodel_ControlHeaderNrCGIProtoToASN1(t *testing.T) {

	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}

	cgi, err := pdubuilder.CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetCGI(cgi).SetRicControlMessagePriority(controlMessagePriority)

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlHeader to bytes")

	asn1Bytes, err := rcPreTestSm.ControlHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-ControlHeader (with NrCGI) are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderNrCGIASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x30, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x09, 0x00, 0x01}

	protoBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testCH := &e2sm_rc_pre_go.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(protoBytes, testCH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-ControlHeader is \n%v", testCH)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, testCH.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().GetValue())
	assert.Equal(t, uint32(36), testCH.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}

func TestServicemodel_ControlMessageProtoToASN1(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue int64 = 20
	var ranParameterID int32 = 1

	ranParameter, err := pdubuilder.CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlMessage to bytes")

	asn1Bytes, err := rcPreTestSm.ControlMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-ControlMessage are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlMessageASN1toProto(t *testing.T) {
	ControlMessageAsn1 := []byte{0x00, 0x00, 0x01, 0x01, 0x00, 0x50, 0x43, 0x49, 0x00, 0x00, 0x14}

	protoBytes, err := rcPreTestSm.ControlMessageASN1toProto(ControlMessageAsn1)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 19, len(protoBytes))
	testIM := &e2sm_rc_pre_go.E2SmRcPreControlMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-ControlMessage is \n%v", testIM)
	assert.Equal(t, int64(20), testIM.GetControlMessage().GetParameterVal().GetValueInt())
}

func TestServicemodel_ControlOutcomeProtoToASN1(t *testing.T) {
	var ranParameterID int32 = 20

	elementList := make([]*e2sm_rc_pre_go.RanparameterItem, 0)
	elementList = append(elementList, pdubuilder.CreateRanParameterItem(ranParameterID))

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlOutcome(elementList)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlOutcome to bytes")

	asn1Bytes, err := rcPreTestSm.ControlOutcomeProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-IndicationOutcome are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlOutcomeASN1toProto(t *testing.T) {

	ControlOutcomeAsn1 := []byte{0x20, 0x00, 0x00, 0x00, 0x00, 0x14}
	protoBytes, err := rcPreTestSm.ControlOutcomeASN1toProto(ControlOutcomeAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2sm_rc_pre_go.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-ControlOutcome is \n%v", testIM)

	assert.Equal(t, 1, len(testIM.GetControlOutcomeFormat1().GetOutcomeElementList()))
	outcome := testIM.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, int32(20), outcome.GetRanParameterId().GetValue())
}
