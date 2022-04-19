//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package e2smrcies

func (m *E2SmRcRanfunctionDefinition) SetRanfunctionDefinitionEventTrigger(ranFunctionDefinitionEventTrigger *RanfunctionDefinitionEventTrigger) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionEventTrigger = ranFunctionDefinitionEventTrigger
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanfunctionDefinitionReport(ranFunctionDefinitionReport *RanfunctionDefinitionReport) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionReport = ranFunctionDefinitionReport
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanfunctionDefinitionInsert(ranFunctionDefinitionInsert *RanfunctionDefinitionInsert) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionInsert = ranFunctionDefinitionInsert
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanfunctionDefinitionControl(ranFunctionDefinitionControl *RanfunctionDefinitionControl) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionControl = ranFunctionDefinitionControl
	return m
}

func (m *E2SmRcRanfunctionDefinition) SetRanfunctionDefinitionPolicy(ranFunctionDefinitionPolicy *RanfunctionDefinitionPolicy) *E2SmRcRanfunctionDefinition {
	m.RanFunctionDefinitionPolicy = ranFunctionDefinitionPolicy
	return m
}
