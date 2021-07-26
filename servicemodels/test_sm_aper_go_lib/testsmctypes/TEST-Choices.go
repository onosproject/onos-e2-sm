// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-Choices.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestChoices(testChoices *test_sm_ies.TestChoices) ([]byte, error) {
	testChoicesCP, err := newTestChoices(testChoices)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestChoices() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_Choices, unsafe.Pointer(testChoicesCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestChoices() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestChoices(testChoices *test_sm_ies.TestChoices) ([]byte, error) {
	testChoicesCP, err := newTestChoices(testChoices)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestChoices() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_Choices, unsafe.Pointer(testChoicesCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestChoices() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestChoices(bytes []byte) (*test_sm_ies.TestChoices, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_Choices)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestChoices((*C.TEST_Choices_t)(unsafePtr))
}

func perDecodeTestChoices(bytes []byte) (*test_sm_ies.TestChoices, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_Choices)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestChoices((*C.TEST_Choices_t)(unsafePtr))
}

func newTestChoices(testChoices *test_sm_ies.TestChoices) (*C.TEST_Choices_t, error) {

	var err error
	testChoicesC := C.TEST_Choices_t{}

	otherAttrC, err := newPrintableString(testChoices.OtherAttr)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	choice1C, err := newChoice1(testChoices.Choice1)
	if err != nil {
		return nil, fmt.Errorf("newChoice1() %s", err.Error())
	}

	choice2C, err := newChoice2(testChoices.Choice2)
	if err != nil {
		return nil, fmt.Errorf("newChoice2() %s", err.Error())
	}

	choice3C, err := newChoice3(testChoices.Choice3)
	if err != nil {
		return nil, fmt.Errorf("newChoice3() %s", err.Error())
	}

	choice4C, err := newChoice4(testChoices.Choice4)
	if err != nil {
		return nil, fmt.Errorf("newChoice4() %s", err.Error())
	}

	testChoicesC.otherAttr = *otherAttrC
	testChoicesC.choice1 = *choice1C
	testChoicesC.choice2 = *choice2C
	testChoicesC.choice3 = *choice3C
	testChoicesC.choice4 = *choice4C

	return &testChoicesC, nil
}

func decodeTestChoices(testChoicesC *C.TEST_Choices_t) (*test_sm_ies.TestChoices, error) {

	var err error
	testChoices := test_sm_ies.TestChoices{}

	testChoices.OtherAttr, err = decodePrintableString(&testChoicesC.otherAttr)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testChoices.Choice1, err = decodeChoice1(&testChoicesC.choice1)
	if err != nil {
		return nil, fmt.Errorf("decodeChoice1() %s", err.Error())
	}

	testChoices.Choice2, err = decodeChoice2(&testChoicesC.choice2)
	if err != nil {
		return nil, fmt.Errorf("decodeChoice2() %s", err.Error())
	}

	testChoices.Choice3, err = decodeChoice3(&testChoicesC.choice3)
	if err != nil {
		return nil, fmt.Errorf("decodeChoice3() %s", err.Error())
	}

	testChoices.Choice4, err = decodeChoice4(&testChoicesC.choice4)
	if err != nil {
		return nil, fmt.Errorf("decodeChoice4() %s", err.Error())
	}

	return &testChoices, nil
}
