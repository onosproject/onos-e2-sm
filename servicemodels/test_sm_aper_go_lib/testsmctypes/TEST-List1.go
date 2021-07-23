// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-List1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
//#include ".h" //ToDo - include correct .h file for corresponding C-struct of "Repeated" field or other anonymous structure defined in .h file
import "C"
import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestList1(testList1 *test_sm_ies.TestList1) ([]byte, error) {
	testList1CP, err := newTestList1(testList1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_List1, unsafe.Pointer(testList1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestList1(testList1 *test_sm_ies.TestList1) ([]byte, error) {
	testList1CP, err := newTestList1(testList1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_List1, unsafe.Pointer(testList1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestList1(bytes []byte) (*test_sm_ies.TestList1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_List1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestList1((*C.TEST_List1_t)(unsafePtr))
}

func perDecodeTestList1(bytes []byte) (*test_sm_ies.TestList1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_List1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestList1((*C.TEST_List1_t)(unsafePtr))
}

func newTestList1(testList1 *test_sm_ies.TestList1) (*C.TEST_List1_t, error) {

	testList1C := new(C.TEST_List1_t)         //ToDo - verify correctness of the variable's name
	for _, ie := range testList1.GetValue() { //ToDo - Verify if GetSmth() function is called correctly
		ieC, err := newItem(ie)
		if err != nil {
			return nil, fmt.Errorf("newItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testList1C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &testList1C, nil
}

func decodeTestList1(testList1C *C.TEST_List1_t) (*test_sm_ies.TestList1, error) {

	var ieCount int

	testList1 := test_sm_ies.TestList1{}

	ieCount = int(testList1C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testList1C.list.array)) * uintptr(i)
		ieC := *(**C.Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testList1C.list.array)) + offset))
		ie, err := decodeItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeItem() %s", err.Error())
		}
		testList1.Value = append(testList1.Value, ie)
	}

	return &testList1, nil
}

func decodeTestList1Bytes(array [8]byte) (*test_sm_ies.TestList1, error) {
	testList1C := (*C.TEST_List1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeTestList1(testList1C)
}
