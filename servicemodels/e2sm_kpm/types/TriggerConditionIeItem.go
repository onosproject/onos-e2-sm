// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type TriggerConditionIeItem struct {
	ReportPeriodIe *RtPeriodIe
}

type TriggerConditionIeItemBuilder interface {
	NewTriggerConditionIeItem(rtPeriodIe *RtPeriodIe)
	GetRtPeriodIe()
	GetTriggerConditionIeItem()
}

func NewTriggerConditionIeItem(rtPeriodIe *RtPeriodIe) *TriggerConditionIeItem {
	return &TriggerConditionIeItem{
		ReportPeriodIe: rtPeriodIe,
	}
}

func (b *TriggerConditionIeItem) GetRtPeriodIe() *RtPeriodIe {
	return b.ReportPeriodIe
}

func (b *TriggerConditionIeItem) GetTriggerConditionIeItem() *TriggerConditionIeItem {
	return b
}
