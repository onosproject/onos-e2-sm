// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-UnconstrainedReal.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestUnconstrainedReal(testUnconstrainedReal *test_sm_ies.TestUnconstrainedReal) ([]byte, error) {
	testUnconstrainedRealCP, err := newTestUnconstrainedReal(testUnconstrainedReal)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestUnconstrainedReal() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_UnconstrainedReal, unsafe.Pointer(testUnconstrainedRealCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestUnconstrainedReal() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestUnconstrainedReal(testUnconstrainedReal *test_sm_ies.TestUnconstrainedReal) ([]byte, error) {
	testUnconstrainedRealCP, err := newTestUnconstrainedReal(testUnconstrainedReal)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestUnconstrainedReal() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_UnconstrainedReal, unsafe.Pointer(testUnconstrainedRealCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestUnconstrainedReal() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestUnconstrainedReal(bytes []byte) (*test_sm_ies.TestUnconstrainedReal, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_UnconstrainedReal)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestUnconstrainedReal((*C.TEST_UnconstrainedReal_t)(unsafePtr))
}

func perDecodeTestUnconstrainedReal(bytes []byte) (*test_sm_ies.TestUnconstrainedReal, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_UnconstrainedReal)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestUnconstrainedReal((*C.TEST_UnconstrainedReal_t)(unsafePtr))
}

func newTestUnconstrainedReal(testUnconstrainedReal *test_sm_ies.TestUnconstrainedReal) (*C.TEST_UnconstrainedReal_t, error) {

	testUnconstrainedRealC := C.TEST_UnconstrainedReal_t{
		attrUcrA: C.double(testUnconstrainedReal.AttrUcrA),
		attrUcrB: C.double(testUnconstrainedReal.AttrUcrB),
	}

	return &testUnconstrainedRealC, nil
}

func decodeTestUnconstrainedReal(testUnconstrainedRealC *C.TEST_UnconstrainedReal_t) (*test_sm_ies.TestUnconstrainedReal, error) {

	testUnconstrainedReal := test_sm_ies.TestUnconstrainedReal{
		AttrUcrA: float64(testUnconstrainedRealC.attrUcrA),
		AttrUcrB: float64(testUnconstrainedRealC.attrUcrB),
	}

	return &testUnconstrainedReal, nil
}
