// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-PrintableString.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestPrintableString(testPrintableString *test_sm_ies.TestPrintableString) ([]byte, error) {
	testPrintableStringCP, err := newTestPrintableString(testPrintableString)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestPrintableString() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_PrintableString, unsafe.Pointer(testPrintableStringCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestPrintableString() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestPrintableString(testPrintableString *test_sm_ies.TestPrintableString) ([]byte, error) {
	testPrintableStringCP, err := newTestPrintableString(testPrintableString)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestPrintableString() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_PrintableString, unsafe.Pointer(testPrintableStringCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestPrintableString() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestPrintableString(bytes []byte) (*test_sm_ies.TestPrintableString, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_PrintableString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestPrintableString((*C.TEST_PrintableString_t)(unsafePtr))
}

func perDecodeTestPrintableString(bytes []byte) (*test_sm_ies.TestPrintableString, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_PrintableString)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestPrintableString((*C.TEST_PrintableString_t)(unsafePtr))
}

func newTestPrintableString(testPrintableString *test_sm_ies.TestPrintableString) (*C.TEST_PrintableString_t, error) {

	var err error
	testPrintableStringC := C.TEST_PrintableString_t{}

	attrPs1C, err := newPrintableString(testPrintableString.AttrPs1)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	attrPs2C, err := newPrintableString(testPrintableString.AttrPs2)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	attrPs3C, err := newPrintableString(testPrintableString.AttrPs3)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	attrPs4C, err := newPrintableString(testPrintableString.AttrPs4)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	attrPs5C, err := newPrintableString(testPrintableString.AttrPs5)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	attrPs6C, err := newPrintableString(testPrintableString.AttrPs6)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	testPrintableStringC.attrPs1 = *attrPs1C
	testPrintableStringC.attrPs2 = *attrPs2C
	testPrintableStringC.attrPs3 = *attrPs3C
	testPrintableStringC.attrPs4 = *attrPs4C
	testPrintableStringC.attrPs5 = *attrPs5C
	testPrintableStringC.attrPs6 = *attrPs6C

	if testPrintableString.AttrPs7 != nil {
		attrPs7C, err := newPrintableString(*testPrintableString.AttrPs7)
		if err != nil {
			return nil, fmt.Errorf("newPrintableString() %s", err.Error())
		}
		testPrintableStringC.attrPs7 = attrPs7C
	}

	return &testPrintableStringC, nil
}

func decodeTestPrintableString(testPrintableStringC *C.TEST_PrintableString_t) (*test_sm_ies.TestPrintableString, error) {

	var err error
	testPrintableString := test_sm_ies.TestPrintableString{}

	testPrintableString.AttrPs1, err = decodePrintableString(&testPrintableStringC.attrPs1)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testPrintableString.AttrPs2, err = decodePrintableString(&testPrintableStringC.attrPs2)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testPrintableString.AttrPs3, err = decodePrintableString(&testPrintableStringC.attrPs3)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testPrintableString.AttrPs4, err = decodePrintableString(&testPrintableStringC.attrPs4)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testPrintableString.AttrPs5, err = decodePrintableString(&testPrintableStringC.attrPs5)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testPrintableString.AttrPs6, err = decodePrintableString(&testPrintableStringC.attrPs6)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	if testPrintableStringC.attrPs7 != nil {
		res, err := decodePrintableString(testPrintableStringC.attrPs7)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
		testPrintableString.AttrPs7 = &res
	}
	return &testPrintableString, nil
}
