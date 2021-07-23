// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ConstrainedInt.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) ([]byte, error) {
	testConstrainedIntCP, err := newTestConstrainedInt(testConstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedInt() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ConstrainedInt, unsafe.Pointer(testConstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) ([]byte, error) {
	testConstrainedIntCP, err := newTestConstrainedInt(testConstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedInt() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ConstrainedInt, unsafe.Pointer(testConstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestConstrainedInt(bytes []byte) (*test_sm_ies.TestConstrainedInt, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ConstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestConstrainedInt((*C.TEST_ConstrainedInt_t)(unsafePtr))
}

func perDecodeTestConstrainedInt(bytes []byte) (*test_sm_ies.TestConstrainedInt, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ConstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestConstrainedInt((*C.TEST_ConstrainedInt_t)(unsafePtr))
}

func newTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) (*C.TEST_ConstrainedInt_t, error) {

	var err error
	testConstrainedIntC := C.TEST_ConstrainedInt_t{}

	attrCiAC := C.long(testConstrainedInt.AttrCiA)
	attrCiBC := C.long(testConstrainedInt.AttrCiB)
	attrCiCC := C.long(testConstrainedInt.AttrCiC)
	attrCiDC := C.long(testConstrainedInt.AttrCiD)
	attrCiEC := C.long(testConstrainedInt.AttrCiE)
	attrCiFC := C.long(testConstrainedInt.AttrCiF)
	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	testConstrainedIntC.attrCiA = attrCiAC
	testConstrainedIntC.attrCiB = attrCiBC
	testConstrainedIntC.attrCiC = attrCiCC
	testConstrainedIntC.attrCiD = attrCiDC
	testConstrainedIntC.attrCiE = attrCiEC
	testConstrainedIntC.attrCiF = attrCiFC

	return &testConstrainedIntC, nil
}

func decodeTestConstrainedInt(testConstrainedIntC *C.TEST_ConstrainedInt_t) (*test_sm_ies.TestConstrainedInt, error) {

	var err error
	testConstrainedInt := test_sm_ies.TestConstrainedInt{
	}

	testConstrainedInt.AttrCiA = int32(testConstrainedIntC.attrCiA)
	testConstrainedInt.AttrCiB = int32(testConstrainedIntC.attrCiB)
	testConstrainedInt.AttrCiC = int32(testConstrainedIntC.attrCiC)
	testConstrainedInt.AttrCiD = int32(testConstrainedIntC.attrCiD)
	testConstrainedInt.AttrCiE = int32(testConstrainedIntC.attrCiE)
	testConstrainedInt.AttrCiF = int32(testConstrainedIntC.attrCiF)

	return &testConstrainedInt, nil
}

func decodeTestConstrainedIntBytes(array [8]byte) (*test_sm_ies.TestConstrainedInt, error) {
	testConstrainedIntC := (*C.TEST_ConstrainedInt_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeTestConstrainedInt(testConstrainedIntC)
}
