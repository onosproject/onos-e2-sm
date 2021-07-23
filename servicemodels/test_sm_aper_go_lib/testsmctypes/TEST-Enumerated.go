// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-Enumerated.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestEnumerated(testEnumerated *test_sm_ies.TestEnumerated) ([]byte, error) {
	testEnumeratedCP, err := newTestEnumerated(testEnumerated)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_Enumerated, unsafe.Pointer(testEnumeratedCP)) //ToDo - change name of C-encoder tag
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestEnumerated() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestEnumerated(testEnumerated *test_sm_ies.TestEnumerated) ([]byte, error) {
	testEnumeratedCP, err := newTestEnumerated(testEnumerated)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_Enumerated, unsafe.Pointer(testEnumeratedCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestEnumerated() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestEnumerated(bytes []byte) (*test_sm_ies.TestEnumerated, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_Enumerated)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestEnumerated((*C.TEST_Enumerated_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeTestEnumerated(bytes []byte) (*test_sm_ies.TestEnumerated, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_Enumerated)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestEnumerated((*C.TEST_Enumerated_t)(unsafePtr))
}

func newTestEnumerated(testEnumerated *test_sm_ies.TestEnumerated) (*C.TEST_Enumerated_t, error) {
	var ret C.TEST_Enumerated_t
	switch *testEnumerated {
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM1:
		ret = C.TEST_Enumerated_enum1 //ToDo - double-check correctness of the name
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM2:
		ret = C.TEST_Enumerated_enum2 //ToDo - double-check correctness of the name
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM3:
		ret = C.TEST_Enumerated_enum3 //ToDo - double-check correctness of the name
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM4:
		ret = C.TEST_Enumerated_enum4 //ToDo - double-check correctness of the name
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM5:
		ret = C.TEST_Enumerated_enum5 //ToDo - double-check correctness of the name
	case test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM6:
		ret = C.TEST_Enumerated_enum6 //ToDo - double-check correctness of the name
	default:
		return nil, fmt.Errorf("unexpected TestEnumerated %v", testEnumerated)
	}

	return &ret, nil
}

func decodeTestEnumerated(testEnumeratedC *C.TEST_Enumerated_t) (*test_sm_ies.TestEnumerated, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	testEnumerated := test_sm_ies.TestEnumerated(int32(*testEnumeratedC))

	return &testEnumerated, nil
}

func decodeTestEnumeratedBytes(array [8]byte) (*test_sm_ies.TestEnumerated, error) { //ToDo - Check addressing correct structure in Protobuf
	testEnumeratedC := (*C.TestEnumerated_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))

	return decodeTestEnumerated(testEnumeratedC)
}
