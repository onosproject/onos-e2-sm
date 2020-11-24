// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"gotest.tools/assert"
	"testing"
)

func Test_newInteger(t *testing.T) {

	var msg int64 = 212747

	intC := newInteger(msg)
	assert.Equal(t, 8, int(intC.size), "unexpected size")
}

func Test_decodeInteger(t *testing.T) {

	var msg int64 = 2

	intC := newInteger(msg)

	result := decodeInteger(intC)
	assert.Equal(t, result, msg, "Something went wrong")

	//value := []byte{0x9A, 0xBC, 0xDE, 0xF0} // 28 bits
	//bsC := newBitStringFromBytes(value, 4, 4)
	//
	//protoBitString, err := decodeBitString(bsC)
	//assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	//assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	//assert.Equal(t, protoBitString.Value, uint64(0xf0debc9a), "unexpected bit string value")
}

func Test_perEncodeDecodeInteger(t *testing.T) {

	var msg int64 = 74721569

	intPer, err := perEncodeInteger(msg)
	assert.NilError(t, err)
	assert.Assert(t, intPer != nil)
	t.Logf("INTEGER PER\n%v", intPer)

	result, err := perDecodeInteger(intPer)
	assert.NilError(t, err)
	assert.Assert(t, result != 0)
	t.Logf("INTEGER PER decoded is \n%v", result)

	assert.Equal(t, msg, result, "Error between the keyboard and the chair?")
}