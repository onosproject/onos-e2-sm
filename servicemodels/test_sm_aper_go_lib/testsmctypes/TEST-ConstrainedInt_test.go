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

func createTestConstrainedIntMsg() (*test_sm_ies.TestConstrainedInt, error) {

	testConstrainedInt := test_sm_ies.TestConstrainedInt{
		AttrCiA: 99,
		AttrCiB: 11,
		AttrCiC: -153,
		AttrCiD: 15,
		AttrCiE: 10,
		AttrCiF: 10,
	}

	return &testConstrainedInt, nil
}

func Test_xerEncodingTestConstrainedInt(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	xer, err := xerEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedInt(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt XER - decoded\n%v", result)
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())
}

func Test_perEncodingTestConstrainedInt(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	per, err := perEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt PER\n%v", hex.Dump(per))

	result, err := perDecodeTestConstrainedInt(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt PER - decoded\n%v", result)
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())
}
