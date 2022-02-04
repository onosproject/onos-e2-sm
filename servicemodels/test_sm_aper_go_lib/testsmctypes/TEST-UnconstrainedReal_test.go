// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTestUnconstrainedRealMsg() (*test_sm_ies.TestUnconstrainedReal, error) {

	testUnconstrainedReal := test_sm_ies.TestUnconstrainedReal{
		AttrUcrA: 21.7,
		AttrUcrB: -653.43,
	}

	return &testUnconstrainedReal, nil
}

func Test_xerEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	xer, err := xerEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedReal XER\n%s", string(xer))

	result, err := xerDecodeTestUnconstrainedReal(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal XER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())
}

func Test_perEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	per, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedReal PER\n%v", hex.Dump(per))

	result, err := perDecodeTestUnconstrainedReal(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal PER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())
}
