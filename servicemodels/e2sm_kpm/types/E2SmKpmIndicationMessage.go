// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmIndicationMessage struct {
	//RicStyleType *RicStyleType
	IndicationMessageFormat1 *E2SmKpmIndicationMessageFormat1
}

type E2SmKpmIndicationMessageBuilder interface {
	NewE2SmKpmIndicationMessage()
	SetE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1 *E2SmKpmIndicationMessageFormat1)
	//SetRicStyleType(ricStyleType *RicStyleType)
	GetE2SmKpmIndicationMessageFormat1()
	//GetRicStyleType()
	GetE2SmKpmIndicationMessage()
}

func NewE2SmKpmIndicationMessage() *E2SmKpmIndicationMessage {
	return &E2SmKpmIndicationMessage{}
}

func (b *E2SmKpmIndicationMessage) SetE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1 *E2SmKpmIndicationMessageFormat1) *E2SmKpmIndicationMessage {
	b.IndicationMessageFormat1 = e2SmKpmIndicationMessageFormat1
	//b.RicStyleType = nil
	return b
}

//func (b *E2SmKpmIndicationHeader) SetRicStyleType(ricStyleType *RicStyleType) *E2SmKpmIndicationHeader {
//	b.IndicationMessageFormat1 = nil
//	b.E2SmKpmIndicationHeaderFormatX = e2SmKpmIndicationHeaderFormatX
//	return b
//}

func (b *E2SmKpmIndicationMessage) GetE2SmKpmIndicationMessageFormat1() *E2SmKpmIndicationMessageFormat1 {
	return b.IndicationMessageFormat1
}

//func (b *E2SmKpmIndicationHeader) GetRicStyleType() *RicStyleType {
//	return b.RicStyleType
//}

func (b *E2SmKpmIndicationMessage) GetE2SmKpmIndicationMessage() *E2SmKpmIndicationMessage {
	return b
}
