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

func createTestFullyOptionalSequenceMsg() (*test_sm.TestFullyOptionalSequence, error) {

	// testFullyOptionalSequence := pdubuilder.CreateTestFullyOptionalSequence() //ToDo - fill in arguments here(if this function exists

	testFullyOptionalSequence := test_sm.TestFullyOptionalSequence{
		Item1: nil,
		Item2: nil,
		Item3: nil,
		Item4: nil,
		Item5: nil,
	}

	if err := testFullyOptionalSequence.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestFullyOptionalSequence %s", err.Error())
	}
	return &testFullyOptionalSequence, nil
}

func Test_xerEncodingTestFullyOptionalSequence(t *testing.T) {

	testFullyOptionalSequence, err := createTestFullyOptionalSequenceMsg()
	assert.NilError(t, err, "Error creating TestFullyOptionalSequence PDU")

	xer, err := xerEncodeTestFullyOptionalSequence(testFullyOptionalSequence)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestFullyOptionalSequence XER\n%s", string(xer))

	result, err := xerDecodeTestFullyOptionalSequence(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestFullyOptionalSequence XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testFullyOptionalSequence.GetItem1(), result.GetItem1())
	assert.Equal(t, testFullyOptionalSequence.GetItem2(), result.GetItem2())
	assert.Equal(t, testFullyOptionalSequence.GetItem3(), result.GetItem3())
	assert.Equal(t, testFullyOptionalSequence.GetItem4(), result.GetItem4())
	assert.Equal(t, testFullyOptionalSequence.GetItem5(), result.GetItem5())

}

func Test_perEncodingTestFullyOptionalSequence(t *testing.T) {

	testFullyOptionalSequence, err := createTestFullyOptionalSequenceMsg()
	assert.NilError(t, err, "Error creating TestFullyOptionalSequence PDU")

	per, err := perEncodeTestFullyOptionalSequence(testFullyOptionalSequence)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestFullyOptionalSequence PER\n%v", hex.Dump(per))

	result, err := perDecodeTestFullyOptionalSequence(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestFullyOptionalSequence PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testFullyOptionalSequence.GetItem1(), result.GetItem1())
	assert.Equal(t, testFullyOptionalSequence.GetItem2(), result.GetItem2())
	assert.Equal(t, testFullyOptionalSequence.GetItem3(), result.GetItem3())
	assert.Equal(t, testFullyOptionalSequence.GetItem4(), result.GetItem4())
	assert.Equal(t, testFullyOptionalSequence.GetItem5(), result.GetItem5())

}
