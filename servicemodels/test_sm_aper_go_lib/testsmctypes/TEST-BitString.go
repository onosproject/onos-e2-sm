// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-BitString.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestBitString(testBitString *test_sm_ies.TestBitString) ([]byte, error) {
	testBitStringCP, err := newTestBitString(testBitString)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestBitString() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_BitString, unsafe.Pointer(testBitStringCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestBitString() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestBitString(testBitString *test_sm_ies.TestBitString) ([]byte, error) {
	testBitStringCP, err := newTestBitString(testBitString)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestBitString() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_BitString, unsafe.Pointer(testBitStringCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestBitString() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestBitString(bytes []byte) (*test_sm_ies.TestBitString, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_BitString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestBitString((*C.TEST_BitString_t)(unsafePtr))
}

func perDecodeTestBitString(bytes []byte) (*test_sm_ies.TestBitString, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_BitString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestBitString((*C.TEST_BitString_t)(unsafePtr))
}

func newTestBitString(testBitString *test_sm_ies.TestBitString) (*C.TEST_BitString_t, error) {

	var err error
	testBitStringC := C.TEST_BitString_t{}

	attrBs1C, err := newBitString(testBitString.AttrBs1)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	attrBs2C, err := newBitString(testBitString.AttrBs2)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	attrBs3C, err := newBitString(testBitString.AttrBs3)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	attrBs4C, err := newBitString(testBitString.AttrBs4)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	attrBs5C, err := newBitString(testBitString.AttrBs5)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	attrBs6C, err := newBitString(testBitString.AttrBs6)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	testBitStringC.attrBs1 = *attrBs1C
	testBitStringC.attrBs2 = *attrBs2C
	testBitStringC.attrBs3 = *attrBs3C
	testBitStringC.attrBs4 = *attrBs4C
	testBitStringC.attrBs5 = *attrBs5C
	testBitStringC.attrBs6 = *attrBs6C

	if testBitString.AttrBs7 != nil {
		attrBs7C, err := newBitString(testBitString.AttrBs7)
		if err != nil {
			return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
		}
		testBitStringC.attrBs7 = attrBs7C
	}
	return &testBitStringC, nil
}

func decodeTestBitString(testBitStringC *C.TEST_BitString_t) (*test_sm_ies.TestBitString, error) {

	var err error
	testBitString := test_sm_ies.TestBitString{}

	testBitString.AttrBs1, err = decodeBitString(&testBitStringC.attrBs1)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	testBitString.AttrBs2, err = decodeBitString(&testBitStringC.attrBs2)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	testBitString.AttrBs3, err = decodeBitString(&testBitStringC.attrBs3)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	testBitString.AttrBs4, err = decodeBitString(&testBitStringC.attrBs4)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	testBitString.AttrBs5, err = decodeBitString(&testBitStringC.attrBs5)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	testBitString.AttrBs6, err = decodeBitString(&testBitStringC.attrBs6)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	if testBitStringC.attrBs7 != nil {
		testBitString.AttrBs7, err = decodeBitString(testBitStringC.attrBs7)
		if err != nil {
			return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
		}
	}

	return &testBitString, nil
}
