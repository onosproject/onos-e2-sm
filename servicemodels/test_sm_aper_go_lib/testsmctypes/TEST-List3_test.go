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

func createTestList3Msg() (*test_sm.TestList3, error) {

	// testList3 := pdubuilder.CreateTestList3() //ToDo - fill in arguments here(if this function exists

	testList3 := test_sm.TestList3{}

	if err := testList3.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestList3 %s", err.Error())
	}
	return &testList3, nil
}

func Test_xerEncodingTestList3(t *testing.T) {

	testList3, err := createTestList3Msg()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	xer, err := xerEncodeTestList3(testList3)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestList3 XER\n%s", string(xer))

	result, err := xerDecodeTestList3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList3 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList3.GetValue(), result.GetValue())

}

func Test_perEncodingTestList3(t *testing.T) {

	testList3, err := createTestList3Msg()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	per, err := perEncodeTestList3(testList3)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestList3 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList3 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList3.GetValue(), result.GetValue())

}
