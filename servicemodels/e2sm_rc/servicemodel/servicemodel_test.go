//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/pdubuilder"
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var rcTestSm RCServiceModel

func TestRCServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	ueID, err := pdubuilder.CreateUeIDNgEnbDu(123)
	assert.NilError(t, err)

	msg, err := pdubuilder.CreateE2SmRcIndicationHeaderFormat2(ueID, 1, 1)
	assert.NilError(t, err)

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcIndicationHeader to bytes")

	asn1Bytes, err := rcTestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-PRE-IndicationHeader are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	indicationHeaderAsn1Bytes := []byte{0x24, 0x00, 0x7b, 0x01, 0x01, 0x00, 0x00, 0x00}

	protoBytes, err := rcTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)

	testIH := &e2smrcv1.E2SmRcIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-IndicationHeader is \n%v", testIH)
}

func TestRCServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	ueParameterList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat2Item, 0)

	ueID1, err := pdubuilder.CreateUeIDGNb(1, []byte{0xAB, 0xCD, 0xEF}, []byte{0x21}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)
	ueID1.GetGNbUeid().SetRanUeID([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}).SetMNgRanUeXnApID(1)

	ueID2, err := pdubuilder.CreateUeIDNgEnb(1, []byte{0x1B, 0xD9, 0xEF}, []byte{0x21}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem, 0)

	ranParameterValue, err := pdubuilder.CreateRanparameterValueBoolean(true)
	assert.NilError(t, err)

	ranParameterValueType, err := pdubuilder.CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranPItem, err := pdubuilder.CreateE2SmRcIndicationMessageFormat2RanparameterItem(1, ranParameterValueType)
	assert.NilError(t, err)
	ranPList = append(ranPList, ranPItem)

	item1, err := pdubuilder.CreateE2SmRcIndicationMessageFormat2Item(ueID1, ranPList)
	assert.NilError(t, err)
	item2, err := pdubuilder.CreateE2SmRcIndicationMessageFormat2Item(ueID2, ranPList)
	assert.NilError(t, err)

	ueParameterList = append(ueParameterList, item1)
	ueParameterList = append(ueParameterList, item2)

	msg, err := pdubuilder.CreateE2SmRcIndicationMessageFormat2(ueParameterList)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, msg != nil)

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcIndicationMessage to bytes")

	asn1Bytes, err := rcTestSm.IndicationMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-IndicationMessage (Format2) are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	indicationMessageAsn1 := []byte{0x08, 0x00, 0x01, 0x00, 0xc0, 0x01, 0x00, 0xab, 0xcd, 0xef, 0x21, 0xff, 0xff, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8c, 0x00, 0x01, 0x00,
		0x1b, 0xd9, 0xef, 0x21, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}

	protoBytes, err := rcTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-IndicationMessage is \n%v", testIM)
}

func TestRCServicemodel_RanFuncDefinitionProtoToASN1(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC Insert"

	ranFunctionDefinitionInsert := &e2smrcv1.RanfunctionDefinitionInsert{
		RicInsertStyleList: make([]*e2smrcv1.RanfunctionDefinitionInsertItem, 0),
	}

	// adding RanReportParametersList
	insertIndicationList := make([]*e2smrcv1.RanfunctionDefinitionInsertIndicationItem, 0)
	item, err := pdubuilder.CreateRanfunctionDefinitionInsertIndicationItem(1, "E2SM-RC Report Item")
	assert.NilError(t, err)

	insertIndicationItemList := make([]*e2smrcv1.InsertIndicationRanparameterItem, 0)
	insertIndicationItem, err := pdubuilder.CreateInsertIndicationRanparameterItem(1, "Insert Indication Item")
	assert.NilError(t, err)
	insertIndicationItemList = append(insertIndicationItemList, insertIndicationItem)

	item.SetRanInsertIndicationParametersList(insertIndicationItemList)
	insertIndicationList = append(insertIndicationList, item)

	insertItem, err := pdubuilder.CreateRanfunctionDefinitionInsertItem(1, "E2SM-RC Report", 1, 1, 1, 1, 1)
	assert.NilError(t, err)
	insertItem.SetRicInsertIndicationList(insertIndicationList)

	ranFunctionDefinitionInsert.RicInsertStyleList = append(ranFunctionDefinitionInsert.RicInsertStyleList, insertItem)

	msg, err := pdubuilder.CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg != nil)
	msg.SetRanFunctionDefinitionInsert(ranFunctionDefinitionInsert).GetRanFunctionName().SetRanFunctionInstance(4)

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcRanfunctionDescription to bytes")

	asn1Bytes, err := rcTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-RanFunctionDescription are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	// This message is taken as an output from the function above
	ranFuncDescriptionAsn1 := []byte{0x11, 0x03, 0x00, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x00, 0x00, 0x18, 0x31, 0x2e, 0x33,
		0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31,
		0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x33, 0x06, 0x80, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x20,
		0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x00, 0x00, 0x04, 0x02, 0x80, 0x00, 0x01, 0x06, 0x80, 0x45,
		0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x00, 0x40, 0x00, 0x00, 0x09, 0x00, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x20,
		0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x49, 0x74, 0x65, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x0a,
		0x80, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69,
		0x6f, 0x6e, 0x20, 0x49, 0x74, 0x65, 0x6d, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01}

	protoBytes, err := rcTestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testRFD := &e2smrcv1.E2SmRcRanfunctionDefinition{}
	err = proto.Unmarshal(protoBytes, testRFD)
	t.Logf("Decoded RC-RanFunctionDescription is \n%v", testRFD)
	assert.NilError(t, err)
	assert.Equal(t, string(rcTestSm.ServiceModelData().OID), testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(4), testRFD.GetRanFunctionName().GetRanFunctionInstance())
}

func TestRCServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	associatedUeInfo := make([]*e2smrcv1.EventTriggerUeInfoItem, 0)

	ranParameterTesting := make([]*e2smrcv1.RanparameterTestingItem, 0)

	ranParameterValue, err := pdubuilder.CreateRanparameterValueInt(42)
	assert.NilError(t, err)

	ranParameterType, err := pdubuilder.CreateRanParameterTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranParameterTestingItem, err := pdubuilder.CreateRanparameterTestingItem(1, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting = append(ranParameterTesting, ranParameterTestingItem)

	ueType, err := pdubuilder.CreateUeTypeGroup(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})
	assert.NilError(t, err)

	ueInfo, err := pdubuilder.CreateEventTriggerUeInfoItem(2, ueType)
	assert.NilError(t, err)
	ueInfo.SetLogicalOr(pdubuilder.CreateLogicalOrTrue())

	associatedUeInfo = append(associatedUeInfo, ueInfo)

	messageList := make([]*e2smrcv1.E2SmRcEventTriggerFormat1Item, 0)

	rrcType, err := pdubuilder.CreateRrcTypeLte(e2smcommonies.RrcclassLte_RRCCLASS_LTE_B_CCH_BCH)
	assert.NilError(t, err)

	msgTypeChoice, err := pdubuilder.CreateMessageTypeChoiceRrc(rrcType, 32)
	assert.NilError(t, err)

	messageItem, err := pdubuilder.CreateE2SmRcEventTriggerFormat1Item(1, msgTypeChoice, nil, nil, nil, nil)
	assert.NilError(t, err)
	messageItem.SetAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	}).SetMessageDirection(pdubuilder.CreateMessageDirectionOutgoing()).SetLogicalOr(pdubuilder.CreateLogicalOrFalse())

	messageList = append(messageList, messageItem)

	msg, err := pdubuilder.CreateE2SmRcEventTriggerFormat1(messageList)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, msg != nil, "Created E2SmPDU is nil")
	msg.GetRicEventTriggerFormats().GetEventTriggerFormat1().SetGlobalAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	})

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcEventTriggerDefinition to bytes")

	asn1Bytes, err := rcTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	t.Logf("ASN1 bytes for RC-EventTriggerDefinition are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x02, 0x00, 0x00, 0x68, 0x00, 0x00, 0x40, 0x00, 0x01, 0x20, 0x40, 0x00, 0x00, 0x40, 0x00, 0x01,
		0x40, 0x00, 0x00, 0x41, 0x01, 0x2a, 0x10, 0x00, 0x00, 0x40, 0x00, 0x01, 0x40, 0x00, 0x00, 0x41,
		0x01, 0x2a, 0x00}

	protoBytes, err := rcTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcEventTrigger{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-EventTriggerDefinition is \n%v", testIM)
}

func TestRCServicemodel_ControlHeaderProtoToASN1(t *testing.T) {

	ueID, err := pdubuilder.CreateUeIDGnbCuUp(15)
	assert.NilError(t, err)

	msg, err := pdubuilder.CreateE2SmRcControlHeaderFormat1(ueID, 12, 5)
	assert.NilError(t, err)
	msg.GetRicControlHeaderFormats().GetControlHeaderFormat1().SetRicControlDecision(pdubuilder.CreateRicControlDecisionAccept())

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcControlHeader to bytes")

	asn1Bytes, err := rcTestSm.ControlHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-ControlHeader are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x12, 0x00, 0x0f, 0x01, 0x0c, 0x00, 0x00, 0x04, 0x00}

	protoBytes, err := rcTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)

	testCH := &e2smrcv1.E2SmRcControlHeader{}
	err = proto.Unmarshal(protoBytes, testCH)
	assert.NilError(t, err)
	t.Logf("Decoded RC-ControlHeader is \n%v", testCH)
}

func TestRCServicemodel_ControlMessageProtoToASN1(t *testing.T) {
	ranParameterValue, err := pdubuilder.CreateRanparameterValueBitS(&asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		Len:   32,
	})
	assert.NilError(t, err)

	ranParameterValueType, err := pdubuilder.CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	item, err := pdubuilder.CreateE2SmRcControlMessageFormat1Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlMessageFormat1Item, 0)
	ranPList = append(ranPList, item)

	msg, err := pdubuilder.CreateE2SmRcControlMessageFormat1(ranPList)
	assert.NilError(t, err)

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcControlMessage to bytes")

	asn1Bytes, err := rcTestSm.ControlMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-ControlMessage are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_ControlMessageASN1toProto(t *testing.T) {
	ControlMessageAsn1 := []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x29, 0x80, 0x20, 0xff, 0xff, 0xff, 0xff}

	protoBytes, err := rcTestSm.ControlMessageASN1toProto(ControlMessageAsn1)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcControlMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-ControlMessage is \n%v", testIM)
}

func TestRCServicemodel_ControlOutcomeProtoToASN1(t *testing.T) {
	ranParameterValue, err := pdubuilder.CreateRanparameterValuePrintableString("Ran Parameter Value")
	assert.NilError(t, err)

	item, err := pdubuilder.CreateE2SmRcControlOutcomeFormat1Item(1, ranParameterValue)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat1Item, 0)
	ranPList = append(ranPList, item)

	msg, err := pdubuilder.CreateE2SmRcControlOutcomeFormat1(ranPList)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcControlOutcome to bytes")

	asn1Bytes, err := rcTestSm.ControlOutcomeProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-IndicationOutcome are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_ControlOutcomeASN1toProto(t *testing.T) {
	controlOutcomeAsn1 := []byte{0x00, 0x01, 0x00, 0x00, 0x50, 0x13, 0x52, 0x61, 0x6e, 0x20, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
		0x74, 0x65, 0x72, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65}

	protoBytes, err := rcTestSm.ControlOutcomeASN1toProto(controlOutcomeAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcControlOutcome{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-ControlOutcome is \n%v", testIM)
}

func TestRCServiceModel_ActionDefinitionProtoToASN1(t *testing.T) {
	actionDefinitionItemFormat2List := make([]*e2smrcv1.E2SmRcActionDefinitionFormat2Item, 0)

	ricPolicyActionList := make([]*e2smrcv1.RicPolicyActionRanparameterItem, 0)

	ranParameterValue, err := pdubuilder.CreateRanparameterValueBitS(&asn1.BitString{
		Value: []byte{0xC0},
		Len:   8,
	})
	assert.NilError(t, err)
	ranParameterValueType, err := pdubuilder.CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	actionItem1, err := pdubuilder.CreateRicPolicyActionRanParameterItem(12, ranParameterValueType)
	assert.NilError(t, err)
	ricPolicyActionList = append(ricPolicyActionList, actionItem1)

	item1, err := pdubuilder.CreateE2SmRcActionDefinitionFormat2Item(25, ricPolicyActionList)
	assert.NilError(t, err)
	actionDefinitionItemFormat2List = append(actionDefinitionItemFormat2List, item1)

	ranParameterTestingList := &e2smrcv1.RanparameterTestingList{
		Value: make([]*e2smrcv1.RanparameterTestingItem, 0),
	}

	innerRanParameterValue, err := pdubuilder.CreateRanparameterValueInt(37)
	assert.NilError(t, err)

	ranParameterTestingCondition, err := pdubuilder.CreateRanparameterTestingConditionComparison(pdubuilder.CreateRanPChoiceComparisonStartsWith())
	assert.NilError(t, err)

	innerRanParameterType, err := pdubuilder.CreateRanParameterTypeChoiceElementFalse(ranParameterTestingCondition)
	assert.NilError(t, err)
	innerRanParameterType.GetRanPChoiceElementFalse().SetRanParameterValue(innerRanParameterValue).SetLogicalOr(pdubuilder.CreateLogicalOrFalse())

	innerRanParameterTestingItem, err := pdubuilder.CreateRanparameterTestingItem(94, innerRanParameterType)
	assert.NilError(t, err)
	ranParameterTestingList.Value = append(ranParameterTestingList.Value, innerRanParameterTestingItem)

	ranParameterType, err := pdubuilder.CreateRanParameterTypeChoiceList(ranParameterTestingList)
	assert.NilError(t, err)

	ranParameterTestingItem, err := pdubuilder.CreateRanparameterTestingItem(69, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting := &e2smrcv1.RanparameterTesting{
		Value: make([]*e2smrcv1.RanparameterTestingItem, 0),
	}
	ranParameterTesting.Value = append(ranParameterTesting.Value, ranParameterTestingItem)

	actionItem2, err := pdubuilder.CreateE2SmRcActionDefinitionFormat2Item(27, ricPolicyActionList)
	assert.NilError(t, err)
	actionItem2.SetRicPolicyConditionDefinition(ranParameterTesting)
	actionDefinitionItemFormat2List = append(actionDefinitionItemFormat2List, actionItem2)

	msg, err := pdubuilder.CreateE2SmRcActionDefinitionFormat2(11, actionDefinitionItemFormat2List)
	assert.NilError(t, err)

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcActionDefinition to bytes")

	asn1Bytes, err := rcTestSm.ActionDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-ActionDefintiion are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_ActionDefinitionASN1toProto(t *testing.T) {
	actionDefinitionAsn1 := []byte{0x00, 0x01, 0x0b, 0x20, 0x00, 0x01, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x0b, 0x29, 0x80, 0x08,
		0xc0, 0x40, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x0b, 0x29, 0x80, 0x08, 0xc0, 0x00, 0x00, 0x44, 0x00,
		0x00, 0x00, 0x00, 0x5d, 0x6c, 0x51, 0x01, 0x25, 0x40}

	protoBytes, err := rcTestSm.ActionDefinitionASN1toProto(actionDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcActionDefinition{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-ActionDefinition is \n%v", testIM)
}

func TestRCServiceModel_CallProcessIDProtoToASN1(t *testing.T) {
	msg, err := pdubuilder.CreateE2SmRcCallProcessIDFormat1(11)
	assert.NilError(t, err, "error creating E2SmPDU")

	protoBytes, err := proto.Marshal(msg)
	assert.NilError(t, err, "unexpected error marshalling E2SmRcCallProcessID to bytes")

	asn1Bytes, err := rcTestSm.CallProcessIDProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for RC-CallProcessID are \n%v", hex.Dump(asn1Bytes))
}

func TestRCServicemodel_CallProcessIDASN1toProto(t *testing.T) {
	callProcessIDAsn1 := []byte{0x00, 0xa0}

	protoBytes, err := rcTestSm.CallProcessIDASN1toProto(callProcessIDAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)

	testIM := &e2smrcv1.E2SmRcCallProcessId{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded RC-CallProcessID is \n%v", testIM)
	assert.Equal(t, testIM.GetRicCallProcessIdFormats().GetCallProcessIdFormat1().GetRicCallProcessId().GetValue(), int32(11))
}
