// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Choice3.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeChoice3(choice3 *test_sm_ies.Choice3) ([]byte, error) {
	choice3CP, err := newChoice3(choice3)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice3() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Choice3, unsafe.Pointer(choice3CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice3() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeChoice3(choice3 *test_sm_ies.Choice3) ([]byte, error) {
	choice3CP, err := newChoice3(choice3)
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice3() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Choice3, unsafe.Pointer(choice3CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice3() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeChoice3(bytes []byte) (*test_sm_ies.Choice3, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Choice3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeChoice3((*C.Choice3_t)(unsafePtr))
}

func perDecodeChoice3(bytes []byte) (*test_sm_ies.Choice3, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Choice3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeChoice3((*C.Choice3_t)(unsafePtr))
}

func newChoice3(choice3 *test_sm_ies.Choice3) (*C.Choice3_t, error) {

	var pr C.Choice3_PR
	choiceC := [8]byte{}

	switch choice := choice3.Choice3.(type) {
	case *test_sm_ies.Choice3_Choice3A:
		pr = C.Choice3_PR_choice3A

		im := C.long(choice.Choice3A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.Choice3_Choice3B:
		pr = C.Choice3_PR_choice3B

		im := C.long(choice.Choice3B)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.Choice3_Choice3C:
		pr = C.Choice3_PR_choice3C

		im := C.long(choice.Choice3C)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newChoice3() %T not yet implemented", choice)
	}

	choice3C := C.Choice3_t{
		present: pr,
		choice:  choiceC,
	}

	return &choice3C, nil
}

func decodeChoice3(choice3C *C.Choice3_t) (*test_sm_ies.Choice3, error) {

	choice3 := new(test_sm_ies.Choice3)

	switch choice3C.present {
	case C.Choice3_PR_choice3A:
		choice3structC := int32(binary.LittleEndian.Uint64(choice3C.choice[0:8]))
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3A{
			Choice3A: choice3structC,
		}
	case C.Choice3_PR_choice3B:
		choice3structC := int32(binary.LittleEndian.Uint64(choice3C.choice[0:8]))
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3B{
			Choice3B: choice3structC,
		}
	case C.Choice3_PR_choice3C:
		choice3structC := int32(binary.LittleEndian.Uint64(choice3C.choice[0:8]))
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3C{
			Choice3C: choice3structC,
		}
	default:
		return nil, fmt.Errorf("decodeChoice3() %v not yet implemented", choice3C.present)
	}

	return choice3, nil
}
