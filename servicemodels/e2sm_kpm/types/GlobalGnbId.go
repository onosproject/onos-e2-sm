// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GlobalGnbID struct {
	PlmnID *PlmnIdentity
	GnbID  *GnbIDChoice
}

type GlobalGnbIDBuilder interface {
	NewGlobalGnbID()
	SetPlmnID(plmnID PlmnIdentity)
	SetGnbID(gnbID GnbIDChoice)
	GetPlmnID()
	GetGnbID()
	GetGlobalGnbID()
}

func NewGlobalGnbID() *GlobalGnbID {
	return &GlobalGnbID{}
}

func (b *GlobalGnbID) SetPlmnID(plmnID *PlmnIdentity) *GlobalGnbID {
	b.PlmnID = plmnID
	return b
}

func (b *GlobalGnbID) SetGnbID(gnbID *GnbIDChoice) *GlobalGnbID {
	b.GnbID = gnbID
	return b
}

func (b *GlobalGnbID) GetPlmnID() *PlmnIdentity {
	return b.PlmnID
}

func (b *GlobalGnbID) GetGnbID() *GnbIDChoice {
	return b.GnbID
}

func (b *GlobalGnbID) GetGlobalGnbID() *GlobalGnbID {
	return b
}
