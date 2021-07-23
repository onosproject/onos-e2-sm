// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-FullyOptionalSequence.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestFullyOptionalSequence(testFullyOptionalSequence *test_sm_ies.TestFullyOptionalSequence) ([]byte, error) {
	testFullyOptionalSequenceCP, err := newTestFullyOptionalSequence(testFullyOptionalSequence)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestFullyOptionalSequence() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_FullyOptionalSequence, unsafe.Pointer(testFullyOptionalSequenceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestFullyOptionalSequence() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestFullyOptionalSequence(testFullyOptionalSequence *test_sm_ies.TestFullyOptionalSequence) ([]byte, error) {
	testFullyOptionalSequenceCP, err := newTestFullyOptionalSequence(testFullyOptionalSequence)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestFullyOptionalSequence() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_FullyOptionalSequence, unsafe.Pointer(testFullyOptionalSequenceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestFullyOptionalSequence() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestFullyOptionalSequence(bytes []byte) (*test_sm_ies.TestFullyOptionalSequence, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_FullyOptionalSequence)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestFullyOptionalSequence((*C.TEST_FullyOptionalSequence_t)(unsafePtr))
}

func perDecodeTestFullyOptionalSequence(bytes []byte) (*test_sm_ies.TestFullyOptionalSequence, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_FullyOptionalSequence)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestFullyOptionalSequence((*C.TEST_FullyOptionalSequence_t)(unsafePtr))
}

func newTestFullyOptionalSequence(testFullyOptionalSequence *test_sm_ies.TestFullyOptionalSequence) (*C.TEST_FullyOptionalSequence_t, error) {

	var err error
	testFullyOptionalSequenceC := C.TEST_FullyOptionalSequence_t{}

	item1C := C.long(testFullyOptionalSequence.Item1)
	item2C, err := newOctetString(testFullyOptionalSequence.Item2)
	if err != nil {
		return nil, err
	}
	item3C := newBoolean(*testFullyOptionalSequence.Item3)

	item4C := C.long(testFullyOptionalSequence.Item4)
	item5C := C.long(testFullyOptionalSequence.Item5)
	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	testFullyOptionalSequenceC.item1 = item1C
	testFullyOptionalSequenceC.item2 = item2C
	testFullyOptionalSequenceC.item3 = item3C
	testFullyOptionalSequenceC.item4 = item4C
	testFullyOptionalSequenceC.item5 = item5C

	return &testFullyOptionalSequenceC, nil
}

func decodeTestFullyOptionalSequence(testFullyOptionalSequenceC *C.TEST_FullyOptionalSequence_t) (*test_sm_ies.TestFullyOptionalSequence, error) {

	var err error
	testFullyOptionalSequence := test_sm_ies.TestFullyOptionalSequence{}

	testFullyOptionalSequence.Item1 = int32(testFullyOptionalSequenceC.item1)
	testFullyOptionalSequence.Item2, err = decodeOctetString(testFullyOptionalSequenceC.item2)
	if err != nil {
		return nil, err
	}
	testFullyOptionalSequence.Item3 = decodeBoolean(testFullyOptionalSequenceC.item3)
	if err != nil {
		return nil, fmt.Errorf("decodeBoolean() %s", err.Error())
	}

	testFullyOptionalSequence.Item4 = int32(testFullyOptionalSequenceC.item4)
	testFullyOptionalSequence.Item5 = int32(testFullyOptionalSequenceC.item5)

	return &testFullyOptionalSequence, nil
}

func decodeTestFullyOptionalSequenceBytes(array [8]byte) (*test_sm_ies.TestFullyOptionalSequence, error) {
	testFullyOptionalSequenceC := (*C.TEST_FullyOptionalSequence_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeTestFullyOptionalSequence(testFullyOptionalSequenceC)
}
