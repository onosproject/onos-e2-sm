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

func createTestPrintableStringMsg() (*test_sm.TestPrintableString, error) {

	// testPrintableString := pdubuilder.CreateTestPrintableString() //ToDo - fill in arguments here(if this function exists

	testPrintableString := test_sm.TestPrintableString{
		AttrPs1: nil,
		AttrPs2: nil,
		AttrPs3: nil,
		AttrPs4: nil,
		AttrPs5: nil,
		AttrPs6: nil,
		AttrPs7: nil,
	}

	if err := testPrintableString.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestPrintableString %s", err.Error())
	}
	return &testPrintableString, nil
}

func Test_xerEncodingTestPrintableString(t *testing.T) {

	testPrintableString, err := createTestPrintableStringMsg()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	xer, err := xerEncodeTestPrintableString(testPrintableString)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestPrintableString XER\n%s", string(xer))

	result, err := xerDecodeTestPrintableString(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestPrintableString XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testPrintableString.GetAttrPs1(), result.GetAttrPs1())
	assert.Equal(t, testPrintableString.GetAttrPs2(), result.GetAttrPs2())
	assert.Equal(t, testPrintableString.GetAttrPs3(), result.GetAttrPs3())
	assert.Equal(t, testPrintableString.GetAttrPs4(), result.GetAttrPs4())
	assert.Equal(t, testPrintableString.GetAttrPs5(), result.GetAttrPs5())
	assert.Equal(t, testPrintableString.GetAttrPs6(), result.GetAttrPs6())
	assert.Equal(t, testPrintableString.GetAttrPs7(), result.GetAttrPs7())

}

func Test_perEncodingTestPrintableString(t *testing.T) {

	testPrintableString, err := createTestPrintableStringMsg()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	per, err := perEncodeTestPrintableString(testPrintableString)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestPrintableString PER\n%v", hex.Dump(per))

	result, err := perDecodeTestPrintableString(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestPrintableString PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testPrintableString.GetAttrPs1(), result.GetAttrPs1())
	assert.Equal(t, testPrintableString.GetAttrPs2(), result.GetAttrPs2())
	assert.Equal(t, testPrintableString.GetAttrPs3(), result.GetAttrPs3())
	assert.Equal(t, testPrintableString.GetAttrPs4(), result.GetAttrPs4())
	assert.Equal(t, testPrintableString.GetAttrPs5(), result.GetAttrPs5())
	assert.Equal(t, testPrintableString.GetAttrPs6(), result.GetAttrPs6())
	assert.Equal(t, testPrintableString.GetAttrPs7(), result.GetAttrPs7())

}
