// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: LicenseRef-Intel

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-Extension2.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestExtension2(testExt1 *test_sm_ies.TestExtension2) ([]byte, error) {
	testExtension2CP, err := newTestExtension2(testExt1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_Extension2, unsafe.Pointer(testExtension2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension2() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeTestExtension2(testExt1 *test_sm_ies.TestExtension2) ([]byte, error) {
	testExt1CP, err := newTestExtension2(testExt1)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_Extension2, unsafe.Pointer(testExt1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestExtension2(bytes []byte) (*test_sm_ies.TestExtension2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_Extension2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestExtension2((*C.TEST_Extension2_t)(unsafePtr))
}

func PerDecodeTestExtension2(bytes []byte) (*test_sm_ies.TestExtension2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_Extension2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestExtension2((*C.TEST_Extension2_t)(unsafePtr))
}

func newTestExtension2(testExtension2 *test_sm_ies.TestExtension2) (*C.TEST_Extension2_t, error) {

	testExtension2C := C.TEST_Extension2_t{
		item1: C.long(testExtension2.GetItem1()),
	}

	if testExtension2.GetItem2() != nil {
		item2C, err := newOctetString(testExtension2.GetItem2())
		if err != nil {
			return nil, err
		}
		testExtension2C.item2 = item2C
	}

	if testExtension2.Ext1 != nil {
		var ext1C C.TEST_Extension2__ext1_t
		switch *testExtension2.Ext1 {
		case test_sm_ies.EnumThree_ENUM_THREE_ONE:
			ext1C = C.TEST_Extension2__ext1_one
		case test_sm_ies.EnumThree_ENUM_THREE_TWO:
			ext1C = C.TEST_Extension2__ext1_two
		case test_sm_ies.EnumThree_ENUM_THREE_THREE:
			ext1C = C.TEST_Extension2__ext1_three
		default:
			return nil, fmt.Errorf("unexpected TestExtension2 Ext1 %v", testExtension2.Ext1)
		}

		testExtension2C.ext1 = &ext1C
	}

	if testExtension2.Ext2 != nil {
		var ext2C C.TEST_Extension2__ext2_t
		switch *testExtension2.Ext2 {
		case test_sm_ies.EnumThree_ENUM_THREE_ONE:
			ext2C = C.TEST_Extension2__ext2_one
		case test_sm_ies.EnumThree_ENUM_THREE_TWO:
			ext2C = C.TEST_Extension2__ext2_two
		case test_sm_ies.EnumThree_ENUM_THREE_THREE:
			ext2C = C.TEST_Extension2__ext2_three
		default:
			return nil, fmt.Errorf("unexpected TestExtension2 Ext2 %v", testExtension2.Ext2)
		}

		testExtension2C.ext2 = &ext2C
	}

	if testExtension2.Ext3 != nil {
		var ext3C C.TEST_Extension2__ext3_t
		switch *testExtension2.Ext3 {
		case test_sm_ies.EnumThree_ENUM_THREE_ONE:
			ext3C = C.TEST_Extension2__ext3_one
		case test_sm_ies.EnumThree_ENUM_THREE_TWO:
			ext3C = C.TEST_Extension2__ext3_two
		case test_sm_ies.EnumThree_ENUM_THREE_THREE:
			ext3C = C.TEST_Extension2__ext3_three
		default:
			return nil, fmt.Errorf("unexpected TestExtension2 Ext3 %v", testExtension2.Ext3)
		}

		testExtension2C.ext3 = &ext3C
	}

	return &testExtension2C, nil
}

func decodeTestExtension2(testExtension2C *C.TEST_Extension2_t) (*test_sm_ies.TestExtension2, error) {

	var err error
	testExtension2 := test_sm_ies.TestExtension2{}

	testExtension2.Item1 = int32(testExtension2C.item1)

	if testExtension2C.item2 != nil {
		testExtension2.Item2, err = decodeOctetString(testExtension2C.item2)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext1 != nil {
		ext1 := test_sm_ies.EnumThree(int32(*testExtension2C.ext1))
		testExtension2.Ext1 = &ext1
	}

	if testExtension2C.ext2 != nil {
		ext2 := test_sm_ies.EnumThree(int32(*testExtension2C.ext2))
		testExtension2.Ext2 = &ext2
	}

	if testExtension2C.ext3 != nil {
		ext3 := test_sm_ies.EnumThree(int32(*testExtension2C.ext3))
		testExtension2.Ext3 = &ext3
	}

	return &testExtension2, nil
}
