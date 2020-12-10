// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import "fmt"

type RtPeriodIe struct {
	RtPeriodIe int32
}

type RtPeriodIeBuilder interface {
	NewRtPeriodIe(rtPeriodIe int32)
	GetRtPeriodIe()
}

func NewRtPeriodIe(rtPeriodIe int32) (*RtPeriodIe, error) {
	if rtPeriodIe >= 0 && rtPeriodIe <= 19 {
		return &RtPeriodIe{
			RtPeriodIe: rtPeriodIe,
		}, nil
	}
	return nil, fmt.Errorf("RtPeriodIe should should be in range <0, 19>")
}

func (b *RtPeriodIe) GetRtPeriodIe() int32 {
	return b.RtPeriodIe
}
