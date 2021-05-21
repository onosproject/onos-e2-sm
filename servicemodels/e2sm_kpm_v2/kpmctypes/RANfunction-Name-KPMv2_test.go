// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createRanfunctionName() *e2sm_kpm_v2.RanfunctionName {

	return &e2sm_kpm_v2.RanfunctionName{
		RanFunctionShortName:   "onf",
		RanFunctionE2SmOid:     "oid123",
		RanFunctionDescription: "someDescription",
		RanFunctionInstance:    25,
	}
}

func Test_xerEncodeRanfunctionName(t *testing.T) {

	rfn := createRanfunctionName()

	xer, err := xerEncodeRanfunctionName(rfn)
	assert.NilError(t, err)
	//assert.Equal(t, 273, len(xer))
	t.Logf("RanfunctionName XER\n%s", string(xer))
}

func Test_xerDecodeRanfunctionName(t *testing.T) {

	rfn := createRanfunctionName()

	xer, err := xerEncodeRanfunctionName(rfn)
	assert.NilError(t, err)
	//assert.Equal(t, 273, len(xer))
	t.Logf("RanfunctionName XER\n%s", string(xer))

	result, err := xerDecodeRanfunctionName(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RanfunctionName XER - decoded\n%v", result)
}

func Test_perEncodeRanfunctionName(t *testing.T) {

	rfn := createRanfunctionName()

	per, err := perEncodeRanfunctionName(rfn)
	assert.NilError(t, err)
	//assert.Equal(t, 33, len(per))
	t.Logf("RanfunctionName PER\n%s", string(per))
}

func Test_perDecodeRanfunctionName(t *testing.T) {

	rfn := createRanfunctionName()

	per, err := perEncodeRanfunctionName(rfn)
	assert.NilError(t, err)
	//assert.Equal(t, 33, len(per))
	t.Logf("RanfunctionName PER\n%s", string(per))

	result, err := perDecodeRanfunctionName(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RanfunctionName PER - decoded\n%v", result)
}
