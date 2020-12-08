// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type OcucpPfContainerCuCpResourceStatus001 struct {
	NumberOfActiveUes int32
}

type OcucpPfContainerCuCpResourceStatus001Builder interface {
	NewOcucpPfContainerCuCpResourceStatus001(nUEs int32)
	GetNumberOfActiveUes()
	GetOcucpPfContainerCuCpResourceStatus001()
}

func NewOcucpPfContainerCuCpResourceStatus001(nUEs int32) *OcucpPfContainerCuCpResourceStatus001 {
	return &OcucpPfContainerCuCpResourceStatus001{
		NumberOfActiveUes: nUEs,
	}
}

func (b *OcucpPfContainerCuCpResourceStatus001) GetNumberOfActiveUes() int32 {
	return b.NumberOfActiveUes
}

func (b *OcucpPfContainerCuCpResourceStatus001) GetOcucpPfContainerCuCpResourceStatus001() *OcucpPfContainerCuCpResourceStatus001 {
	return b
}
