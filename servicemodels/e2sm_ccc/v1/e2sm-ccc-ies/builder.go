// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smcccv1

import (
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"
)

func (m *PlmnInfo) SetSnssai(snssai *e2smcommoniesv1.SNSsai) *PlmnInfo {
	m.Snssai = snssai
	return m
}

func (m *ConfigurationStructure) SetOldValuesOfAttributes(oldValuesOfAttributes *ValuesOfAttributes) *ConfigurationStructure {
	m.OldValuesOfAttributes = oldValuesOfAttributes
	return m
}

func (m *E2SmCCcRAnfunctionDefinition) SetListOfSupportedNodeLevelConfigurationStructures(listOfSupportedNodeLevelConfigurationStructures *ListOfSupportedRanconfigurationStructures) *E2SmCCcRAnfunctionDefinition {
	m.ListOfSupportedNodeLevelConfigurationStructures = listOfSupportedNodeLevelConfigurationStructures
	return m
}

func (m *E2SmCCcRAnfunctionDefinition) SetListOfCellsForRanfunctionDefinition(listOfCellsForRanfunctionDefinition *ListOfCellsForRanfunctionDefinition) *E2SmCCcRAnfunctionDefinition {
	m.ListOfCellsForRanfunctionDefinition = listOfCellsForRanfunctionDefinition
	return m
}

func (m *RanconfigurationStructure) SetListOfSupportedAttributes(listOfSupportedAttributes *ListOfSupportedAttributes) *RanconfigurationStructure {
	m.ListOfSupportedAttributes = listOfSupportedAttributes
	return m
}

func (m *Ricservices) SetEventTrigger(eventTrigger *EventTrigger) *Ricservices {
	m.EventTrigger = eventTrigger
	return m
}

func (m *Ricservices) SetReportService(reportService *ReportService) *Ricservices {
	m.ReportService = reportService
	return m
}

func (m *Ricservices) SetInsertService(insertService *InsertService) *Ricservices {
	m.InsertService = insertService
	return m
}

func (m *Ricservices) SetControlService(controlService *ControlService) *Ricservices {
	m.ControlService = controlService
	return m
}

func (m *Ricservices) SetPolicyService(policyService *PolicyService) *Ricservices {
	m.PolicyService = policyService
	return m
}

func (m *ControlStyle) SetRicCallProcessIDformatType(ricCallProcessIDformatType *e2smcommoniesv1.RicFormatType) *ControlStyle {
	m.RicCallProcessIdformatType = ricCallProcessIDformatType
	return m
}

func (m *CellForRanfunctionDefinition) SetListOfSupportedCellLevelRanconfigurationStructures(listOfSupportedCellLevelRanconfigurationStructures *ListOfSupportedRanconfigurationStructures) *CellForRanfunctionDefinition {
	m.ListOfSupportedCellLevelRanconfigurationStructures = listOfSupportedCellLevelRanconfigurationStructures
	return m
}

func (m *RanconfigurationStructureForEventTrigger) SetListOfAttributes(listOfAttributes *ListOfAttributes) *RanconfigurationStructureForEventTrigger {
	m.ListOfAttributes = listOfAttributes
	return m
}

func (m *RanconfigurationStructureForAdf) SetListOfAttributes(listOfAttributes *ListOfAttributes) *RanconfigurationStructureForAdf {
	m.ListOfAttributes = listOfAttributes
	return m
}

func (m *CellConfigurationToBeReportedForAdf) SetCellGlobalID(cellGlobalID *CellGlobalId) *CellConfigurationToBeReportedForAdf {
	m.CellGlobalId = cellGlobalID
	return m
}

func (m *CellConfigurationToBeReportedForAdf) SetListOfCellLevelRanconfigurationStructuresForAdf(listOfCellLevelRanconfigurationStructuresForAdf *ListOfRanconfigurationStructuresForAdf) *CellConfigurationToBeReportedForAdf {
	m.ListOfCellLevelRanconfigurationStructuresForAdf = listOfCellLevelRanconfigurationStructuresForAdf
	return m
}
