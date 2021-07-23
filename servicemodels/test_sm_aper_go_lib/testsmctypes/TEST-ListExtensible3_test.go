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

func createTestListExtensible3Msg() (*test_sm.TestListExtensible3, error) {

	// testListExtensible3 := pdubuilder.CreateTestListExtensible3() //ToDo - fill in arguments here(if this function exists

	testListExtensible3 := test_sm.TestListExtensible3{}

	if err := testListExtensible3.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestListExtensible3 %s", err.Error())
	}
	return &testListExtensible3, nil
}

func Test_xerEncodingTestListExtensible3(t *testing.T) {

	testListExtensible3, err := createTestListExtensible3Msg()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	xer, err := xerEncodeTestListExtensible3(testListExtensible3)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestListExtensible3 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible3 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible3.GetValue(), result.GetValue())

}

func Test_perEncodingTestListExtensible3(t *testing.T) {

	testListExtensible3, err := createTestListExtensible3Msg()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	per, err := perEncodeTestListExtensible3(testListExtensible3)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestListExtensible3 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible3 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible3.GetValue(), result.GetValue())

}
