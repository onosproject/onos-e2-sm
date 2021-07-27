// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTestUnconstrainedIntMsg() (*test_sm_ies.TestUnconstrainedInt, error) {

	testUnconstrainedInt := test_sm_ies.TestUnconstrainedInt{
		AttrUciA: -153,
		AttrUciB: 21,
	}

	return &testUnconstrainedInt, nil
}

func Test_xerEncodingTestUnconstrainedInt(t *testing.T) {

	testUnconstrainedInt, err := createTestUnconstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedInt PDU")

	xer, err := xerEncodeTestUnconstrainedInt(testUnconstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedInt XER\n%s", string(xer))

	result, err := xerDecodeTestUnconstrainedInt(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedInt XER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedInt.GetAttrUciA(), result.GetAttrUciA())
	assert.Equal(t, testUnconstrainedInt.GetAttrUciB(), result.GetAttrUciB())
}

func Test_perEncodingTestUnconstrainedInt(t *testing.T) {

	testUnconstrainedInt, err := createTestUnconstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedInt PDU")

	per, err := perEncodeTestUnconstrainedInt(testUnconstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedInt PER\n%v", hex.Dump(per))

	result, err := perDecodeTestUnconstrainedInt(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedInt PER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedInt.GetAttrUciA(), result.GetAttrUciA())
	assert.Equal(t, testUnconstrainedInt.GetAttrUciB(), result.GetAttrUciB())
}
