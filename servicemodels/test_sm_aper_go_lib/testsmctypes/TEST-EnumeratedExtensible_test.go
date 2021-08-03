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

func createTestEnum1Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM1
}

func createTestEnum2Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM2
}

func createTestEnum3Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM3
}

func createTestEnum4Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM4
}

func createTestEnum5Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM5
}

func createTestEnum6Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6
}

func Test_xerEncodingTestEnumeratedExtensible(t *testing.T) {

	enum1 := createTestEnum1Ext()

	xer1, err := xerEncodeTestEnumeratedExtensible(&enum1)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum1) XER\n%s", string(xer1))

	result1, err := xerDecodeTestEnumeratedExtensible(xer1)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum1) XER - decoded\n%v", result1)
	assert.Equal(t, enum1.Number(), result1.Number())

	enum2 := createTestEnum2Ext()

	xer2, err := xerEncodeTestEnumeratedExtensible(&enum2)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum2) XER\n%s", string(xer2))

	result2, err := xerDecodeTestEnumeratedExtensible(xer2)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum2) XER - decoded\n%v", result2)
	assert.Equal(t, enum2.Number(), result2.Number())

	enum3 := createTestEnum3Ext()

	xer3, err := xerEncodeTestEnumeratedExtensible(&enum3)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum3) XER\n%s", string(xer3))

	result3, err := xerDecodeTestEnumeratedExtensible(xer3)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum3) XER - decoded\n%v", result3)
	assert.Equal(t, enum3.Number(), result3.Number())

	enum4 := createTestEnum4Ext()

	xer4, err := xerEncodeTestEnumeratedExtensible(&enum4)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum4) XER\n%s", string(xer4))

	result4, err := xerDecodeTestEnumeratedExtensible(xer4)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum4) XER - decoded\n%v", result4)
	assert.Equal(t, enum4.Number(), result4.Number())

	enum5 := createTestEnum5Ext()

	xer5, err := xerEncodeTestEnumeratedExtensible(&enum5)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum5) XER\n%s", string(xer5))

	result5, err := xerDecodeTestEnumeratedExtensible(xer5)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum5) XER - decoded\n%v", result5)
	assert.Equal(t, enum5.Number(), result5.Number())

	enum6 := createTestEnum6Ext()

	xer6, err := xerEncodeTestEnumeratedExtensible(&enum6)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum6) XER\n%s", string(xer6))

	result6, err := xerDecodeTestEnumeratedExtensible(xer6)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum6) XER - decoded\n%v", result6)
	assert.Equal(t, enum6.Number(), result6.Number())
}

func Test_perEncodingTestEnumeratedExtensible(t *testing.T) {

	enum1 := createTestEnum1Ext()

	per1, err := perEncodeTestEnumeratedExtensible(&enum1)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum1) PER\n%v", hex.Dump(per1))

	result1, err := perDecodeTestEnumeratedExtensible(per1)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum1) PER - decoded\n%v", result1)
	assert.Equal(t, enum1.Number(), result1.Number())

	enum2 := createTestEnum2Ext()

	per2, err := perEncodeTestEnumeratedExtensible(&enum2)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum2) PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestEnumeratedExtensible(per2)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum2) PER - decoded\n%v", result2)
	assert.Equal(t, enum2.Number(), result2.Number())

	enum3 := createTestEnum3Ext()

	per3, err := perEncodeTestEnumeratedExtensible(&enum3)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum3) PER\n%v", hex.Dump(per3))

	result3, err := perDecodeTestEnumeratedExtensible(per3)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum3) XER - decoded\n%v", result3)
	assert.Equal(t, enum3.Number(), result3.Number())

	enum4 := createTestEnum4Ext()

	per4, err := perEncodeTestEnumeratedExtensible(&enum4)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum4) PER\n%v", hex.Dump(per4))

	result4, err := perDecodeTestEnumeratedExtensible(per4)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum4) PER - decoded\n%v", result4)
	assert.Equal(t, enum4.Number(), result4.Number())

	enum5 := createTestEnum5Ext()

	per5, err := perEncodeTestEnumeratedExtensible(&enum5)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum5) PER\n%v", hex.Dump(per5))

	result5, err := perDecodeTestEnumeratedExtensible(per5)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum5) PER - decoded\n%v", result5)
	assert.Equal(t, enum5.Number(), result5.Number())

	enum6 := createTestEnum6Ext()

	per6, err := perEncodeTestEnumeratedExtensible(&enum6)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum6) PER\n%v", hex.Dump(per6))

	result6, err := perDecodeTestEnumeratedExtensible(per6)
	assert.NilError(t, err)
	t.Logf("TestEnumeratedExtensible (Enum6) PER - decoded\n%v", result6)
	assert.Equal(t, enum6.Number(), result6.Number())
}
