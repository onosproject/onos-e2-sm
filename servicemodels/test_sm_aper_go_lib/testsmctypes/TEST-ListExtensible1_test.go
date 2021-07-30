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

func createTestListExtensible1Msg0Items() (*test_sm_ies.TestListExtensible1, error) {

	testListExtensible1 := test_sm_ies.TestListExtensible1{}

	return &testListExtensible1, nil
}

func createTestListExtensible1MsgFull() (*test_sm_ies.TestListExtensible1, error) {

	testListExtensible1 := test_sm_ies.TestListExtensible1{
		Value: make([]*test_sm_ies.Item, 0),
	}

	var ie11 int32 = -56

	ie1 := test_sm_ies.Item{
		Item1: &ie11,
		Item2: &asn1.BitString{
			Value: []byte{0xDE},
			Len:   7,
		},
	}
	testListExtensible1.Value = append(testListExtensible1.Value, &ie1)

	ie2 := test_sm_ies.Item{
		Item2: &asn1.BitString{
			Value: []byte{0xAE},
			Len:   7,
		},
	}
	testListExtensible1.Value = append(testListExtensible1.Value, &ie2)

	var ie31 int32 = 56

	ie3 := test_sm_ies.Item{
		Item1: &ie31,
		Item2: &asn1.BitString{
			Value: []byte{0xA8},
			Len:   5,
		},
	}
	testListExtensible1.Value = append(testListExtensible1.Value, &ie3)

	var ie41 int32 = 56065

	ie4 := test_sm_ies.Item{
		Item1: &ie41,
		Item2: &asn1.BitString{
			Value: []byte{0xAA},
			Len:   7,
		},
	}
	testListExtensible1.Value = append(testListExtensible1.Value, &ie4)

	return &testListExtensible1, nil
}

func Test_xerEncodingTestListExtensible1(t *testing.T) {

	testListExtensible1, err := createTestListExtensible1Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	xer, err := xerEncodeTestListExtensible1(testListExtensible1)
	assert.NilError(t, err)
	t.Logf("TestListExtensible1 XER\n%s", string(xer))

	result, err := xerDecodeTestListExtensible1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible1 XER - decoded\n%v", result)

	testListExtensible11, err := createTestListExtensible1MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	xer1, err := xerEncodeTestListExtensible1(testListExtensible11)
	assert.NilError(t, err)
	t.Logf("TestListExtensible1 XER\n%s", string(xer1))

	result1, err := xerDecodeTestListExtensible1(xer1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestListExtensible1 XER - decoded\n%v", result1)
	assert.Equal(t, 4, len(result1.GetValue()))
	assert.Equal(t, testListExtensible11.GetValue()[0].GetItem1(), result1.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[0].GetItem2().GetValue(), result1.GetValue()[0].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[0].GetItem2().GetLen(), result1.GetValue()[0].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[1].GetItem1(), result1.GetValue()[1].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[1].GetItem2().GetValue(), result1.GetValue()[1].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[1].GetItem2().GetLen(), result1.GetValue()[1].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[2].GetItem1(), result1.GetValue()[2].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[2].GetItem2().GetValue(), result1.GetValue()[2].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[2].GetItem2().GetLen(), result1.GetValue()[2].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[3].GetItem1(), result1.GetValue()[3].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[3].GetItem2().GetValue(), result1.GetValue()[3].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[3].GetItem2().GetLen(), result1.GetValue()[3].GetItem2().GetLen())
}

func Test_perEncodingTestListExtensible1(t *testing.T) {

	testListExtensible1, err := createTestListExtensible1Msg0Items()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	per, err := perEncodeTestListExtensible1(testListExtensible1)
	assert.NilError(t, err)
	t.Logf("TestListExtensible1 PER\n%v", hex.Dump(per))

	result, err := perDecodeTestListExtensible1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestListExtensible1 PER - decoded\n%v", result)

	testListExtensible11, err := createTestListExtensible1MsgFull()
	assert.NilError(t, err, "Error creating TestListExtensible1 PDU")

	per1, err := perEncodeTestListExtensible1(testListExtensible11)
	assert.NilError(t, err)
	t.Logf("TestListExtensible1 PER\n%v", hex.Dump(per1))

	result1, err := perDecodeTestListExtensible1(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestListExtensible1 PER - decoded\n%v", result1)
	assert.Equal(t, 4, len(result1.GetValue()))
	assert.Equal(t, testListExtensible11.GetValue()[0].GetItem1(), result1.GetValue()[0].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[0].GetItem2().GetValue(), result1.GetValue()[0].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[0].GetItem2().GetLen(), result1.GetValue()[0].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[1].GetItem1(), result1.GetValue()[1].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[1].GetItem2().GetValue(), result1.GetValue()[1].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[1].GetItem2().GetLen(), result1.GetValue()[1].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[2].GetItem1(), result1.GetValue()[2].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[2].GetItem2().GetValue(), result1.GetValue()[2].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[2].GetItem2().GetLen(), result1.GetValue()[2].GetItem2().GetLen())
	assert.Equal(t, testListExtensible11.GetValue()[3].GetItem1(), result1.GetValue()[3].GetItem1())
	assert.DeepEqual(t, testListExtensible11.GetValue()[3].GetItem2().GetValue(), result1.GetValue()[3].GetItem2().GetValue())
	assert.Equal(t, testListExtensible11.GetValue()[3].GetItem2().GetLen(), result1.GetValue()[3].GetItem2().GetLen())
}
