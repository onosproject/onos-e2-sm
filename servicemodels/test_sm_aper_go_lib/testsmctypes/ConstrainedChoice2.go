// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ConstrainedChoice2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeConstrainedChoice2(constrainedChoice2 *test_sm_ies.ConstrainedChoice2) ([]byte, error) {
	constrainedChoice2CP, err := newConstrainedChoice2(constrainedChoice2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ConstrainedChoice2, unsafe.Pointer(constrainedChoice2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeConstrainedChoice2(constrainedChoice2 *test_sm_ies.ConstrainedChoice2) ([]byte, error) {
	constrainedChoice2CP, err := newConstrainedChoice2(constrainedChoice2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ConstrainedChoice2, unsafe.Pointer(constrainedChoice2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeConstrainedChoice2(bytes []byte) (*test_sm_ies.ConstrainedChoice2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ConstrainedChoice2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeConstrainedChoice2((*C.ConstrainedChoice2_t)(unsafePtr))
}

func perDecodeConstrainedChoice2(bytes []byte) (*test_sm_ies.ConstrainedChoice2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ConstrainedChoice2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeConstrainedChoice2((*C.ConstrainedChoice2_t)(unsafePtr))
}

func newConstrainedChoice2(constrainedChoice2 *test_sm_ies.ConstrainedChoice2) (*C.ConstrainedChoice2_t, error) {

	var pr C.ConstrainedChoice2_PR
	choiceC := [8]byte{}

	switch choice := constrainedChoice2.ConstrainedChoice2.(type) {
	case *test_sm_ies.ConstrainedChoice2_ConstrainedChoice2A:
		pr = C.ConstrainedChoice2_PR_constrainedChoice2A

		im := C.long(choice.ConstrainedChoice2A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.ConstrainedChoice2_ConstrainedChoice2B:
		pr = C.ConstrainedChoice2_PR_constrainedChoice2B

		im := C.ulong(choice.ConstrainedChoice2B)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newConstrainedChoice2() %T not yet implemented", choice)
	}

	constrainedChoice2C := C.ConstrainedChoice2_t{
		present: pr,
		choice:  choiceC,
	}

	return &constrainedChoice2C, nil
}

func decodeConstrainedChoice2(constrainedChoice2C *C.ConstrainedChoice2_t) (*test_sm_ies.ConstrainedChoice2, error) {

	constrainedChoice2 := new(test_sm_ies.ConstrainedChoice2)

	switch constrainedChoice2C.present {
	case C.ConstrainedChoice2_PR_constrainedChoice2A:
		constrainedChoice2structC := int32(binary.LittleEndian.Uint64(constrainedChoice2C.choice[0:8]))
		constrainedChoice2.ConstrainedChoice2 = &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2A{
			ConstrainedChoice2A: constrainedChoice2structC,
		}
	case C.ConstrainedChoice2_PR_constrainedChoice2B:
		constrainedChoice2structC:= int32(binary.LittleEndian.Uint64(constrainedChoice2C.choice[0:8]))
		constrainedChoice2.ConstrainedChoice2 = &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2B{
			ConstrainedChoice2B: constrainedChoice2structC,
		}
	default:
		return nil, fmt.Errorf("decodeConstrainedChoice2() %v not yet implemented", constrainedChoice2C.present)
	}

	return constrainedChoice2, nil
}
