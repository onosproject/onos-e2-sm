// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-UnconstrainedInt.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestUnconstrainedInt(testUnconstrainedInt *test_sm_ies.TestUnconstrainedInt) ([]byte, error) {
	testUnconstrainedIntCP, err := newTestUnconstrainedInt(testUnconstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestUnconstrainedInt() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_UnconstrainedInt, unsafe.Pointer(testUnconstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestUnconstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestUnconstrainedInt(testUnconstrainedInt *test_sm_ies.TestUnconstrainedInt) ([]byte, error) {
	testUnconstrainedIntCP, err := newTestUnconstrainedInt(testUnconstrainedInt)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestUnconstrainedInt() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_UnconstrainedInt, unsafe.Pointer(testUnconstrainedIntCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestUnconstrainedInt() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestUnconstrainedInt(bytes []byte) (*test_sm_ies.TestUnconstrainedInt, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_UnconstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestUnconstrainedInt((*C.TEST_UnconstrainedInt_t)(unsafePtr))
}

func perDecodeTestUnconstrainedInt(bytes []byte) (*test_sm_ies.TestUnconstrainedInt, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_UnconstrainedInt)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestUnconstrainedInt((*C.TEST_UnconstrainedInt_t)(unsafePtr))
}

func newTestUnconstrainedInt(testUnconstrainedInt *test_sm_ies.TestUnconstrainedInt) (*C.TEST_UnconstrainedInt_t, error) {

	testUnconstrainedIntC := C.TEST_UnconstrainedInt_t{
		attrUciA: C.long(testUnconstrainedInt.AttrUciA),
		attrUciB: C.long(testUnconstrainedInt.AttrUciB),
	}

	return &testUnconstrainedIntC, nil
}

func decodeTestUnconstrainedInt(testUnconstrainedIntC *C.TEST_UnconstrainedInt_t) (*test_sm_ies.TestUnconstrainedInt, error) {

	testUnconstrainedInt := test_sm_ies.TestUnconstrainedInt{
		AttrUciA: int32(testUnconstrainedIntC.attrUciA),
		AttrUciB: int32(testUnconstrainedIntC.attrUciB),
	}

	return &testUnconstrainedInt, nil
}
