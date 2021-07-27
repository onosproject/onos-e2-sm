// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-ConstrainedChoices.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestConstrainedChoices(testConstrainedChoices *test_sm_ies.TestConstrainedChoices) ([]byte, error) {
	testConstrainedChoicesCP, err := newTestConstrainedChoices(testConstrainedChoices)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedChoices() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_ConstrainedChoices, unsafe.Pointer(testConstrainedChoicesCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestConstrainedChoices() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestConstrainedChoices(testConstrainedChoices *test_sm_ies.TestConstrainedChoices) ([]byte, error) {
	testConstrainedChoicesCP, err := newTestConstrainedChoices(testConstrainedChoices)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedChoices() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_ConstrainedChoices, unsafe.Pointer(testConstrainedChoicesCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestConstrainedChoices() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestConstrainedChoices(bytes []byte) (*test_sm_ies.TestConstrainedChoices, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_ConstrainedChoices)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestConstrainedChoices((*C.TEST_ConstrainedChoices_t)(unsafePtr))
}

func perDecodeTestConstrainedChoices(bytes []byte) (*test_sm_ies.TestConstrainedChoices, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_ConstrainedChoices)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestConstrainedChoices((*C.TEST_ConstrainedChoices_t)(unsafePtr))
}

func newTestConstrainedChoices(testConstrainedChoices *test_sm_ies.TestConstrainedChoices) (*C.TEST_ConstrainedChoices_t, error) {

	var err error
	testConstrainedChoicesC := C.TEST_ConstrainedChoices_t{}

	otherCattrC, err := newPrintableString(testConstrainedChoices.OtherCattr)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	constrainedChoice1C, err := newConstrainedChoice1(testConstrainedChoices.ConstrainedChoice1)
	if err != nil {
		return nil, fmt.Errorf("newConstrainedChoice1() %s", err.Error())
	}

	constrainedChoice2C, err := newConstrainedChoice2(testConstrainedChoices.ConstrainedChoice2)
	if err != nil {
		return nil, fmt.Errorf("newConstrainedChoice2() %s", err.Error())
	}

	constrainedChoice3C, err := newConstrainedChoice3(testConstrainedChoices.ConstrainedChoice3)
	if err != nil {
		return nil, fmt.Errorf("newConstrainedChoice3() %s", err.Error())
	}

	constrainedChoice4C, err := newConstrainedChoice4(testConstrainedChoices.ConstrainedChoice4)
	if err != nil {
		return nil, fmt.Errorf("newConstrainedChoice4() %s", err.Error())
	}

	testConstrainedChoicesC.otherCAttr = *otherCattrC
	testConstrainedChoicesC.constrainedChoice1 = *constrainedChoice1C
	testConstrainedChoicesC.constrainedChoice2 = *constrainedChoice2C
	testConstrainedChoicesC.constrainedChoice3 = *constrainedChoice3C
	testConstrainedChoicesC.constrainedChoice4 = *constrainedChoice4C

	return &testConstrainedChoicesC, nil
}

func decodeTestConstrainedChoices(testConstrainedChoicesC *C.TEST_ConstrainedChoices_t) (*test_sm_ies.TestConstrainedChoices, error) {

	var err error
	testConstrainedChoices := test_sm_ies.TestConstrainedChoices{}

	testConstrainedChoices.OtherCattr, err = decodePrintableString(&testConstrainedChoicesC.otherCAttr)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	testConstrainedChoices.ConstrainedChoice1, err = decodeConstrainedChoice1(&testConstrainedChoicesC.constrainedChoice1)
	if err != nil {
		return nil, fmt.Errorf("decodeConstrainedChoice1() %s", err.Error())
	}

	testConstrainedChoices.ConstrainedChoice2, err = decodeConstrainedChoice2(&testConstrainedChoicesC.constrainedChoice2)
	if err != nil {
		return nil, fmt.Errorf("decodeConstrainedChoice2() %s", err.Error())
	}

	testConstrainedChoices.ConstrainedChoice3, err = decodeConstrainedChoice3(&testConstrainedChoicesC.constrainedChoice3)
	if err != nil {
		return nil, fmt.Errorf("decodeConstrainedChoice3() %s", err.Error())
	}

	testConstrainedChoices.ConstrainedChoice4, err = decodeConstrainedChoice4(&testConstrainedChoicesC.constrainedChoice4)
	if err != nil {
		return nil, fmt.Errorf("decodeConstrainedChoice4() %s", err.Error())
	}

	return &testConstrainedChoices, nil
}
