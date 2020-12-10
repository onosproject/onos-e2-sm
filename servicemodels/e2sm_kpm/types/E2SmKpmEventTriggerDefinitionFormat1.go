// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmEventTriggerDefinitionFormat1 struct {
	PolicyTestList []*TriggerConditionIeItem
}

type E2SmKpmEventTriggerDefinitionFormat1Builder interface {
	NewE2SmKpmEventTriggerDefinitionFormat1()
	AddTriggerConditionIeItem(item *TriggerConditionIeItem)
	RetrieveTriggerConditionIeItemByRtPeriod(ie *RtPeriodIe)
	GetE2SmKpmEventTriggerDefinitionFormat1()
}

func NewE2SmKpmEventTriggerDefinitionFormat1() *E2SmKpmEventTriggerDefinitionFormat1 {
	return &E2SmKpmEventTriggerDefinitionFormat1{}
}

func (b *E2SmKpmEventTriggerDefinitionFormat1) AddTriggerConditionIeItem(item *TriggerConditionIeItem) *E2SmKpmEventTriggerDefinitionFormat1 {
	b.PolicyTestList = append(b.PolicyTestList, item)
	return b
}

func (b *E2SmKpmEventTriggerDefinitionFormat1) RetrieveTriggerConditionIeItemByRtPeriod(ie *RtPeriodIe) *TriggerConditionIeItem {
	for _, tCIe := range b.PolicyTestList {
		if tCIe.ReportPeriodIe.RtPeriodIe == ie.RtPeriodIe {
			return tCIe
		}
	}
	return nil
}

func (b *E2SmKpmEventTriggerDefinitionFormat1) GetE2SmKpmEventTriggerDefinitionFormat1() *E2SmKpmEventTriggerDefinitionFormat1 {
	return b
}
