// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-Value-RCPRE.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeRanparameterValue(ranparameterValue *e2sm_rc_pre_v2.RanparameterValue) ([]byte, error) {
	ranparameterValueCP, err := newRanparameterValue(ranparameterValue)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RANparameter_Value_RCPRE, unsafe.Pointer(ranparameterValueCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanparameterValue(ranparameterValue *e2sm_rc_pre_v2.RanparameterValue) ([]byte, error) {
	ranparameterValueCP, err := newRanparameterValue(ranparameterValue)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_Value_RCPRE, unsafe.Pointer(ranparameterValueCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanparameterValue() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanparameterValue(bytes []byte) (*e2sm_rc_pre_v2.RanparameterValue, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_Value_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanparameterValue((*C.RANparameter_Value_RCPRE_t)(unsafePtr))
}

func perDecodeRanparameterValue(bytes []byte) (*e2sm_rc_pre_v2.RanparameterValue, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_Value_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanparameterValue((*C.RANparameter_Value_RCPRE_t)(unsafePtr))
}

func newRanparameterValue(ranparameterValue *e2sm_rc_pre_v2.RanparameterValue) (*C.RANparameter_Value_RCPRE_t, error) {

	var pr C.RANparameter_Value_RCPRE_PR
	choiceC := [48]byte{}

	switch choice := ranparameterValue.RanparameterValue.(type) {
	case *e2sm_rc_pre_v2.RanparameterValue_ValueInt:
		pr = C.RANparameter_Value_RCPRE_PR_valueInt
		binary.LittleEndian.PutUint32(choiceC[:8], choice.ValueInt)
	case *e2sm_rc_pre_v2.RanparameterValue_ValueEnum:
		pr = C.RANparameter_Value_RCPRE_PR_valueEnum
		binary.LittleEndian.PutUint32(choiceC[:8], uint32(choice.ValueEnum))
	case *e2sm_rc_pre_v2.RanparameterValue_ValueBool:
		pr = C.RANparameter_Value_RCPRE_PR_valueBool

		bC := newBoolean(choice.ValueBool)
		binary.LittleEndian.PutUint32(choiceC[:], uint32(*bC))
	case *e2sm_rc_pre_v2.RanparameterValue_ValueBitS:
		pr = C.RANparameter_Value_RCPRE_PR_valueBitS

		im, err := newBitString(choice.ValueBitS)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
		binary.LittleEndian.PutUint32(choiceC[16:24], uint32(im.bits_unused))
	case *e2sm_rc_pre_v2.RanparameterValue_ValueOctS:
		pr = C.RANparameter_Value_RCPRE_PR_valueOctS

		im := newOctetString(choice.ValueOctS)
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
	case *e2sm_rc_pre_v2.RanparameterValue_ValuePrtS:
		pr = C.RANparameter_Value_RCPRE_PR_valuePrtS

		im := newPrintableString(choice.ValuePrtS)
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
	default:
		return nil, fmt.Errorf("newRanparameterValue() %T not yet implemented", choice)
	}

	ranparameterValueC := C.RANparameter_Value_RCPRE_t{
		present: pr,
		choice:  choiceC,
	}

	return &ranparameterValueC, nil
}

func decodeRanparameterValue(ranparameterValueC *C.RANparameter_Value_RCPRE_t) (*e2sm_rc_pre_v2.RanparameterValue, error) {
	ranparameterValue := new(e2sm_rc_pre_v2.RanparameterValue)

	switch ranparameterValueC.present {
	case C.RANparameter_Value_RCPRE_PR_valueInt:
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValueInt{
			ValueInt: binary.LittleEndian.Uint32(ranparameterValueC.choice[:8]),
		}
	case C.RANparameter_Value_RCPRE_PR_valueEnum:
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValueEnum{
			ValueEnum: int32(binary.LittleEndian.Uint32(ranparameterValueC.choice[:8])),
		}
	case C.RANparameter_Value_RCPRE_PR_valueBool:
		var a [8]byte
		copy(a[:], ranparameterValueC.choice[:8])
		b := decodeBooleanBytes(a)

		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValueBool{
			ValueBool: b,
		}
	case C.RANparameter_Value_RCPRE_PR_valueBitS:
		ranparameterValuestructC := newBitStringFromArray(ranparameterValueC.choice)

		ranparameterValuestruct, err := decodeBitString(ranparameterValuestructC)
		if err != nil {
			return nil, fmt.Errorf("decodeRanparameterValue() %s", err.Error())
		}
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValueBitS{
			ValueBitS: ranparameterValuestruct,
		}
	case C.RANparameter_Value_RCPRE_PR_valueOctS:
		var a [16]byte
		copy(a[:], ranparameterValueC.choice[:16])
		octS, err := decodeOctetStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeOctetStringBytes() %s", err.Error())
		}
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValueOctS{
			ValueOctS: octS,
		}
	case C.RANparameter_Value_RCPRE_PR_valuePrtS:
		var a [16]byte
		copy(a[:], ranparameterValueC.choice[:16])
		prtS, err := decodePrintableStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableStringBytes() %s", err.Error())
		}
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_v2.RanparameterValue_ValuePrtS{
			ValuePrtS: prtS,
		}
	default:
		return nil, fmt.Errorf("decodeRanparameterValue() %v not yet implemented", ranparameterValueC.present)
	}

	return ranparameterValue, nil
}
