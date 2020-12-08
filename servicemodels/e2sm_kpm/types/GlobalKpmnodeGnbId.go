// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GlobalKpmnodeGnbID struct {
	GlobalGnbID *GlobalGnbID
	GnbCuUpID   *GnbCuUpID
	GnbDuID     *GnbDuID
}

type GlobalKpmnodeGnbIDBuilder interface {
	NewGlobalKpmnodeGnbID()
	SetGlobalGnbID(globalGNbID GlobalGnbID)
	SetGnbCuUpID(gnbCuUpID GnbCuUpID)
	SetGnbDuID(gnbDuID GnbDuID)
	GetGlobalGnbID()
	GetGnbCuUpID()
	GetGnbDuID()
	GetGlobalKpmnodeGnbID()
}

func NewGlobalKpmnodeGnbID() *GlobalKpmnodeGnbID {
	return &GlobalKpmnodeGnbID{}
}

func (b *GlobalKpmnodeGnbID) SetGlobalGnbID(globalGnbID *GlobalGnbID) *GlobalKpmnodeGnbID {
	b.GlobalGnbID = globalGnbID
	return b
}

func (b *GlobalKpmnodeGnbID) SetGnbCuUpID(gnbCuUpID *GnbCuUpID) *GlobalKpmnodeGnbID {
	b.GnbCuUpID = gnbCuUpID
	return b
}

func (b *GlobalKpmnodeGnbID) SetGnbDuID(gnbDuID *GnbDuID) *GlobalKpmnodeGnbID {
	b.GnbDuID = gnbDuID
	return b
}

func (b *GlobalKpmnodeGnbID) GetGlobalGnbID() *GlobalGnbID {
	return b.GlobalGnbID
}

func (b *GlobalKpmnodeGnbID) GetGnbCuUpID() *GnbCuUpID {
	return b.GnbCuUpID
}

func (b *GlobalKpmnodeGnbID) GetGnbDuID() *GnbDuID {
	return b.GnbDuID
}

func (b *GlobalKpmnodeGnbID) GetGlobalKpmnodeGnbID() *GlobalKpmnodeGnbID {
	return b
}
