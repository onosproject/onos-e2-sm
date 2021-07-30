// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-NestedChoice.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestNestedChoice(testNestedChoice *test_sm_ies.TestNestedChoice) ([]byte, error) {
	testNestedChoiceCP, err := newTestNestedChoice(testNestedChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestNestedChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_NestedChoice, unsafe.Pointer(testNestedChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestNestedChoice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestNestedChoice(testNestedChoice *test_sm_ies.TestNestedChoice) ([]byte, error) {
	testNestedChoiceCP, err := newTestNestedChoice(testNestedChoice)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestNestedChoice() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_NestedChoice, unsafe.Pointer(testNestedChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestNestedChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestNestedChoice(bytes []byte) (*test_sm_ies.TestNestedChoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_NestedChoice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestNestedChoice((*C.TEST_NestedChoice_t)(unsafePtr))
}

func perDecodeTestNestedChoice(bytes []byte) (*test_sm_ies.TestNestedChoice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_NestedChoice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestNestedChoice((*C.TEST_NestedChoice_t)(unsafePtr))
}

func newTestNestedChoice(testNestedChoice *test_sm_ies.TestNestedChoice) (*C.TEST_NestedChoice_t, error) {

	var pr C.TEST_NestedChoice_PR
	choiceC := [8]byte{}

	switch choice := testNestedChoice.TestNestedChoice.(type) {
	case *test_sm_ies.TestNestedChoice_Option1:
		pr = C.TEST_NestedChoice_PR_option1

		im, err := newChoice3(testNestedChoice.GetOption1())
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *test_sm_ies.TestNestedChoice_Option2:
		pr = C.TEST_NestedChoice_PR_option2

		im, err := newConstrainedChoice3(testNestedChoice.GetOption2())
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *test_sm_ies.TestNestedChoice_Option3:
		pr = C.TEST_NestedChoice_PR_option3

		im, err := newConstrainedChoice4(testNestedChoice.GetOption3())
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newTestNestedChoice() %T not yet implemented", choice)
	}

	nestedChoiceC := C.TEST_NestedChoice_t{
		present: pr,
		choice:  choiceC,
	}

	return &nestedChoiceC, nil
}

func decodeTestNestedChoice(testNestedChoiceC *C.TEST_NestedChoice_t) (*test_sm_ies.TestNestedChoice, error) {

	testNestedChoice := new(test_sm_ies.TestNestedChoice)

	switch testNestedChoiceC.present {
	case C.TEST_NestedChoice_PR_option1:
		im3C, err := decodeChoice3Bytes(testNestedChoiceC.choice)
		if err != nil {
			return nil, err
		}
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option1{
			Option1: im3C,
		}
	case C.TEST_NestedChoice_PR_option2:
		im3C, err := decodeConstrainedChoice3Bytes(testNestedChoiceC.choice)
		if err != nil {
			return nil, err
		}
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option2{
			Option2: im3C,
		}
	case C.TEST_NestedChoice_PR_option3:
		im4C, err := decodeConstrainedChoice4Bytes(testNestedChoiceC.choice)
		if err != nil {
			return nil, err
		}
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option3{
			Option3: im4C,
		}
	default:
		return nil, fmt.Errorf("decodeTestNestedChoice() %v not yet implemented", testNestedChoiceC.present)
	}

	return testNestedChoice, nil
}
