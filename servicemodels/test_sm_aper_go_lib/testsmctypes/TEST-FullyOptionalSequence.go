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

	testFullyOptionalSequenceC := C.TEST_FullyOptionalSequence_t{}

	if testFullyOptionalSequence.Item1 != nil {
		item1C := C.long(*testFullyOptionalSequence.Item1)
		testFullyOptionalSequenceC.item1 = &item1C
	}

	if testFullyOptionalSequence.Item2 != nil {
		item2C, err := newOctetString(testFullyOptionalSequence.Item2)
		if err != nil {
			return nil, err
		}
		testFullyOptionalSequenceC.item2 = item2C
	}

	if testFullyOptionalSequence.Item3 != nil {
		item3C := newBoolean(*testFullyOptionalSequence.Item3)
		testFullyOptionalSequenceC.item3 = item3C
	}

	if testFullyOptionalSequence.Item4 != nil {
		var item4C C.TEST_FullyOptionalSequence__item4_t
		switch *testFullyOptionalSequence.Item4 {
		case test_sm_ies.TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE:
			item4C = C.TEST_FullyOptionalSequence__item4_one
		case test_sm_ies.TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_TWO:
			item4C = C.TEST_FullyOptionalSequence__item4_two
		default:
			return nil, fmt.Errorf("unexpected FullyOptionalSequence Item4 %v", testFullyOptionalSequence.Item4)
		}

		//item4C := C.long(*testFullyOptionalSequence.Item4)
		testFullyOptionalSequenceC.item4 = &item4C
	}

	if testFullyOptionalSequence.Item5 != nil {
		item5C := C.int(*testFullyOptionalSequence.Item5)
		testFullyOptionalSequenceC.item5 = &item5C
	}

	return &testFullyOptionalSequenceC, nil
}

func decodeTestFullyOptionalSequence(testFullyOptionalSequenceC *C.TEST_FullyOptionalSequence_t) (*test_sm_ies.TestFullyOptionalSequence, error) {

	var err error
	testFullyOptionalSequence := test_sm_ies.TestFullyOptionalSequence{}

	if testFullyOptionalSequenceC.item1 != nil {
		ie1 := int32(*testFullyOptionalSequenceC.item1)
		testFullyOptionalSequence.Item1 = &ie1
	}

	if testFullyOptionalSequenceC.item2 != nil {
		testFullyOptionalSequence.Item2, err = decodeOctetString(testFullyOptionalSequenceC.item2)
		if err != nil {
			return nil, err
		}
	}

	if testFullyOptionalSequenceC.item3 != nil {
		bl := decodeBoolean(testFullyOptionalSequenceC.item3)
		testFullyOptionalSequence.Item3 = &bl
	}

	if testFullyOptionalSequenceC.item4 != nil {
		ie4 := test_sm_ies.TestFullyOptionalSequenceItem4(int32(*testFullyOptionalSequenceC.item4))
		testFullyOptionalSequence.Item4 = &ie4
	}

	if testFullyOptionalSequenceC.item5 != nil {
		ie5 := int32(*testFullyOptionalSequenceC.item5)
		testFullyOptionalSequence.Item5 = &ie5
	}

	return &testFullyOptionalSequence, nil
}
