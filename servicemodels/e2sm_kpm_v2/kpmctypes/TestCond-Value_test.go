// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondValue(t *testing.T) {

	testCondValue := &e2sm_kpm_v2.TestCondValue{
		TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
			ValueInt: 21,
		},
	}

	xer, err := xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(xer))
	t.Logf("TestCondValue XER\n%s", string(xer))
}

func Test_xerDecodeTestCondValue(t *testing.T) {

	testCondValue := &e2sm_kpm_v2.TestCondValue{
		TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
			ValueInt: 21,
		},
	}

	xer, err := xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(xer))
	t.Logf("TestCondValue XER\n%s", string(xer))

	result, err := xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue XER - decoded\n%s", result)
}

func Test_perEncodeTestCondValue(t *testing.T) {

	testCondValue := &e2sm_kpm_v2.TestCondValue{
		TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
			ValueInt: 21,
		},
	}

	per, err := perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondValue PER\n%s", string(per))
}

func Test_perDecodeTestCondValue(t *testing.T) {

	testCondValue := &e2sm_kpm_v2.TestCondValue{
		TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
			ValueInt: 21,
		},
	}

	per, err := perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondInfo PER\n%s", string(per))

	result, err := perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue PER - decoded\n%v", result)
}
