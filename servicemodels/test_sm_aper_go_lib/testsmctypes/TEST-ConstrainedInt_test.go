// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	"fmt"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/test_sm/pdubuilder"
	test_sm "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm" //ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createTestConstrainedIntMsg() (*test_sm.TestConstrainedInt, error) {

	// testConstrainedInt := pdubuilder.CreateTestConstrainedInt() //ToDo - fill in arguments here(if this function exists

	testConstrainedInt := test_sm.TestConstrainedInt{
		AttrCiA: nil,
		AttrCiB: nil,
		AttrCiC: nil,
		AttrCiD: nil,
		AttrCiE: nil,
		AttrCiF: nil,
	}

	if err := testConstrainedInt.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestConstrainedInt %s", err.Error())
	}
	return &testConstrainedInt, nil
}

func Test_xerEncodingTestConstrainedInt(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	xer, err := xerEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestConstrainedInt XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedInt(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt XER - decoded\n%v", result)
	//ToDo - adjust field's verification
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
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestConstrainedInt PER\n%v", hex.Dump(per))

	result, err := perDecodeTestConstrainedInt(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())

}
