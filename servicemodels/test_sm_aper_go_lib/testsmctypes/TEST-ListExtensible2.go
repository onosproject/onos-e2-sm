// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ListExtensible2.h"
//#include "ItemExtensible.h"
import "C"
import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestListExtensible2(testListExtensible2 *test_sm_ies.TestListExtensible2) ([]byte, error) {
	testListExtensible2CP, err := newTestListExtensible2(testListExtensible2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ListExtensible2, unsafe.Pointer(testListExtensible2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestListExtensible2(testListExtensible2 *test_sm_ies.TestListExtensible2) ([]byte, error) {
	testListExtensible2CP, err := newTestListExtensible2(testListExtensible2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ListExtensible2, unsafe.Pointer(testListExtensible2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestListExtensible2(bytes []byte) (*test_sm_ies.TestListExtensible2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ListExtensible2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestListExtensible2((*C.TEST_ListExtensible2_t)(unsafePtr))
}

func perDecodeTestListExtensible2(bytes []byte) (*test_sm_ies.TestListExtensible2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ListExtensible2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestListExtensible2((*C.TEST_ListExtensible2_t)(unsafePtr))
}

func newTestListExtensible2(testListExtensible2 *test_sm_ies.TestListExtensible2) (*C.TEST_ListExtensible2_t, error) {

	testListExtensible2C := new(C.TEST_ListExtensible2_t)
	for _, ie := range testListExtensible2.GetValue() {
		ieC, err := newItemExtensible(ie)
		if err != nil {
			return nil, fmt.Errorf("newItemExtensible() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testListExtensible2C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return testListExtensible2C, nil
}

func decodeTestListExtensible2(testListExtensible2C *C.TEST_ListExtensible2_t) (*test_sm_ies.TestListExtensible2, error) {

	var ieCount int
	testListExtensible2 := test_sm_ies.TestListExtensible2{}

	ieCount = int(testListExtensible2C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testListExtensible2C.list.array)) * uintptr(i)
		ieC := *(**C.ItemExtensible_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testListExtensible2C.list.array)) + offset))
		ie, err := decodeItemExtensible(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeItemExtensible() %s", err.Error())
		}
		testListExtensible2.Value = append(testListExtensible2.Value, ie)
	}

	return &testListExtensible2, nil
}
