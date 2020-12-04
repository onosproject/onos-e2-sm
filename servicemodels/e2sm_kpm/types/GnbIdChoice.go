// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GnbIDChoice struct {
	ID *BitString
	//Type string
}

type GnbIDChoiceBuilder interface {
	NewGnbIdChoice()
	SetId(id BitString)
	SetTypeGnb()
	SetTypeNgGNb()
	GetId()
	GetType()
	GetGnbIdChoice()
}

func NewGnbIDChoice() *GnbIDChoice {
	return &GnbIDChoice{}
}

func (b *GnbIDChoice) SetID(id BitString) {
	b.ID = &id
}

//func (b *GnbIdChoice) SetTypeGNb() {
//	b.Type = "GNb"
//}
//
//func (b *GnbIdChoice) SetTypeNgGNb() {
//	b.Type = "NgGNb"
//}

func (b GnbIDChoice) GetID() BitString {
	return *b.ID
}

//func (b GnbIdChoice) GetType() string {
//	return b.Type
//}

func (b GnbIDChoice) GnbIDChoice() GnbIDChoice {
	return GnbIDChoice{
		ID: b.ID,
		//Type: b.Type,
	}
}
