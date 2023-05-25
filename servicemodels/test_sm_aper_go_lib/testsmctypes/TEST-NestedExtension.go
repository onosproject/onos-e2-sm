// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-NestedExtension.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestNestedExtension(testNestedExt *test_sm_ies.TestNestedExtension) ([]byte, error) {
	testNestedExtCP, err := newTestNestedExtension(testNestedExt)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestNestedExtension() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_NestedExtension, unsafe.Pointer(testNestedExtCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestNestedExtension() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeTestNestedExtension(testNestedExt *test_sm_ies.TestNestedExtension) ([]byte, error) {
	testNestedExtCP, err := newTestNestedExtension(testNestedExt)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestNestedExtension() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_NestedExtension, unsafe.Pointer(testNestedExtCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestNestedExtension() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestNestedExtension(bytes []byte) (*test_sm_ies.TestNestedExtension, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_NestedExtension)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestNestedExtension((*C.TEST_NestedExtension_t)(unsafePtr))
}

func PerDecodeTestNestedExtension(bytes []byte) (*test_sm_ies.TestNestedExtension, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_NestedExtension)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestNestedExtension((*C.TEST_NestedExtension_t)(unsafePtr))
}

func newTestNestedExtension(testNestedExt *test_sm_ies.TestNestedExtension) (*C.TEST_NestedCExtension_t, error) {

	testNestedExtC := C.TEST_NestedChoice_t{
		item1: C.long(testNestedExt.GetItem1()),
	}

	item2C, err := newOctetString(testNestedExt.GetItem2())
	if err != nil {
		return nil, err
	}
	testNestedExtC.item2 = item2C

	if testNestedExt.Ch != nil {
		chC, err := newAChoice(testNestedExt.GetCh())
		if err != nil {
			return nil, err
		}
		testNestedExtC.ch = chC
	}

	return &testNestedExtC, nil
}

func decodeTestNestedExtension(testNestedExtC *C.TEST_NestedExtension_t) (*test_sm_ies.TestNestedExtension, error) {

	var err error
	testNestedExt := new(test_sm_ies.TestNestedExtension)

	testNestedExt.Item1 = int32(testNestedExtC.item1)

	testNestedExt.Item2, err = decodeOctetString(testNestedExtC.item2)
	if err != nil {
		return nil, err
	}

	if testNestedExtC.ch != nil {
		testNestedExt.Ch, err = decodeAChoice(testNestedExtC.item2)
		if err != nil {
			return nil, err
		}

	}

	return testNestedExtension, nil
}
