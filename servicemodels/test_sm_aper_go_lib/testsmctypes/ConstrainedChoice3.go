// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ConstrainedChoice3.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeConstrainedChoice3(constrainedChoice3 *test_sm_ies.ConstrainedChoice3) ([]byte, error) {
	constrainedChoice3CP, err := newConstrainedChoice3(constrainedChoice3)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice3() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ConstrainedChoice3, unsafe.Pointer(constrainedChoice3CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice3() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeConstrainedChoice3(constrainedChoice3 *test_sm_ies.ConstrainedChoice3) ([]byte, error) {
	constrainedChoice3CP, err := newConstrainedChoice3(constrainedChoice3)
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice3() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ConstrainedChoice3, unsafe.Pointer(constrainedChoice3CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice3() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeConstrainedChoice3(bytes []byte) (*test_sm_ies.ConstrainedChoice3, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ConstrainedChoice3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeConstrainedChoice3((*C.ConstrainedChoice3_t)(unsafePtr))
}

func perDecodeConstrainedChoice3(bytes []byte) (*test_sm_ies.ConstrainedChoice3, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ConstrainedChoice3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeConstrainedChoice3((*C.ConstrainedChoice3_t)(unsafePtr))
}

func newConstrainedChoice3(constrainedChoice3 *test_sm_ies.ConstrainedChoice3) (*C.ConstrainedChoice3_t, error) {

	var pr C.ConstrainedChoice3_PR
	choiceC := [8]byte{}

	switch choice := constrainedChoice3.ConstrainedChoice3.(type) {
	case *test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A:
		pr = C.ConstrainedChoice3_PR_constrainedChoice3A

		im := C.long(choice.ConstrainedChoice3A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.ConstrainedChoice3_ConstrainedChoice3B:
		pr = C.ConstrainedChoice3_PR_constrainedChoice3B

		im := C.long(choice.ConstrainedChoice3B)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C:
		pr = C.ConstrainedChoice3_PR_constrainedChoice3C

		im := C.ulong(choice.ConstrainedChoice3C)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D:
		pr = C.ConstrainedChoice3_PR_constrainedChoice3D

		im := C.long(choice.ConstrainedChoice3D)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newConstrainedChoice3() %T not yet implemented", choice)
	}

	constrainedChoice3C := C.ConstrainedChoice3_t{
		present: pr,
		choice:  choiceC,
	}

	return &constrainedChoice3C, nil
}

func decodeConstrainedChoice3(constrainedChoice3C *C.ConstrainedChoice3_t) (*test_sm_ies.ConstrainedChoice3, error) {

	constrainedChoice3 := new(test_sm_ies.ConstrainedChoice3)

	switch constrainedChoice3C.present {
	case C.ConstrainedChoice3_PR_constrainedChoice3A:
		constrainedChoice3structC := int32(binary.LittleEndian.Uint64(constrainedChoice3C.choice[0:8]))
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A{
			ConstrainedChoice3A: constrainedChoice3structC,
		}
	case C.ConstrainedChoice3_PR_constrainedChoice3B:
		constrainedChoice3structC := int32(binary.LittleEndian.Uint64(constrainedChoice3C.choice[0:8]))
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3B{
			ConstrainedChoice3B: constrainedChoice3structC,
		}
	case C.ConstrainedChoice3_PR_constrainedChoice3C:
		constrainedChoice3structC := int32(binary.LittleEndian.Uint64(constrainedChoice3C.choice[0:8]))
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C{
			ConstrainedChoice3C: constrainedChoice3structC,
		}
	case C.ConstrainedChoice3_PR_constrainedChoice3D:
		constrainedChoice3structC := int32(binary.LittleEndian.Uint64(constrainedChoice3C.choice[0:8]))
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D{
			ConstrainedChoice3D: constrainedChoice3structC,
		}
	default:
		return nil, fmt.Errorf("decodeConstrainedChoice3() %v not yet implemented", constrainedChoice3C.present)
	}

	return constrainedChoice3, nil
}

func decodeConstrainedChoice3Bytes(array [8]byte) (*test_sm_ies.ConstrainedChoice3, error) {
	ch3C := (*C.ConstrainedChoice3_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeConstrainedChoice3(ch3C)
}
