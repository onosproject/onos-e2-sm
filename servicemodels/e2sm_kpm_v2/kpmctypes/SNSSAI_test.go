// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: []byte{0x01},
		SD:  []byte{0x01, 0x02, 0x03},
	}
	xer, err := xerEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 59, len(xer))
	t.Logf("Snssai XER\n%s", string(xer))
}

func Test_xerDecodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: []byte{0x01},
		SD:  []byte{0x01, 0x02, 0x03},
	}
	xer, err := xerEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 59, len(xer))
	t.Logf("Snssai XER\n%s", string(xer))

	result, err := xerDecodeSnssai(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Snssai XER - decoded\n%s", result)
}

func Test_perEncodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: []byte{0x01},
		SD:  []byte{0x01, 0x02, 0x03},
	}
	per, err := perEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("Snssai PER\n%s", string(per))
}

func Test_perDecodeSnssai(t *testing.T) {

	snssai := &e2sm_kpm_v2.Snssai{
		SSt: []byte{0x01},
		SD:  []byte{0x01, 0x02, 0x03},
	}
	per, err := perEncodeSnssai(snssai)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("Snssai PER\n%s", string(per))

	result, err := perDecodeSnssai(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Snssai PER - decoded\n%v", result)
}
