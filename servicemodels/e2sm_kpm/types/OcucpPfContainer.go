// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type OcucpPfContainer struct {
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

func NewOcucpPfContainer() *OcucpPfContainer {
	return &OcucpPfContainer{}
}

func (b *OcucpPfContainer) SetGNbCuCpName(gbbCuCpName *GnbCuCpName) *OcucpPfContainer {
	b.GNbCuCpName = gbbCuCpName
	return b
}

func (b *OcucpPfContainer) SetCuCpResourceStatus(cuCpResourceStatus *OcucpPfContainerCuCpResourceStatus001) *OcucpPfContainer {
	b.CuCpResourceStatus = cuCpResourceStatus
	return b
}

func (b *OcucpPfContainer) GetGNbCuCpName() *GnbCuCpName {
	return b.GNbCuCpName
}

func (b *OcucpPfContainer) GetCuCpResourceStatus() *OcucpPfContainerCuCpResourceStatus001 {
	return b.CuCpResourceStatus
}

func (b *OcucpPfContainer) GetOcucpPfContainer() *OcucpPfContainer {
	return b
}
