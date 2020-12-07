// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbIDChoice struct {
	ID *BitString
}

type GnbIDChoiceBuilder interface {
	NewGnbIDChoice(id *BitString)
	GetID()
	GetGnbIDChoice()
}

func NewGnbIDChoice(id *BitString) *GnbIDChoice {
	return &GnbIDChoice{
		ID: id,
	}
}

func (b *GnbIDChoice) GetID() *BitString {
	return b.ID
}

func (b *GnbIDChoice) GetGnbIDChoice() *GnbIDChoice {
	return b
}
