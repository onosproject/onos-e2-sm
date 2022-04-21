//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRcRanfunctionDefinition(t *testing.T) {

	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription1 := "empty E2SM-RC"

	msg1, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription1)
	assert.NilError(t, err)
	assert.Assert(t, msg1 != nil)
	msg1.GetRanFunctionName().SetRanFunctionInstance(1)

	aper1, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg1)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper1))

	result1, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper1)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result1)
	assert.Equal(t, msg1.String(), result1.String())
}

func Test_E2SmRcRanfunctionDefinitionEventTrigger(t *testing.T) {
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

func Test_E2SmRcRanfunctionDefinitionReport(t *testing.T) {
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

func Test_E2SmRcRanfunctionDefinitionInsert(t *testing.T) {
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

func Test_E2SmRcRanfunctionDefinitionControl(t *testing.T) {
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

func Test_E2SmRcRanfunctionDefinitionPolicy(t *testing.T) {
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
