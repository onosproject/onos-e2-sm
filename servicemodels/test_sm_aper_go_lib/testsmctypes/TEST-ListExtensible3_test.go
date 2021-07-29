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

func createTestListExtensible3Msg0Items() (*test_sm_ies.TestListExtensible3, error) {

	testListExtensible3 := test_sm_ies.TestListExtensible3{}

	return &testListExtensible3, nil
}

func createTestListExtensible3Msg1EmptyItem() (*test_sm_ies.TestListExtensible3, error) {

	testListExtensible3 := test_sm_ies.TestListExtensible3{
		Value: make([]*test_sm_ies.TestFullyOptionalSequence, 0),
	}

	item := test_sm_ies.TestFullyOptionalSequence{}
	testListExtensible3.Value = append(testListExtensible3.Value, &item)

	return &testListExtensible3, nil
}

func createTestListExtensible3MsgFull() (*test_sm_ies.TestListExtensible3, error) {

	testListExtensible3 := test_sm_ies.TestListExtensible3{
		Value: make([]*test_sm_ies.TestFullyOptionalSequence, 0),
	}

	var ie11 int32 = 153
	var ie12 = []byte{0x02, 0x3F, 0x5D, 0x9A}
	var ie13 bool = true
	ie14 := test_sm_ies.TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE
	var ie15 int32 = 0

	item1 := test_sm_ies.TestFullyOptionalSequence{
		Item1: &ie11,
		Item2: ie12,
		Item3: &ie13,
		Item4: &ie14,
		Item5: &ie15,
	}

	testListExtensible3.Value = append(testListExtensible3.Value, &item1)

	item2 := test_sm_ies.TestFullyOptionalSequence{}
	testListExtensible3.Value = append(testListExtensible3.Value, &item2)

	var ie32 = []byte{0xC2, 0xF3, 0xD3, 0x9A}
	var ie33 bool = true
	var ie35 int32 = 0

	item3 := test_sm_ies.TestFullyOptionalSequence{
		Item2: ie32,
		Item3: &ie33,
		Item5: &ie35,
	}
	testListExtensible3.Value = append(testListExtensible3.Value, &item3)

	var ie43 bool = false

	item4 := test_sm_ies.TestFullyOptionalSequence{
		Item3: &ie43,
	}
	testListExtensible3.Value = append(testListExtensible3.Value, &item4)

	return &testListExtensible3, nil
}

func Test_xerEncodingTestListExtensible3(t *testing.T) {

	testListExtensible3, err := createTestListExtensible3Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	xer, err := xerEncodeTestListExtensible3(testListExtensible3)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible3 XER - decoded\n%v", result)

	testListExtensible31, err := createTestListExtensible3Msg1EmptyItem()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	xer1, err := xerEncodeTestListExtensible3(testListExtensible31)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 XER\n%s", string(xer1))

	result1, err := xerDecodeTestListExtensible3(xer1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestListExtensible3 XER - decoded\n%v", result1)
	assert.Equal(t, 1, len(result1.GetValue()))

	testListExtensible32, err := createTestListExtensible3MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	xer2, err := xerEncodeTestListExtensible3(testListExtensible32)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 XER\n%s", string(xer2))

	result2, err := xerDecodeTestListExtensible3(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestListExtensible3 XER - decoded\n%v", result2)
	assert.Equal(t, 4, len(result2.GetValue()))
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible32.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem3(), result2.GetValue()[0].GetItem3())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem4(), result2.GetValue()[0].GetItem4())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem5(), result2.GetValue()[0].GetItem5())
	assert.DeepEqual(t, testListExtensible32.GetValue()[2].GetItem2(), result2.GetValue()[2].GetItem2())
	assert.Equal(t, testListExtensible32.GetValue()[2].GetItem3(), result2.GetValue()[2].GetItem3())
	assert.Equal(t, testListExtensible32.GetValue()[2].GetItem5(), result2.GetValue()[2].GetItem5())
	assert.Equal(t, testListExtensible32.GetValue()[3].GetItem4(), result2.GetValue()[3].GetItem4())
}

func Test_perEncodingTestListExtensible3(t *testing.T) {

	testListExtensible3, err := createTestListExtensible3Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	per, err := perEncodeTestListExtensible3(testListExtensible3)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible3 PER - decoded\n%v", result)

	testListExtensible31, err := createTestListExtensible3Msg1EmptyItem()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	per1, err := perEncodeTestListExtensible3(testListExtensible31)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 PER\n%v", hex.Dump(per1))

	result1, err := perDecodeTestListExtensible3(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestListExtensible3 PER - decoded\n%v", result1)
	assert.Equal(t, 1, len(result1.GetValue()))

	testListExtensible32, err := createTestListExtensible3MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible3 PDU")

	per2, err := perEncodeTestListExtensible3(testListExtensible32)
	assert.NilError(t, err)
	t.Logf("TestListExtensible3 PER\n%v", hex.Dump(per1))

	result2, err := perDecodeTestListExtensible3(per2)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestListExtensible3 PER - decoded\n%v", result2)
	assert.Equal(t, 4, len(result2.GetValue()))
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible32.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem3(), result2.GetValue()[0].GetItem3())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem4(), result2.GetValue()[0].GetItem4())
	assert.Equal(t, testListExtensible32.GetValue()[0].GetItem5(), result2.GetValue()[0].GetItem5())
	assert.DeepEqual(t, testListExtensible32.GetValue()[2].GetItem2(), result2.GetValue()[2].GetItem2())
	assert.Equal(t, testListExtensible32.GetValue()[2].GetItem3(), result2.GetValue()[2].GetItem3())
	assert.Equal(t, testListExtensible32.GetValue()[2].GetItem5(), result2.GetValue()[2].GetItem5())
	assert.Equal(t, testListExtensible32.GetValue()[3].GetItem4(), result2.GetValue()[3].GetItem4())
}
