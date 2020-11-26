// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"gotest.tools/assert"
	"testing"
)

const (
	zero                            = 0
	twenty                          = 20
	twoHundredAndFive               = 205
	threeHundredAndThirtyThree      = 333
	sixtyThousand                   = 60000
	minusTwentyOne                  = -21
	minusTwoHundredAndFive          = -205
	minusThreeHundredAndThirtyThree = -333
	minusSixtyThousand              = -60000
)

func Test_IntegerEncode(t *testing.T) {
	xer20, err := xerEncodeInteger(twenty) // 1 byte because less than 128
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>20</INTEGER>\n", string(xer20))
	// Now do the reverse
	int20X, err := xerDecodeInteger(xer20)
	assert.NilError(t, err)
	assert.Equal(t, int64(twenty), int20X)
	// Now do the same in PER
	per20, err := perEncodeInteger(twenty) // 1 byte because less than 128
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x1, 0x14}, per20)
	// Now do the reverse
	int20P, err := perDecodeInteger(per20)
	assert.NilError(t, err)
	assert.Equal(t, int64(twenty), int20P)

	xer205, err := xerEncodeInteger(twoHundredAndFive) // 2 bytes because > 128
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>205</INTEGER>\n", string(xer205))
	// Now do the reverse
	int205X, err := xerDecodeInteger(xer205)
	assert.NilError(t, err)
	assert.Equal(t, int64(twoHundredAndFive), int205X)
	// Now do the same in PER
	per205, err := perEncodeInteger(twoHundredAndFive) // 2 bytes because > 128
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x2, 0x00, 0xcd}, per205)
	// Now do the reverse
	int205P, err := perDecodeInteger(per205)
	assert.NilError(t, err)
	assert.Equal(t, int64(twoHundredAndFive), int205P)

	xer333, err := xerEncodeInteger(threeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>333</INTEGER>\n", string(xer333))
	// Now do the reverse
	int333, err := xerDecodeInteger(xer333)
	assert.NilError(t, err)
	assert.Equal(t, int64(threeHundredAndThirtyThree), int333)
	// Now do the same in PER
	per333, err := perEncodeInteger(threeHundredAndThirtyThree) // 2 bytes because > 128
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x02, 0x1, 0x4d}, per333)
	// Now do the reverse
	int333P, err := perDecodeInteger(per333)
	assert.NilError(t, err)
	assert.Equal(t, int64(threeHundredAndThirtyThree), int333P)

	xer60000, err := xerEncodeInteger(sixtyThousand) // 3 bytes because > 32767
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>60000</INTEGER>\n", string(xer60000))
	// Now do the reverse
	int60000X, err := xerDecodeInteger(xer60000)
	assert.NilError(t, err)
	assert.Equal(t, int64(sixtyThousand), int60000X)
	// Now do the same in PER
	per60000, err := perEncodeInteger(sixtyThousand) // 3 bytes because > 32767
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x03, 0x00, 0xea, 0x60}, per60000)
	// Now do the reverse
	int60000P, err := perDecodeInteger(per60000)
	assert.NilError(t, err)
	assert.Equal(t, int64(sixtyThousand), int60000P)

	xer0, err := xerEncodeInteger(zero)
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>0</INTEGER>\n", string(xer0))
	// Now do the reverse
	int0X, err := xerDecodeInteger(xer0)
	assert.NilError(t, err)
	assert.Equal(t, int64(zero), int0X)
	// Now do the same in PER
	per0, err := perEncodeInteger(zero)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x01, 0x00}, per0)
	// Now do the reverse
	int0P, err := perDecodeInteger(per0)
	assert.NilError(t, err)
	assert.Equal(t, int64(zero), int0P)

	xerN21, err := xerEncodeInteger(minusTwentyOne) // 1 byte because less than 128
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>-21</INTEGER>\n", string(xerN21))
	// Now do the reverse
	intN21, err := xerDecodeInteger(xerN21)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusTwentyOne), intN21)
	// Now do the same in PER
	perN21, err := perEncodeInteger(minusTwentyOne) // 1 byte because less than 128
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x01, 0xeb}, perN21)
	// Now do the reverse
	intN21P, err := perDecodeInteger(perN21)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusTwentyOne), intN21P)

	xerN205, err := xerEncodeInteger(minusTwoHundredAndFive) // 2 bytes because > 128
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>-205</INTEGER>\n", string(xerN205))
	// Now do the reverse
	intN205, err := xerDecodeInteger(xerN205)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusTwoHundredAndFive), intN205)
	// Now do the same in PER
	perN205, err := perEncodeInteger(minusTwoHundredAndFive) // 1 byte because less than 128
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x02, 0xff, 0x33}, perN205)
	// Now do the reverse
	intN205P, err := perDecodeInteger(perN205)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusTwoHundredAndFive), intN205P)

	xerN333, err := xerEncodeInteger(minusThreeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>-333</INTEGER>\n", string(xerN333))
	// Now do the reverse
	intN333, err := xerDecodeInteger(xerN333)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusThreeHundredAndThirtyThree), intN333)
	// Now do the same in PER
	perN333, err := perEncodeInteger(minusThreeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x02, 0xfe, 0xb3}, perN333)
	// Now do the reverse
	intN333P, err := perDecodeInteger(perN333)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusThreeHundredAndThirtyThree), intN333P)

	xerN60000, err := xerEncodeInteger(minusSixtyThousand) // 3 bytes because > 32767
	assert.NilError(t, err)
	assert.Equal(t, "<INTEGER>-60000</INTEGER>\n", string(xerN60000))
	// Now do the reverse
	intN60000X, err := xerDecodeInteger(xerN60000)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusSixtyThousand), intN60000X)
	// Now do the same in PER
	perN60000, err := perEncodeInteger(minusSixtyThousand) // 3 bytes because > 32767
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x03, 0xff, 0x15, 0xa0}, perN60000)
	// Now do the reverse
	intN60000P, err := perDecodeInteger(perN60000)
	assert.NilError(t, err)
	assert.Equal(t, int64(minusSixtyThousand), intN60000P)
}
