// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "INTEGER.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func perEncodeInteger(msg int64) ([]byte, error) {
	intCP := newInteger(msg)

	bytes, err := encodePerBuffer(&C.asn_DEF_INTEGER, unsafe.Pointer(intCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeInteger() %s", err.Error())
	}
	//fmt.Printf("perEncodeInteger -- bytes -- %v\n", bytes)
	return bytes, nil
}

func perDecodeInteger(bytes []byte) (int64, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_INTEGER)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from PER is nil")
	}
	//fmt.Printf("PerDecodeE2SmKpmIndicationHeader -- unsafePtr -- %v\n", unsafePtr)
	return decodeInteger((*C.INTEGER_t)(unsafePtr)), nil
}

func newInteger(msg int64) *C.INTEGER_t {

	intC := newAsnCodecsPrim(msg)

	return intC
}

func decodeInteger(intC *C.INTEGER_t) int64 {

	bytes := decodeAsnCodecsPrim(intC)
	// There could possibly occur an issue in parsing str-to-int. It happens, output value would be 0.
	//byteToInt, _ := strconv.Atoi(string(bytes))
	//fmt.Printf("decodeInteger -- bytes -- %v\n", bytes)

	return bytes //int64(byteToInt)

}