// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ListExtensible3.h"
//#include "TEST-FullyOptionalSequence.h"
import "C"
import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestListExtensible3(testListExtensible3 *test_sm_ies.TestListExtensible3) ([]byte, error) {
	testListExtensible3CP, err := newTestListExtensible3(testListExtensible3)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible3() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ListExtensible3, unsafe.Pointer(testListExtensible3CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible3() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestListExtensible3(testListExtensible3 *test_sm_ies.TestListExtensible3) ([]byte, error) {
	testListExtensible3CP, err := newTestListExtensible3(testListExtensible3)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible3() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ListExtensible3, unsafe.Pointer(testListExtensible3CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible3() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestListExtensible3(bytes []byte) (*test_sm_ies.TestListExtensible3, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ListExtensible3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestListExtensible3((*C.TEST_ListExtensible3_t)(unsafePtr))
}

func perDecodeTestListExtensible3(bytes []byte) (*test_sm_ies.TestListExtensible3, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ListExtensible3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestListExtensible3((*C.TEST_ListExtensible3_t)(unsafePtr))
}

func newTestListExtensible3(testListExtensible3 *test_sm_ies.TestListExtensible3) (*C.TEST_ListExtensible3_t, error) {

	testListExtensible3C := new(C.TEST_ListExtensible3_t)
	for _, ie := range testListExtensible3.GetValue() {
		ieC, err := newTestFullyOptionalSequence(ie)
		if err != nil {
			return nil, fmt.Errorf("newTestFullyOptionalSequence() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testListExtensible3C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return testListExtensible3C, nil
}

func decodeTestListExtensible3(testListExtensible3C *C.TEST_ListExtensible3_t) (*test_sm_ies.TestListExtensible3, error) {

	var ieCount int
	testListExtensible3 := test_sm_ies.TestListExtensible3{}

	ieCount = int(testListExtensible3C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testListExtensible3C.list.array)) * uintptr(i)
		ieC := *(**C.TEST_FullyOptionalSequence_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testListExtensible3C.list.array)) + offset))
		ie, err := decodeTestFullyOptionalSequence(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeTestFullyOptionalSequence() %s", err.Error())
		}
		testListExtensible3.Value = append(testListExtensible3.Value, ie)
	}

	return &testListExtensible3, nil
}
