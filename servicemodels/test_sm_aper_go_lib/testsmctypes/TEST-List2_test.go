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

func createTestList2Msg0Items() (*test_sm_ies.TestList2, error) {

	testList2 := test_sm_ies.TestList2{
		Value: make([]*test_sm_ies.ItemExtensible, 0),
	}

	return &testList2, nil
}

func createTestList2Msg2Items() (*test_sm_ies.TestList2, error) {

	testList2 := test_sm_ies.TestList2{
		Value: make([]*test_sm_ies.ItemExtensible, 0),
	}

	ie1 := test_sm_ies.ItemExtensible{
		Item1: 153,
		Item2: []byte{0xFF, 0xBD, 0x85},
	}

	ie2 := test_sm_ies.ItemExtensible{
		Item1: 215698,
		Item2: []byte{0x21, 0xBA, 0xDF, 0x17, 0x56, 0xFD, 0x3C},
	}

	testList2.Value = append(testList2.Value, &ie1)
	testList2.Value = append(testList2.Value, &ie2)

	//if err := testList2.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating TestList2 %s", err.Error())
	//}
	return &testList2, nil
}

func Test_xerEncodingTestList2(t *testing.T) {

	testList21, err := createTestList2Msg0Items()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	xer, err := xerEncodeTestList2(testList21)
	assert.NilError(t, err)
	t.Logf("TestList2 XER\n%s", string(xer))

	result, err := xerDecodeTestList2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList2 XER - decoded\n%v", result)

	testList22, err := createTestList2Msg2Items()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	xer2, err := xerEncodeTestList2(testList22)
	assert.NilError(t, err)
	t.Logf("TestList2 XER\n%s", string(xer2))

	result2, err := xerDecodeTestList2(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList2 XER - decoded\n%v", result2)
	assert.Equal(t, 2, len(result2.GetValue()))
	assert.DeepEqual(t, testList22.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testList22.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList22.GetValue()[1].GetItem2(), result2.GetValue()[1].GetItem2())
	assert.Equal(t, testList22.GetValue()[1].GetItem1(), result2.GetValue()[1].GetItem1())
}

func Test_perEncodingTestList2(t *testing.T) {

	testList21, err := createTestList2Msg0Items()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	per, err := perEncodeTestList2(testList21)
	assert.NilError(t, err)
	t.Logf("TestList2 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList2 PER - decoded\n%v", result)

	testList22, err := createTestList2Msg2Items()
	assert.NilError(t, err, "Error creating TestList2 PDU")

	per2, err := perEncodeTestList2(testList22)
	assert.NilError(t, err)
	t.Logf("TestList2 PER\n%v", hex.Dump(per))

	result2, err := perDecodeTestList2(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList2 PER - decoded\n%v", result2)
	assert.Equal(t, 2, len(result2.GetValue()))
	assert.DeepEqual(t, testList22.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testList22.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList22.GetValue()[1].GetItem2(), result2.GetValue()[1].GetItem2())
	assert.Equal(t, testList22.GetValue()[1].GetItem1(), result2.GetValue()[1].GetItem1())
}
