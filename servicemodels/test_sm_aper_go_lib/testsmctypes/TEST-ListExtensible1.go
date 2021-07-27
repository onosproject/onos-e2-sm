// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ListExtensible1.h"
//#include "Item.h"
import "C"
import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestListExtensible1(testListExtensible1 *test_sm_ies.TestListExtensible1) ([]byte, error) {
	testListExtensible1CP, err := newTestListExtensible1(testListExtensible1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ListExtensible1, unsafe.Pointer(testListExtensible1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestListExtensible1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestListExtensible1(testListExtensible1 *test_sm_ies.TestListExtensible1) ([]byte, error) {
	testListExtensible1CP, err := newTestListExtensible1(testListExtensible1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ListExtensible1, unsafe.Pointer(testListExtensible1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestListExtensible1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestListExtensible1(bytes []byte) (*test_sm_ies.TestListExtensible1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ListExtensible1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestListExtensible1((*C.TEST_ListExtensible1_t)(unsafePtr))
}

func perDecodeTestListExtensible1(bytes []byte) (*test_sm_ies.TestListExtensible1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ListExtensible1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestListExtensible1((*C.TEST_ListExtensible1_t)(unsafePtr))
}

func newTestListExtensible1(testListExtensible1 *test_sm_ies.TestListExtensible1) (*C.TEST_ListExtensible1_t, error) {

	testListExtensible1C := new(C.TEST_ListExtensible1_t)
	for _, ie := range testListExtensible1.GetValue() {
		ieC, err := newItem(ie)
		if err != nil {
			return nil, fmt.Errorf("newItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testListExtensible1C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return testListExtensible1C, nil
}

func decodeTestListExtensible1(testListExtensible1C *C.TEST_ListExtensible1_t) (*test_sm_ies.TestListExtensible1, error) {

	var ieCount int
	testListExtensible1 := test_sm_ies.TestListExtensible1{}

	ieCount = int(testListExtensible1C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testListExtensible1C.list.array)) * uintptr(i)
		ieC := *(**C.Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testListExtensible1C.list.array)) + offset))
		ie, err := decodeItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeItem() %s", err.Error())
		}
		testListExtensible1.Value = append(testListExtensible1.Value, ie)
	}

	return &testListExtensible1, nil
}
