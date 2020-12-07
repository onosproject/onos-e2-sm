// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbDuID struct {
	Value int64
}

type GnbDuIDBuilder interface {
	NewGnbDuID()
	SetValue(value int64)
	GetValue()
	GetGnbDuID()
}

func NewGnbDuID() *GnbDuID {
	return &GnbDuID{}
}

func (b *GnbDuID) SetValue(value int64) *GnbDuID {
	b.Value = value
	return b
}

func (b GnbDuID) GetValue() int64 {
	return b.Value
}

func (b *GnbDuID) GetGnbDuID() *GnbDuID {
	return b
}