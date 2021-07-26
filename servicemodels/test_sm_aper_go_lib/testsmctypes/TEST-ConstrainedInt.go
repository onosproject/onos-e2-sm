// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ConstrainedInt.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) ([]byte, error) {
	testConstrainedIntCP, err := newTestConstrainedInt(testConstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedInt() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ConstrainedInt, unsafe.Pointer(testConstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) ([]byte, error) {
	testConstrainedIntCP, err := newTestConstrainedInt(testConstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedInt() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ConstrainedInt, unsafe.Pointer(testConstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestConstrainedInt(bytes []byte) (*test_sm_ies.TestConstrainedInt, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ConstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestConstrainedInt((*C.TEST_ConstrainedInt_t)(unsafePtr))
}

func perDecodeTestConstrainedInt(bytes []byte) (*test_sm_ies.TestConstrainedInt, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ConstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestConstrainedInt((*C.TEST_ConstrainedInt_t)(unsafePtr))
}

func newTestConstrainedInt(testConstrainedInt *test_sm_ies.TestConstrainedInt) (*C.TEST_ConstrainedInt_t, error) {

	testConstrainedIntC := C.TEST_ConstrainedInt_t{
		attrCiA: C.long(testConstrainedInt.AttrCiA),
		attrCiB: C.ulong(testConstrainedInt.AttrCiB),
		attrCiC: C.long(testConstrainedInt.AttrCiC),
		attrCiD: C.long(testConstrainedInt.AttrCiD),
		attrCiE: C.long(testConstrainedInt.AttrCiE),
		attrCiF: C.long(testConstrainedInt.AttrCiF),
	}

	return &testConstrainedIntC, nil
}

func decodeTestConstrainedInt(testConstrainedIntC *C.TEST_ConstrainedInt_t) (*test_sm_ies.TestConstrainedInt, error) {

	testConstrainedInt := test_sm_ies.TestConstrainedInt{
		AttrCiA: int32(testConstrainedIntC.attrCiA),
		AttrCiB: int32(testConstrainedIntC.attrCiB),
		AttrCiC: int32(testConstrainedIntC.attrCiC),
		AttrCiD: int32(testConstrainedIntC.attrCiD),
		AttrCiE: int32(testConstrainedIntC.attrCiE),
		AttrCiF: int32(testConstrainedIntC.attrCiF),
	}

	return &testConstrainedInt, nil
}
