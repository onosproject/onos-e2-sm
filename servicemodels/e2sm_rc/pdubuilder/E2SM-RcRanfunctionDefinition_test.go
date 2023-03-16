//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcRanfunctionDefinition(t *testing.T) {

	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription1 := "empty E2SM-RC"

	msg, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription1)
	assert.NilError(t, err)
	assert.Assert(t, msg != nil)
	msg.GetRanFunctionName().SetRanFunctionInstance(1)

	aper, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}

func TestE2SmRcRanfunctionDefinitionEventTrigger(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC EventTrigger"

	eventTriggerList := &e2smrcv1.RanfunctionDefinitionEventTrigger{
		RicEventTriggerStyleList: make([]*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, 0),
		// other parameters are optional
		RanL2ParametersList:                 make([]*e2smrcv1.L2ParametersRanparameterItem, 0),
		RanCellIdentificationParametersList: make([]*e2smrcv1.CellIdentificationRanparameterItem, 0),
	}

	ricEventTriggerStyleItem, err := CreateRanfunctionDefinitionEventTriggerStyleItem(1, "SomeName", 1)
	assert.NilError(t, err)
	eventTriggerList.RicEventTriggerStyleList = append(eventTriggerList.RicEventTriggerStyleList, ricEventTriggerStyleItem)

	// adding optional L2RanParameterItem
	l2RanParameterItem, err := CreateL2ParametersRanparameterItem(1, "L2 parameter")
	assert.NilError(t, err)

	ranParameterStructureItemList := make([]*e2smrcv1.RanparameterDefinitionChoiceStructureItem, 0)
	item, err := CreateRanparameterDefinitionChoiceStructureItem(1, "Structure ID")
	assert.NilError(t, err)
	ranParameterStructureItemList = append(ranParameterStructureItemList, item)

	ranParameterDefinitionChoice, err := CreateRanparameterDefinitionChoiceStructure(ranParameterStructureItemList)
	assert.NilError(t, err)

	ranParameterDefinition, err := CreateRanparameterDefinition(ranParameterDefinitionChoice)
	assert.NilError(t, err)

	l2RanParameterItem.SetRanParameterDefinition(ranParameterDefinition)
	eventTriggerList.RanL2ParametersList = append(eventTriggerList.RanL2ParametersList, l2RanParameterItem)

	// adding optional CellIdentificationParameterItem
	cellIdentificationItem, err := CreateCellIdentificationRanparameterItem(25, "Serving Cell")
	assert.NilError(t, err)
	eventTriggerList.RanCellIdentificationParametersList = append(eventTriggerList.RanCellIdentificationParametersList, cellIdentificationItem)

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionEventTrigger(eventTriggerList).GetRanFunctionName().SetRanFunctionInstance(2)

	aper2, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg2)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper2))

	result2, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper2)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())
}

func TestE2SmRcRanfunctionDefinitionReport(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC Report"

	ranFunctionDefinitionReport := &e2smrcv1.RanfunctionDefinitionReport{
		RicReportStyleList: make([]*e2smrcv1.RanfunctionDefinitionReportItem, 0),
	}

	// adding RanReportParametersList
	reportParametersList := make([]*e2smrcv1.ReportRanparameterItem, 0)
	item, err := CreateReportRanparameterItem(1, "E2SM-RC Report Item")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item)

	reportItem, err := CreateRanfunctionDefinitionReportItem(1, "E2SM-RC Report", 1, 1, 1, 1)
	assert.NilError(t, err)
	reportItem.SetRanReportParametersList(reportParametersList)

	ranFunctionDefinitionReport.RicReportStyleList = append(ranFunctionDefinitionReport.RicReportStyleList, reportItem)

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionReport(ranFunctionDefinitionReport).GetRanFunctionName().SetRanFunctionInstance(3)

	aper2, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg2)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper2))

	result2, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper2)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())
}

func TestE2SmRcRanfunctionDefinitionInsert(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC Insert"

	ranFunctionDefinitionInsert := &e2smrcv1.RanfunctionDefinitionInsert{
		RicInsertStyleList: make([]*e2smrcv1.RanfunctionDefinitionInsertItem, 0),
	}

	// adding RanReportParametersList
	insertIndicationList := make([]*e2smrcv1.RanfunctionDefinitionInsertIndicationItem, 0)
	item, err := CreateRanfunctionDefinitionInsertIndicationItem(1, "E2SM-RC Report Item")
	assert.NilError(t, err)

	insertIndicationItemList := make([]*e2smrcv1.InsertIndicationRanparameterItem, 0)
	insertIndicationItem, err := CreateInsertIndicationRanparameterItem(1, "Insert Indication Item")
	assert.NilError(t, err)
	insertIndicationItemList = append(insertIndicationItemList, insertIndicationItem)

	item.SetRanInsertIndicationParametersList(insertIndicationItemList)
	insertIndicationList = append(insertIndicationList, item)

	insertItem, err := CreateRanfunctionDefinitionInsertItem(1, "E2SM-RC Report", 1, 1, 1, 1, 1)
	assert.NilError(t, err)
	insertItem.SetRicInsertIndicationList(insertIndicationList)

	ranFunctionDefinitionInsert.RicInsertStyleList = append(ranFunctionDefinitionInsert.RicInsertStyleList, insertItem)

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionInsert(ranFunctionDefinitionInsert).GetRanFunctionName().SetRanFunctionInstance(4)

	aper2, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg2)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper2))

	result2, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper2)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())
}

func TestE2SmRcRanfunctionDefinitionControl(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC Control"

	controlItemList := make([]*e2smrcv1.RanfunctionDefinitionControlItem, 0)

	controlActionList := make([]*e2smrcv1.RanfunctionDefinitionControlActionItem, 0)
	controlActionItem, err := CreateRanfunctionDefinitionControlActionItem(1, "Control Action Item")
	assert.NilError(t, err)
	controlActionList = append(controlActionList, controlActionItem)

	controlOutcomeList := make([]*e2smrcv1.ControlOutcomeRanparameterItem, 0)
	controlOutcomeItem, err := CreateControlOutcomeRanparameterItem(1, "Control Outcome Item")
	assert.NilError(t, err)
	controlOutcomeList = append(controlOutcomeList, controlOutcomeItem)

	controlItem, err := CreateRanfunctionDefinitionControlItem(1, "Control Item", 1, 1, 1)
	assert.NilError(t, err)
	controlItem.SetRicCallProcessIDformatType(1).SetRicControlActionList(controlActionList).SetRanControlOutcomeParametersList(controlOutcomeList)
	controlItemList = append(controlItemList, controlItem)

	ranFunctionDefinitionControl, err := CreateRanfunctionDefinitionControl(controlItemList)
	assert.NilError(t, err)

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionControl(ranFunctionDefinitionControl).GetRanFunctionName().SetRanFunctionInstance(5)

	aper2, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg2)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper2))

	result2, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper2)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())
}

func TestE2SmRcRanfunctionDefinitionPolicy(t *testing.T) {
	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "E2SM-RC Policy"

	policyConditionParameterList := make([]*e2smrcv1.PolicyConditionRanparameterItem, 0)
	policyConditionParameterItem, err := CreatePolicyConditionRanparameterItem(1, "Policy Condition Item")
	assert.NilError(t, err)
	policyConditionParameterList = append(policyConditionParameterList, policyConditionParameterItem)

	policyActionParameterList := make([]*e2smrcv1.PolicyActionRanparameterItem, 0)
	policyActionParameterItem, err := CreatePolicyActionRanparameterItem(1, "Policy Action Parameter Item")
	assert.NilError(t, err)
	policyActionParameterList = append(policyActionParameterList, policyActionParameterItem)

	actionList := make([]*e2smrcv1.RanfunctionDefinitionPolicyActionItem, 0)
	actionItem, err := CreateRanfunctionDefinitionPolicyActionItem(1, "Policy Action Item", 1)
	assert.NilError(t, err)
	actionItem.SetRanPolicyConditionParametersList(policyConditionParameterList).SetRanPolicyActionParametersList(policyActionParameterList)
	actionList = append(actionList, actionItem)

	ranFunctionDefinitionPolicyList := make([]*e2smrcv1.RanfunctionDefinitionPolicyItem, 0)
	ranFunctionPolicyItem, err := CreateRanfunctionDefinitionPolicyItem(1, "Policy Item", 1)
	assert.NilError(t, err)
	ranFunctionPolicyItem.SetRicPolicyActionList(actionList)
	ranFunctionDefinitionPolicyList = append(ranFunctionDefinitionPolicyList, ranFunctionPolicyItem)

	ranFunctionDefinitionPolicy, err := CreateRanfunctionDefinitionPolicy(ranFunctionDefinitionPolicyList)
	assert.NilError(t, err)

	msg, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg != nil)
	msg.SetRanFunctionDefinitionPolicy(ranFunctionDefinitionPolicy).GetRanFunctionName().SetRanFunctionInstance(6)

	aper, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}

var rSysBytes1 = "7805804f52414e2d4532534d2d5243000018312e332e362e312e342e312e35333134382e312e312e322e3300805243202000020b0043616c6c2050726f6365737320427265616b706f696e74000200000106004d657373616765204576656e74000100000000050a805044552053657373696f6e204d616e6167656d656e7400024000000c805044552053657373696f6e205265736f7572636520536574757000021075320e004c697374206f6620516f5320466c6f777320746f2062652073657475701075330d00516f5320666c6f772073657475702072657175657374206974656d1075350380516f5320666c6f7740000110005044552053657373696f6e205265736f75726365204d6f64696669636174696f6e000210791a10804c697374206f6620516f5320466c6f777320746f20616464206f72206d6f6469667910791b1100516f5320666c6f7720616464206f72206d6f646966792072657175657374206974656d10791d0380516f5320666c6f770000020d805044552053657373696f6e205265736f757263652052656c65617365028000020b0043616c6c2050726f6365737320427265616b706f696e740002000100010002000000000580445242204964656e7469747940000105804d65737361676520436f70790001000100010001000000020500525243204d657373616765008000010d80526164696f2042656172657220436f6e74726f6c20526571756573740002000300004000030f00526164696f2061646d697373696f6e20636f6e74726f6c2072657175657374000000000580445242204964656e7469747900020005000102c000010980526164696f2042656172657220436f6e74726f6c00000000030b00526164696f2061646d697373696f6e20636f6e74726f6c00010001000100016026ad1080484f2044657374696e6174696f6e20446973747269627574656420436f6e74726f6c00000000001680484f2044657374696e6174696f6e20446973747269627574656420436f6e74726f6c20496e737472756374696f6e0001000100010001"

var rSysBytes = []byte{
	0x78, 0x05, 0x80, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x00, 0x00, 0x18, 0x31, 0x2e, 0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e,
	0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x33, 0x00, 0x80, 0x52, 0x43, 0x20, 0x20, 0x00, 0x02, 0x0b,
	0x00, 0x43, 0x61, 0x6c, 0x6c, 0x20, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x20, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x00, 0x02,
	0x00, 0x00, 0x01, 0x06, 0x00, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x05, 0x0a,
	0x80, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x02, 0x40,
	0x00, 0x00, 0x0c, 0x80, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x00, 0x02, 0x90, 0x75, 0x32, 0x0e, 0x00, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x51, 0x6f, 0x53, 0x20, 0x46, 0x6c, 0x6f,
	0x77, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x73, 0x65, 0x74, 0x75, 0x70, 0x00, 0x90, 0x75, 0x33, 0x0d, 0x00, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c,
	0x6f, 0x77, 0x20, 0x73, 0x65, 0x74, 0x75, 0x70, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x00, 0x90, 0x75, 0x35, 0x03,
	0x80, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x00, 0x40, 0x00, 0x01, 0x10, 0x00, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x00, 0x02, 0x90, 0x79,
	0x1a, 0x10, 0x80, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x51, 0x6f, 0x53, 0x20, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64,
	0x64, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x00, 0x90, 0x79, 0x1b, 0x11, 0x00, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x20,
	0x61, 0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x74, 0x65, 0x6d,
	0x00, 0x90, 0x79, 0x1d, 0x03, 0x80, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x00, 0x00, 0x00, 0x02, 0x0d, 0x80, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x02, 0x80, 0x00, 0x02,
	0x0b, 0x00, 0x43, 0x61, 0x6c, 0x6c, 0x20, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x20, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x00,
	0x02, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0x00, 0x00, 0x80, 0x00, 0x05, 0x80, 0x44, 0x52, 0x42, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x00,
	0x40, 0x00, 0x01, 0x05, 0x80, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x43, 0x6f, 0x70, 0x79, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
	0x00, 0x80, 0x02, 0x05, 0x00, 0x52, 0x52, 0x43, 0x20, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0x00, 0x80, 0x00, 0x01, 0x0d, 0x80, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x00,
	0x02, 0x00, 0x03, 0x00, 0x00, 0x40, 0x00, 0x03, 0x0f, 0x00, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x00, 0x00, 0x80, 0x00, 0x05, 0x80, 0x44, 0x52, 0x42, 0x20, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x00, 0x00, 0x02, 0x00, 0x05, 0x00, 0x01, 0x02, 0xc0, 0x00, 0x01, 0x09, 0x80, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x03, 0x0b, 0x00, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x20, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	0x60, 0x26, 0xad, 0x10, 0x80, 0x48, 0x4f, 0x20, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x80, 0x48, 0x4f, 0x20, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01}

func TestE2SMRcRanFunctionDefinitionRadisys(t *testing.T) {
	aper, err := hexlib.Asn1BytesToByte(rSysBytes1)
	assert.NilError(t, err)
	result, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-RC-RanFunctionDefinition message is\n%v\n", result)
}
