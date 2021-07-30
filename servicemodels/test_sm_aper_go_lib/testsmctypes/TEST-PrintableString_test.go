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

func createTestPrintableStringMsg() (*test_sm_ies.TestPrintableString, error) {

	optStr := "omgls"
	testPrintableString := test_sm_ies.TestPrintableString{
		AttrPs1: "Yay",
		AttrPs2: "on",
		AttrPs3: "onONonON",
		AttrPs4: "abc", // it doesn't like anything less than 3 chars
		AttrPs5: "ONF", // It doesn't like anything less than 3 chars again...
		AttrPs6: "X1N", // It doesn't like anything less than 3 chars again... and again...
		AttrPs7: &optStr,
	}

	return &testPrintableString, nil
}

func createTestPrintableStringExcludeOptional() (*test_sm_ies.TestPrintableString, error) {

	//optStr := "omgls"
	testPrintableString := test_sm_ies.TestPrintableString{
		AttrPs1: "Yay",
		AttrPs2: "on",
		AttrPs3: "on",
		AttrPs4: "abc",    // it doesn't like anything less than 3 chars
		AttrPs5: "ONF",    // It doesn't like anything less than 3 chars again...
		AttrPs6: "X1N5fg", // It doesn't like anything less than 3 chars again... and again...
		//AttrPs7: &optStr,
	}

	return &testPrintableString, nil
}

func Test_xerEncodingTestPrintableString(t *testing.T) {

	testPrintableString, err := createTestPrintableStringMsg()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	xer, err := xerEncodeTestPrintableString(testPrintableString)
	assert.NilError(t, err)
	t.Logf("TestPrintableString XER\n%s", string(xer))

	result, err := xerDecodeTestPrintableString(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestPrintableString XER - decoded\n%v", result)
	assert.Equal(t, testPrintableString.GetAttrPs1(), result.GetAttrPs1())
	assert.Equal(t, testPrintableString.GetAttrPs2(), result.GetAttrPs2())
	assert.Equal(t, testPrintableString.GetAttrPs3(), result.GetAttrPs3())
	assert.Equal(t, testPrintableString.GetAttrPs4(), result.GetAttrPs4())
	assert.Equal(t, testPrintableString.GetAttrPs5(), result.GetAttrPs5())
	assert.Equal(t, testPrintableString.GetAttrPs6(), result.GetAttrPs6())
	assert.Equal(t, testPrintableString.GetAttrPs7(), result.GetAttrPs7())

	testPrintableStringExcludeOptional, err := createTestPrintableStringExcludeOptional()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	xer2, err := xerEncodeTestPrintableString(testPrintableStringExcludeOptional)
	assert.NilError(t, err)
	t.Logf("TestPrintableString XER\n%s", string(xer2))

	result2, err := xerDecodeTestPrintableString(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestPrintableString XER - decoded\n%v", result2)
	assert.Equal(t, testPrintableString.GetAttrPs1(), result2.GetAttrPs1())
	assert.Equal(t, testPrintableString.GetAttrPs2(), result2.GetAttrPs2())
	assert.Equal(t, testPrintableString.GetAttrPs3(), result2.GetAttrPs3())
	assert.Equal(t, testPrintableString.GetAttrPs4(), result2.GetAttrPs4())
	assert.Equal(t, testPrintableString.GetAttrPs5(), result2.GetAttrPs5())
	assert.Equal(t, testPrintableString.GetAttrPs6(), result2.GetAttrPs6())
}

func Test_perEncodingTestPrintableString(t *testing.T) {

	testPrintableString, err := createTestPrintableStringMsg()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	per, err := perEncodeTestPrintableString(testPrintableString)
	assert.NilError(t, err)
	t.Logf("TestPrintableString PER\n%v", hex.Dump(per))

	result, err := perDecodeTestPrintableString(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestPrintableString PER - decoded\n%v", result)
	assert.Equal(t, testPrintableString.GetAttrPs1(), result.GetAttrPs1())
	assert.Equal(t, testPrintableString.GetAttrPs2(), result.GetAttrPs2())
	assert.Equal(t, testPrintableString.GetAttrPs3(), result.GetAttrPs3())
	assert.Equal(t, testPrintableString.GetAttrPs4(), result.GetAttrPs4())
	assert.Equal(t, testPrintableString.GetAttrPs5(), result.GetAttrPs5())
	assert.Equal(t, testPrintableString.GetAttrPs6(), result.GetAttrPs6())
	assert.Equal(t, testPrintableString.GetAttrPs7(), result.GetAttrPs7())

	testPrintableStringExcludeOptional, err := createTestPrintableStringExcludeOptional()
	assert.NilError(t, err, "Error creating TestPrintableString PDU")

	per2, err := perEncodeTestPrintableString(testPrintableStringExcludeOptional)
	assert.NilError(t, err)
	t.Logf("TestPrintableString PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestPrintableString(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestPrintableString PER - decoded\n%v", result2)
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs1(), result2.GetAttrPs1())
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs2(), result2.GetAttrPs2())
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs3(), result2.GetAttrPs3())
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs4(), result2.GetAttrPs4())
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs5(), result2.GetAttrPs5())
	assert.Equal(t, testPrintableStringExcludeOptional.GetAttrPs6(), result2.GetAttrPs6())
}
