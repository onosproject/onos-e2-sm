// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmRanfunctionDescription struct {
	RanFunctionName        *RanfunctionName
	E2SmKpmRanfunctionItem *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001
}

type E2SmKpmRanfunctionDescriptionBuilder interface {
	NewE2SmKpmRanfunctionDescription()
	SetRanfunctionName(ranfunctionName *RanfunctionName)
	SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001)
	GetRanfunctionName()
	GetE2SmKpmRanfunctionItem()
	GetE2SmKpmRanfunctionDescription()
}

func NewE2SmKpmRanfunctionDescription() *E2SmKpmRanfunctionDescription {
	return &E2SmKpmRanfunctionDescription{}
}

func (b *E2SmKpmRanfunctionDescription) SetRanfunctionName(ranfunctionName *RanfunctionName) *E2SmKpmRanfunctionDescription {
	b.RanFunctionName = ranfunctionName
	return b
}

func (b *E2SmKpmRanfunctionDescription) SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) *E2SmKpmRanfunctionDescription {
	b.E2SmKpmRanfunctionItem = e2SmKpmRanfunctionItem
	return b
}

func (b *E2SmKpmRanfunctionDescription) GetRanfunctionName() *RanfunctionName {
	return b.RanFunctionName
}

func (b *E2SmKpmRanfunctionDescription) GetE2SmKpmRanfunctionItem() *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 {
	return b.E2SmKpmRanfunctionItem
}

func (b *E2SmKpmRanfunctionDescription) GetE2SmKpmRanfunctionDescription() *E2SmKpmRanfunctionDescription {
	return b
}
