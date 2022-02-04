// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

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
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeENbIDChoice(enbIDchoice *e2sm_kpm_ies.EnbIdChoice) ([]byte, error) {

	enbIDchoiceCP, err := newENbIDChoice(enbIDchoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENbIDChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID_Choice, unsafe.Pointer(enbIDchoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeENbIDChoice(bytes []byte) (*e2sm_kpm_ies.EnbIdChoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeENbIDChoice() pointer decoded from XER is nil")
	}
	return decodeENbIDChoice((*C.ENB_ID_Choice_t)(unsafePtr))
}

func newENbIDChoice(enbIDchoice *e2sm_kpm_ies.EnbIdChoice) (*C.ENB_ID_Choice_t, error) {
	var pr C.ENB_ID_Choice_PR
	choiceC := [48]byte{}

	switch choice := enbIDchoice.EnbIdChoice.(type) {
	case *e2sm_kpm_ies.EnbIdChoice_EnbIdMacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_macro
		macroENbIDchoiceC := newBitString(choice.EnbIdMacro)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(macroENbIDchoiceC))))
	case *e2sm_kpm_ies.EnbIdChoice_EnbIdShortmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_shortmacro
		shortMacroEnbIDchoiceC := newBitString(choice.EnbIdShortmacro)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(shortMacroEnbIDchoiceC))))
	case *e2sm_kpm_ies.EnbIdChoice_EnbIdLongmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_longmacro
		longMacroEnbIDchoiceC := newBitString(choice.EnbIdLongmacro)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(longMacroEnbIDchoiceC))))
	default:
		return nil, fmt.Errorf("newENbIDChoice() unexpected type %T", choice)
	}

	enbIDC := C.ENB_ID_Choice_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDC, nil
}

func decodeENbIDChoice(enbIDchoiceC *C.ENB_ID_Choice_t) (*e2sm_kpm_ies.EnbIdChoice, error) {
	result := new(e2sm_kpm_ies.EnbIdChoice)

	switch enbIDchoiceC.present {
	case C.ENB_ID_Choice_PR_enb_ID_macro:

		enbIDstructC := newBitStringFromArray(enbIDchoiceC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbIdChoice = &e2sm_kpm_ies.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: enb,
		}
	case C.ENB_ID_Choice_PR_enb_ID_shortmacro:

		enbIDstructC := newBitStringFromArray(enbIDchoiceC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbIdChoice = &e2sm_kpm_ies.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: enb,
		}
	case C.ENB_ID_Choice_PR_enb_ID_longmacro:

		enbIDstructC := newBitStringFromArray(enbIDchoiceC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbIdChoice = &e2sm_kpm_ies.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: enb,
		}
	default:
		return nil, fmt.Errorf("decodeENbIDChoice() %v unknown type", enbIDchoiceC.present)
	}

	return result, nil
}
