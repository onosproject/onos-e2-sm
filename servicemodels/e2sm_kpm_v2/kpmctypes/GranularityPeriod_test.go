// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGranularityPeriod(t *testing.T) {

	gp := &e2sm_kpm_v2.GranularityPeriod{
		Value: 21,
	}
	xer, err := xerEncodeGranularityPeriod(gp)
	assert.NilError(t, err)
	assert.Equal(t, 42, len(xer))
	t.Logf("GranularityPeriod XER\n%s", string(xer))
}

func Test_xerDecodeGranularityPeriod(t *testing.T) {

	gp := &e2sm_kpm_v2.GranularityPeriod{
		Value: 21,
	}
	xer, err := xerEncodeGranularityPeriod(gp)
	assert.NilError(t, err)
	assert.Equal(t, 42, len(xer))
	t.Logf("GranularityPeriod XER\n%s", string(xer))

	result, err := xerDecodeGranularityPeriod(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GranularityPeriod XER - decoded\n%s", result)
}

func Test_perEncodeGranularityPeriod(t *testing.T) {

	gp := &e2sm_kpm_v2.GranularityPeriod{
		Value: 21,
	}
	per, err := perEncodeGranularityPeriod(gp)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("GranularityPeriod PER\n%s", string(per))
}

func Test_perDecodeGranularityPeriod(t *testing.T) {

	gp := &e2sm_kpm_v2.GranularityPeriod{
		Value: 21,
	}
	per, err := perEncodeGranularityPeriod(gp)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("GranularityPeriod PER\n%s", string(per))

	result, err := perDecodeGranularityPeriod(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GranularityPeriod PER - decoded\n%s", result)
}
