// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createFiveQi() *e2sm_kpm_v2.FiveQi {

	return &e2sm_kpm_v2.FiveQi{
		Value: 12,
	}
}

func Test_xerEncodeFiveQi(t *testing.T) {

	fqi := createFiveQi()

	xer, err := xerEncodeFiveQi(fqi)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("FiveQi XER\n%s", string(xer))
}

func Test_xerDecodeFiveQi(t *testing.T) {

	fqi := createFiveQi()

	xer, err := xerEncodeFiveQi(fqi)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("FiveQi XER\n%s", string(xer))

	result, err := xerDecodeFiveQi(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeFiveQi(t *testing.T) {

	fqi := createFiveQi()

	per, err := perEncodeFiveQi(fqi)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("FiveQi PER\n%s", string(per))
}

func Test_perDecodeFiveQi(t *testing.T) {

	fqi := createFiveQi()

	per, err := perEncodeFiveQi(fqi)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("FiveQi PER\n%s", string(per))

	result, err := perDecodeFiveQi(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
