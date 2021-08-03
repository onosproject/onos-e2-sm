// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-TopLevelPDU.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestTopLevelPDU(testTopLevelPDU *test_sm_ies.TestTopLevelPdu) ([]byte, error) {
	testTopLevelPDUCP, err := newTestTopLevelPDU(testTopLevelPDU)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestTopLevelPDU() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_TopLevelPDU, unsafe.Pointer(testTopLevelPDUCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestTopLevelPDU() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestTopLevelPDU(testTopLevelPDU *test_sm_ies.TestTopLevelPdu) ([]byte, error) {
	testTopLevelPDUCP, err := newTestTopLevelPDU(testTopLevelPDU)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestTopLevelPDU() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_TopLevelPDU, unsafe.Pointer(testTopLevelPDUCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestTopLevelPDU() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestTopLevelPDU(bytes []byte) (*test_sm_ies.TestTopLevelPdu, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_TopLevelPDU)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestTopLevelPDU((*C.TEST_TopLevelPDU_t)(unsafePtr))
}

func perDecodeTestTopLevelPDU(bytes []byte) (*test_sm_ies.TestTopLevelPdu, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_TopLevelPDU)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestTopLevelPDU((*C.TEST_TopLevelPDU_t)(unsafePtr))
}

func newTestTopLevelPDU(testTopLevelPDU *test_sm_ies.TestTopLevelPdu) (*C.TEST_TopLevelPDU_t, error) {

	opt1C, err := newTestUnconstrainedInt(testTopLevelPDU.GetOpt1())
	if err != nil {
		return nil, err
	}
	opt3C, err := newTestNestedChoice(testTopLevelPDU.GetOpt3())
	if err != nil {
		return nil, err
	}
	opt4C, err := newTestBitString(testTopLevelPDU.GetOpt4())
	if err != nil {
		return nil, err
	}
	opt6C, err := newTestListExtensible3(testTopLevelPDU.GetOpt6())
	if err != nil {
		return nil, err
	}
	opt7C, err := newTestEnumeratedExtensible(&testTopLevelPDU.Opt7)
	if err != nil {
		return nil, err
	}

	nestedChoiceC := C.TEST_TopLevelPDU_t{
		opt1: *opt1C,
		opt3: *opt3C,
		opt4: *opt4C,
		opt6: *opt6C,
		opt7: *opt7C,
	}

	if testTopLevelPDU.GetOpt2() != nil {
		opt2C, err := newTestConstrainedReal(testTopLevelPDU.GetOpt2())
		if err != nil {
			return nil, err
		}
		nestedChoiceC.opt2 = opt2C
	}
	if testTopLevelPDU.GetOpt5() != nil {
		opt5C, err := newTestOctetString(testTopLevelPDU.GetOpt5())
		if err != nil {
			return nil, err
		}
		nestedChoiceC.opt5 = opt5C
	}

	return &nestedChoiceC, nil
}

func decodeTestTopLevelPDU(testTopLevelPDUC *C.TEST_TopLevelPDU_t) (*test_sm_ies.TestTopLevelPdu, error) {

	var err error
	testNestedChoice := test_sm_ies.TestTopLevelPdu{}

	testNestedChoice.Opt1, err = decodeTestUnconstrainedInt(&testTopLevelPDUC.opt1)
	if err != nil {
		return nil, err
	}
	testNestedChoice.Opt3, err = decodeTestNestedChoice(&testTopLevelPDUC.opt3)
	if err != nil {
		return nil, err
	}
	testNestedChoice.Opt4, err = decodeTestBitString(&testTopLevelPDUC.opt4)
	if err != nil {
		return nil, err
	}
	testNestedChoice.Opt6, err = decodeTestListExtensible3(&testTopLevelPDUC.opt6)
	if err != nil {
		return nil, err
	}
	opt7, err := decodeTestEnumeratedExtensible(&testTopLevelPDUC.opt7)
	if err != nil {
		return nil, err
	}
	testNestedChoice.Opt7 = *opt7

	if testTopLevelPDUC.opt2 != nil {
		testNestedChoice.Opt2, err = decodeTestConstrainedReal(testTopLevelPDUC.opt2)
		if err != nil {
			return nil, err
		}
	}

	if testTopLevelPDUC.opt5 != nil {
		testNestedChoice.Opt5, err = decodeTestOctetString(testTopLevelPDUC.opt5)
		if err != nil {
			return nil, err
		}
	}

	return &testNestedChoice, nil
}
