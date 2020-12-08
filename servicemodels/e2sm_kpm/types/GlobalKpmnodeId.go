// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type GlobalKpmnodeID struct {
	Gnb *GlobalKpmnodeGnbID
	//EnGnb *GlobalKpmnodeEnGnbID
	//NgEnb *GlobalKpmnodeNgEnbID
	//Enb   *GlobalKpmnodeEnbID
}

type GlobalKpmnodeIDBuilder interface {
	NewGlobalKpmnodeID()
	SetGlobalKpmnodeGnbID(globalKpmGnbID GlobalKpmnodeGnbID)
	//SetGlobalKpmnodeEnGnbID(globalKpmEnGnbID GlobalKpmnodeEnGnbID)
	//SetGlobalKpmnodeNgEnbID(globalKpmNgEnbID GlobalKpmnodeNgEnbID)
	//SetGlobalKpmnodeEnbID(globalKpmEnbID GlobalKpmnodeEnbID)
	GetGlobalKpmnodeGnbID()
	//GetGlobalKpmnodeEnGnbID()
	//GetGlobalKpmnodeNgEnbID()
	//GetGlobalKpmnodeEnbID()
	GetGlobalKpmnodeID()
}

func NewGlobalKpmnodeID() *GlobalKpmnodeID {
	return &GlobalKpmnodeID{}
}

func (b *GlobalKpmnodeID) SetGlobalKpmnodeGnbID(globalKpmGnbID *GlobalKpmnodeGnbID) *GlobalKpmnodeID {
	b.Gnb = globalKpmGnbID
	//b.EnGnb = nil
	//b.NgEnb = nil
	//b.Enb = nil
	return b
}

//func (b *GlobalKpmnodeID) SetGlobalKpmnodeEnGnbID(globalKpmEnGnbID *GlobalKpmnodeEnGnbID) *GlobalKpmnodeID {
//	//b.Gnb = nil
//	b.EnGnb = globalKpmEnGnbID
//	//b.NgEnb = nil
//	//b.Enb = nil
//	return b
//}

//func (b *GlobalKpmnodeID) SetGlobalKpmnodeNgEnbID(globalKpmNgEnbID *GlobalKpmnodeNgEnbID) *GlobalKpmnodeID {
//	//b.Gnb = nil
//	//b.EnGnb = nil
//	b.NgEnb = globalKpmNgEnbID
//	//b.Enb = nil
//	return b
//}

//func (b *GlobalKpmnodeID) SetGlobalKpmnodeEnbID(globalKpmEnbID *GlobalKpmnodeEnbID) *GlobalKpmnodeID {
//	//b.Gnb = nil
//	//b.EnGnb = nil
//	//b.NgEnb = nil
//	b.Enb = globalKpmEnbID
//	return b
//}

func (b *GlobalKpmnodeID) GetGlobalKpmnodeGnbID() *GlobalKpmnodeGnbID {
	return b.Gnb
}

//func (b *GlobalKpmnodeID) GetGlobalKpmnodeEnGnbID() *GlobalKpmnodeEnGnbID {
//	return b.EnGnb
//}

//func (b *GlobalKpmnodeID) GetGlobalKpmnodeNgEnbID() *GlobalKpmnodeNgEnbID {
//	return b.NgEnb
//}

//func (b *GlobalKpmnodeID) GetGlobalKpmnodeEnbID() *GlobalKpmnodeEnbID {
//	return b.Enb
//}

func (b *GlobalKpmnodeID) GetGlobalKpmnodeID() *GlobalKpmnodeID {
	return b
}
