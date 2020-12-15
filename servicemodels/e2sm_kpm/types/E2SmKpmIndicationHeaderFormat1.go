// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmIndicationHeaderFormat1 struct {
	IDglobalKpmnodeID *GlobalKpmnodeID
	NRcgi             *NRcgi
	PlmnID            *PlmnIdentity
	SliceID           *Snssai
	FiveQi            int32
	Qci               int32
}

type E2SmKpmIndicationHeaderFormat1Builder interface {
	NewE2SmKpmIndicationHeaderFormat1()
	SetIDglobalKpmnodeID(globalKpmnodeID *GlobalKpmnodeID)
	SetNRcgi(nrcgi *NRcgi)
	SetPlmnID(plmnID *PlmnIdentity)
	SetSliceID(sliceID *Snssai)
	SetFiveQi(fiveQi int32)
	SetQci(qci int32)
	GetIDlobalKpmnodeID()
	GetNRcgi()
	GetPlmnID()
	GetSliceID()
	GetFiveQi()
	GetQci()
	GetE2SmKpmIndicationHeaderFormat1()
}

func NewE2SmKpmIndicationHeaderFormat1() *E2SmKpmIndicationHeaderFormat1 {
	return &E2SmKpmIndicationHeaderFormat1{}
}

func (b *E2SmKpmIndicationHeaderFormat1) SetIDglobalKpmnodeID(globalKpmID *GlobalKpmnodeID) *E2SmKpmIndicationHeaderFormat1 {
	b.IDglobalKpmnodeID = globalKpmID
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) SetNRcgi(nrcgi *NRcgi) *E2SmKpmIndicationHeaderFormat1 {
	b.NRcgi = nrcgi
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) SetPlmnID(plmnID *PlmnIdentity) *E2SmKpmIndicationHeaderFormat1 {
	b.PlmnID = plmnID
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) SetSliceID(sliceID *Snssai) *E2SmKpmIndicationHeaderFormat1 {
	b.SliceID = sliceID
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) SetFiveQi(fiveQi int32) *E2SmKpmIndicationHeaderFormat1 {
	b.FiveQi = fiveQi
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) SetQci(qci int32) *E2SmKpmIndicationHeaderFormat1 {
	b.Qci = qci
	return b
}

func (b *E2SmKpmIndicationHeaderFormat1) GetIDlobalKpmnodeID() *GlobalKpmnodeID {
	return b.IDglobalKpmnodeID
}

func (b *E2SmKpmIndicationHeaderFormat1) GetNRcgi() *NRcgi {
	return b.NRcgi
}

func (b *E2SmKpmIndicationHeaderFormat1) GetPlmnID() *PlmnIdentity {
	return b.PlmnID
}

func (b *E2SmKpmIndicationHeaderFormat1) GetSliceID() *Snssai {
	return b.SliceID
}

func (b *E2SmKpmIndicationHeaderFormat1) GetFiveQi() int32 {
	return b.FiveQi
}

func (b *E2SmKpmIndicationHeaderFormat1) GetQci() int32 {
	return b.Qci
}

func (b *E2SmKpmIndicationHeaderFormat1) GetE2SmKpmIndicationHeaderFormat1() *E2SmKpmIndicationHeaderFormat1 {
	return b
}
