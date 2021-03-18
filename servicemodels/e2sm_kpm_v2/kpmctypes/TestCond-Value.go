// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TestCond-Value.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeTestCondValue(testCondValue *e2sm_kpm_v2.TestCondValue) ([]byte, error) {
	testCondValueCP, err := newTestCondValue(testCondValue)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondValue() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TestCond_Value, unsafe.Pointer(testCondValueCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondValue() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestCondValue(testCondValue *e2sm_kpm_v2.TestCondValue) ([]byte, error) {
	testCondValueCP, err := newTestCondValue(testCondValue)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondValue() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TestCond_Value, unsafe.Pointer(testCondValueCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondValue() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestCondValue(bytes []byte) (*e2sm_kpm_v2.TestCondValue, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TestCond_Value)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestCondValue((*C.TestCond_Value_t)(unsafePtr))
}

func perDecodeTestCondValue(bytes []byte) (*e2sm_kpm_v2.TestCondValue, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TestCond_Value)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestCondValue((*C.TestCond_Value_t)(unsafePtr))
}

func newTestCondValue(testCondValue *e2sm_kpm_v2.TestCondValue) (*C.TestCond_Value_t, error) {

	var pr C.TestCond_Value_PR
	choiceC := [48]byte{}

	switch choice := testCondValue.TestCondValue.(type) {
	case *e2sm_kpm_v2.TestCondValue_ValueInt:
		pr = C.TestCond_Value_PR_valueInt

		binary.LittleEndian.PutUint64(choiceC[:], uint64(choice.ValueInt))
	case *e2sm_kpm_v2.TestCondValue_ValueEnum:
		pr = C.TestCond_Value_PR_valueEnum

		binary.LittleEndian.PutUint64(choiceC[:], uint64(choice.ValueEnum))
		//ToDo - Implement Boolean.go
	case *e2sm_kpm_v2.TestCondValue_ValueBool:
		pr = C.TestCond_Value_PR_valueBool

		fmt.Printf("Boolean is %v\n", choice.ValueBool)
		bC := newBoolean(choice.ValueBool)
		binary.LittleEndian.PutUint32(choiceC[:], uint32(*bC))
	case *e2sm_kpm_v2.TestCondValue_ValueBitS:
		pr = C.TestCond_Value_PR_valueBitS

		im, err := newBitString(choice.ValueBitS)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
		binary.LittleEndian.PutUint32(choiceC[16:24], uint32(im.bits_unused))
	case *e2sm_kpm_v2.TestCondValue_ValueOctS:
		pr = C.TestCond_Value_PR_valueOctS

		im, err := newOctetString(choice.ValueOctS)
		if err != nil {
			return nil, fmt.Errorf("newOctetString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
	case *e2sm_kpm_v2.TestCondValue_ValuePrtS:
		pr = C.TestCond_Value_PR_valuePrtS

		im, err := newPrintableString(choice.ValuePrtS)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
	default:
		return nil, fmt.Errorf("newTestCondValue() %T not yet implemented", choice)
	}

	testCondValueC := C.TestCond_Value_t{
		present: pr,
		choice:  choiceC,
	}

	return &testCondValueC, nil
}

func decodeTestCondValue(testCondValueC *C.TestCond_Value_t) (*e2sm_kpm_v2.TestCondValue, error) {

	testCondValue := new(e2sm_kpm_v2.TestCondValue)

	switch testCondValueC.present {
	case C.TestCond_Value_PR_valueInt:
		intgr := int64(binary.LittleEndian.Uint64(testCondValueC.choice[:]))
		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValueInt{
			ValueInt: intgr,
		}
	case C.TestCond_Value_PR_valueEnum:
		enm := int64(binary.LittleEndian.Uint64(testCondValueC.choice[:]))
		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValueEnum{
			ValueEnum: enm,
		}
		//ToDo - Implement Boolean.go
	case C.TestCond_Value_PR_valueBool:
		var a [8]byte
		copy(a[:], testCondValueC.choice[:8])
		b := decodeBooleanBytes(a)

		fmt.Printf("Decoded Boolean is %v\n", b)
		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValueBool{
			ValueBool: b,
		}
	case C.TestCond_Value_PR_valueBitS:
		bsC := newBitStringFromArray(testCondValueC.choice)

		bs, err := decodeBitString(bsC)
		if err != nil {
			return nil, fmt.Errorf("decodeBitString() %s", err.Error())
		}

		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValueBitS{
			ValueBitS: bs,
		}
	case C.TestCond_Value_PR_valueOctS:
		//ToDo: Fix decoding
		var a [16]byte
		copy(a[:], testCondValueC.choice[:16])
		octS, err := decodeOctetStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeOctetStringBytes() %s", err.Error())
		}

		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValueOctS{
			ValueOctS: octS,
		}
	case C.TestCond_Value_PR_valuePrtS:
		//ToDo - Fix decoding
		var a [16]byte
		copy(a[:], testCondValueC.choice[:16])
		prtS, err := decodePrintableStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableStringBytes() %s", err.Error())
		}
		testCondValue.TestCondValue = &e2sm_kpm_v2.TestCondValue_ValuePrtS{
			ValuePrtS: prtS,
		}
	default:
		return nil, fmt.Errorf("decodeTestCondValue() %v not yet implemented", testCondValueC.present)
	}

	return testCondValue, nil
}
