// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type OcucpPfContainer_CuCpResourceStatus001 struct {
	NumberOfActiveUes int32
}

type OcucpPfContainer_CuCpResourceStatus001Builder interface {
	NewOcucpPfContainer_CuCpResourceStatus001(nUEs int32)
	GetNumberOfActiveUes()
	GetOcucpPfContainer_CuCpResourceStatus001()
}

func NewOcucpPfContainer_CuCpResourceStatus001(nUEs int32) *OcucpPfContainer_CuCpResourceStatus001 {
	return &OcucpPfContainer_CuCpResourceStatus001{
		NumberOfActiveUes: nUEs,
	}
}

func (b *OcucpPfContainer_CuCpResourceStatus001) GetNumberOfActiveUes() int32 {
	return b.NumberOfActiveUes
}

func (b *OcucpPfContainer_CuCpResourceStatus001) GetOcucpPfContainer_CuCpResourceStatus001() *OcucpPfContainer_CuCpResourceStatus001 {
	return b
}