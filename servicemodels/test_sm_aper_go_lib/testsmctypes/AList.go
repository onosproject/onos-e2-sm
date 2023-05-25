// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "AList.h"
//#include "TEST-NestedExtension.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeAList(alist *test_sm_ies.Alist) ([]byte, error) {
	alistCP, err := newAList(alist)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_AList, unsafe.Pointer(alistCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeAList(alist *test_sm_ies.Alist) ([]byte, error) {
	alistCP, err := newAList(alist)
	if err != nil {
		return nil, fmt.Errorf("perEncodeAList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_AList, unsafe.Pointer(alistCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeAList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeAList(bytes []byte) (*test_sm_ies.Alist, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_AList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeAList((*C.AList_t)(unsafePtr))
}

func perDecodeAList(bytes []byte) (*test_sm_ies.Alist, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_AList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeAList((*C.AList_t)(unsafePtr))
}

func newAList(alist *test_sm_ies.Alist) (*C.AList_t, error) {

	alistC := new(C.AList_t)

	itemC := new(C.struct_AList__item)
	for _, ie := range alist.GetItem() {
		ieC, err := newTestNestedExtension(ie)
		if err != nil {
			return nil, fmt.Errorf("newTestNestedExtension() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(itemC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}
	alistC.item = *itemC

	return alistC, nil
}

func decodeAList(alistC *C.AList_t) (*test_sm_ies.Alist, error) {

	var ieCount int
	alist := test_sm_ies.Alist{}

	ieCount = int(alistC.item.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(alistC.item.list.array)) * uintptr(i)
		ieC := *(**C.TEST_NestedExtension_t)(unsafe.Pointer(uintptr(unsafe.Pointer(alistC.item.list.array)) + offset))
		ie, err := decodeTestNestedExtension(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeTestNestedExtension() %s", err.Error())
		}
		alist.Item = append(alist.Item, ie)
	}

	return &alist, nil
}

func decodeAListBytes(array [8]byte) (*test_sm_ies.Alist, error) {
	choiceC := (*C.AList_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeAList(choiceC)
}
