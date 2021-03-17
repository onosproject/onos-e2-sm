// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeQfi(t *testing.T) {

	qfi := &e2sm_kpm_v2.Qfi{
		Value: 32,
	}
	xer, err := xerEncodeQfi(qfi)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(xer))
	t.Logf("QFI XER\n%s", string(xer))
}

func Test_xerDecodeQfi(t *testing.T) {

	qfi := &e2sm_kpm_v2.Qfi{
		Value: 32,
	}
	xer, err := xerEncodeQfi(qfi)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(xer))
	t.Logf("QFI XER\n%s", string(xer))

	result, err := xerDecodeQfi(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("QFI XER - decoded\n%s", result)
}

func Test_perEncodeQfi(t *testing.T) {

	qfi := &e2sm_kpm_v2.Qfi{
		Value: 32,
	}
	per, err := perEncodeQfi(qfi)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("QFI PER\n%s", string(per))
}

func Test_perDecodeQfi(t *testing.T) {

	qfi := &e2sm_kpm_v2.Qfi{
		Value: 32,
	}
	per, err := perEncodeQfi(qfi)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("QFI PER\n%s", string(per))

	result, err := perDecodeQfi(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("QFI PER - decoded\n%v", result)
}
