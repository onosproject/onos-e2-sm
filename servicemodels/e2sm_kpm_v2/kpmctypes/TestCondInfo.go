// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TestCondInfo.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeTestCondInfo(testCondInfo *e2sm_kpm_v2.TestCondInfo) ([]byte, error) {
	testCondInfoCP, err := newTestCondInfo(testCondInfo)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondInfo() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TestCondInfo, unsafe.Pointer(testCondInfoCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondInfo() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestCondInfo(testCondInfo *e2sm_kpm_v2.TestCondInfo) ([]byte, error) {
	testCondInfoCP, err := newTestCondInfo(testCondInfo)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondInfo() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TestCondInfo, unsafe.Pointer(testCondInfoCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondInfo() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestCondInfo(bytes []byte) (*e2sm_kpm_v2.TestCondInfo, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TestCondInfo)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestCondInfo((*C.TestCondInfo_t)(unsafePtr))
}

func perDecodeTestCondInfo(bytes []byte) (*e2sm_kpm_v2.TestCondInfo, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TestCondInfo)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestCondInfo((*C.TestCondInfo_t)(unsafePtr))
}

func newTestCondInfo(testCondInfo *e2sm_kpm_v2.TestCondInfo) (*C.TestCondInfo_t, error) {

	testTypeC, err := newTestCondType(testCondInfo.TestType)
	if err != nil {
		return nil, fmt.Errorf("newTestCondType() %s", err.Error())
	}

	testExprC, err := newTestCondExpression(testCondInfo.TestExpr)
	if err != nil {
		return nil, fmt.Errorf("newTestCondExpression() %s", err.Error())
	}

	testValueC, err := newTestCondValue(testCondInfo.TestValue)
	if err != nil {
		return nil, fmt.Errorf("newTestCondValue() %s", err.Error())
	}

	testCondInfoC := C.TestCondInfo_t{
		testType:  *testTypeC,
		testExpr:  *testExprC,
		testValue: *testValueC,
	}

	return &testCondInfoC, nil
}

func decodeTestCondInfo(testCondInfoC *C.TestCondInfo_t) (*e2sm_kpm_v2.TestCondInfo, error) {

	testType, err := decodeTestCondType(&testCondInfoC.testType)
	if err != nil {
		return nil, fmt.Errorf("decodeTestCondType() %s", err.Error())
	}

	testExpr, err := decodeTestCondExpression(&testCondInfoC.testExpr)
	if err != nil {
		return nil, fmt.Errorf("decodeTestCondExpression() %s", err.Error())
	}

	testValue, err := decodeTestCondValue(&testCondInfoC.testValue)
	if err != nil {
		return nil, fmt.Errorf("decodeTestCondValue() %s", err.Error())
	}

	testCondInfo := e2sm_kpm_v2.TestCondInfo{
		TestType:  testType,
		TestExpr:  *testExpr,
		TestValue: testValue,
	}

	return &testCondInfo, nil
}

func decodeTestCondInfoBytes(array [8]byte) (*e2sm_kpm_v2.TestCondInfo, error) {
	testCondInfoC := (*C.TestCondInfo_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeTestCondInfo(testCondInfoC)
}
