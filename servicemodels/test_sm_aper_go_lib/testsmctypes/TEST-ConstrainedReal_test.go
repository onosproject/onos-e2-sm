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

func createTestConstrainedRealMsg() (*test_sm.TestConstrainedReal, error) {

	// testConstrainedReal := pdubuilder.CreateTestConstrainedReal() //ToDo - fill in arguments here(if this function exists

	testConstrainedReal := test_sm.TestConstrainedReal{
		AttrCrA: nil,
		AttrCrB: nil,
		AttrCrC: nil,
		AttrCrD: nil,
		AttrCrE: nil,
		AttrCrF: nil,
	}

	if err := testConstrainedReal.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestConstrainedReal %s", err.Error())
	}
	return &testConstrainedReal, nil
}

func Test_xerEncodingTestConstrainedReal(t *testing.T) {

	testConstrainedReal, err := createTestConstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestConstrainedReal PDU")

	xer, err := xerEncodeTestConstrainedReal(testConstrainedReal)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestConstrainedReal XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedReal(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedReal XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testConstrainedReal.GetAttrCrA(), result.GetAttrCrA())
	assert.Equal(t, testConstrainedReal.GetAttrCrB(), result.GetAttrCrB())
	assert.Equal(t, testConstrainedReal.GetAttrCrC(), result.GetAttrCrC())
	assert.Equal(t, testConstrainedReal.GetAttrCrD(), result.GetAttrCrD())
	assert.Equal(t, testConstrainedReal.GetAttrCrE(), result.GetAttrCrE())
	assert.Equal(t, testConstrainedReal.GetAttrCrF(), result.GetAttrCrF())

}

func Test_perEncodingTestConstrainedReal(t *testing.T) {

	testConstrainedReal, err := createTestConstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestConstrainedReal PDU")

	per, err := perEncodeTestConstrainedReal(testConstrainedReal)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestConstrainedReal PER\n%v", hex.Dump(per))

	result, err := perDecodeTestConstrainedReal(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedReal PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testConstrainedReal.GetAttrCrA(), result.GetAttrCrA())
	assert.Equal(t, testConstrainedReal.GetAttrCrB(), result.GetAttrCrB())
	assert.Equal(t, testConstrainedReal.GetAttrCrC(), result.GetAttrCrC())
	assert.Equal(t, testConstrainedReal.GetAttrCrD(), result.GetAttrCrD())
	assert.Equal(t, testConstrainedReal.GetAttrCrE(), result.GetAttrCrE())
	assert.Equal(t, testConstrainedReal.GetAttrCrF(), result.GetAttrCrF())

}
