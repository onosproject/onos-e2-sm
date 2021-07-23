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

func createTestOctetStringMsg() (*test_sm.TestOctetString, error) {

	// testOctetString := pdubuilder.CreateTestOctetString() //ToDo - fill in arguments here(if this function exists

	testOctetString := test_sm.TestOctetString{
		AttrOs1: nil,
		AttrOs2: nil,
		AttrOs3: nil,
		AttrOs4: nil,
		AttrOs5: nil,
		AttrOs6: nil,
		AttrOs7: nil,
	}

	if err := testOctetString.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestOctetString %s", err.Error())
	}
	return &testOctetString, nil
}

func Test_xerEncodingTestOctetString(t *testing.T) {

	testOctetString, err := createTestOctetStringMsg()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	xer, err := xerEncodeTestOctetString(testOctetString)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestOctetString XER\n%s", string(xer))

	result, err := xerDecodeTestOctetString(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestOctetString XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testOctetString.GetAttrOs1(), result.GetAttrOs1())
	assert.Equal(t, testOctetString.GetAttrOs2(), result.GetAttrOs2())
	assert.Equal(t, testOctetString.GetAttrOs3(), result.GetAttrOs3())
	assert.Equal(t, testOctetString.GetAttrOs4(), result.GetAttrOs4())
	assert.Equal(t, testOctetString.GetAttrOs5(), result.GetAttrOs5())
	assert.Equal(t, testOctetString.GetAttrOs6(), result.GetAttrOs6())
	assert.Equal(t, testOctetString.GetAttrOs7(), result.GetAttrOs7())

}

func Test_perEncodingTestOctetString(t *testing.T) {

	testOctetString, err := createTestOctetStringMsg()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	per, err := perEncodeTestOctetString(testOctetString)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestOctetString PER\n%v", hex.Dump(per))

	result, err := perDecodeTestOctetString(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestOctetString PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testOctetString.GetAttrOs1(), result.GetAttrOs1())
	assert.Equal(t, testOctetString.GetAttrOs2(), result.GetAttrOs2())
	assert.Equal(t, testOctetString.GetAttrOs3(), result.GetAttrOs3())
	assert.Equal(t, testOctetString.GetAttrOs4(), result.GetAttrOs4())
	assert.Equal(t, testOctetString.GetAttrOs5(), result.GetAttrOs5())
	assert.Equal(t, testOctetString.GetAttrOs6(), result.GetAttrOs6())
	assert.Equal(t, testOctetString.GetAttrOs7(), result.GetAttrOs7())

}
