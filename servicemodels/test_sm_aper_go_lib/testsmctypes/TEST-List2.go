// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-List2.h"
//#include "ItemExtensible.h"
import "C"
import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestList2(testList2 *test_sm_ies.TestList2) ([]byte, error) {
	testList2CP, err := newTestList2(testList2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_List2, unsafe.Pointer(testList2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestList2(testList2 *test_sm_ies.TestList2) ([]byte, error) {
	testList2CP, err := newTestList2(testList2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_List2, unsafe.Pointer(testList2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestList2(bytes []byte) (*test_sm_ies.TestList2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_List2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestList2((*C.TEST_List2_t)(unsafePtr))
}

func perDecodeTestList2(bytes []byte) (*test_sm_ies.TestList2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_List2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestList2((*C.TEST_List2_t)(unsafePtr))
}

func newTestList2(testList2 *test_sm_ies.TestList2) (*C.TEST_List2_t, error) {

	testList2C := new(C.TEST_List2_t)
	for _, ie := range testList2.GetValue() {
		ieC, err := newItemExtensible(ie)
		if err != nil {
			return nil, fmt.Errorf("newItemExtensible() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testList2C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return testList2C, nil
}

func decodeTestList2(testList2C *C.TEST_List2_t) (*test_sm_ies.TestList2, error) {

	var ieCount int
	testList2 := test_sm_ies.TestList2{}

	ieCount = int(testList2C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testList2C.list.array)) * uintptr(i)
		ieC := *(**C.ItemExtensible_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testList2C.list.array)) + offset))
		ie, err := decodeItemExtensible(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeItemExtensible() %s", err.Error())
		}
		testList2.Value = append(testList2.Value, ie)
	}

	return &testList2, nil
}
