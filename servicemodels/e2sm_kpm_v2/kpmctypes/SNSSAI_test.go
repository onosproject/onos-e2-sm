// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: 21,
		SD:  22,
	}
	xer, err := xerEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("Snssai XER\n%s", string(xer))
}

func Test_xerDecodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: 21,
		SD:  22,
	}
	xer, err := xerEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("Snssai XER\n%s", string(xer))

	result, err := xerDecodeSnssai(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: 21,
		SD:  22,
	}
	per, err := perEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("Snssai PER\n%s", string(per))
}

func Test_perDecodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: 21,
		SD:  22,
	}
	per, err := perEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("Snssai PER\n%s", string(per))

	result, err := perDecodeSnssai(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
