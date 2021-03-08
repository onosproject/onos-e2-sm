// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "REAL.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"
)

func xerEncodeReal(real float64) ([]byte, error) {
	realCP, err := newReal(real)
	if err != nil {
		return nil, fmt.Errorf("newReal() %s", err.Error())
	}
	defer freeReal(realCP)
	bytes, err := encodeXer(&C.asn_DEF_REAL, unsafe.Pointer(realCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeReal() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeReal(real float64) ([]byte, error) {
	realCP, err := newReal(real)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}
	defer freeReal(realCP)
	bytes, err := encodePerBuffer(&C.asn_DEF_REAL, unsafe.Pointer(realCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeReal() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeReal(realBytes []byte) (float64, error) {
	unsafePtr, err := decodeXer(realBytes, &C.asn_DEF_REAL)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	realC := (*C.REAL_t)(unsafePtr)
	return decodeReal(realC)
}

func perDecodeReal(realBytes []byte) (float64, error) {
	unsafePtr, err := decodePer(realBytes, len(realBytes), &C.asn_DEF_REAL)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	realC := (*C.REAL_t)(unsafePtr)
	return decodeReal(realC)
}

// It is like a two's complement encoding of a signed number, but not quite the same
func newReal(msg float64) (*C.REAL_t, error) {
	if msg == 0 {
		return &C.REAL_t{
			buf:  (*C.uchar)(C.CBytes([]byte{0x0})),
			size: C.ulong(1),
		}, nil
	}

	absMsg := uint64(math.Abs(msg))
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
func decodeReal(realC *C.REAL_t) (float64, error) {
	bytes := decodeAsnCodecsPrim(realC)

	a := binary.LittleEndian.Uint64(bytes)
	b := math.Float64frombits(a)

	return b, nil
}

// Input value should always be 8 bytes. If you have more than 8 bytes,
// please split it on slices of 8 bytes and run this function on each slice
func decodeRealBytes(array [8]byte) (float64, error) {
	intC := (*C.REAL_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))
	return decodeReal(intC)
}

func freeReal(realC *C.REAL_t) {
	freeAsnCodecsPrim(realC)
}
