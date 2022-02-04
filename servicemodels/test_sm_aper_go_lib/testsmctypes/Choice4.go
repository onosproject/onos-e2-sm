// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Choice4.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeChoice4(choice4 *test_sm_ies.Choice4) ([]byte, error) {
	choice4CP, err := newChoice4(choice4)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice4() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Choice4, unsafe.Pointer(choice4CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeChoice4() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeChoice4(choice4 *test_sm_ies.Choice4) ([]byte, error) {
	choice4CP, err := newChoice4(choice4)
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice4() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Choice4, unsafe.Pointer(choice4CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeChoice4() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeChoice4(bytes []byte) (*test_sm_ies.Choice4, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Choice4)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeChoice4((*C.Choice4_t)(unsafePtr))
}

func perDecodeChoice4(bytes []byte) (*test_sm_ies.Choice4, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Choice4)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeChoice4((*C.Choice4_t)(unsafePtr))
}

func newChoice4(choice4 *test_sm_ies.Choice4) (*C.Choice4_t, error) {

	var pr C.Choice4_PR
	choiceC := [8]byte{}

	switch choice := choice4.Choice4.(type) {
	case *test_sm_ies.Choice4_Choice4A:
		pr = C.Choice4_PR_choice4A

		im := C.long(choice.Choice4A)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newChoice4() %T not yet implemented", choice)
	}

	choice4C := C.Choice4_t{
		present: pr,
		choice:  choiceC,
	}

	return &choice4C, nil
}

func decodeChoice4(choice4C *C.Choice4_t) (*test_sm_ies.Choice4, error) {

	choice4 := new(test_sm_ies.Choice4)

	switch choice4C.present {
	case C.Choice4_PR_choice4A:
		choice4structC := int32(binary.LittleEndian.Uint64(choice4C.choice[0:8]))
		choice4.Choice4 = &test_sm_ies.Choice4_Choice4A{
			Choice4A: choice4structC,
		}
	default:
		return nil, fmt.Errorf("decodeChoice4() %v not yet implemented", choice4C.present)
	}

	return choice4, nil
}
