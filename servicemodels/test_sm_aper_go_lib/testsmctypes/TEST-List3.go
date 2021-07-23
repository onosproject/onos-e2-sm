// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-List3.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
//#include ".h" //ToDo - include correct .h file for corresponding C-struct of "Repeated" field or other anonymous structure defined in .h file
import "C"
import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestList3(testList3 *test_sm_ies.TestList3) ([]byte, error) {
	testList3CP, err := newTestList3(testList3)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList3() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_List3, unsafe.Pointer(testList3CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestList3() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestList3(testList3 *test_sm_ies.TestList3) ([]byte, error) {
	testList3CP, err := newTestList3(testList3)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList3() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_List3, unsafe.Pointer(testList3CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestList3() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestList3(bytes []byte) (*test_sm_ies.TestList3, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_List3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestList3((*C.TEST_List3_t)(unsafePtr))
}

func perDecodeTestList3(bytes []byte) (*test_sm_ies.TestList3, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_List3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestList3((*C.TEST_List3_t)(unsafePtr))
}

func newTestList3(testList3 *test_sm_ies.TestList3) (*C.TEST_List3_t, error) {

	testList3C := new(C.TEST_List3_t)         //ToDo - verify correctness of the variable's name
	for _, ie := range testList3.GetValue() { //ToDo - Verify if GetSmth() function is called correctly
		ieC, err := newTestFullyOptionalSequence(ie)
		if err != nil {
			return nil, fmt.Errorf("newTestFullyOptionalSequence() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(testList3C), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &testList3C, nil
}

func decodeTestList3(testList3C *C.TEST_List3_t) (*test_sm_ies.TestList3, error) {

	var ieCount int

	testList3 := test_sm_ies.TestList3{}

	ieCount = int(testList3C.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(testList3C.list.array)) * uintptr(i)
		ieC := *(**C.TestFullyOptionalSequence_t)(unsafe.Pointer(uintptr(unsafe.Pointer(testList3C.list.array)) + offset))
		ie, err := decodeTestFullyOptionalSequence(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeTestFullyOptionalSequence() %s", err.Error())
		}
		testList3.Value = append(testList3.Value, ie)
	}

	return &testList3, nil
}

func decodeTestList3Bytes(array [8]byte) (*test_sm_ies.TestList3, error) {
	testList3C := (*C.TEST_List3_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeTestList3(testList3C)
}
