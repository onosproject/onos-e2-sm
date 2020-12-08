// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type PfContainer struct {
	//ODu  *ODuPfContainer
	OCuCp *OCuCpPfContainer
	//OCuUp *OCuUpPfContainer
}

type PfContainerBuilder interface {
	NewPfContainer()
	//SetODu(ODu *ODuPfContainer)
	SetOCuCp(OCuCp *OCuCpPfContainer)
	//SetOCuUp(OCuUp *OCuUpPfContainer)
	//GetODu()
	GetOCuCp()
	//GetOCuUp()
	GetPfContainer()
}

func NewPfContainer() *PfContainer {
	return &PfContainer{}
}

//func (b *PfContainer) SetODu(ODu *ODuPfContainer) *PfContainer {
//	b.ODu = globalKpmGnbID
//	return b
//}

func (b *PfContainer) SetOCuCp(OCuCp *OCuCpPfContainer) *PfContainer {
	b.OCuCp = OCuCp
	return b
}

//func (b *PfContainer) SetOCuUp(OCuUp *OCuUpPfContainer) *PfContainer {
//	b.OCuUp = OCuUp
//	return b
//}

//func (b *PfContainer) GetODu() *ODuPfContainer {
//	return b.ODu
//}

func (b *PfContainer) GetOCuCp() *OCuCpPfContainer {
	return b.OCuCp
}

//func (b *PfContainer) GetOCuUp() *OCuUpPfContainer {
//	return b.OCuUp
//}

func (b *PfContainer) GetPfContainer() *PfContainer {
	return b
}
