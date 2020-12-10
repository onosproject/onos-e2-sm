// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RicReportStyleList struct {
	RicReportStyleType             *RicStyleType
	RicReportStyleName             *RicStyleName
	RicIndicationHeaderFormatType  *RicFormatType
	RicIndicationMessageFormatType *RicFormatType
}

type RicReportStyleListBuilder interface {
	NewRicReportStyleList()
	SetRicReportStyleType(ricStyleType *RicStyleType)
	SetRicReportStyleName(ricStyleName *RicStyleName)
	SetRicIndicationHeaderFormatType(indHdrFormatType *RicFormatType)
	SetRicIndicationMessageFormatType(indMsgFormatType *RicFormatType)
	GetRicReportStyleType()
	GetRicReportStyleName()
	GetRicIndicationHeaderFormatType()
	GetRicIndicationMessageFormatType()
	GetRicReportStyleList()
}

func NewRicReportStyleList() *RicReportStyleList {
	return &RicReportStyleList{}
}

func (b *RicReportStyleList) SetRicReportStyleType(ricStyleType *RicStyleType) *RicReportStyleList {
	b.RicReportStyleType = ricStyleType
	return b
}

func (b *RicReportStyleList) SetRicReportStyleName(ricStyleName *RicStyleName) *RicReportStyleList {
	b.RicReportStyleName = ricStyleName
	return b
}

func (b *RicReportStyleList) SetRicIndicationHeaderFormatType(indHdrFormatType *RicFormatType) *RicReportStyleList {
	b.RicIndicationHeaderFormatType = indHdrFormatType
	return b
}

func (b *RicReportStyleList) SetRicIndicationMessageFormatType(indMsgFormatType *RicFormatType) *RicReportStyleList {
	b.RicIndicationMessageFormatType = indMsgFormatType
	return b
}

func (b *RicReportStyleList) GetRicReportStyleType() *RicStyleType {
	return b.RicReportStyleType
}

func (b *RicReportStyleList) GetRicReportStyleName() *RicStyleName {
	return b.RicReportStyleName
}

func (b *RicReportStyleList) GetRicIndicationHeaderFormatType() *RicFormatType {
	return b.RicIndicationHeaderFormatType
}

func (b *RicReportStyleList) GetRicIndicationMessageFormatType() *RicFormatType {
	return b.RicIndicationMessageFormatType
}

func (b *RicReportStyleList) GetRicReportStyleList() *RicReportStyleList {
	return b
}
