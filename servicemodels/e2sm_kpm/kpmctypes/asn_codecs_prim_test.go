// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"gotest.tools/assert"
	"testing"
)

func Test_newAsnCodecsPrim(t *testing.T) {

	var integer int64 = 212747

	byteArrayC := newAsnCodecsPrim(integer)
	assert.Equal(t, 8, int(byteArrayC.size), "unexpected size")
}

func Test_decodeAsnCodecsPrim(t *testing.T) {

	var integer int64 = 2

	byteArrayC := newAsnCodecsPrim(integer)

	result := decodeAsnCodecsPrim(byteArrayC)
	assert.Equal(t, result, integer, "Something went wrong")

	//value := []byte{0x9A, 0xBC, 0xDE, 0xF0} // 28 bits
	//bsC := newBitStringFromBytes(value, 4, 4)
	//
	//protoBitString, err := decodeBitString(bsC)
	//assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	//assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	//assert.Equal(t, protoBitString.Value, uint64(0xf0debc9a), "unexpected bit string value")
}