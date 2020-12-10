// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RanfunctionName struct {
	RanFunctionShortName   string
	RanFunctionE2SmOid     string
	RanFunctionDescription string
	RanFunctionInstance    int32
}

type RanfunctionNameBuilder interface {
	NewRanfunctionName(shortName string, oid string, description string, instance int32)
	GetRanFunctionShortName()
	GetRanFunctionE2SmOid()
	GetRanFunctionDescription()
	GetRanFunctionInstance()
	GetRanfunctionName()
}

func NewRanfunctionName(shortName string, oid string, description string, instance int32) *RanfunctionName {
	return &RanfunctionName{
		RanFunctionShortName:   shortName,
		RanFunctionE2SmOid:     oid,
		RanFunctionDescription: description,
		RanFunctionInstance:    instance,
	}
}

func (b *RanfunctionName) GetRanFunctionShortName() string {
	return b.RanFunctionShortName
}

func (b *RanfunctionName) GetRanFunctionE2SmOid() string {
	return b.RanFunctionE2SmOid
}

func (b *RanfunctionName) GetRanFunctionDescription() string {
	return b.RanFunctionDescription
}

func (b *RanfunctionName) GetRanFunctionInstance() int32 {
	return b.RanFunctionInstance
}

func (b *RanfunctionName) GetRanfunctionName() *RanfunctionName {
	return b
}
