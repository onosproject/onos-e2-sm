// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	//"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	//"io/ioutil"
	"reflect"
	"testing"
)

func TestBitString_NewBitString(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bs := NewBitString(value, length)
	assert.Equal(t, reflect.TypeOf(BitString{}), reflect.TypeOf(*bs), "BitString{} types are mismatched")
	assert.Equal(t, bs.Value, value, "BitString{} values are mismatched")
	assert.Equal(t, bs.Len, length, "BitString{} lengths are mismatched")
}

func TestBitString_GetValue(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bs := NewBitString(value, length)

	assert.Equal(t, bs.GetValue(), value, "Test_BitString GetValue BitString values mismatch")
}

func TestBitString_GetLen(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bs := NewBitString(value, length)

	assert.Equal(t, bs.GetLen(), length, "Test_BitString GetLen BitString lengths mismatch")
}

func TestBitString_GetBitString(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bs1 := NewBitString(value, length)
	bs2 := bs1.GetBitString()

	assert.Equal(t, bs1.GetValue(), bs2.GetValue(), "Test_BitString GetBitString BitString value mismatch")
	assert.Equal(t, bs1.GetLen(), bs2.GetLen(), "Test_BitString GetBitString BitString lengths mismatch")
}
