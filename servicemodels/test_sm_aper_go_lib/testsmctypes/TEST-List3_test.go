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

func createTestList3Msg0Items() (*test_sm_ies.TestList3, error) {

	testList3 := test_sm_ies.TestList3{
		Value: make([]*test_sm_ies.TestFullyOptionalSequence, 0),
	}

	return &testList3, nil
}

func createTestList3Msg3Items() (*test_sm_ies.TestList3, error) {

	testList3 := test_sm_ies.TestList3{
		Value: make([]*test_sm_ies.TestFullyOptionalSequence, 0),
	}

	var ie11 int32 = 153
	var ie12 = []byte{0x02, 0x3F, 0x5D, 0x9A}
	var ie13 bool = true
	var ie14 int32 = 584236
	var ie15 int32 = -654

	item1 := test_sm_ies.TestFullyOptionalSequence{
		Item1: &ie11,
		Item2: ie12,
		Item3: &ie13,
		Item4: &ie14,
		Item5: &ie15,
	}

	testList3.Value = append(testList3.Value, &item1)

	item2 := test_sm_ies.TestFullyOptionalSequence{
		//Item1: &ie11,
		//Item2: ie12,
		//Item3: &ie13,
		//Item4: &ie14,
		//Item5: &ie15,
	}
	testList3.Value = append(testList3.Value, &item2)

	var ie32 = []byte{0xC2, 0xF3, 0xD3, 0x9A}
	var ie33 bool = true
	var ie35 int32 = -153

	item3 := test_sm_ies.TestFullyOptionalSequence{
		//Item1: &ie11,
		Item2: ie32,
		Item3: &ie33,
		//Item4: &ie14,
		Item5: &ie35,
	}
	testList3.Value = append(testList3.Value, &item3)

	//if err := testList3.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating TestList3 %s", err.Error())
	//}
	return &testList3, nil
}

func Test_xerEncodingTestList3(t *testing.T) {

	testList31, err := createTestList3Msg0Items()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	xer, err := xerEncodeTestList3(testList31)
	assert.NilError(t, err)
	t.Logf("TestList3 XER\n%s", string(xer))

	result, err := xerDecodeTestList3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList3 XER - decoded\n%v", result)

	testList32, err := createTestList3Msg3Items()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	xer2, err := xerEncodeTestList3(testList32)
	assert.NilError(t, err)
	t.Logf("TestList3 XER\n%s", string(xer2))

	result2, err := xerDecodeTestList3(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList3 XER - decoded\n%v", result2)
	assert.Equal(t, 3, len(result2.GetValue()))
	assert.Equal(t, testList32.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList32.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testList32.GetValue()[0].GetItem3(), result2.GetValue()[0].GetItem3())
	assert.Equal(t, testList32.GetValue()[0].GetItem4(), result2.GetValue()[0].GetItem4())
	assert.Equal(t, testList32.GetValue()[0].GetItem5(), result2.GetValue()[0].GetItem5())
	assert.DeepEqual(t, testList32.GetValue()[3].GetItem2(), result2.GetValue()[3].GetItem2())
	assert.Equal(t, testList32.GetValue()[3].GetItem3(), result2.GetValue()[3].GetItem3())
	assert.Equal(t, testList32.GetValue()[3].GetItem5(), result2.GetValue()[3].GetItem5())
}

func Test_perEncodingTestList3(t *testing.T) {

	testList31, err := createTestList3Msg0Items()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	per, err := perEncodeTestList3(testList31)
	assert.NilError(t, err)
	t.Logf("TestList3 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList3 PER - decoded\n%v", result)

	testList32, err := createTestList3Msg3Items()
	assert.NilError(t, err, "Error creating TestList3 PDU")

	per2, err := perEncodeTestList3(testList32)
	assert.NilError(t, err)
	t.Logf("TestList3 PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestList3(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList3 PER - decoded\n%v", result2)
	assert.Equal(t, 3, len(result2.GetValue()))
	assert.Equal(t, testList32.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList32.GetValue()[0].GetItem2(), result2.GetValue()[0].GetItem2())
	assert.Equal(t, testList32.GetValue()[0].GetItem3(), result2.GetValue()[0].GetItem3())
	assert.Equal(t, testList32.GetValue()[0].GetItem4(), result2.GetValue()[0].GetItem4())
	assert.Equal(t, testList32.GetValue()[0].GetItem5(), result2.GetValue()[0].GetItem5())
	assert.DeepEqual(t, testList32.GetValue()[3].GetItem2(), result2.GetValue()[3].GetItem2())
	assert.Equal(t, testList32.GetValue()[3].GetItem3(), result2.GetValue()[3].GetItem3())
	assert.Equal(t, testList32.GetValue()[3].GetItem5(), result2.GetValue()[3].GetItem5())
}
