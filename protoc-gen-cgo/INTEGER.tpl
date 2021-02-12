// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "INTEGER.h"
import "C"
import (
	"fmt"
	"math"
	"math/big"
	"unsafe"
)

func xerEncodeInteger(integer int64) ([]byte, error) {
	integerCP, err := newInteger(integer)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}
	defer freeInteger(integerCP)
	bytes, err := encodeXer(&C.asn_DEF_INTEGER, unsafe.Pointer(integerCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeInteger() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeInteger(integer int64) ([]byte, error) {
	integerCP, err := newInteger(integer)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}
	defer freeInteger(integerCP)
	bytes, err := encodePerBuffer(&C.asn_DEF_INTEGER, unsafe.Pointer(integerCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeInteger() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeInteger(integerBytes []byte) (int64, error) {
	unsafePtr, err := decodeXer(integerBytes, &C.asn_DEF_INTEGER)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	intC := (*C.INTEGER_t)(unsafePtr)
	return decodeInteger(intC)
}

func perDecodeInteger(integerBytes []byte) (int64, error) {
	unsafePtr, err := decodePer(integerBytes, len(integerBytes), &C.asn_DEF_INTEGER)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	intC := (*C.INTEGER_t)(unsafePtr)
	return decodeInteger(intC)
}

// It is like a two's complement encoding of a signed number, but not quite the same
func newInteger(msg int64) (*C.INTEGER_t, error) {
	if msg == 0 {
		return &C.INTEGER_t{
			buf:  (*C.uchar)(C.CBytes([]byte{0x0})),
			size: C.ulong(1),
		}, nil
	}

	absMsg := uint64(math.Abs(float64(msg)))
	isNegative := false
	if msg < 0 {
		isNegative = true
	}
	places := math.Log2(float64(absMsg))
	var extraByte uint8 = 0
	if math.Mod(places, 8) > 7 {
		extraByte = 1
	}
	size := uint8(math.Ceil(places/8.0)) + extraByte
	asBytes := make([]byte, size)

	var i uint8 = 0
	for ; i < size-extraByte; i++ {
		asBytes[size-i-1] = byte(absMsg >> (i * 8))
	}
	if isNegative {
		for idx := range asBytes {
			asBytes[idx] = ^asBytes[idx] // Get the inverse (NOT) of it
			if idx == int(size-1) {
				asBytes[idx]++ // Increase the last byte by 1
			}
		}
	}

	return newAsnCodecsPrim(asBytes), nil
}

// It is like a two's complement decoding of a signed number, but not quite
// the same for negative values
func decodeInteger(intC *C.INTEGER_t) (int64, error) {
	bytes := decodeAsnCodecsPrim(intC)

	ni := big.NewInt(0)
	ni.SetBytes(bytes)

	if bytes[0]&0x80 > 0 { // If negative
		bigBytes := make([]byte, 0)
		for idx, b := range bytes {
			b2 := ^b
			if idx == len(bytes)-1 {
				b2++
			}
			bigBytes = append(bigBytes, b2)
		}
		ni.SetBytes(bigBytes)
		ni.Neg(ni)
	}

	return ni.Int64(), nil
}

func freeInteger(intC *C.INTEGER_t) {
	freeAsnCodecsPrim(intC)
}
