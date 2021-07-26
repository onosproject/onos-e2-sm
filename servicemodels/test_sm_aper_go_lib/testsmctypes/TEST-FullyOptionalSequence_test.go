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

func createTestFullyOptionalSequenceMsg() (*test_sm_ies.TestFullyOptionalSequence, error) {

	// testFullyOptionalSequence := pdubuilder.CreateTestFullyOptionalSequence() //ToDo - fill in arguments here(if this function exists

	var ie1 int32 = 32
	ie2 := []byte{0xF0}
	var ie3 bool = true
	var ie4 int32 = 32
	var ie5 int32 = 32
	testFullyOptionalSequence := test_sm_ies.TestFullyOptionalSequence{
		Item1: &ie1,
		Item2: ie2,
		Item3: &ie3,
		Item4: &ie4,
		Item5: &ie5,
	}

	//if err := testFullyOptionalSequence.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating TestFullyOptionalSequence %s", err.Error())
	//}
	return &testFullyOptionalSequence, nil
}

func Test_xerEncodingTestFullyOptionalSequence(t *testing.T) {

	testFullyOptionalSequence, err := createTestFullyOptionalSequenceMsg()
	assert.NilError(t, err, "Error creating TestFullyOptionalSequence PDU")

	xer, err := xerEncodeTestFullyOptionalSequence(testFullyOptionalSequence)
	assert.NilError(t, err)
	t.Logf("TestFullyOptionalSequence XER\n%s", string(xer))

	result, err := xerDecodeTestFullyOptionalSequence(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestFullyOptionalSequence XER - decoded\n%v", result)
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
	t.Logf("TestFullyOptionalSequence PER\n%v", hex.Dump(per))

	result, err := perDecodeTestFullyOptionalSequence(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestFullyOptionalSequence PER - decoded\n%v", result)
	assert.Equal(t, testFullyOptionalSequence.GetItem1(), result.GetItem1())
	assert.Equal(t, testFullyOptionalSequence.GetItem2(), result.GetItem2())
	assert.Equal(t, testFullyOptionalSequence.GetItem3(), result.GetItem3())
	assert.Equal(t, testFullyOptionalSequence.GetItem4(), result.GetItem4())
	assert.Equal(t, testFullyOptionalSequence.GetItem5(), result.GetItem5())
}
