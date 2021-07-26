// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ConstrainedReal.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestConstrainedReal(testConstrainedReal *test_sm_ies.TestConstrainedReal) ([]byte, error) {
	testConstrainedRealCP, err := newTestConstrainedReal(testConstrainedReal)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedReal() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ConstrainedReal, unsafe.Pointer(testConstrainedRealCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedReal() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestConstrainedReal(testConstrainedReal *test_sm_ies.TestConstrainedReal) ([]byte, error) {
	testConstrainedRealCP, err := newTestConstrainedReal(testConstrainedReal)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedReal() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ConstrainedReal, unsafe.Pointer(testConstrainedRealCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedReal() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestConstrainedReal(bytes []byte) (*test_sm_ies.TestConstrainedReal, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ConstrainedReal)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestConstrainedReal((*C.TEST_ConstrainedReal_t)(unsafePtr))
}

func perDecodeTestConstrainedReal(bytes []byte) (*test_sm_ies.TestConstrainedReal, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ConstrainedReal)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestConstrainedReal((*C.TEST_ConstrainedReal_t)(unsafePtr))
}

func newTestConstrainedReal(testConstrainedReal *test_sm_ies.TestConstrainedReal) (*C.TEST_ConstrainedReal_t, error) {

	testConstrainedRealC := C.TEST_ConstrainedReal_t{
		attrCrA: C.double(testConstrainedReal.AttrCrA),
		attrCrB: C.double(testConstrainedReal.AttrCrB),
		attrCrC: C.double(testConstrainedReal.AttrCrC),
		attrCrD: C.double(testConstrainedReal.AttrCrD),
		attrCrE: C.double(testConstrainedReal.AttrCrE),
		attrCrF: C.double(testConstrainedReal.AttrCrF),
	}

	return &testConstrainedRealC, nil
}

func decodeTestConstrainedReal(testConstrainedRealC *C.TEST_ConstrainedReal_t) (*test_sm_ies.TestConstrainedReal, error) {

	testConstrainedReal := test_sm_ies.TestConstrainedReal{
		AttrCrA: float64(testConstrainedRealC.attrCrA),
		AttrCrB: float64(testConstrainedRealC.attrCrB),
		AttrCrC: float64(testConstrainedRealC.attrCrC),
		AttrCrD: float64(testConstrainedRealC.attrCrD),
		AttrCrE: float64(testConstrainedRealC.attrCrE),
		AttrCrF: float64(testConstrainedRealC.attrCrF),
	}

	return &testConstrainedReal, nil
}
