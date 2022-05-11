// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smrcv1

import (
e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)


func (m *EventTriggerCellInfoItem) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *EventTriggerCellInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *EventTriggerUeInfoItem) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *EventTriggerUeInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *EventTriggerUeeventInfoItem) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *EventTriggerUeeventInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *RanparameterDefinitionChoiceListItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *RanparameterDefinitionChoiceListItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanparameterDefinitionChoiceStructureItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *RanparameterDefinitionChoiceStructureItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanparameterValueTypeChoiceElementFalse) SetRanParameterValue(ranParameterValue *e2smrcv1.RanparameterValue) *RanparameterValueTypeChoiceElementFalse {
	m.RanParameterValue = ranParameterValue
	return m
}

func (m *RanparameterTestingItemChoiceElementFalse) SetRanParameterValue(ranParameterValue *e2smrcv1.RanparameterValue) *RanparameterTestingItemChoiceElementFalse {
	m.RanParameterValue = ranParameterValue
	return m
}

func (m *RanparameterTestingItemChoiceElementFalse) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *RanparameterTestingItemChoiceElementFalse {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat1) SetGlobalAssociatedUeinfo(globalAssociatedUeinfo *e2smrcv1.EventTriggerUeInfo) *E2SmRcEventTriggerFormat1 {
	m.GlobalAssociatedUeinfo = globalAssociatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetMessageDirection(messageDirection e2smrcv1.MessageDirection) *E2SmRcEventTriggerFormat1Item {
	m.MessageDirection = &messageDirection
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeinfo(associatedUeinfo *e2smrcv1.EventTriggerUeInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeevent(associatedUeevent *e2smrcv1.EventTriggerUeeventInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeevent = associatedUeevent
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *E2SmRcEventTriggerFormat1Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *MessageTypeChoiceNi) SetNIIDentifier(nIIDentifier *InterfaceIdentifier) *MessageTypeChoiceNi {
	m.NIIdentifier = nIIDentifier
	return m
}

func (m *MessageTypeChoiceNi) SetNIMessage(nIMessage *InterfaceMessageId) *MessageTypeChoiceNi {
	m.NIMessage = nIMessage
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedE2NodeInfo(associatedE2NodeInfo *e2smrcv1.RanparameterTesting) *E2SmRcEventTriggerFormat2 {
	m.AssociatedE2NodeInfo = associatedE2NodeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedUeinfo(associatedUeinfo *e2smrcv1.EventTriggerUeInfo) *E2SmRcEventTriggerFormat2 {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetAssociatedCellInfo(associatedCellInfo *e2smrcv1.EventTriggerCellInfo) *E2SmRcEventTriggerFormat3Item {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *E2SmRcEventTriggerFormat3Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetAssociatedUeinfo(associatedUeinfo *e2smrcv1.EventTriggerUeInfo) *E2SmRcEventTriggerFormat4Item {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *E2SmRcEventTriggerFormat4Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *TriggerTypeChoiceRrcstateItem) SetLogicalOr(logicalOr e2smrcv1.LogicalOr) *TriggerTypeChoiceRrcstateItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedUeinfo(associatedUeinfo *e2smrcv1.EventTriggerUeInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedCellInfo(associatedCellInfo *e2smrcv1.EventTriggerCellInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcActionDefinitionFormat2Item) SetRicPolicyConditionDefinition(ricPolicyConditionDefinition *e2smrcv1.RanparameterTesting) *E2SmRcActionDefinitionFormat2Item {
	m.RicPolicyConditionDefinition = ricPolicyConditionDefinition
	return m
}

func (m *E2SmRcActionDefinitionFormat3) SetUeID(ueID *Ueid) *E2SmRcActionDefinitionFormat3 {
	m.UeId = ueID
	return m
}

func (m *E2SmRcIndicationHeaderFormat1) SetRicEventTriggerConditionID(ricEventTriggerConditionID *e2smrcv1.RicEventTriggerConditionId) *E2SmRcIndicationHeaderFormat1 {
	m.RicEventTriggerConditionId = ricEventTriggerConditionID
	return m
}

func (m *E2SmRcIndicationMessageFormat3Item) SetCellContextInfo(cellContextInfo []byte) *E2SmRcIndicationMessageFormat3Item {
	m.CellContextInfo = cellContextInfo
	return m
}

func (m *E2SmRcIndicationMessageFormat3Item) SetCellDeleted(cellDeleted bool) *E2SmRcIndicationMessageFormat3Item {
	m.CellDeleted = &cellDeleted
	return m
}

func (m *E2SmRcIndicationMessageFormat3Item) SetNeighborRelationTable(neighborRelationTable *e2smrcv1.NeighborRelationInfo) *E2SmRcIndicationMessageFormat3Item {
	m.NeighborRelationTable = neighborRelationTable
	return m
}

func (m *E2SmRcIndicationMessageFormat4ItemUe) SetUeContextInfo(ueContextInfo []byte) *E2SmRcIndicationMessageFormat4ItemUe {
	m.UeContextInfo = ueContextInfo
	return m
}

func (m *E2SmRcIndicationMessageFormat4ItemCell) SetCellContextInfo(cellContextInfo []byte) *E2SmRcIndicationMessageFormat4ItemCell {
	m.CellContextInfo = cellContextInfo
	return m
}

func (m *E2SmRcIndicationMessageFormat4ItemCell) SetNeighborRelationTable(neighborRelationTable *e2smrcv1.NeighborRelationInfo) *E2SmRcIndicationMessageFormat4ItemCell {
	m.NeighborRelationTable = neighborRelationTable
	return m
}

func (m *E2SmRcControlHeaderFormat1) SetRicControlDecision(ricControlDecision e2smrcv1.RicControlDecision) *E2SmRcControlHeaderFormat1 {
	m.RicControlDecision = &ricControlDecision
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionEventTrigger(ranFunctionDefinitionEventTrigger *e2smrcv1.RanfunctionDefinitionEventTrigger) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionEventTrigger = ranFunctionDefinitionEventTrigger
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionReport(ranFunctionDefinitionReport *e2smrcv1.RanfunctionDefinitionReport) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionReport = ranFunctionDefinitionReport
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionInsert(ranFunctionDefinitionInsert *e2smrcv1.RanfunctionDefinitionInsert) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionInsert = ranFunctionDefinitionInsert
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionControl(ranFunctionDefinitionControl *e2smrcv1.RanfunctionDefinitionControl) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionControl = ranFunctionDefinitionControl
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionPolicy(ranFunctionDefinitionPolicy *e2smrcv1.RanfunctionDefinitionPolicy) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionPolicy = ranFunctionDefinitionPolicy
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanL2ParametersList(ranL2ParametersList []*e2smrcv1.L2ParametersRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanL2ParametersList = ranL2ParametersList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanCallProcessTypesList(ranCallProcessTypesList []*e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem) *RanfunctionDefinitionEventTrigger {
	m.RanCallProcessTypesList = ranCallProcessTypesList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanUeIDentificationParametersList(ranUeIDentificationParametersList []*e2smrcv1.UeidentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanUeidentificationParametersList = ranUeIDentificationParametersList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanCellIDentificationParametersList(ranCellIDentificationParametersList []*e2smrcv1.CellIdentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanCellIdentificationParametersList = ranCellIDentificationParametersList
	return m
}

func (m *L2ParametersRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *L2ParametersRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *UeidentificationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *UeidentificationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *CellIdentificationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *CellIdentificationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionEventTriggerBreakpointItem) SetRanCallProcessBreakpointParametersList(ranCallProcessBreakpointParametersList []*e2smrcv1.CallProcessBreakpointRanparameterItem) *RanfunctionDefinitionEventTriggerBreakpointItem {
	m.RanCallProcessBreakpointParametersList = ranCallProcessBreakpointParametersList
	return m
}

func (m *CallProcessBreakpointRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *CallProcessBreakpointRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionReportItem) SetRanReportParametersList(ranReportParametersList []*e2smrcv1.ReportRanparameterItem) *RanfunctionDefinitionReportItem {
	m.RanReportParametersList = ranReportParametersList
	return m
}

func (m *ReportRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *ReportRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionInsertItem) SetRicInsertIndicationList(ricInsertIndicationList []*e2smrcv1.RanfunctionDefinitionInsertIndicationItem) *RanfunctionDefinitionInsertItem {
	m.RicInsertIndicationList = ricInsertIndicationList
	return m
}

func (m *RanfunctionDefinitionInsertIndicationItem) SetRanInsertIndicationParametersList(ranInsertIndicationParametersList []*e2smrcv1.InsertIndicationRanparameterItem) *RanfunctionDefinitionInsertIndicationItem {
	m.RanInsertIndicationParametersList = ranInsertIndicationParametersList
	return m
}

func (m *InsertIndicationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *InsertIndicationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRicControlActionList(ricControlActionList []*e2smrcv1.RanfunctionDefinitionControlActionItem) *RanfunctionDefinitionControlItem {
	m.RicControlActionList = ricControlActionList
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRicCallProcessIDformatType(ricCallProcessIDformatType *RicFormatType) *RanfunctionDefinitionControlItem {
	m.RicCallProcessIdformatType = ricCallProcessIDformatType
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRanControlOutcomeParametersList(ranControlOutcomeParametersList []*e2smrcv1.ControlOutcomeRanparameterItem) *RanfunctionDefinitionControlItem {
	m.RanControlOutcomeParametersList = ranControlOutcomeParametersList
	return m
}

func (m *ControlOutcomeRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *ControlOutcomeRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionControlActionItem) SetRanControlActionParametersList(ranControlActionParametersList []*e2smrcv1.ControlActionRanparameterItem) *RanfunctionDefinitionControlActionItem {
	m.RanControlActionParametersList = ranControlActionParametersList
	return m
}

func (m *ControlActionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *ControlActionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionPolicyItem) SetRicPolicyActionList(ricPolicyActionList []*e2smrcv1.RanfunctionDefinitionPolicyActionItem) *RanfunctionDefinitionPolicyItem {
	m.RicPolicyActionList = ricPolicyActionList
	return m
}

func (m *RanfunctionDefinitionPolicyActionItem) SetRanPolicyActionParametersList(ranPolicyActionParametersList []*e2smrcv1.PolicyActionRanparameterItem) *RanfunctionDefinitionPolicyActionItem {
	m.RanPolicyActionParametersList = ranPolicyActionParametersList
	return m
}

func (m *RanfunctionDefinitionPolicyActionItem) SetRanPolicyConditionParametersList(ranPolicyConditionParametersList []*e2smrcv1.PolicyConditionRanparameterItem) *RanfunctionDefinitionPolicyActionItem {
	m.RanPolicyConditionParametersList = ranPolicyConditionParametersList
	return m
}

func (m *PolicyActionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *PolicyActionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *PolicyConditionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *e2smrcv1.RanparameterDefinition) *PolicyConditionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}
