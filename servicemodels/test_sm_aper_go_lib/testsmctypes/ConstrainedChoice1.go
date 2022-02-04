// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ConstrainedChoice1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeConstrainedChoice1(constrainedChoice1 *test_sm_ies.ConstrainedChoice1) ([]byte, error) {
	constrainedChoice1CP, err := newConstrainedChoice1(constrainedChoice1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ConstrainedChoice1, unsafe.Pointer(constrainedChoice1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeConstrainedChoice1(constrainedChoice1 *test_sm_ies.ConstrainedChoice1) ([]byte, error) {
	constrainedChoice1CP, err := newConstrainedChoice1(constrainedChoice1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ConstrainedChoice1, unsafe.Pointer(constrainedChoice1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeConstrainedChoice1(bytes []byte) (*test_sm_ies.ConstrainedChoice1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ConstrainedChoice1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeConstrainedChoice1((*C.ConstrainedChoice1_t)(unsafePtr))
}

func perDecodeConstrainedChoice1(bytes []byte) (*test_sm_ies.ConstrainedChoice1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ConstrainedChoice1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeConstrainedChoice1((*C.ConstrainedChoice1_t)(unsafePtr))
}

func newConstrainedChoice1(constrainedChoice1 *test_sm_ies.ConstrainedChoice1) (*C.ConstrainedChoice1_t, error) {

	var pr C.ConstrainedChoice1_PR
	choiceC := [8]byte{}

	switch choice := constrainedChoice1.ConstrainedChoice1.(type) {
	case *test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A:
		pr = C.ConstrainedChoice1_PR_constrainedChoice1A

		im := C.long(choice.ConstrainedChoice1A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newConstrainedChoice1() %T not yet implemented", choice)
	}

	constrainedChoice1C := C.ConstrainedChoice1_t{
		present: pr,
		choice:  choiceC,
	}

	return &constrainedChoice1C, nil
}

func decodeConstrainedChoice1(constrainedChoice1C *C.ConstrainedChoice1_t) (*test_sm_ies.ConstrainedChoice1, error) {

	constrainedChoice1 := new(test_sm_ies.ConstrainedChoice1)

	switch constrainedChoice1C.present {
	case C.ConstrainedChoice1_PR_constrainedChoice1A:
		constrainedChoice1structC := int32(binary.LittleEndian.Uint64(constrainedChoice1C.choice[0:8]))
		constrainedChoice1.ConstrainedChoice1 = &test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A{
			ConstrainedChoice1A: constrainedChoice1structC,
		}
	default:
		return nil, fmt.Errorf("decodeConstrainedChoice1() %v not yet implemented", constrainedChoice1C.present)
	}

	return constrainedChoice1, nil
}
