// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RicStyleName struct {
	Value string
}

type RicStyleNameBuilder interface {
	NewRicStyleName(value string)
	GetValue()
	GetRicStyleName()
}

func NewRicStyleName(value string) *RicStyleName {
	return &RicStyleName{
		Value: value,
	}
}

func (b *RicStyleName) GetValue() string {
	return b.Value
}

func (b *RicStyleName) GetRicStyleName() *RicStyleName {
	return b
}
