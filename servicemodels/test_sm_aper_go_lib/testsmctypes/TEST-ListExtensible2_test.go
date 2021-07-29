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

func createTestListExtensible2Msg0Items() (*test_sm_ies.TestListExtensible2, error) {

	testListExtensible2 := test_sm_ies.TestListExtensible2{}

	return &testListExtensible2, nil
}

func createTestListExtensible2MsgFull() (*test_sm_ies.TestListExtensible2, error) {

	testListExtensible2 := test_sm_ies.TestListExtensible2{
		Value: make([]*test_sm_ies.ItemExtensible, 0),
	}

	ie1 := test_sm_ies.ItemExtensible{
		Item1: 153,
		Item2: []byte{0xFF, 0xBD, 0x85},
	}
	testListExtensible2.Value = append(testListExtensible2.Value, &ie1)

	ie2 := test_sm_ies.ItemExtensible{
		Item1: -2198,
		Item2: []byte{0xBA, 0xDF, 0x56, 0xFD, 0x3C},
	}
	testListExtensible2.Value = append(testListExtensible2.Value, &ie2)

	ie3 := test_sm_ies.ItemExtensible{
		Item1: 215698,
		Item2: []byte{0x21, 0xBA, 0xDF, 0x17, 0x56, 0xFD, 0x3C},
	}
	testListExtensible2.Value = append(testListExtensible2.Value, &ie3)

	ie4 := test_sm_ies.ItemExtensible{
		Item1: 0,
		Item2: []byte{0x21, 0xBA, 0x17, 0x56, 0xFD, 0x3C},
	}
	testListExtensible2.Value = append(testListExtensible2.Value, &ie4)

	return &testListExtensible2, nil
}

func Test_xerEncodingTestListExtensible2(t *testing.T) {

	testListExtensible2, err := createTestListExtensible2Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	xer, err := xerEncodeTestListExtensible2(testListExtensible2)
	assert.NilError(t, err)
	t.Logf("TestListExtensible2 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible2 XER - decoded\n%v", result)

	testListExtensible22, err := createTestListExtensible2MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	xer2, err := xerEncodeTestListExtensible2(testListExtensible22)
	assert.NilError(t, err)
	t.Logf("TestListExtensible2 XER\n%s", string(xer2))

	result2, err := xerDecodeTestListExtensible2(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestListExtensible2 XER - decoded\n%v", result2)
	assert.Equal(t, 4, len(result2.GetValue()))
	assert.Equal(t, testListExtensible22.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[1].GetItem1(), result2.GetValue()[1].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[1].GetItem2(), result2.GetValue()[1].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[2].GetItem1(), result2.GetValue()[2].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[2].GetItem2(), result2.GetValue()[2].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[3].GetItem1(), result2.GetValue()[3].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[3].GetItem2(), result2.GetValue()[3].GetItem2())
}

func Test_perEncodingTestListExtensible2(t *testing.T) {

	testListExtensible2, err := createTestListExtensible2Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	per, err := perEncodeTestListExtensible2(testListExtensible2)
	assert.NilError(t, err)
	t.Logf("TestListExtensible2 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible2 PER - decoded\n%v", result)

	testListExtensible22, err := createTestListExtensible2MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible2 PDU")

	per2, err := perEncodeTestListExtensible2(testListExtensible22)
	assert.NilError(t, err)
	t.Logf("TestListExtensible2 PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestListExtensible2(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestListExtensible2 PER - decoded\n%v", result2)
	assert.Equal(t, 4, len(result2.GetValue()))
	assert.Equal(t, testListExtensible22.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[1].GetItem1(), result2.GetValue()[1].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[1].GetItem2(), result2.GetValue()[1].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[2].GetItem1(), result2.GetValue()[2].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[2].GetItem2(), result2.GetValue()[2].GetItem2())
	assert.Equal(t, testListExtensible22.GetValue()[3].GetItem1(), result2.GetValue()[3].GetItem1())
	assert.DeepEqual(t, testListExtensible22.GetValue()[3].GetItem2(), result2.GetValue()[3].GetItem2())
}
