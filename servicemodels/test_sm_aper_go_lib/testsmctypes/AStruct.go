// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "AStruct.h"
//#include "TEST-NestedExtension.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeAStruct(astruct *test_sm_ies.Astruct) ([]byte, error) {
	astructCP, err := newAStruct(astruct)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAStruct() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_AStruct, unsafe.Pointer(astructCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAStruct() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeAStruct(astruct *test_sm_ies.Astruct) ([]byte, error) {
	astructCP, err := newAStruct(astruct)
	if err != nil {
		return nil, fmt.Errorf("perEncodeAStruct() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_AStruct, unsafe.Pointer(astructCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeAStruct() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeAStruct(bytes []byte) (*test_sm_ies.Astruct, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_AStruct)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeAStruct((*C.AStruct_t)(unsafePtr))
}

func perDecodeAStruct(bytes []byte) (*test_sm_ies.Astruct, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_AStruct)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeAStruct((*C.AStruct_t)(unsafePtr))
}

func newAStruct(astruct *test_sm_ies.Astruct) (*C.AStruct_t, error) {

	astructC := new(C.AStruct_t)

	itemC := new(C.struct_AStruct__item)
	for _, ie := range astruct.GetItem() {
		ieC, err := newTestNestedExtension(ie)
		if err != nil {
			return nil, fmt.Errorf("newTestNestedExtension() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(itemC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}
	astructC.item = *itemC

	return astructC, nil
}

func decodeAStruct(astructC *C.AStruct_t) (*test_sm_ies.Astruct, error) {

	var ieCount int
	astruct := test_sm_ies.Astruct{}

	ieCount = int(astructC.item.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(astructC.item.list.array)) * uintptr(i)
		ieC := *(**C.TEST_NestedExtension_t)(unsafe.Pointer(uintptr(unsafe.Pointer(astructC.item.list.array)) + offset))
		ie, err := decodeTestNestedExtension(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeTestNestedExtension() %s", err.Error())
		}
		astruct.Item = append(astruct.Item, ie)
	}

	return &astruct, nil
}

func decodeAStructBytes(array [8]byte) (*test_sm_ies.Astruct, error) {
	choiceC := (*C.AStruct_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeAStruct(choiceC)
}
