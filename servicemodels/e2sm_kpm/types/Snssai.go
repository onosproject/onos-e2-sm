// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import "fmt"

type Snssai struct {
	SSt []byte
	SD  []byte
}

type SnssaiBuilder interface {
	NewSnssai(sst []byte, sd []byte)
	GetSSt()
	GetSD()
	GetSnssai()
}

func NewSnssai(sst []byte, sd []byte) (*Snssai, error) {
	if len(sst) != 1 {
		return nil, fmt.Errorf("SSt size must be 1 byte")
	}
	if len(sd) != 3 {
		return nil, fmt.Errorf("SD size must be 3 bytes")
	}
	return &Snssai{
		SSt: sst,
		SD:  sd,
	}, nil
}

func (b *Snssai) GetSSt() []byte {
	return b.SSt
}

func (b *Snssai) GetSD() []byte {
	return b.SD
}

func (b *Snssai) GetSnssai() *Snssai {
	return b
}
