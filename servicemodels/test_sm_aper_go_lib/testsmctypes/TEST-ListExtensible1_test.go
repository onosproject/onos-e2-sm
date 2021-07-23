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

func createTestListExtensible1Msg() (*test_sm.TestListExtensible1, error) {

	// testListExtensible1 := pdubuilder.CreateTestListExtensible1() //ToDo - fill in arguments here(if this function exists

	testListExtensible1 := test_sm.TestListExtensible1{}

	if err := testListExtensible1.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestListExtensible1 %s", err.Error())
	}
	return &testListExtensible1, nil
}

func Test_xerEncodingTestListExtensible1(t *testing.T) {

	testListExtensible1, err := createTestListExtensible1Msg()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	xer, err := xerEncodeTestListExtensible1(testListExtensible1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestListExtensible1 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible1 XER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible1.GetValue(), result.GetValue())

}

func Test_perEncodingTestListExtensible1(t *testing.T) {

	testListExtensible1, err := createTestListExtensible1Msg()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	per, err := perEncodeTestListExtensible1(testListExtensible1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestListExtensible1 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible1 PER - decoded\n%v", result)
	//ToDo - adjust field's verification

	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
	assert.DeepEqual(t, testListExtensible1.GetValue(), result.GetValue())

}
