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

func createTestList2Msg() (*test_sm.TestList2, error) {

	// testList2 := pdubuilder.CreateTestList2() //ToDo - fill in arguments here(if this function exists

	testList2 := test_sm.TestList2{}

	if err := testList2.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestList2 %s", err.Error())
	}
	return &testList2, nil
}

func Test_xerEncodingTestList2(t *testing.T) {

	testList2, err := createTestList2Msg()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	xer, err := xerEncodeTestList2(testList2)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestList2 XER\n%s", string(xer))

	result, err := xerDecodeTestList2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList2 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList2.GetValue(), result.GetValue())

}

func Test_perEncodingTestList2(t *testing.T) {

	testList2, err := createTestList2Msg()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	per, err := perEncodeTestList2(testList2)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestList2 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList2 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList2.GetValue(), result.GetValue())

}
