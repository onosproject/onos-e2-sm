// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func createTestList1Msg0Items() (*test_sm_ies.TestList1, error) {

	testList1 := test_sm_ies.TestList1{
		Value: make([]*test_sm_ies.Item, 0),
	}

	return &testList1, nil
}

func createTestList1Msg12Items() (*test_sm_ies.TestList1, error) {

	testList1 := test_sm_ies.TestList1{
		Value: make([]*test_sm_ies.Item, 0),
	}
	var ie11 int32 = 32
	item1 := &test_sm_ies.Item{
		Item1: &ie11,
		Item2: &asn1.BitString{
			Value: []byte{0xF0},
			Len:   4,
		},
	}
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)
	testList1.Value = append(testList1.Value, item1)

	return &testList1, nil
}

func Test_xerEncodingTestList1(t *testing.T) {

	testList11, err := createTestList1Msg0Items()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	xer, err := xerEncodeTestList1(testList11)
	assert.NilError(t, err)
	t.Logf("TestList1 (0 items) XER\n%s", string(xer))

	result, err := xerDecodeTestList1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList1 (0 items) XER - decoded\n%v", result)

	testList12, err := createTestList1Msg12Items()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	xer2, err := xerEncodeTestList1(testList12)
	assert.NilError(t, err)
	t.Logf("TestList1 (12 items) XER\n%s", string(xer2))

	result2, err := xerDecodeTestList1(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList1 (12 items) XER - decoded\n%v", result2)
	assert.Equal(t, 12, len(result2.GetValue()))
	assert.Equal(t, testList12.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[0].GetItem2().GetValue(), result2.GetValue()[0].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[0].GetItem2().GetLen(), result2.GetValue()[0].GetItem2().GetLen())
	assert.Equal(t, testList12.GetValue()[5].GetItem1(), result2.GetValue()[5].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[5].GetItem2().GetValue(), result2.GetValue()[5].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[5].GetItem2().GetLen(), result2.GetValue()[5].GetItem2().GetLen())
	assert.Equal(t, testList12.GetValue()[11].GetItem1(), result2.GetValue()[11].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[11].GetItem2().GetValue(), result2.GetValue()[11].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[11].GetItem2().GetLen(), result2.GetValue()[11].GetItem2().GetLen())
}

func Test_perEncodingTestList1(t *testing.T) {

	testList11, err := createTestList1Msg0Items()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	per, err := perEncodeTestList1(testList11)
	assert.NilError(t, err)
	t.Logf("TestList1 (0 items) PER\n%v", hex.Dump(per))

	result, err := perDecodeTestList1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestList1 (0 items) PER - decoded\n%v", result)

	testList12, err := createTestList1Msg12Items()
	assert.NilError(t, err, "Error creating TestList1 PDU")

	per2, err := perEncodeTestList1(testList12)
	assert.NilError(t, err)
	t.Logf("TestList1 (12 items) PER\n%v", hex.Dump(per2))

	result2, err := perDecodeTestList1(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("TestList1 (12 items) PER - decoded\n%v", result2)
	assert.Equal(t, 12, len(result2.GetValue()))
	assert.Equal(t, testList12.GetValue()[0].GetItem1(), result2.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[0].GetItem2().GetValue(), result2.GetValue()[0].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[0].GetItem2().GetLen(), result2.GetValue()[0].GetItem2().GetLen())
	assert.Equal(t, testList12.GetValue()[5].GetItem1(), result2.GetValue()[5].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[5].GetItem2().GetValue(), result2.GetValue()[5].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[5].GetItem2().GetLen(), result2.GetValue()[5].GetItem2().GetLen())
	assert.Equal(t, testList12.GetValue()[11].GetItem1(), result2.GetValue()[11].GetItem1())
	assert.DeepEqual(t, testList12.GetValue()[11].GetItem2().GetValue(), result2.GetValue()[11].GetItem2().GetValue())
	assert.Equal(t, testList12.GetValue()[11].GetItem2().GetLen(), result2.GetValue()[11].GetItem2().GetLen())
}
