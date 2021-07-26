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

func createTestOctetStringMsg() (*test_sm_ies.TestOctetString, error) {

	testOctetString := test_sm_ies.TestOctetString{
		AttrOs1: []byte{0x12, 0x34, 0x56, 0xA4},
		AttrOs2: []byte{0xFF, 0xFF},
		AttrOs3: []byte{0xFF, 0xFF},
		AttrOs4: []byte{},
		AttrOs5: []byte{0xDE, 0xC7, 0x23},
		AttrOs6: []byte{0x01},
		AttrOs7: []byte{0x21, 0x44, 0xA8, 0xDF, 0x11},
	}

	return &testOctetString, nil
}

func createTestOctetStringExcludeOptional() (*test_sm_ies.TestOctetString, error) {

	testOctetString := test_sm_ies.TestOctetString{
		AttrOs1: []byte{0x12, 0x34, 0x56, 0xA4},
		AttrOs2: []byte{0xFF, 0xFF},
		AttrOs3: []byte{0xFF, 0xFF},
		AttrOs4: []byte{},
		AttrOs5: []byte{0xDE, 0xC7, 0x23},
		AttrOs6: []byte{0x01},
		//AttrOs7: []byte{0x21, 0x44, 0xA8, 0xDF, 0x11},
	}

	return &testOctetString, nil
}

func Test_xerEncodingTestOctetString(t *testing.T) {

	testOctetString, err := createTestOctetStringMsg()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	xer, err := xerEncodeTestOctetString(testOctetString)
	assert.NilError(t, err)
	t.Logf("TestOctetString XER\n%s", string(xer))

	result, err := xerDecodeTestOctetString(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestOctetString XER - decoded\n%v", result)
	assert.DeepEqual(t, testOctetString.GetAttrOs1(), result.GetAttrOs1())
	assert.DeepEqual(t, testOctetString.GetAttrOs2(), result.GetAttrOs2())
	assert.DeepEqual(t, testOctetString.GetAttrOs3(), result.GetAttrOs3())
	assert.DeepEqual(t, testOctetString.GetAttrOs4(), result.GetAttrOs4())
	assert.DeepEqual(t, testOctetString.GetAttrOs5(), result.GetAttrOs5())
	assert.DeepEqual(t, testOctetString.GetAttrOs6(), result.GetAttrOs6())
	assert.DeepEqual(t, testOctetString.GetAttrOs7(), result.GetAttrOs7())

	testOctetStringExcludeOptional, err := createTestOctetStringExcludeOptional()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	xer2, err := xerEncodeTestOctetString(testOctetStringExcludeOptional)
	assert.NilError(t, err)
	t.Logf("TestOctetString XER\n%s", string(xer2))

	result2, err := xerDecodeTestOctetString(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestOctetString XER - decoded\n%v", result2)
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs1(), result2.GetAttrOs1())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs2(), result2.GetAttrOs2())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs3(), result2.GetAttrOs3())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs4(), result2.GetAttrOs4())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs5(), result2.GetAttrOs5())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs6(), result2.GetAttrOs6())
}

func Test_perEncodingTestOctetString(t *testing.T) {

	testOctetString, err := createTestOctetStringMsg()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	per, err := perEncodeTestOctetString(testOctetString)
	assert.NilError(t, err)
	t.Logf("TestOctetString PER\n%v", hex.Dump(per))

	result, err := perDecodeTestOctetString(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestOctetString PER - decoded\n%v", result)
	assert.DeepEqual(t, testOctetString.GetAttrOs1(), result.GetAttrOs1())
	assert.DeepEqual(t, testOctetString.GetAttrOs2(), result.GetAttrOs2())
	assert.DeepEqual(t, testOctetString.GetAttrOs3(), result.GetAttrOs3())
	assert.DeepEqual(t, testOctetString.GetAttrOs4(), result.GetAttrOs4())
	assert.DeepEqual(t, testOctetString.GetAttrOs5(), result.GetAttrOs5())
	assert.DeepEqual(t, testOctetString.GetAttrOs6(), result.GetAttrOs6())
	assert.DeepEqual(t, testOctetString.GetAttrOs7(), result.GetAttrOs7())

	testOctetStringExcludeOptional, err := createTestOctetStringExcludeOptional()
	assert.NilError(t, err, "Error creating TestOctetString PDU")

	per2, err := perEncodeTestOctetString(testOctetStringExcludeOptional)
	assert.NilError(t, err)
	t.Logf("TestOctetString PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestOctetString(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestOctetString PER - decoded\n%v", result2)
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs1(), result2.GetAttrOs1())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs2(), result2.GetAttrOs2())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs3(), result2.GetAttrOs3())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs4(), result2.GetAttrOs4())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs5(), result2.GetAttrOs5())
	assert.DeepEqual(t, testOctetStringExcludeOptional.GetAttrOs6(), result2.GetAttrOs6())
}
