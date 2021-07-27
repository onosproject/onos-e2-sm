// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-EnumeratedExtensible.h"
import "C"
import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestEnumeratedExtensible(testEnumeratedExtensible *test_sm_ies.TestEnumeratedExtensible) ([]byte, error) {
	testEnumeratedExtensibleCP, err := newTestEnumeratedExtensible(testEnumeratedExtensible)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_EnumeratedExtensible, unsafe.Pointer(testEnumeratedExtensibleCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestEnumeratedExtensible() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestEnumeratedExtensible(testEnumeratedExtensible *test_sm_ies.TestEnumeratedExtensible) ([]byte, error) {
	testEnumeratedExtensibleCP, err := newTestEnumeratedExtensible(testEnumeratedExtensible)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_EnumeratedExtensible, unsafe.Pointer(testEnumeratedExtensibleCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestEnumeratedExtensible() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestEnumeratedExtensible(bytes []byte) (*test_sm_ies.TestEnumeratedExtensible, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_EnumeratedExtensible)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestEnumeratedExtensible((*C.TEST_EnumeratedExtensible_t)(unsafePtr))
}

func perDecodeTestEnumeratedExtensible(bytes []byte) (*test_sm_ies.TestEnumeratedExtensible, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_EnumeratedExtensible)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestEnumeratedExtensible((*C.TEST_EnumeratedExtensible_t)(unsafePtr))
}

func newTestEnumeratedExtensible(testEnumeratedExtensible *test_sm_ies.TestEnumeratedExtensible) (*C.TEST_EnumeratedExtensible_t, error) {
	var ret C.TEST_EnumeratedExtensible_t
	switch *testEnumeratedExtensible {
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM1:
		ret = C.TEST_EnumeratedExtensible_enum1
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM2:
		ret = C.TEST_EnumeratedExtensible_enum2
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM3:
		ret = C.TEST_EnumeratedExtensible_enum3
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM4:
		ret = C.TEST_EnumeratedExtensible_enum4
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM5:
		ret = C.TEST_EnumeratedExtensible_enum5
	case test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6:
		ret = C.TEST_EnumeratedExtensible_enum6
	default:
		return nil, fmt.Errorf("unexpected TestEnumeratedExtensible %v", testEnumeratedExtensible)
	}

	return &ret, nil
}

func decodeTestEnumeratedExtensible(testEnumeratedExtensibleC *C.TEST_EnumeratedExtensible_t) (*test_sm_ies.TestEnumeratedExtensible, error) {

	testEnumeratedExtensible := test_sm_ies.TestEnumeratedExtensible(int32(*testEnumeratedExtensibleC))

	return &testEnumeratedExtensible, nil
}
