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

func createTestUnconstrainedRealMsg() (*test_sm.TestUnconstrainedReal, error) {

	// testUnconstrainedReal := pdubuilder.CreateTestUnconstrainedReal() //ToDo - fill in arguments here(if this function exists

	testUnconstrainedReal := test_sm.TestUnconstrainedReal{
		AttrUcrA: nil,
		AttrUcrB: nil,
	}

	if err := testUnconstrainedReal.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestUnconstrainedReal %s", err.Error())
	}
	return &testUnconstrainedReal, nil
}

func Test_xerEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	xer, err := xerEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestUnconstrainedReal XER\n%s", string(xer))

	result, err := xerDecodeTestUnconstrainedReal(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())

}

func Test_perEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	per, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestUnconstrainedReal PER\n%v", hex.Dump(per))

	result, err := perDecodeTestUnconstrainedReal(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())

}
