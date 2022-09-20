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

var rsysBytes = "7805804f52414e2d4532534d2d5243000018312e332e362e312e342e312e35333134382e312e312e322e330080524330000004194532204e6f6465" +
	"20496e666f726d6174696f6e204368616e676500010000000000068043656c6c20476c6f62616c20494400004000000a8052524320636f6e6e656374696f6e20726" +
	"56c65617365000080000b0052656469726563746564204361727269657220496e666f0000048000068043656c6c20476c6f62616c20494400800102804e52204347" +
	"490080020600504c4d4e204964656e7469747900800307804e522043656c6c204964656e7469747900801d09004578706c69636974205545204c697374204944000" +
	"280000414554520496e666f726d6174696f6e204368616e67650004000100010002000480c8080043757272656e74205252432053746174650080c90980525243205" +
	"374617465204368616e67656420546f0080ca0500525243204d6573736167650090012b04004f6c642055452049440090012c060043757272656e74205545204944" +
	"00400003194532204e6f646520496e666f726d6174696f6e204368616e67650003000100010003000080000b8043656c6c20436f6e7465787420496e666f726d61746" +
	"96f6e00008000041b526164696f2041636365737320436f6e74726f6c20526571756573740002000300004000030a8052524320636f6e6e656374696f6e2072656c6" +
	"5617365000180000b0052656469726563746564204361727269657220496e666f00800300804e520000020005000100c0000413526164696f2061636365737320636f" +
	"6e74726f6c00004000030e8052524320436f6e6e656374696f6e2052656c6561736520636f6e74726f6c000180000b005265646972656374656420436172726965722" +
	"0496e666f00800300804e52000001000100010001"

func TestRadisysBytes(t *testing.T) {
	aper, err := hexlib.Asn1BytesToByte(rsysBytes)
	assert.NilError(t, err)

	result, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestReferenceMsg(t *testing.T) {
	ranFunctionShortName := "ORAN-E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription2 := "RC"

	eventTriggerList := &e2smrcv1.RanfunctionDefinitionEventTrigger{
		RicEventTriggerStyleList:          make([]*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, 0),
		RanCallProcessTypesList:           make([]*e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem, 0),
		RanUeidentificationParametersList: make([]*e2smrcv1.UeidentificationRanparameterItem, 0),
	}

	ricEventTriggerStyleItem, err := CreateRanfunctionDefinitionEventTriggerStyleItem(4, "E2 Node Information Change", 1)
	assert.NilError(t, err)
	eventTriggerList.RicEventTriggerStyleList = append(eventTriggerList.RicEventTriggerStyleList, ricEventTriggerStyleItem)

	// adding optional CallProcessType
	ranCallProcessBreakpointsList := make([]*e2smrcv1.CallProcessBreakpointRanparameterItem, 0)
	ranParameterDefinition, err := CreateCallProcessBreakpointRanparameterItem(1, "Redirected Carrier Info")
	assert.NilError(t, err)
	ranCallProcessBreakpointsList = append(ranCallProcessBreakpointsList, ranParameterDefinition)

	callProcessBreakpointsList := make([]*e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem, 0)
	brItem, err := CreateRanfunctionDefinitionEventTriggerBreakpointItem(1, "RRC connection release")
	assert.NilError(t, err)
	brItem.SetRanCallProcessBreakpointParametersList(ranCallProcessBreakpointsList)
	callProcessBreakpointsList = append(callProcessBreakpointsList, brItem)

	callProcessTypeItem, err := CreateRanfunctionDefinitionEventTriggerCallProcessItem(1, "Cell Global ID", callProcessBreakpointsList)
	assert.NilError(t, err)
	eventTriggerList.RanCallProcessTypesList = append(eventTriggerList.RanCallProcessTypesList, callProcessTypeItem)

	// adding optional UEIdentificationParameters
	ueIdentificationItem1, err := CreateUeidentificationRanparameterItem(1, "Cell Global ID")
	assert.NilError(t, err)
	eventTriggerList.RanUeidentificationParametersList = append(eventTriggerList.RanUeidentificationParametersList, ueIdentificationItem1)

	ueIdentificationItem2, err := CreateUeidentificationRanparameterItem(2, "NR CGI")
	assert.NilError(t, err)
	eventTriggerList.RanUeidentificationParametersList = append(eventTriggerList.RanUeidentificationParametersList, ueIdentificationItem2)

	ueIdentificationItem3, err := CreateUeidentificationRanparameterItem(3, "PLMN Identity")
	assert.NilError(t, err)
	eventTriggerList.RanUeidentificationParametersList = append(eventTriggerList.RanUeidentificationParametersList, ueIdentificationItem3)

	ueIdentificationItem4, err := CreateUeidentificationRanparameterItem(4, "NR Cell Identity")
	assert.NilError(t, err)
	eventTriggerList.RanUeidentificationParametersList = append(eventTriggerList.RanUeidentificationParametersList, ueIdentificationItem4)

	ueIdentificationItem5, err := CreateUeidentificationRanparameterItem(30, "Explicit UE List ID")
	assert.NilError(t, err)
	eventTriggerList.RanUeidentificationParametersList = append(eventTriggerList.RanUeidentificationParametersList, ueIdentificationItem5)

	// adding report part
	ranFunctionDefinitionReport := &e2smrcv1.RanfunctionDefinitionReport{
		RicReportStyleList: make([]*e2smrcv1.RanfunctionDefinitionReportItem, 0),
	}

	// adding RanReportParametersList
	reportParametersList := make([]*e2smrcv1.ReportRanparameterItem, 0)
	item1, err := CreateReportRanparameterItem(201, "Current RRC State")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item1)

	item2, err := CreateReportRanparameterItem(202, "RRC State Changed To")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item2)

	item3, err := CreateReportRanparameterItem(203, "RRC Message")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item3)

	item4, err := CreateReportRanparameterItem(300, "Old UE ID")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item4)

	item5, err := CreateReportRanparameterItem(301, "Current UE ID")
	assert.NilError(t, err)
	reportParametersList = append(reportParametersList, item5)

	reportItem1, err := CreateRanfunctionDefinitionReportItem(4, "UE Information Change", 4, 1, 1, 2)
	assert.NilError(t, err)
	reportItem1.SetRanReportParametersList(reportParametersList)
	ranFunctionDefinitionReport.RicReportStyleList = append(ranFunctionDefinitionReport.RicReportStyleList, reportItem1)

	reportParametersList2 := make([]*e2smrcv1.ReportRanparameterItem, 0)
	item21, err := CreateReportRanparameterItem(1, "Cell Context Information")
	assert.NilError(t, err)
	reportParametersList2 = append(reportParametersList2, item21)

	reportItem2, err := CreateRanfunctionDefinitionReportItem(3, "E2 Node Information Change", 3, 1, 1, 3)
	assert.NilError(t, err)
	reportItem2.SetRanReportParametersList(reportParametersList2)
	ranFunctionDefinitionReport.RicReportStyleList = append(ranFunctionDefinitionReport.RicReportStyleList, reportItem2)

	// adding RanFunctionDefinitionInsert
	ranFunctionDefinitionInsert := &e2smrcv1.RanfunctionDefinitionInsert{
		RicInsertStyleList: make([]*e2smrcv1.RanfunctionDefinitionInsertItem, 0),
	}

	// adding RanReportParametersList
	insertIndicationList := make([]*e2smrcv1.RanfunctionDefinitionInsertIndicationItem, 0)
	item, err := CreateRanfunctionDefinitionInsertIndicationItem(4, "RRC connection release")
	assert.NilError(t, err)

	insertIndicationItemList := make([]*e2smrcv1.InsertIndicationRanparameterItem, 0)
	insertIndicationItem1, err := CreateInsertIndicationRanparameterItem(1, "Redirected Carrier Info")
	assert.NilError(t, err)
	insertIndicationItemList = append(insertIndicationItemList, insertIndicationItem1)

	insertIndicationItem2, err := CreateInsertIndicationRanparameterItem(4, "NR")
	assert.NilError(t, err)
	insertIndicationItemList = append(insertIndicationItemList, insertIndicationItem2)

	item.SetRanInsertIndicationParametersList(insertIndicationItemList)
	insertIndicationList = append(insertIndicationList, item)

	insertItem, err := CreateRanfunctionDefinitionInsertItem(4, "Radio Access Control Request", 2, 3, 2, 5, 1)
	assert.NilError(t, err)
	insertItem.SetRicInsertIndicationList(insertIndicationList)
	ranFunctionDefinitionInsert.RicInsertStyleList = append(ranFunctionDefinitionInsert.RicInsertStyleList, insertItem)

	// adding RanFunctionControl part
	controlItemList := make([]*e2smrcv1.RanfunctionDefinitionControlItem, 0)

	controlActionParameterList := make([]*e2smrcv1.ControlActionRanparameterItem, 0)

	controlActionParameterItem1, err := CreateControlActionRanparameterItem(1, "Redirected Carrier Info")
	assert.NilError(t, err)
	controlActionParameterList = append(controlActionParameterList, controlActionParameterItem1)

	controlActionParameterItem2, err := CreateControlActionRanparameterItem(4, "NR")
	assert.NilError(t, err)
	controlActionParameterList = append(controlActionParameterList, controlActionParameterItem2)

	controlActionList := make([]*e2smrcv1.RanfunctionDefinitionControlActionItem, 0)
	controlActionItem, err := CreateRanfunctionDefinitionControlActionItem(4, "RRC Connection Release control")
	assert.NilError(t, err)
	controlActionItem.SetRanControlActionParametersList(controlActionParameterList)
	controlActionList = append(controlActionList, controlActionItem)

	controlItem, err := CreateRanfunctionDefinitionControlItem(4, "Radio access control", 1, 1, 1)
	assert.NilError(t, err)
	controlItem.SetRicCallProcessIDformatType(1).SetRicControlActionList(controlActionList)
	controlItemList = append(controlItemList, controlItem)

	ranFunctionDefinitionControl, err := CreateRanfunctionDefinitionControl(controlItemList)
	assert.NilError(t, err)

	// composing main message
	msg, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg != nil)
	msg.SetRanFunctionDefinitionEventTrigger(eventTriggerList).SetRanFunctionDefinitionReport(ranFunctionDefinitionReport).
		SetRanFunctionDefinitionInsert(ranFunctionDefinitionInsert).SetRanFunctionDefinitionControl(ranFunctionDefinitionControl)

	aper, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}
