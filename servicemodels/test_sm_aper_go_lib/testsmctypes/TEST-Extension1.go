// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-Extension1.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestExtension1(testExt1 *test_sm_ies.TestExtension1) ([]byte, error) {
	testExtension1CP, err := newTestExtension1(testExt1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_Extension1, unsafe.Pointer(testExtension1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeTestExtension1(testExt1 *test_sm_ies.TestExtension1) ([]byte, error) {
	testExt1CP, err := newTestExtension1(testExt1)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_Extension1, unsafe.Pointer(testExt1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestExtension1(bytes []byte) (*test_sm_ies.TestExtension1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_Extension1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestExtension1((*C.TEST_Extension1_t)(unsafePtr))
}

func PerDecodeTestExtension1(bytes []byte) (*test_sm_ies.TestExtension1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_Extension1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestExtension1((*C.TEST_Extension1_t)(unsafePtr))
}

func newTestExtension1(testExtension1 *test_sm_ies.TestExtension1) (*C.TEST_Extension1_t, error) {

	testExtension1C := C.TEST_Extension1_t{
		item1: C.long(testExtension1.GetItem1()),
	}

	if testExtension1.GetItem2() != nil {
		item2C, err := newOctetString(testExtension1.GetItem2())
		if err != nil {
			return nil, err
		}
		testExtension1C.item2 = item2C
	}

	if testExtension1.Ext1 != nil {
		ext1C, err := newOctetString(testExtension1.GetExt1())
		if err != nil {
			return nil, err
		}

		testExtension1C.ext1 = ext1C
	}

	return &testExtension1C, nil
}

func decodeTestExtension1(testExtension1C *C.TEST_Extension1_t) (*test_sm_ies.TestExtension1, error) {

	var err error
	testExtension1 := test_sm_ies.TestExtension1{}

	testExtension1.Item1 = int32(testExtension1C.item1)

	if testExtension1C.item2 != nil {
		testExtension1.Item2, err = decodeOctetString(testExtension1C.item2)
		if err != nil {
			return nil, err
		}
	}

	if testExtension1C.ext1 != nil {
		testExtension1.Ext1, err = decodeOctetString(testExtension1C.ext1)
		if err != nil {
			return nil, err
		}
	}

	return &testExtension1, nil
}
