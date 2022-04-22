//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package e2smrcies

import e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"

func (m *EventTriggerCellInfoItem) SetLogicalOr(lor LogicalOr) *EventTriggerCellInfoItem {
	m.LogicalOr = &lor
	return m
}

func (m *EventTriggerUeInfoItem) SetLogicalOr(lor LogicalOr) *EventTriggerUeInfoItem {
	m.LogicalOr = &lor
	return m
}

func (m *EventTriggerUeeventInfoItem) SetLogicalOr(lor LogicalOr) *EventTriggerUeeventInfoItem {
	m.LogicalOr = &lor
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

func (m *RanparameterTestingItemChoiceElementFalse) SetRanParameterValue(ranparameterValue *RanparameterValue) *RanparameterTestingItemChoiceElementFalse {
	m.RanParameterValue = ranparameterValue
	return m
}

func (m *RanparameterTestingItemChoiceElementFalse) SetLogicalOr(lor LogicalOr) *RanparameterTestingItemChoiceElementFalse {
	m.LogicalOr = &lor
	return m
}

func (m *E2SmRcEventTriggerFormat1) SetGlobalAssociatedUeinfo(eventTriggerUeInfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat1 {
	m.GlobalAssociatedUeinfo = eventTriggerUeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetMessageDirection(messageDirection MessageDirection) *E2SmRcEventTriggerFormat1Item {
	m.MessageDirection = &messageDirection
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeinfo(eventTriggerUeInfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeinfo = eventTriggerUeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetAssociatedUeevent(eventTriggerUeeventInfo *EventTriggerUeeventInfo) *E2SmRcEventTriggerFormat1Item {
	m.AssociatedUeevent = eventTriggerUeeventInfo
	return m
}

func (m *E2SmRcEventTriggerFormat1Item) SetLogicalOr(lor LogicalOr) *E2SmRcEventTriggerFormat1Item {
	m.LogicalOr = &lor
	return m
}

func (m *MessageTypeChoiceNi) SetNIIdentifier(interfaceIdentifier *e2smcommonies.InterfaceIdentifier) *MessageTypeChoiceNi {
	m.NIIdentifier = interfaceIdentifier
	return m
}

func (m *MessageTypeChoiceNi) SetNIMessage(interfaceMessageID *e2smcommonies.InterfaceMessageId) *MessageTypeChoiceNi {
	m.NIMessage = interfaceMessageID
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedE2NodeInfo(associatedE2NodeInfo *RanparameterTesting) *E2SmRcEventTriggerFormat2 {
	m.AssociatedE2NodeInfo = associatedE2NodeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat2) SetAssociatedUeinfo(associatedUeInfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat2 {
	m.AssociatedUeinfo = associatedUeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetAssociatedE2NodeInfo(associatedCellInfo *EventTriggerCellInfo) *E2SmRcEventTriggerFormat3Item {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcEventTriggerFormat3Item) SetLogicalOr(lor LogicalOr) *E2SmRcEventTriggerFormat3Item {
	m.LogicalOr = &lor
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetAssociatedUeinfo(associatedUeInfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat4Item {
	m.AssociatedUeinfo = associatedUeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat4Item) SetLogicalOr(lor LogicalOr) *E2SmRcEventTriggerFormat4Item {
	m.LogicalOr = &lor
	return m
}

func (m *TriggerTypeChoiceRrcstateItem) SetLogicalOr(lor LogicalOr) *TriggerTypeChoiceRrcstateItem {
	m.LogicalOr = &lor
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedUeinfo(associatedUeInfo *EventTriggerUeInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedUeinfo = associatedUeInfo
	return m
}

func (m *E2SmRcEventTriggerFormat5) SetAssociatedCellInfo(associatedCellInfo *EventTriggerCellInfo) *E2SmRcEventTriggerFormat5 {
	m.AssociatedCellInfo = associatedCellInfo
	return m
}

func (m *E2SmRcActionDefinitionFormat2Item) SetAssociatedCellInfo(ricPolicyConditionDefinition *RanparameterTesting) *E2SmRcActionDefinitionFormat2Item {
	m.RicPolicyConditionDefinition = ricPolicyConditionDefinition
	return m
}

func (m *E2SmRcActionDefinitionFormat3) SetUeID(ueID *e2smcommonies.Ueid) *E2SmRcActionDefinitionFormat3 {
	m.UeId = ueID
	return m
}

func (m *E2SmRcIndicationHeaderFormat1) SetRicEventTriggerConditionID(ricEventConditionID int32) *E2SmRcIndicationHeaderFormat1 {
	m.RicEventTriggerConditionId = &RicEventTriggerConditionId{
		Value: ricEventConditionID,
	}
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

func (m *E2SmRcControlHeaderFormat1) SetRicControlDecision(ricControlDecision *RicControlDecision) *E2SmRcControlHeaderFormat1 {
	m.RicControlDecision = ricControlDecision
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

func (m *RanfunctionDefinitionEventTrigger) SetRanUeidentificationParametersList(ranUeidentificationParametersList []*UeidentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanUeidentificationParametersList = ranUeidentificationParametersList
	return m
}

func (m *RanfunctionDefinitionEventTrigger) SetRanCellIdentificationParametersList(ranCellIdentificationParametersList []*CellIdentificationRanparameterItem) *RanfunctionDefinitionEventTrigger {
	m.RanCellIdentificationParametersList = ranCellIdentificationParametersList
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

func (m *RanfunctionDefinitionControlItem) SetRicCallProcessIDformatType(ricCallProcessIDformatType int32) *RanfunctionDefinitionControlItem {
	m.RicCallProcessIdformatType = &e2smcommonies.RicFormatType{
		Value: ricCallProcessIDformatType,
	}
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
