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

func createTestEnum1() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM1
}

func createTestEnum2() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM2
}

func createTestEnum3() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM3
}

func createTestEnum4() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM4
}

func createTestEnum5() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM5
}

func createTestEnum6() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM6
}

func Test_xerEncodingTestEnumerated(t *testing.T) {

	enum1 := createTestEnum1()

	xer1, err := xerEncodeTestEnumerated(&enum1)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum1) XER\n%s", string(xer1))

	result1, err := xerDecodeTestEnumerated(xer1)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum1) XER - decoded\n%v", result1)
	assert.Equal(t, enum1.Number(), result1.Number())

	enum2 := createTestEnum2()

	xer2, err := xerEncodeTestEnumerated(&enum2)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum2) XER\n%s", string(xer2))

	result2, err := xerDecodeTestEnumerated(xer2)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum2) XER - decoded\n%v", result2)
	assert.Equal(t, enum2.Number(), result2.Number())

	enum3 := createTestEnum3()

	xer3, err := xerEncodeTestEnumerated(&enum3)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum3) XER\n%s", string(xer3))

	result3, err := xerDecodeTestEnumerated(xer3)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum3) XER - decoded\n%v", result3)
	assert.Equal(t, enum3.Number(), result3.Number())

	enum4 := createTestEnum4()

	xer4, err := xerEncodeTestEnumerated(&enum4)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum4) XER\n%s", string(xer4))

	result4, err := xerDecodeTestEnumerated(xer4)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum4) XER - decoded\n%v", result4)
	assert.Equal(t, enum4.Number(), result4.Number())

	enum5 := createTestEnum5()

	xer5, err := xerEncodeTestEnumerated(&enum5)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum5) XER\n%s", string(xer5))

	result5, err := xerDecodeTestEnumerated(xer5)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum5) XER - decoded\n%v", result5)
	assert.Equal(t, enum5.Number(), result5.Number())

	enum6 := createTestEnum6()

	xer6, err := xerEncodeTestEnumerated(&enum6)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum6) XER\n%s", string(xer6))

	result6, err := xerDecodeTestEnumerated(xer6)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum6) XER - decoded\n%v", result6)
	assert.Equal(t, enum6.Number(), result6.Number())
}

func Test_perEncodingTestEnumerated(t *testing.T) {

	enum1 := createTestEnum1()

	per1, err := perEncodeTestEnumerated(&enum1)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum1) PER\n%v", hex.Dump(per1))

	result1, err := perDecodeTestEnumerated(per1)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum1) PER - decoded\n%v", result1)
	assert.Equal(t, enum1.Number(), result1.Number())

	enum2 := createTestEnum2()

	per2, err := perEncodeTestEnumerated(&enum2)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum2) PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestEnumerated(per2)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum2) PER - decoded\n%v", result2)
	assert.Equal(t, enum2.Number(), result2.Number())

	enum3 := createTestEnum3()

	per3, err := perEncodeTestEnumerated(&enum3)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum3) PER\n%v", hex.Dump(per3))

	result3, err := perDecodeTestEnumerated(per3)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum3) XER - decoded\n%v", result3)
	assert.Equal(t, enum3.Number(), result3.Number())

	enum4 := createTestEnum4()

	per4, err := perEncodeTestEnumerated(&enum4)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum4) PER\n%v", hex.Dump(per4))

	result4, err := perDecodeTestEnumerated(per4)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum4) PER - decoded\n%v", result4)
	assert.Equal(t, enum4.Number(), result4.Number())

	enum5 := createTestEnum5()

	per5, err := perEncodeTestEnumerated(&enum5)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum5) PER\n%v", hex.Dump(per5))

	result5, err := perDecodeTestEnumerated(per5)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum5) PER - decoded\n%v", result5)
	assert.Equal(t, enum5.Number(), result5.Number())

	enum6 := createTestEnum6()

	per6, err := perEncodeTestEnumerated(&enum6)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum6) PER\n%v", hex.Dump(per6))

	result6, err := perDecodeTestEnumerated(per6)
	assert.NilError(t, err)
	t.Logf("TestEnumerated (Enum6) PER - decoded\n%v", result6)
	assert.Equal(t, enum6.Number(), result6.Number())
}
