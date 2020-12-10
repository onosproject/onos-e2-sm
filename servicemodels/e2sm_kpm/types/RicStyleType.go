// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RicStyleType struct {
	Value int32
}

type RicStyleTypeBuilder interface {
	NewRicStyleType(value int32)
	GetValue()
	GetRicStyleType()
}

func NewRicStyleType(value int32) *RicStyleType {
	return &RicStyleType{
		Value: value,
	}
}

func (b *RicStyleType) GetValue() int32 {
	return b.Value
}

func (b *RicStyleType) GetRicStyleType() *RicStyleType {
	return b
}
