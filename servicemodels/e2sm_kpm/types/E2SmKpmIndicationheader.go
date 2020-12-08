// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmIndicationHeader struct {
	E2SmKpmIndicationHeaderFormat1 *E2SmKpmIndicationHeaderFormat1
	//E2SmKpmIndicationHeaderFormatX *E2SmKpmIndicationHeaderFormatX
}

type E2SmKpmIndicationHeaderBuilder interface {
	NewE2SmKpmIndicationHeader()
	SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1 *E2SmKpmIndicationHeaderFormat1)
	//SetE2SmKpmIndicationHeaderFormatX(e2SmKpmIndicationHeader *E2SmKpmIndicationHeaderFormatX)
	GetE2SmKpmIndicationHeaderFormat1()
	//GetE2SmKpmIndicationHeaderFormatX()
	GetE2SmKpmIndicationHeader()
}

func NewE2SmKpmIndicationHeader() *E2SmKpmIndicationHeader {
	return &E2SmKpmIndicationHeader{}
}

func (b *E2SmKpmIndicationHeader) SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1 *E2SmKpmIndicationHeaderFormat1) *E2SmKpmIndicationHeader {
	b.E2SmKpmIndicationHeaderFormat1 = e2SmKpmIndicationHeaderFormat1
	//b.E2SmKpmIndicationHeaderFormatX = nil
	return b
}

//func (b *E2SmKpmIndicationHeader) SetE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormatX *E2SmKpmIndicationHeaderFormatX) *E2SmKpmIndicationHeader {
//	b.E2SmKpmIndicationHeaderFormat1 = nil
//	b.E2SmKpmIndicationHeaderFormatX = e2SmKpmIndicationHeaderFormatX
//	return b
//}

func (b *E2SmKpmIndicationHeader) GetE2SmKpmIndicationHeaderFormat1() *E2SmKpmIndicationHeaderFormat1 {
	return b.E2SmKpmIndicationHeaderFormat1
}

//func (b *E2SmKpmIndicationHeader) GetE2SmKpmIndicationHeaderFormatX() *E2SmKpmIndicationHeaderFormatX {
//	return b.E2SmKpmIndicationHeaderFormatX
//}

func (b *E2SmKpmIndicationHeader) GetE2SmKpmIndicationHeader() *E2SmKpmIndicationHeader {
	return b
}
