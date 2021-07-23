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

func createTestList1Msg() (*test_sm.TestList1, error) {

	// testList1 := pdubuilder.CreateTestList1() //ToDo - fill in arguments here(if this function exists

	testList1 := test_sm.TestList1{}

	if err := testList1.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestList1 %s", err.Error())
	}
	return &testList1, nil
}

func Test_xerEncodingTestList1(t *testing.T) {

	testList1, err := createTestList1Msg()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	xer, err := xerEncodeTestList1(testList1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestList1 XER\n%s", string(xer))

	result, err := xerDecodeTestList1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList1 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList1.GetValue(), result.GetValue())

}

func Test_perEncodingTestList1(t *testing.T) {

	testList1, err := createTestList1Msg()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	per, err := perEncodeTestList1(testList1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestList1 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList1 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testList1.GetValue(), result.GetValue())

}
