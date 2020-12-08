// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type PlmnIdentity struct {
	Value []byte
}

type PlmnIdentityBuilder interface {
	NewPlmnID(value []byte)
	GetValue()
	GetPlmnID()
}

func NewPlmnID(value []byte) *PlmnIdentity {
	return &PlmnIdentity{
		Value: value,
	}
}

func (b *PlmnIdentity) GetValue() []byte {
	return b.Value
}

func (b *PlmnIdentity) GetPlmnID() *PlmnIdentity {
	return b
}
