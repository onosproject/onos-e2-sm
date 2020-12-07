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

	variable := NewBitString()
	assert.Equal(t, reflect.TypeOf(BitString{}), reflect.TypeOf(*variable), "BitString{} types are mismatched")

}

func TestBitString_SetValue(t *testing.T) {

	var value uint64 = 0x9bcd4

	bs := NewBitString().SetValue(value)

	assert.Equal(t, bs.Value, value, "Mismatch of BitString values")

}

func TestBitString_SetLen(t *testing.T) {

	var len uint32 = 22

	bs := NewBitString().SetLen(len)

	assert.Equal(t, bs.Len, len, "Mismatch of BitString length")

}

func TestBitString_GetValue(t *testing.T) {

	var value uint64 = 0x9bcd4

	bs := NewBitString().SetValue(value)

	assert.Equal(t, bs.GetValue(), value, "Mismatch of BitString values")

}

func TestBitString_GetLen(t *testing.T) {

	var len uint32 = 22

	bs := NewBitString().SetLen(len)

	assert.Equal(t, bs.GetLen(), len, "Mismatch of BitString length")

}

func TestBitString_GetBitString(t *testing.T) {

	bs1 := NewBitString().SetValue(0x9bcd4).SetLen(22)

	bs2 := bs1.GetBitString()

	assert.Equal(t, bs1.GetValue(), bs2.GetValue(), "Mismatch of BitString values")
	assert.Equal(t, bs1.GetLen(), bs2.GetLen(), "Mismatch of BitString length")
}
