// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ConstrainedChoice4.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeConstrainedChoice4(constrainedChoice4 *test_sm_ies.ConstrainedChoice4) ([]byte, error) {
	constrainedChoice4CP, err := newConstrainedChoice4(constrainedChoice4)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice4() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ConstrainedChoice4, unsafe.Pointer(constrainedChoice4CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeConstrainedChoice4() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeConstrainedChoice4(constrainedChoice4 *test_sm_ies.ConstrainedChoice4) ([]byte, error) {
	constrainedChoice4CP, err := newConstrainedChoice4(constrainedChoice4)
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice4() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ConstrainedChoice4, unsafe.Pointer(constrainedChoice4CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeConstrainedChoice4() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeConstrainedChoice4(bytes []byte) (*test_sm_ies.ConstrainedChoice4, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ConstrainedChoice4)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeConstrainedChoice4((*C.ConstrainedChoice4_t)(unsafePtr))
}

func perDecodeConstrainedChoice4(bytes []byte) (*test_sm_ies.ConstrainedChoice4, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ConstrainedChoice4)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeConstrainedChoice4((*C.ConstrainedChoice4_t)(unsafePtr))
}

func newConstrainedChoice4(constrainedChoice4 *test_sm_ies.ConstrainedChoice4) (*C.ConstrainedChoice4_t, error) {

	var pr C.ConstrainedChoice4_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}           //ToDo - Check if number of bytes is sufficient

	switch choice := constrainedChoice4.ConstrainedChoice4.(type) {
	case *test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A:
		pr = C.ConstrainedChoice4_PR_constrainedChoice4A //ToDo - Check if it's correct PR's name

		im, err := C.long(choice.ConstrainedChoice4A)
		if err != nil {
			return nil, fmt.Errorf("C.long() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newConstrainedChoice4() %T not yet implemented", choice)
	}

	constrainedChoice4C := C.ConstrainedChoice4_t{
		present: pr,
		choice:  choiceC,
	}

	return &constrainedChoice4C, nil
}

func decodeConstrainedChoice4(constrainedChoice4C *C.ConstrainedChoice4_t) (*test_sm_ies.ConstrainedChoice4, error) {

	constrainedChoice4 := new(test_sm_ies.ConstrainedChoice4)

	switch constrainedChoice4C.present {
	case C.ConstrainedChoice4_PR_constrainedChoice4A:
		constrainedChoice4structC, err := int32Bytes(constrainedChoice4C.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("int32Bytes() %s", err.Error())
		}
		constrainedChoice4.ConstrainedChoice4 = &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
			ConstrainedChoice4A: constrainedChoice4structC,
		}
	default:
		return nil, fmt.Errorf("decodeConstrainedChoice4() %v not yet implemented", constrainedChoice4C.present)
	}

	return constrainedChoice4, nil
}

func decodeConstrainedChoice4Bytes(array [8]byte) (*test_sm_ies.ConstrainedChoice4, error) {
	constrainedChoice4C := (*C.ConstrainedChoice4_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeConstrainedChoice4(constrainedChoice4C)
}