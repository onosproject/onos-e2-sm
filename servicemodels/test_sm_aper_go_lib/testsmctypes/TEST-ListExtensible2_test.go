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

func createTestListExtensible2Msg() (*test_sm.TestListExtensible2, error) {

	// testListExtensible2 := pdubuilder.CreateTestListExtensible2() //ToDo - fill in arguments here(if this function exists

	testListExtensible2 := test_sm.TestListExtensible2{}

	if err := testListExtensible2.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestListExtensible2 %s", err.Error())
	}
	return &testListExtensible2, nil
}

func Test_xerEncodingTestListExtensible2(t *testing.T) {

	testListExtensible2, err := createTestListExtensible2Msg()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	xer, err := xerEncodeTestListExtensible2(testListExtensible2)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestListExtensible2 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible2 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible2.GetValue(), result.GetValue())

}

func Test_perEncodingTestListExtensible2(t *testing.T) {

	testListExtensible2, err := createTestListExtensible2Msg()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	per, err := perEncodeTestListExtensible2(testListExtensible2)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestListExtensible2 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible2 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible2.GetValue(), result.GetValue())

}
