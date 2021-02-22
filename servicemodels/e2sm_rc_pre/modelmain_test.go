// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var rcPreTestSm servicemodel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "ONF"

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(plmnID)
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
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "ONF"

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationMsg(plmnID)
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
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int32 = 12
	e2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmRcPreEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmRcPreEventTriggerDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreEventTriggerDefinition to bytes")
	assert.Equal(t, 4, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 3, len(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x10, 0x01, 0x0C}
	protoBytes, err := rcPreTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 4, len(protoBytes))
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
}

func TestServicemodel_ControlHeaderProtoToASN1(t *testing.T) {

	var controlMessagePriority int32 = 1

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader(controlMessagePriority)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmRcPrePdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcPreControlHeader to bytes")
	assert.Equal(t, 6, len(protoBytes))

	asn1Bytes, err := rcPreTestSm.ControlHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 3, len(asn1Bytes))
}

func TestServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x20, 0x01, 0x01}

	protoBytes, err := rcPreTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 6, len(protoBytes))
	testIH := e2sm_rc_pre_ies.E2SmRcPreControlHeader{}
	err = proto.Unmarshal(protoBytes, &testIH)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(testIH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
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
}

func TestServicemodel_ControlMessageASN1toProto(t *testing.T) {
	ControlMessageAsn1 := []byte{0x00, 0x00, 0x14, 0x01, 0x00, 0x50, 0x43, 0x49, 0x00, 0x06,
		0x00, 0xc0, 0x00, 0x0c, 0xbb, 0x70}

	protoBytes, err := rcPreTestSm.ControlMessageASN1toProto(ControlMessageAsn1)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 19, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreControlMessage{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 20, int(testIM.GetControlMessage().GetParameterVal().GetValueInt()))
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
}

func TestServicemodel_ControlOutcomeASN1toProto(t *testing.T) {

	ControlOutcomeAsn1 := []byte{0x20, 0x00, 0x00, 0x00, 0x00, 0x14, 0x00, 0x06,
		0x00, 0xc0, 0x00, 0x0c, 0xbb, 0x40}
	protoBytes, err := rcPreTestSm.ControlOutcomeASN1toProto(ControlOutcomeAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 12, len(protoBytes))
	testIM := e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(testIM.GetControlOutcomeFormat1().GetOutcomeElementList()))
	outcome := testIM.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, 20, int(outcome.GetRanParameterValue().GetValueInt()))
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
