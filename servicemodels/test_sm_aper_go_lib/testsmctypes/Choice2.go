// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Choice2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeChoice2(choice2 *test_sm_ies.Choice2) ([]byte, error) {
	choice2CP, err := newChoice2(choice2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Choice2, unsafe.Pointer(choice2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeChoice2(choice2 *test_sm_ies.Choice2) ([]byte, error) {
	choice2CP, err := newChoice2(choice2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Choice2, unsafe.Pointer(choice2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeChoice2(bytes []byte) (*test_sm_ies.Choice2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Choice2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeChoice2((*C.Choice2_t)(unsafePtr))
}

func perDecodeChoice2(bytes []byte) (*test_sm_ies.Choice2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Choice2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeChoice2((*C.Choice2_t)(unsafePtr))
}

func newChoice2(choice2 *test_sm_ies.Choice2) (*C.Choice2_t, error) {

	var pr C.Choice2_PR  //ToDo - verify correctness of the name
	choiceC := [8]byte{} //ToDo - Check if number of bytes is sufficient

	switch choice := choice2.Choice2.(type) {
	case *test_sm_ies.Choice2_Choice2A:
		pr = C.Choice2_PR_choice2A //ToDo - Check if it's correct PR's name

		im, err := C.long(choice.Choice2A)
		if err != nil {
			return nil, fmt.Errorf("C.long() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *test_sm_ies.Choice2_Choice2B:
		pr = C.Choice2_PR_choice2B //ToDo - Check if it's correct PR's name

		im, err := C.long(choice.Choice2B)
		if err != nil {
			return nil, fmt.Errorf("C.long() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newChoice2() %T not yet implemented", choice)
	}

	choice2C := C.Choice2_t{
		present: pr,
		choice:  choiceC,
	}

	return &choice2C, nil
}

func decodeChoice2(choice2C *C.Choice2_t) (*test_sm_ies.Choice2, error) {

	choice2 := new(test_sm_ies.Choice2)

	switch choice2C.present {
	case C.Choice2_PR_choice2A:
		choice2structC, err := int32Bytes(choice2C.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("int32Bytes() %s", err.Error())
		}
		choice2.Choice2 = &test_sm_ies.Choice2_Choice2A{
			Choice2A: choice2structC,
		}
	case C.Choice2_PR_choice2B:
		choice2structC, err := int32Bytes(choice2C.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("int32Bytes() %s", err.Error())
		}
		choice2.Choice2 = &test_sm_ies.Choice2_Choice2B{
			Choice2B: choice2structC,
		}
	default:
		return nil, fmt.Errorf("decodeChoice2() %v not yet implemented", choice2C.present)
	}

	return choice2, nil
}

func decodeChoice2Bytes(array [8]byte) (*test_sm_ies.Choice2, error) {
	choice2C := (*C.Choice2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeChoice2(choice2C)
}