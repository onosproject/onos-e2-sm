// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type BitString struct {
	Value uint64
	Len   uint32
}

type BitStringBuilder interface {
	NewBitString(value uint64, len uint32)
	GetValue()
	GetLen()
	GetBitString()
}

func NewBitString(value uint64, len uint32) *BitString {
	return &BitString{
		Value: value,
		Len:   len,
	}
}

func (b BitString) GetValue() uint64 {
	return b.Value
}

func (b BitString) GetLen() uint32 {
	return b.Len
}

func (b *BitString) GetBitString() *BitString {
	return b
}
