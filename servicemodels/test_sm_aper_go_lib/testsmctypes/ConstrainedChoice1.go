// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ConstrainedChoice1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
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

	var pr C.ConstrainedChoice1_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}           //ToDo - Check if number of bytes is sufficient

	switch choice := constrainedChoice1.ConstrainedChoice1.(type) {
	case *test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A:
		pr = C.ConstrainedChoice1_PR_constrainedChoice1A //ToDo - Check if it's correct PR's name

		im, err := C.long(choice.ConstrainedChoice1A)
		if err != nil {
			return nil, fmt.Errorf("C.long() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
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
		constrainedChoice1structC, err := int32Bytes(constrainedChoice1C.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("int32Bytes() %s", err.Error())
		}
		constrainedChoice1.ConstrainedChoice1 = &test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A{
			ConstrainedChoice1A: constrainedChoice1structC,
		}
	default:
		return nil, fmt.Errorf("decodeConstrainedChoice1() %v not yet implemented", constrainedChoice1C.present)
	}

	return constrainedChoice1, nil
}

func decodeConstrainedChoice1Bytes(array [8]byte) (*test_sm_ies.ConstrainedChoice1, error) {
	constrainedChoice1C := (*C.ConstrainedChoice1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeConstrainedChoice1(constrainedChoice1C)
}