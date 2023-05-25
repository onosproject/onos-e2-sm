// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "AChoice.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeAChoice(AChoice *test_sm_ies.Achoice) ([]byte, error) {
	AChoiceCP, err := newAChoice(AChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_AChoice, unsafe.Pointer(AChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeAChoice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeAChoice(AChoice *test_sm_ies.Achoice) ([]byte, error) {
	AChoiceCP, err := newAChoice(AChoice)
	if err != nil {
		return nil, fmt.Errorf("perEncodeAChoice() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_AChoice, unsafe.Pointer(AChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeAChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeAChoice(bytes []byte) (*test_sm_ies.Achoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_AChoice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeAChoice((*C.AChoice_t)(unsafePtr))
}

func perDecodeAChoice(bytes []byte) (*test_sm_ies.Achoice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_AChoice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeAChoice((*C.AChoice_t)(unsafePtr))
}

func newAChoice(achoice *test_sm_ies.Achoice) (*C.AChoice_t, error) {

	var pr C.AChoice_PR
	choiceC := [8]byte{}

	switch t := achoice.Achoice.(type) {
	case *test_sm_ies.Achoice_Ch1:
		pr = C.AChoice_PR_ch1

		im, err := newAList(achoice.GetCh1())
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	case *test_sm_ies.Achoice_Ch2:
		pr = C.AChoice_PR_ch2

		im, err := newAStruct(achoice.GetCh2())
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(im))
	default:
		return nil, fmt.Errorf("newAChoice() %T not yet implemented", t)
	}

	AChoiceC := C.AChoice_t{
		present: pr,
		choice:  choiceC,
	}

	return &AChoiceC, nil
}

func decodeAChoice(achoiceC *C.AChoice_t) (*test_sm_ies.Achoice, error) {

	achoice := new(test_sm_ies.Achoice)

	switch achoiceC.present {
	case C.AChoice_PR_ch1:
		ch1C, err := decodeAList(achoiceC.choice)
		if err != nil {
			return nil, err
		}
		achoice.Achoice = &test_sm_ies.Achoice_Ch1{
			Ch1: ch1C,
		}
	case C.AChoice_PR_ch2:
		ch2C, err := decodeAStruct(achoiceC.choice)
		if err != nil {
			return nil, err
		}
		achoice.Achoice = &test_sm_ies.Achoice_Ch2{
			Ch2: ch2C,
		}
	default:
		return nil, fmt.Errorf("decodeAChoice() %v not yet implemented", achoiceC.present)
	}

	return achoice, nil
}
