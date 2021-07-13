// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_newBitString(t *testing.T) {
	bs1 := e2sm_kpm_v2.BitString{
		Value: []byte{0xD4, 0xBC, 0x9A},
		Len:   22,
	}

	xer1, err := xerEncodeBitString(&bs1)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer1)

	per1, err := perEncodeBitString(&bs1)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per1))

	cBitString, err := newBitString(&bs1)
	assert.NilError(t, err)

	assert.Equal(t, 3, int(cBitString.size), "unexpected number of bits")
	assert.Equal(t, 2, int(cBitString.bits_unused), "unexpected number of bits_unused")
	// Can't do any further analysis as we can't have C in tests

	// Now reverse it to get proto back out
	//ToDo - incorrect decoding since value is d4bc9a0000000000 now - investigate why
	bs2, err := decodeBitString(cBitString)
	assert.NilError(t, err)
	assert.Equal(t, uint32(22), bs2.Len)
	assert.DeepEqual(t, []byte{0xD4, 0xBC, 0x9A, 0x00, 0x00, 0x00, 0x00, 0x00}, bs2.Value)

	xer2, err := xerEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, xer1, xer2)
	t.Logf("XER Bit String \n%s", xer1)

	per2, err := perEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per1, per2)
	t.Logf("PER Bit String \n%v", hex.Dump(per2))
}

func Test_decodeBitString(t *testing.T) {
	//value := []byte{0x9A, 0xBC, 0xDE, 0xF0} // 28 bits
	bs2 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x9A, 0xBC, 0xDE, 0xF0},
		Len:   28,
	}

	bsC, err := newBitString(bs2)
	assert.NilError(t, err)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	assert.DeepEqual(t, protoBitString.Value, []byte{0x9a, 0xbc, 0xde, 0xf0, 0x00, 0x00, 0x00, 0x00})

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per))
}

func Test_decodeBitString2(t *testing.T) {
	//value := []byte{0x9A, 0xBC, 0xD4} // 22 bits
	bs3 := &e2sm_kpm_v2.BitString{
		Value: []byte{0x9A, 0xBC, 0xDE},
		Len:   22,
	}
	bsC, err := newBitString(bs3)
	assert.NilError(t, err)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	//assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 22, "unexpected bit string length")
	assert.DeepEqual(t, protoBitString.Value, []byte{0x9a, 0xbc, 0xde, 0x00, 0x00, 0x00, 0x00, 0x00})

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", hex.Dump(per))
}
