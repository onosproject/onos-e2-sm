// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondExpression(t *testing.T) {

	testCondExpression := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	xer, err := xerEncodeTestCondExpression(testCondExpression)
	assert.NilError(t, err)
	//assert.Equal(t, 54, len(xer))
	t.Logf("TestCondExpression XER\n%s", string(xer))
}

func Test_xerDecodeTestCondExpression(t *testing.T) {

	testCondExpression := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	xer, err := xerEncodeTestCondExpression(testCondExpression)
	assert.NilError(t, err)
	//assert.Equal(t, 54, len(xer))
	t.Logf("TestCondExpression XER\n%s", string(xer))

	result, err := xerDecodeTestCondExpression(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondExpression XER - decoded\n%s", result)
}

func Test_perEncodeTestCondExpression(t *testing.T) {

	testCondExpression := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	per, err := perEncodeTestCondExpression(testCondExpression)
	assert.NilError(t, err)
	//assert.Equal(t, 1, len(per))
	t.Logf("TestCondExpression PER\n%v", hex.Dump(per))
}

func Test_perDecodeTestCondExpression(t *testing.T) {

	testCondExpression := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	per, err := perEncodeTestCondExpression(testCondExpression)
	assert.NilError(t, err)
	//assert.Equal(t, 1, len(per))
	t.Logf("TestCondExpression PER\n%v", hex.Dump(per))

	result, err := perDecodeTestCondExpression(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondExpression PER - decoded\n%v", result)
}
