// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Choice1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeChoice1(choice1 *test_sm_ies.Choice1) ([]byte, error) {
	choice1CP, err := newChoice1(choice1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Choice1, unsafe.Pointer(choice1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeChoice1(choice1 *test_sm_ies.Choice1) ([]byte, error) {
	choice1CP, err := newChoice1(choice1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Choice1, unsafe.Pointer(choice1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeChoice1(bytes []byte) (*test_sm_ies.Choice1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Choice1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeChoice1((*C.Choice1_t)(unsafePtr))
}

func perDecodeChoice1(bytes []byte) (*test_sm_ies.Choice1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Choice1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeChoice1((*C.Choice1_t)(unsafePtr))
}

func newChoice1(choice1 *test_sm_ies.Choice1) (*C.Choice1_t, error) {

	var pr C.Choice1_PR
	choiceC := [8]byte{}

	switch choice := choice1.Choice1.(type) {
	case *test_sm_ies.Choice1_Choice1A:
		pr = C.Choice1_PR_choice1A

		im := C.long(choice.Choice1A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newChoice1() %T not yet implemented", choice)
	}

	choice1C := C.Choice1_t{
		present: pr,
		choice:  choiceC,
	}

	return &choice1C, nil
}

func decodeChoice1(choice1C *C.Choice1_t) (*test_sm_ies.Choice1, error) {

	choice1 := new(test_sm_ies.Choice1)

	switch choice1C.present {
	case C.Choice1_PR_choice1A:
		choice1structC := int32(binary.LittleEndian.Uint64(choice1C.choice[0:8]))
		choice1.Choice1 = &test_sm_ies.Choice1_Choice1A{
			Choice1A: choice1structC,
		}
	default:
		return nil, fmt.Errorf("decodeChoice1() %v not yet implemented", choice1C.present)
	}

	return choice1, nil
}
