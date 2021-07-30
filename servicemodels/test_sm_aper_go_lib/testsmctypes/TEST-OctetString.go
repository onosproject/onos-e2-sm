// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-OctetString.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestOctetString(testOctetString *test_sm_ies.TestOctetString) ([]byte, error) {
	testOctetStringCP, err := newTestOctetString(testOctetString)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestOctetString() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_OctetString, unsafe.Pointer(testOctetStringCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestOctetString() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestOctetString(testOctetString *test_sm_ies.TestOctetString) ([]byte, error) {
	testOctetStringCP, err := newTestOctetString(testOctetString)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestOctetString() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_OctetString, unsafe.Pointer(testOctetStringCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestOctetString() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestOctetString(bytes []byte) (*test_sm_ies.TestOctetString, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_OctetString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestOctetString((*C.TEST_OctetString_t)(unsafePtr))
}

func perDecodeTestOctetString(bytes []byte) (*test_sm_ies.TestOctetString, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_OctetString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestOctetString((*C.TEST_OctetString_t)(unsafePtr))
}

func newTestOctetString(testOctetString *test_sm_ies.TestOctetString) (*C.TEST_OctetString_t, error) {

	var err error
	testOctetStringC := C.TEST_OctetString_t{}

	attrOs1C, err := newOctetString(testOctetString.AttrOs1)
	if err != nil {
		return nil, err
	}
	attrOs2C, err := newOctetString(testOctetString.AttrOs2)
	if err != nil {
		return nil, err
	}
	attrOs3C, err := newOctetString(testOctetString.AttrOs3)
	if err != nil {
		return nil, err
	}
	attrOs4C, err := newOctetString(testOctetString.AttrOs4)
	if err != nil {
		return nil, err
	}
	attrOs5C, err := newOctetString(testOctetString.AttrOs5)
	if err != nil {
		return nil, err
	}
	attrOs6C, err := newOctetString(testOctetString.AttrOs6)
	if err != nil {
		return nil, err
	}

	testOctetStringC.attrOs1 = *attrOs1C
	testOctetStringC.attrOs2 = *attrOs2C
	testOctetStringC.attrOs3 = *attrOs3C
	testOctetStringC.attrOs4 = *attrOs4C
	testOctetStringC.attrOs5 = *attrOs5C
	testOctetStringC.attrOs6 = *attrOs6C

	if testOctetString.AttrOs7 != nil {
		attrOs7C, err := newOctetString(testOctetString.AttrOs7)
		if err != nil {
			return nil, err
		}
		testOctetStringC.attrOs7 = attrOs7C
	}

	return &testOctetStringC, nil
}

func decodeTestOctetString(testOctetStringC *C.TEST_OctetString_t) (*test_sm_ies.TestOctetString, error) {

	var err error
	testOctetString := test_sm_ies.TestOctetString{}

	testOctetString.AttrOs1, err = decodeOctetString(&testOctetStringC.attrOs1)
	if err != nil {
		return nil, err
	}
	testOctetString.AttrOs2, err = decodeOctetString(&testOctetStringC.attrOs2)
	if err != nil {
		return nil, err
	}
	testOctetString.AttrOs3, err = decodeOctetString(&testOctetStringC.attrOs3)
	if err != nil {
		return nil, err
	}
	testOctetString.AttrOs4, err = decodeOctetString(&testOctetStringC.attrOs4)
	if err != nil {
		return nil, err
	}
	testOctetString.AttrOs5, err = decodeOctetString(&testOctetStringC.attrOs5)
	if err != nil {
		return nil, err
	}
	testOctetString.AttrOs6, err = decodeOctetString(&testOctetStringC.attrOs6)
	if err != nil {
		return nil, err
	}

	if testOctetStringC.attrOs7 != nil {
		testOctetString.AttrOs7, err = decodeOctetString(testOctetStringC.attrOs7)
		if err != nil {
			return nil, err
		}
	}

	return &testOctetString, nil
}
