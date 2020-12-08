// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbCuCpName struct {
	Value string
}

type GnbCuCpNameBuilder interface {
	NewGnbCuCpName(name string)
	GetValue()
	GetGnbCuCpName()
}

func NewGnbCuCpName(name string) *GnbCuCpName {
	return &GnbCuCpName{
		Value: name,
	}
}

func (b *GnbCuCpName) GetValue() string {
	return b.Value
}

func (b *GnbCuCpName) GetGnbCuCpName() *GnbCuCpName {
	return b
}
