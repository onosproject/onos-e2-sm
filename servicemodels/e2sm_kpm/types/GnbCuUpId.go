// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbCuUpID struct {
	Value int64
}

type GnbCuUpIDBuilder interface {
	NewGnbCuUpID(value int64)
	GetValue()
	GetGnbCuUpID()
}

func NewGnbCuUpID(value int64) *GnbCuUpID {
	return &GnbCuUpID{
		Value: value,
	}
}

func (b GnbCuUpID) GetValue() int64 {
	return b.Value
}

func (b *GnbCuUpID) GetGnbCuUpID() *GnbCuUpID {
	return b
}
