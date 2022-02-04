// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_newBitString(t *testing.T) {
	bs1 := &e2sm_kpm_v2.BitString{
		Value: []byte{0xD4, 0xBC, 0x90},
		Len:   22,
	}

	xer1, err := xerEncodeBitString(bs1)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer1)

	result1Xer, err := xerDecodeBitString(xer1)
	assert.NilError(t, err)
	t.Logf("XER Bit String - decoded\n%v", result1Xer)
	assert.Equal(t, bs1.Len, result1Xer.Len)
	assert.DeepEqual(t, bs1.Value, result1Xer.Value)

	per1, err := perEncodeBitString(bs1)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per1))

	result1Per, err := perDecodeBitString(per1)
	assert.NilError(t, err)
	t.Logf("PER Bit String - decoded\n%v", result1Per)
	assert.Equal(t, bs1.Len, result1Per.Len)
	assert.DeepEqual(t, bs1.Value, result1Per.Value)

	cBitString, err := newBitString(bs1)
	assert.NilError(t, err)

	assert.Equal(t, 3, int(cBitString.size), "unexpected number of bits")
	assert.Equal(t, 2, int(cBitString.bits_unused), "unexpected number of bits_unused")
	// Can't do any further analysis as we can't have C in tests

	// Now reverse it to get proto back out
	bs2, err := decodeBitString(cBitString)
	assert.NilError(t, err)
	assert.Equal(t, uint32(22), bs2.Len)
	assert.DeepEqual(t, []byte{0xD4, 0xBC, 0x90}, bs2.Value)

	xer2, err := xerEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, xer1, xer2)
	t.Logf("XER Bit String \n%s", xer1)

	result2Xer, err := xerDecodeBitString(xer2)
	assert.NilError(t, err)
	t.Logf("XER Bit String - decoded\n%v", result2Xer)
	assert.Equal(t, bs2.Len, result2Xer.Len)
	assert.DeepEqual(t, bs2.Value, result2Xer.Value)
	assert.Equal(t, result1Xer.Len, result2Xer.Len)
	assert.DeepEqual(t, result1Xer.Value, result2Xer.Value)

	per2, err := perEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per1, per2)
	t.Logf("PER Bit String \n%v", hex.Dump(per2))

	result2Per, err := perDecodeBitString(per1)
	assert.NilError(t, err)
	t.Logf("PER Bit String - decoded\n%v", result2Per)
	assert.Equal(t, bs2.Len, result2Per.Len)
	assert.DeepEqual(t, bs2.Value, result2Per.Value)
	assert.Equal(t, result2Per.Len, result2Per.Len)
	assert.DeepEqual(t, result2Per.Value, result2Per.Value)
}

func Test_decodeBitString(t *testing.T) {
	//value := []byte{0x9A, 0xBC, 0xDE, 0xF0} // 28 bits
	bs := &e2sm_kpm_v2.BitString{
		Value: []byte{0x9A, 0xBC, 0xDE, 0xF0},
		Len:   28,
	}

	bsC, err := newBitString(bs)
	assert.NilError(t, err)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	assert.DeepEqual(t, protoBitString.Value, []byte{0x9a, 0xbc, 0xde, 0xf0})

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	resultXer, err := xerDecodeBitString(xer)
	assert.NilError(t, err)
	t.Logf("XER Bit String - decoded\n%v", resultXer)
	assert.Equal(t, bs.Len, resultXer.Len)
	assert.DeepEqual(t, bs.Value, resultXer.Value)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per))

	resultPer, err := perDecodeBitString(per)
	assert.NilError(t, err)
	t.Logf("PER Bit String - decoded\n%v", resultPer)
	assert.Equal(t, bs.Len, resultPer.Len)
	assert.DeepEqual(t, bs.Value, resultPer.Value)
}

func Test_decodeBitString2(t *testing.T) {
	//value := []byte{0x9A, 0xBC, 0xD4} // 22 bits
	bs3 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x9A, 0xBC, 0xDC},
		Len:   22,
	}
	bsC, err := newBitString(bs3)
	assert.NilError(t, err)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 22, "unexpected bit string length")
	assert.DeepEqual(t, protoBitString.Value, []byte{0x9a, 0xbc, 0xdc})

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	resultXer, err := xerDecodeBitString(xer)
	assert.NilError(t, err)
	t.Logf("XER Bit String - decoded\n%v", resultXer)
	assert.Equal(t, bs3.Len, resultXer.Len)
	assert.DeepEqual(t, bs3.Value, resultXer.Value)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per))

	resultPer, err := perDecodeBitString(per)
	assert.NilError(t, err)
	t.Logf("PER Bit String - decoded\n%v", resultPer)
	assert.Equal(t, bs3.Len, resultPer.Len)
	assert.DeepEqual(t, bs3.Value, resultPer.Value)
}

// A unit test to illustrate importance of Octet alignment and bit shifting
func Test_invalidBitStrings(t *testing.T) {

	bs1 := &e2sm_kpm_v2.BitString{
		Value: []byte{0xab, 0xbc},
		//Value: []byte{0xab, 0xbc, 0x00}, // - correct set of bytes to satisfy byte length constraint
		Len: 22,
	}
	_, err := newBitString(bs1)
	//assert.NilError(t, err)
	assert.ErrorContains(t, err, "bytes are required for length")

	bs2 := &e2sm_kpm_v2.BitString{
		Value: []byte{0xab, 0xbc, 0xcf},
		//Value: []byte{0xab, 0xbc, 0xcc}, // - a correct set of bytes to satisfy octet-alignment constraint
		Len: 22,
	}
	_, err = newBitString(bs2)
	//assert.NilError(t, err)
	assert.ErrorContains(t, err, "bit string is NOT octet-aligned")

	// A valid bit string with proper length and proper octet alignment
	// 25 bits need 4 bytes to be encoded successfully
	// 4 bytes is 32 bits, 32-25 = 7 (unused bits)
	// BitString.Value byte array should be shifted on 7 (unused) bits to the left to be Octet-aligned and satisfy APER encoding rules
	bs3 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x09, 0xAB, 0xCD, 0x80},
		Len:   25,
	}
	bsC, err := newBitString(bs3)
	assert.NilError(t, err)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 25, "unexpected bit string length")
	assert.DeepEqual(t, protoBitString.Value, []byte{0x09, 0xab, 0xcd, 0x80})
}

// A unit test to illustrate Octet alignment and bit shifting on the example of encoding 1
func Test_validBitStringsOne(t *testing.T) {
	// Here is a set of valid representations of 1 with different BitString length (see logs)
	bs4 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x40},
		Len:   18, // 6 bits unused
	}

	xer1, err := xerEncodeBitString(bs4)
	assert.NilError(t, err)
	t.Logf("XER Bit String is \n%s", xer1)

	bs5 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22, // 2 bits unused
	}

	xer2, err := xerEncodeBitString(bs5)
	assert.NilError(t, err)
	t.Logf("XER Bit String is \n%s", xer2)

	bs6 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x10},
		Len:   28,
	}

	xer3, err := xerEncodeBitString(bs6)
	assert.NilError(t, err)
	t.Logf("XER Bit String is \n%s", xer3)

	bs7 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x01},
		Len:   32,
	}

	xer4, err := xerEncodeBitString(bs7)
	assert.NilError(t, err)
	t.Logf("XER Bit String is \n%s", xer4)

	bs8 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x00, 0x10},
		Len:   36,
	}

	xer5, err := xerEncodeBitString(bs8)
	assert.NilError(t, err)
	t.Logf("XER Bit String is \n%s", xer5)
}
