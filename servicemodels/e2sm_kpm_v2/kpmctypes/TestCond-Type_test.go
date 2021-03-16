// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondType(t *testing.T) {

	testCondType := &e2sm_kpm_v2.TestCondType{
		TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
		},
	}

	xer, err := xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType XER\n%s", string(xer))
}

func Test_xerDecodeTestCondType(t *testing.T) {

	testCondType := &e2sm_kpm_v2.TestCondType{
		TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
		},
	}

	xer, err := xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType XER\n%s", string(xer))

	result, err := xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType XER - decoded\n%s", result)
}

func Test_perEncodeTestCondType(t *testing.T) {

	testCondType := &e2sm_kpm_v2.TestCondType{
		TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
		},
	}

	per, err := perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType PER\n%s", string(per))
}

func Test_perDecodeTestCondType(t *testing.T) {

	testCondType := &e2sm_kpm_v2.TestCondType{
		TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
		},
	}

	per, err := perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType PER\n%s", string(per))

	result, err := perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType PER - decoded\n%v", result)
}
