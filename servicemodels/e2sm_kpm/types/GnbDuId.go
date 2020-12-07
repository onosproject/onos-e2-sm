// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbDuID struct {
	Value int64
}

type GnbDuIDBuilder interface {
	NewGnbDuID(value int64)
	GetValue()
	GetGnbDuID()
}

func NewGnbDuID(value int64) *GnbDuID {
	return &GnbDuID{
		Value: value,
	}
}

func (b GnbDuID) GetValue() int64 {
	return b.Value
}

func (b *GnbDuID) GetGnbDuID() *GnbDuID {
	return b
}
