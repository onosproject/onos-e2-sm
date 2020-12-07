// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type BitString struct {
	Value uint64
	Len   uint32
}

type BitStringBuilder interface {
	NewBitString()
	SetValue(value uint64)
	SetLen(len uint32)
	GetValue()
	GetLen()
	GetBitString()
}

func NewBitString() *BitString {
	return &BitString{}
}

func (b *BitString) SetValue(value uint64) *BitString {
	b.Value = value
	return b
}

func (b *BitString) SetLen(len uint32) *BitString {
	b.Len = len
	return b
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
