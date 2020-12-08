// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RanContainer struct {
	Value []byte
}

type RanContainerBuilder interface {
	NewRanContainer(value []byte)
	GetValue()
	GetRanContainer()
}

func NewRanContainer(value []byte) *RanContainer {
	return &RanContainer{
		Value: value,
	}
}

func (b *RanContainer) GetValue() []byte {
	return b.Value
}

func (b *RanContainer) GetRanContainer() *RanContainer {
	return b
}
