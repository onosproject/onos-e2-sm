// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENB-ID-Choice.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeEnbIDChoice(enbIDChoice *e2sm_kpm_v2.EnbIdChoice) ([]byte, error) {
	enbIDChoiceCP, err := newEnbIDChoice(enbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbIDChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID_Choice, unsafe.Pointer(enbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEnbIDChoice(enbIDChoice *e2sm_kpm_v2.EnbIdChoice) ([]byte, error) {
	enbIDChoiceCP, err := newEnbIDChoice(enbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbIDChoice() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ENB_ID_Choice, unsafe.Pointer(enbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEnbIDChoice(bytes []byte) (*e2sm_kpm_v2.EnbIdChoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEnbIDChoice((*C.ENB_ID_Choice_t)(unsafePtr))
}

func perDecodeEnbIDChoice(bytes []byte) (*e2sm_kpm_v2.EnbIdChoice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEnbIDChoice((*C.ENB_ID_Choice_t)(unsafePtr))
}

func newEnbIDChoice(enbIDChoice *e2sm_kpm_v2.EnbIdChoice) (*C.ENB_ID_Choice_t, error) {

	var pr C.ENB_ID_Choice_PR
	choiceC := [48]byte{}

	switch choice := enbIDChoice.EnbIdChoice.(type) {
	case *e2sm_kpm_v2.EnbIdChoice_EnbIdMacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_macro

		im, err := newBitString(choice.EnbIdMacro)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.EnbIdChoice_EnbIdShortmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_shortmacro

		im, err := newBitString(choice.EnbIdShortmacro)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.EnbIdChoice_EnbIdLongmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_longmacro

		im, err := newBitString(choice.EnbIdLongmacro)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[16:24], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newEnbIDChoice() %T not yet implemented", choice)
	}

	enbIDChoiceC := C.ENB_ID_Choice_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDChoiceC, nil
}

func decodeEnbIDChoice(enbIDChoiceC *C.ENB_ID_Choice_t) (*e2sm_kpm_v2.EnbIdChoice, error) {

	enbIDChoice := new(e2sm_kpm_v2.EnbIdChoice)

	switch enbIDChoiceC.present {
	case C.ENB_ID_Choice_PR_enb_ID_macro:
		var a [8]byte
		copy(a[:], enbIDChoiceC.choice[0:8])
		enbIDChoicestructC, err := decodeBitStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbIDChoice.EnbIdChoice = &e2sm_kpm_v2.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: enbIDChoicestructC,
		}
	case C.ENB_ID_Choice_PR_enb_ID_shortmacro:
		var a [8]byte
		copy(a[:], enbIDChoiceC.choice[8:16])
		enbIDChoicestructC, err := decodeBitStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbIDChoice.EnbIdChoice = &e2sm_kpm_v2.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: enbIDChoicestructC,
		}
	case C.ENB_ID_Choice_PR_enb_ID_longmacro:
		var a [8]byte
		copy(a[:], enbIDChoiceC.choice[16:24])
		enbIDChoicestructC, err := decodeBitStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbIDChoice.EnbIdChoice = &e2sm_kpm_v2.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: enbIDChoicestructC,
		}
	default:
		return nil, fmt.Errorf("decodeEnbIDChoice() %v not yet implemented", enbIDChoiceC.present)
	}

	return enbIDChoice, nil
}

func decodeEnbIDChoiceBytes(array [48]byte) (*e2sm_kpm_v2.EnbIdChoice, error) {
	enbIDChoiceC := (*C.ENB_ID_Choice_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeEnbIDChoice(enbIDChoiceC)
}
