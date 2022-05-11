// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smrcv1

import (
e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
)


func (m *EventTriggerCellInfoItem) SetLogicalOr(logicalOr LogicalOr) *EventTriggerCellInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *EventTriggerUeInfoItem) SetLogicalOr(logicalOr LogicalOr) *EventTriggerUeInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *EventTriggerUeeventInfoItem) SetLogicalOr(logicalOr LogicalOr) *EventTriggerUeeventInfoItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *RanparameterDefinitionChoiceListItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *RanparameterDefinitionChoiceListItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanparameterDefinitionChoiceStructureItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *RanparameterDefinitionChoiceStructureItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanparameterValueTypeChoiceElementFalse) SetRanParameterValue(ranParameterValue *RanparameterValue) *RanparameterValueTypeChoiceElementFalse {
	m.RanParameterValue = ranParameterValue
	return m
}

func (m *RanparameterTestingItemChoiceElementFalse) SetRanParameterValue(ranParameterValue *RanparameterValue) *RanparameterTestingItemChoiceElementFalse {
	m.RanParameterValue = ranParameterValue
	return m
}

func (m *RanparameterTestingItemChoiceElementFalse) SetLogicalOr(logicalOr LogicalOr) *RanparameterTestingItemChoiceElementFalse {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat1) SetGlobalAssociatedUeinfo(globalAssociatedUeinfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat1 {
	m.GlobalAssociatedUeinfo = globalAssociatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetMessageDirection(messageDirection MessageDirection) *E2SmRcEventTriggerFormat1Item {
	m.MessageDirection = &messageDirection
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeinfo(associatedUeinfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeevent(associatedUeevent *EventTriggerUeeventInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeevent = associatedUeevent
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetLogicalOr(logicalOr LogicalOr) *E2SmRcEventTriggerFormat1Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *MessageTypeChoiceNi) SetNIIDentifier(nIIDentifier *e2smcommoniesv1.InterfaceIdentifier) *MessageTypeChoiceNi {
	m.NIIdentifier = nIIDentifier
	return m
}

func (m *MessageTypeChoiceNi) SetNIMessage(nIMessage *e2smcommoniesv1.InterfaceMessageId) *MessageTypeChoiceNi {
	m.NIMessage = nIMessage
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedE2NodeInfo(associatedE2NodeInfo *RanparameterTesting) *E2SmRcEventTriggerFormat2 {
	m.AssociatedE2NodeInfo = associatedE2NodeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedUeinfo(associatedUeinfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat2 {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetAssociatedCellInfo(associatedCellInfo *EventTriggerCellInfo) *E2SmRcEventTriggerFormat3Item {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetLogicalOr(logicalOr LogicalOr) *E2SmRcEventTriggerFormat3Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetAssociatedUeinfo(associatedUeinfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat4Item {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetLogicalOr(logicalOr LogicalOr) *E2SmRcEventTriggerFormat4Item {
	m.LogicalOr = &logicalOr
	return m
}

func (m *TriggerTypeChoiceRrcstateItem) SetLogicalOr(logicalOr LogicalOr) *TriggerTypeChoiceRrcstateItem {
	m.LogicalOr = &logicalOr
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedUeinfo(associatedUeinfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedUeinfo = associatedUeinfo
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedCellInfo(associatedCellInfo *EventTriggerCellInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcActionDefinitionFormat2Item) SetRicPolicyConditionDefinition(ricPolicyConditionDefinition *RanparameterTesting) *E2SmRcActionDefinitionFormat2Item {
	m.RicPolicyConditionDefinition = ricPolicyConditionDefinition
	return m
}

func (m *E2SmRcActionDefinitionFormat3) SetUeID(ueID *e2smcommoniesv1.Ueid) *E2SmRcActionDefinitionFormat3 {
	m.UeId = ueID
	return m
}

func (m *E2SmRcIndicationHeaderFormat1) SetRicEventTriggerConditionID(ricEventTriggerConditionID *RicEventTriggerConditionId) *E2SmRcIndicationHeaderFormat1 {
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

func (m *E2SmRcIndicationMessageFormat3Item) SetNeighborRelationTable(neighborRelationTable *NeighborRelationInfo) *E2SmRcIndicationMessageFormat3Item {
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

func (m *E2SmRcIndicationMessageFormat4ItemCell) SetNeighborRelationTable(neighborRelationTable *NeighborRelationInfo) *E2SmRcIndicationMessageFormat4ItemCell {
	m.NeighborRelationTable = neighborRelationTable
	return m
}

func (m *E2SmRcControlHeaderFormat1) SetRicControlDecision(ricControlDecision RicControlDecision) *E2SmRcControlHeaderFormat1 {
	m.RicControlDecision = &ricControlDecision
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionEventTrigger(ranFunctionDefinitionEventTrigger *RanfunctionDefinitionEventTrigger) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionEventTrigger = ranFunctionDefinitionEventTrigger
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionReport(ranFunctionDefinitionReport *RanfunctionDefinitionReport) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionReport = ranFunctionDefinitionReport
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionInsert(ranFunctionDefinitionInsert *RanfunctionDefinitionInsert) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionInsert = ranFunctionDefinitionInsert
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionControl(ranFunctionDefinitionControl *RanfunctionDefinitionControl) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionControl = ranFunctionDefinitionControl
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanFunctionDefinitionPolicy(ranFunctionDefinitionPolicy *RanfunctionDefinitionPolicy) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionPolicy = ranFunctionDefinitionPolicy
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanL2ParametersList(ranL2ParametersList []*L2ParametersRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanL2ParametersList = ranL2ParametersList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanCallProcessTypesList(ranCallProcessTypesList []*RanfunctionDefinitionEventTriggerCallProcessItem) *RanfunctionDefinitionEventTrigger {
	m.RanCallProcessTypesList = ranCallProcessTypesList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanUeIDentificationParametersList(ranUeIDentificationParametersList []*UeidentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanUeidentificationParametersList = ranUeIDentificationParametersList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanCellIDentificationParametersList(ranCellIDentificationParametersList []*CellIdentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanCellIdentificationParametersList = ranCellIDentificationParametersList
	return m
}

func (m *L2ParametersRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *L2ParametersRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *UeidentificationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *UeidentificationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *CellIdentificationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *CellIdentificationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionEventTriggerBreakpointItem) SetRanCallProcessBreakpointParametersList(ranCallProcessBreakpointParametersList []*CallProcessBreakpointRanparameterItem) *RanfunctionDefinitionEventTriggerBreakpointItem {
	m.RanCallProcessBreakpointParametersList = ranCallProcessBreakpointParametersList
	return m
}

func (m *CallProcessBreakpointRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *CallProcessBreakpointRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionReportItem) SetRanReportParametersList(ranReportParametersList []*ReportRanparameterItem) *RanfunctionDefinitionReportItem {
	m.RanReportParametersList = ranReportParametersList
	return m
}

func (m *ReportRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *ReportRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionInsertItem) SetRicInsertIndicationList(ricInsertIndicationList []*RanfunctionDefinitionInsertIndicationItem) *RanfunctionDefinitionInsertItem {
	m.RicInsertIndicationList = ricInsertIndicationList
	return m
}

func (m *RanfunctionDefinitionInsertIndicationItem) SetRanInsertIndicationParametersList(ranInsertIndicationParametersList []*InsertIndicationRanparameterItem) *RanfunctionDefinitionInsertIndicationItem {
	m.RanInsertIndicationParametersList = ranInsertIndicationParametersList
	return m
}

func (m *InsertIndicationRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *InsertIndicationRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRicControlActionList(ricControlActionList []*RanfunctionDefinitionControlActionItem) *RanfunctionDefinitionControlItem {
	m.RicControlActionList = ricControlActionList
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRicCallProcessIDformatType(ricCallProcessIDformatType *e2smcommoniesv1.RicFormatType) *RanfunctionDefinitionControlItem {
	m.RicCallProcessIdformatType = ricCallProcessIDformatType
	return m
}

func (m *RanfunctionDefinitionControlItem) SetRanControlOutcomeParametersList(ranControlOutcomeParametersList []*ControlOutcomeRanparameterItem) *RanfunctionDefinitionControlItem {
	m.RanControlOutcomeParametersList = ranControlOutcomeParametersList
	return m
}

func (m *ControlOutcomeRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *ControlOutcomeRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionControlActionItem) SetRanControlActionParametersList(ranControlActionParametersList []*ControlActionRanparameterItem) *RanfunctionDefinitionControlActionItem {
	m.RanControlActionParametersList = ranControlActionParametersList
	return m
}

func (m *ControlActionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *ControlActionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *RanfunctionDefinitionPolicyItem) SetRicPolicyActionList(ricPolicyActionList []*RanfunctionDefinitionPolicyActionItem) *RanfunctionDefinitionPolicyItem {
	m.RicPolicyActionList = ricPolicyActionList
	return m
}

func (m *RanfunctionDefinitionPolicyActionItem) SetRanPolicyActionParametersList(ranPolicyActionParametersList []*PolicyActionRanparameterItem) *RanfunctionDefinitionPolicyActionItem {
	m.RanPolicyActionParametersList = ranPolicyActionParametersList
	return m
}

func (m *RanfunctionDefinitionPolicyActionItem) SetRanPolicyConditionParametersList(ranPolicyConditionParametersList []*PolicyConditionRanparameterItem) *RanfunctionDefinitionPolicyActionItem {
	m.RanPolicyConditionParametersList = ranPolicyConditionParametersList
	return m
}

func (m *PolicyActionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *PolicyActionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}

func (m *PolicyConditionRanparameterItem) SetRanParameterDefinition(ranParameterDefinition *RanparameterDefinition) *PolicyConditionRanparameterItem {
	m.RanParameterDefinition = ranParameterDefinition
	return m
}
