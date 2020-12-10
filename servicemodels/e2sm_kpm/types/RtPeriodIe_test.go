// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestRtPeriodIe_NewRtPeriodIe(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)

	assert.Assert(t, reflect.TypeOf(RtPeriodIe{}) == reflect.TypeOf(*rtPeriod), "RtPeriodIe{} values are mismatched")
	assert.DeepEqual(t, rtPeriod.RtPeriodIe, rtPeriodIe)
}

func TestRtPeriodIe_GetRtPeriodIe(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)

	assert.DeepEqual(t, rtPeriod.GetRtPeriodIe(), rtPeriodIe)
}
