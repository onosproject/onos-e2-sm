// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//const (
//	zero                            = 0.0
//	twenty                          = 20.5
//	twoHundredAndFive               = 205.0
//	threeHundredAndThirtyThree      = 333.01
//	sixtyThousand                   = 60000.24
//	minusTwentyOne                  = -21.0
//	minusTwoHundredAndFive          = -205.35
//	minusThreeHundredAndThirtyThree = -333.33
//	minusSixtyThousand              = -60000.606
//)

//func Test_RealEncode(t *testing.T) {
//	//t.Logf("Format is %T\n", twenty)
//	xer20, err := xerEncodeReal(twenty) // 1 byte because less than 128
//	assert.NilError(t, err)
//	assert.Equal(t, "<REAL>20.5</REAL>\n", string(xer20))
//	// Now do the reverse
//	int20X, err := xerDecodeReal(xer20)
//	assert.NilError(t, err)
//	assert.Equal(t, twenty, int20X)
//	// Now do the same in PER
//	per20, err := perEncodeReal(twenty) // 1 byte because less than 128
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x1, 0x14}, per20)
//	// Now do the reverse
//	int20P, err := perDecodeReal(per20)
//	assert.NilError(t, err)
//	assert.Equal(t, twenty, int20P)
//
//	xer205, err := xerEncodeReal(twoHundredAndFive) // 2 bytes because > 128
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>205.0</INTEGER>\n", string(xer205))
//	// Now do the reverse
//	int205X, err := xerDecodeReal(xer205)
//	assert.NilError(t, err)
//	assert.Equal(t, twoHundredAndFive, int205X)
//	// Now do the same in PER
//	per205, err := perEncodeReal(twoHundredAndFive) // 2 bytes because > 128
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x2, 0x00, 0xcd}, per205)
//	// Now do the reverse
//	int205P, err := perDecodeReal(per205)
//	assert.NilError(t, err)
//	assert.Equal(t, twoHundredAndFive, int205P)
//
//	xer333, err := xerEncodeReal(threeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>333.01</INTEGER>\n", string(xer333))
//	// Now do the reverse
//	int333, err := xerDecodeReal(xer333)
//	assert.NilError(t, err)
//	assert.Equal(t, threeHundredAndThirtyThree, int333)
//	// Now do the same in PER
//	per333, err := perEncodeReal(threeHundredAndThirtyThree) // 2 bytes because > 128
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x02, 0x1, 0x4d}, per333)
//	// Now do the reverse
//	int333P, err := perDecodeReal(per333)
//	assert.NilError(t, err)
//	assert.Equal(t, threeHundredAndThirtyThree, int333P)
//
//	xer60000, err := xerEncodeReal(sixtyThousand) // 3 bytes because > 32767
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>60000.24</INTEGER>\n", string(xer60000))
//	// Now do the reverse
//	int60000X, err := xerDecodeReal(xer60000)
//	assert.NilError(t, err)
//	assert.Equal(t, sixtyThousand, int60000X)
//	// Now do the same in PER
//	per60000, err := perEncodeReal(sixtyThousand) // 3 bytes because > 32767
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x03, 0x00, 0xea, 0x60}, per60000)
//	// Now do the reverse
//	int60000P, err := perDecodeReal(per60000)
//	assert.NilError(t, err)
//	assert.Equal(t, sixtyThousand, int60000P)
//
//	xer0, err := xerEncodeReal(zero)
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>0</INTEGER>\n", string(xer0))
//	// Now do the reverse
//	int0X, err := xerDecodeReal(xer0)
//	assert.NilError(t, err)
//	assert.Equal(t, zero, int0X)
//	// Now do the same in PER
//	per0, err := perEncodeReal(zero)
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x01, 0x00}, per0)
//	// Now do the reverse
//	int0P, err := perDecodeReal(per0)
//	assert.NilError(t, err)
//	assert.Equal(t, zero, int0P)
//
//	xerN21, err := xerEncodeReal(minusTwentyOne) // 1 byte because less than 128
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>-21.0</INTEGER>\n", string(xerN21))
//	// Now do the reverse
//	intN21, err := xerDecodeReal(xerN21)
//	assert.NilError(t, err)
//	assert.Equal(t, minusTwentyOne, intN21)
//	// Now do the same in PER
//	perN21, err := perEncodeReal(minusTwentyOne) // 1 byte because less than 128
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x01, 0xeb}, perN21)
//	// Now do the reverse
//	intN21P, err := perDecodeReal(perN21)
//	assert.NilError(t, err)
//	assert.Equal(t, minusTwentyOne, intN21P)
//
//	xerN205, err := xerEncodeReal(minusTwoHundredAndFive) // 2 bytes because > 128
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>-205.35</INTEGER>\n", string(xerN205))
//	// Now do the reverse
//	intN205, err := xerDecodeReal(xerN205)
//	assert.NilError(t, err)
//	assert.Equal(t, minusTwoHundredAndFive, intN205)
//	// Now do the same in PER
//	perN205, err := perEncodeReal(minusTwoHundredAndFive) // 1 byte because less than 128
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x02, 0xff, 0x33}, perN205)
//	// Now do the reverse
//	intN205P, err := perDecodeReal(perN205)
//	assert.NilError(t, err)
//	assert.Equal(t, minusTwoHundredAndFive, intN205P)
//
//	xerN333, err := xerEncodeReal(minusThreeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>-333.33</INTEGER>\n", string(xerN333))
//	// Now do the reverse
//	intN333, err := xerDecodeReal(xerN333)
//	assert.NilError(t, err)
//	assert.Equal(t, minusThreeHundredAndThirtyThree, intN333)
//	// Now do the same in PER
//	perN333, err := perEncodeReal(minusThreeHundredAndThirtyThree) // 2 bytes because > 128 and < 32767
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x02, 0xfe, 0xb3}, perN333)
//	// Now do the reverse
//	intN333P, err := perDecodeReal(perN333)
//	assert.NilError(t, err)
//	assert.Equal(t, minusThreeHundredAndThirtyThree, intN333P)
//
//	xerN60000, err := xerEncodeReal(minusSixtyThousand) // 3 bytes because > 32767
//	assert.NilError(t, err)
//	assert.Equal(t, "<INTEGER>-60000.606</INTEGER>\n", string(xerN60000))
//	// Now do the reverse
//	intN60000X, err := xerDecodeReal(xerN60000)
//	assert.NilError(t, err)
//	assert.Equal(t, minusSixtyThousand, intN60000X)
//	// Now do the same in PER
//	perN60000, err := perEncodeReal(minusSixtyThousand) // 3 bytes because > 32767
//	assert.NilError(t, err)
//	//assert.DeepEqual(t, []byte{0x03, 0xff, 0x15, 0xa0}, perN60000)
//	// Now do the reverse
//	intN60000P, err := perDecodeReal(perN60000)
//	assert.NilError(t, err)
//	assert.Equal(t, minusSixtyThousand, intN60000P)
//}
