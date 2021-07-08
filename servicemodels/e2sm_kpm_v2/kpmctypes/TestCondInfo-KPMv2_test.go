// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondInfo(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2.TestCondInfo{
		TestType: &e2sm_kpm_v2.TestCondType{
			TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
				AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
			},
		},
		TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
		TestValue: &e2sm_kpm_v2.TestCondValue{
			TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
				ValueInt: 21,
			},
		},
	}

	xer, err := xerEncodeTestCondInfo(testCondInfo)
	assert.NilError(t, err)
	//assert.Equal(t, 196, len(xer))
	t.Logf("TestCondInfo XER\n%s", string(xer))
}

func Test_xerDecodeTestCondInfo(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2.TestCondInfo{
		TestType: &e2sm_kpm_v2.TestCondType{
			TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
				AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
			},
		},
		TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
		TestValue: &e2sm_kpm_v2.TestCondValue{
			TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
				ValueInt: 21,
			},
		},
	}

	xer, err := xerEncodeTestCondInfo(testCondInfo)
	assert.NilError(t, err)
	//assert.Equal(t, 196, len(xer))
	t.Logf("TestCondInfo XER\n%s", string(xer))

	result, err := xerDecodeTestCondInfo(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondInfo XER - decoded\n%s", result)
}

func Test_perEncodeTestCondInfo(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2.TestCondInfo{
		TestType: &e2sm_kpm_v2.TestCondType{
			TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
				AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
			},
		},
		TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
		TestValue: &e2sm_kpm_v2.TestCondValue{
			TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
				ValueInt: 21,
			},
		},
	}

	per, err := perEncodeTestCondInfo(testCondInfo)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("TestCondInfo PER\n%v", hex.Dump(per))
}

func Test_perDecodeTestCondInfo(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2.TestCondInfo{
		TestType: &e2sm_kpm_v2.TestCondType{
			TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
				AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
			},
		},
		TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
		TestValue: &e2sm_kpm_v2.TestCondValue{
			TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
				ValueInt: 21,
			},
		},
	}

	per, err := perEncodeTestCondInfo(testCondInfo)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("TestCondInfo PER\n%v", hex.Dump(per))

	result, err := perDecodeTestCondInfo(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondInfo PER - decoded\n%v", result)
}
