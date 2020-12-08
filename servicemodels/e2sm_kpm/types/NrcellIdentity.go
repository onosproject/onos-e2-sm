// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type NrcellIdentity struct {
	Value *BitString
}

type NrcellIdentityBuilder interface {
	NewNrcellIdentity(value *BitString)
	GetValue()
	GetNrcellIdentity()
}

func NewNrcellIdentity(value *BitString) *NrcellIdentity {
	return &NrcellIdentity{
		Value: value,
	}
}

func (b *NrcellIdentity) GetValue() *BitString {
	return b.Value
}

func (b *NrcellIdentity) GetNrcellIdentity() *NrcellIdentity {
	return b
}
