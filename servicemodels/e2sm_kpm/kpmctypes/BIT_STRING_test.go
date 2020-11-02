// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_newBitString(t *testing.T) {
	bs1 := e2sm_kpm_ies.BitString{
		Value: 0x9ABCD4,
		Len:   22,
	}

	xer1, err := xerEncodeBitString(&bs1)
	assert.NilError(t, err)
	//t.Logf("XER Bit String \n%s", xer1)

	per1, err := perEncodeBitString(&bs1)
	assert.NilError(t, err)
	//t.Logf("PER Bit String \n%v", per1)

	cBitString := newBitString(&bs1)

	assert.Equal(t, 3, int(cBitString.size), "unexpected number of bits")
	assert.Equal(t, 2, int(cBitString.bits_unused), "unexpected number of bits_unused")
	// Can't do any further analysis as we can't have C in tests

	// Now reverse it to get proto back out
	bs2, err := decodeBitString(cBitString)
	assert.NilError(t, err)
	assert.Equal(t, uint32(22), bs2.Len)
	assert.Equal(t, uint64(0x9ABCD4), bs2.Value)

	xer2, err := xerEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, xer1, xer2)
	t.Logf("XER Bit String \n%s", xer1)

	per2, err := perEncodeBitString(bs2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per1, per2)
	t.Logf("PER Bit String \n%v", per2)
}

func Test_decodeBitString(t *testing.T) {
	value := []byte{0x9A, 0xBC, 0xDE, 0xF0} // 28 bits
	bsC := newBitStringFromBytes(value, 4, 4)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	assert.Equal(t, protoBitString.Value, uint64(0xf0debc9a), "unexpected bit string value")

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", per)
}

func Test_decodeBitString2(t *testing.T) {
	value := []byte{0x9A, 0xBC, 0xD4} // 22 bits
	bsC := newBitStringFromBytes(value, 3, 2)

	protoBitString, err := decodeBitString(bsC)
	assert.NilError(t, err)
	assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 22, "unexpected bit string length")
	assert.Equal(t, protoBitString.Value, uint64(0xd4bc9a), "unexpected bit string value")

	xer, err := xerEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("XER Bit String \n%s", xer)

	per, err := perEncodeBitString(protoBitString)
	assert.NilError(t, err)
	t.Logf("PER Bit String \n%v", per)
}
