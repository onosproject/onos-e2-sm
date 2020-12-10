// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RicEventTriggerStyleList struct {
	RicEventTriggerStyleType  *RicStyleType
	RicEventTriggerStyleName  *RicStyleName
	RicEventTriggerFormatType *RicFormatType
}

type RicEventTriggerStyleListBuilder interface {
	NewRicEventTriggerStyleList()
	SetRicEventTriggerStyleType(ricStyleType *RicStyleType)
	SetRicEventTriggerStyleName(ricStyleName *RicStyleName)
	SetRicEventTriggerFormatType(ricFormatType *RicFormatType)
	GetRicEventTriggerStyleType()
	GetRicEventTriggerStyleName()
	GetRicEventTriggerFormatType()
	GetRicEventTriggerStyleList()
}

func NewRicEventTriggerStyleList() *RicEventTriggerStyleList {
	return &RicEventTriggerStyleList{}
}

func (b *RicEventTriggerStyleList) SetRicEventTriggerStyleType(ricStyleType *RicStyleType) *RicEventTriggerStyleList {
	b.RicEventTriggerStyleType = ricStyleType
	return b
}

func (b *RicEventTriggerStyleList) SetRicEventTriggerStyleName(ricStyleName *RicStyleName) *RicEventTriggerStyleList {
	b.RicEventTriggerStyleName = ricStyleName
	return b
}

func (b *RicEventTriggerStyleList) SetRicEventTriggerFormatType(ricFormatType *RicFormatType) *RicEventTriggerStyleList {
	b.RicEventTriggerFormatType = ricFormatType
	return b
}

func (b *RicEventTriggerStyleList) GetRicEventTriggerStyleType() *RicStyleType {
	return b.RicEventTriggerStyleType
}

func (b *RicEventTriggerStyleList) GetRicEventTriggerStyleName() *RicStyleName {
	return b.RicEventTriggerStyleName
}

func (b *RicEventTriggerStyleList) GetRicEventTriggerFormatType() *RicFormatType {
	return b.RicEventTriggerFormatType
}

func (b *RicEventTriggerStyleList) GetRicEventTriggerStyleList() *RicEventTriggerStyleList {
	return b
}
