// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbIDChoice struct {
	ID *BitString
}

type GnbIDChoiceBuilder interface {
	NewGnbIdChoice()
	SetId(id BitString)
	GetId()
	GetGnbIdChoice()
}

func NewGnbIDChoice() *GnbIDChoice {
	return &GnbIDChoice{}
}

func (b *GnbIDChoice) SetID(id *BitString) *GnbIDChoice {
	b.ID = id
	return b
}

func (b *GnbIDChoice) GetID() *BitString {
	return b.ID
}

func (b *GnbIDChoice) GetGnbIDChoice() *GnbIDChoice {
	return b
}
