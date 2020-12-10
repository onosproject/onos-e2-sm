// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RicFormatType struct {
	Value int32
}

type RicFormatTypeBuilder interface {
	NewRicFormatType(value int32)
	GetValue()
	GetRicFormatType()
}

func NewRicFormatType(value int32) *RicFormatType {
	return &RicFormatType{
		Value: value,
	}
}

func (b *RicFormatType) GetValue() int32 {
	return b.Value
}

func (b *RicFormatType) GetRicFormatType() *RicFormatType {
	return b
}
