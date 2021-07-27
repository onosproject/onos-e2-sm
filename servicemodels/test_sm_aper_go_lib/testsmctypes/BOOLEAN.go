// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "BOOLEAN.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func xerEncodeBoolean(b bool) ([]byte, error) {
	bCP := newBoolean(b)

	bytes, err := encodeXer(&C.asn_DEF_BOOLEAN, unsafe.Pointer(bCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeBoolean() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeBoolean(b bool) ([]byte, error) {
	bCP := newBoolean(b)

	bytes, err := encodePerBuffer(&C.asn_DEF_BOOLEAN, unsafe.Pointer(bCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeBoolean() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeBoolean(bytes []byte) (bool, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_BOOLEAN)
	if err != nil {
		return false, err
	}
	if unsafePtr == nil {
		return false, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeBoolean((*C.BOOLEAN_t)(unsafePtr)), nil
}

func perDecodeBoolean(bytes []byte) (bool, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_BOOLEAN)
	if err != nil {
		return false, err
	}
	if unsafePtr == nil {
		return false, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeBoolean((*C.BOOLEAN_t)(unsafePtr)), nil
}

func newBoolean(b bool) *C.BOOLEAN_t {

	bC := C.int(0x00)
	if b {
		bC = C.int(0xff)
	}
	//fmt.Printf("Encoded boolean is %v\n", bC)
	return &bC
}

func decodeBoolean(bC *C.BOOLEAN_t) bool {

	b := int32(*bC)
	//fmt.Printf("Decoded boolean is %v\n", b)

	return b != 0
}

func decodeBooleanBytes(array [8]byte) bool {

	b := int32(binary.LittleEndian.Uint32(array[:8]))
	//fmt.Printf("Decoded from bytes boolean is %v\n", b)

	return b != 0
}
