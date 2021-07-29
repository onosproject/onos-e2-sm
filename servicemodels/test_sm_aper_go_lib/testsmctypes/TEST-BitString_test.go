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

func createTestBitStringMsg() (*test_sm_ies.TestBitString, error) {

	testBitString := test_sm_ies.TestBitString{
		AttrBs1: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
			Len:   35,
		},
		AttrBs2: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x40},
			Len:   20,
		},
		AttrBs3: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x40},
			Len:   20,
		},
		//AttrBs4: &asn1.BitString{  // ToDo - there should be no Octet-Alignment - BitStrings is less than 16!!
		//	Value: []byte{0x01},
		//	Len:   1,
		//},
		AttrBs4: &asn1.BitString{
			Value: []byte{0xCC, 0xC0},
			Len:   16,
		},
		AttrBs5: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x4F, 0x00},
			Len:   32,
		},
		AttrBs6: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x4F, 0x00},
			Len:   32,
		},
		AttrBs7: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
			Len:   35,
		},
	}

	return &testBitString, nil
}

func Test_xerEncodingTestBitString(t *testing.T) {

	testBitString, err := createTestBitStringMsg()
	assert.NilError(t, err, "Error creating TestBitString PDU")

	xer, err := xerEncodeTestBitString(testBitString)
	assert.NilError(t, err)
	t.Logf("TestBitString XER\n%s", string(xer))

	result, err := xerDecodeTestBitString(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestBitString XER - decoded\n%v", result)
	assert.DeepEqual(t, testBitString.GetAttrBs1().GetValue(), result.GetAttrBs1().GetValue())
	assert.Equal(t, testBitString.GetAttrBs1().GetLen(), result.GetAttrBs1().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs2().GetValue(), result.GetAttrBs2().GetValue())
	assert.Equal(t, testBitString.GetAttrBs2().GetLen(), result.GetAttrBs2().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs3().GetValue(), result.GetAttrBs3().GetValue())
	assert.Equal(t, testBitString.GetAttrBs3().GetLen(), result.GetAttrBs3().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs4().GetValue(), result.GetAttrBs4().GetValue())
	assert.Equal(t, testBitString.GetAttrBs4().GetLen(), result.GetAttrBs4().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs5().GetValue(), result.GetAttrBs5().GetValue())
	assert.Equal(t, testBitString.GetAttrBs5().GetLen(), result.GetAttrBs5().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs6().GetValue(), result.GetAttrBs6().GetValue())
	assert.Equal(t, testBitString.GetAttrBs6().GetLen(), result.GetAttrBs6().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs7().GetValue(), result.GetAttrBs7().GetValue())
	assert.Equal(t, testBitString.GetAttrBs7().GetLen(), result.GetAttrBs7().GetLen())
}

func Test_perEncodingTestBitString(t *testing.T) {

	testBitString, err := createTestBitStringMsg()
	assert.NilError(t, err, "Error creating TestBitString PDU")

	per, err := perEncodeTestBitString(testBitString)
	assert.NilError(t, err)
	t.Logf("TestBitString PER\n%v", hex.Dump(per))

	result, err := perDecodeTestBitString(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestBitString PER - decoded\n%v", result)
	assert.DeepEqual(t, testBitString.GetAttrBs1().GetValue(), result.GetAttrBs1().GetValue())
	assert.Equal(t, testBitString.GetAttrBs1().GetLen(), result.GetAttrBs1().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs2().GetValue(), result.GetAttrBs2().GetValue())
	assert.Equal(t, testBitString.GetAttrBs2().GetLen(), result.GetAttrBs2().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs3().GetValue(), result.GetAttrBs3().GetValue())
	assert.Equal(t, testBitString.GetAttrBs3().GetLen(), result.GetAttrBs3().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs4().GetValue(), result.GetAttrBs4().GetValue())
	assert.Equal(t, testBitString.GetAttrBs4().GetLen(), result.GetAttrBs4().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs5().GetValue(), result.GetAttrBs5().GetValue())
	assert.Equal(t, testBitString.GetAttrBs5().GetLen(), result.GetAttrBs5().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs6().GetValue(), result.GetAttrBs6().GetValue())
	assert.Equal(t, testBitString.GetAttrBs6().GetLen(), result.GetAttrBs6().GetLen())
	assert.DeepEqual(t, testBitString.GetAttrBs7().GetValue(), result.GetAttrBs7().GetValue())
	assert.Equal(t, testBitString.GetAttrBs7().GetLen(), result.GetAttrBs7().GetLen())
}
