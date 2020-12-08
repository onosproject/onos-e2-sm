// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type OCuCpPfContainer struct {
	GNbCuCpName        *GnbCuCpName
	CuCpResourceStatus *OcucpPfContainerCuCpResourceStatus001
}

type OcucpPfContainerBuilder interface {
	NewOcucpPfContainer()
	SetGNbCuCpName(gbbCuCpName *GnbCuCpName)
	SetCuCpResourceStatus(cuCpResourceStatus *OcucpPfContainerCuCpResourceStatus001)
	GetGNbCuCpName()
	GetCuCpResourceStatus()
	GetOcucpPfContainer()
}

func NewOCuCpPfContainer() *OCuCpPfContainer {
	return &OCuCpPfContainer{}
}

func (b *OCuCpPfContainer) SetGNbCuCpName(gbbCuCpName *GnbCuCpName) *OCuCpPfContainer {
	b.GNbCuCpName = gbbCuCpName
	return b
}

func (b *OCuCpPfContainer) SetCuCpResourceStatus(cuCpResourceStatus *OcucpPfContainerCuCpResourceStatus001) *OCuCpPfContainer {
	b.CuCpResourceStatus = cuCpResourceStatus
	return b
}

func (b *OCuCpPfContainer) GetGNbCuCpName() *GnbCuCpName {
	return b.GNbCuCpName
}

func (b *OCuCpPfContainer) GetCuCpResourceStatus() *OcucpPfContainerCuCpResourceStatus001 {
	return b.CuCpResourceStatus
}

func (b *OCuCpPfContainer) GetOCuCpPfContainer() *OCuCpPfContainer {
	return b
}
