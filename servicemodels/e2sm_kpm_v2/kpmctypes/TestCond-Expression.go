// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TestCond-Expression.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeTestCondExpression(testCondExpression e2sm_kpm_v2.TestCondExpression) ([]byte, error) {
	testCondExpressionCP, err := newTestCondExpression(testCondExpression)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_TestCond_Expression, unsafe.Pointer(testCondExpressionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondExpression() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestCondExpression(testCondExpression e2sm_kpm_v2.TestCondExpression) ([]byte, error) {
	testCondExpressionCP, err := newTestCondExpression(testCondExpression)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TestCond_Expression, unsafe.Pointer(testCondExpressionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondExpression() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestCondExpression(bytes []byte) (*e2sm_kpm_v2.TestCondExpression, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TestCond_Expression)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestCondExpression((*C.TestCond_Expression_t)(unsafePtr))
}

func perDecodeTestCondExpression(bytes []byte) (*e2sm_kpm_v2.TestCondExpression, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TestCond_Expression)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestCondExpression((*C.TestCond_Expression_t)(unsafePtr))
}

func newTestCondExpression(testCondExpression e2sm_kpm_v2.TestCondExpression) (*C.TestCond_Expression_t, error) {
	var ret C.TestCond_Expression_t
	switch testCondExpression {
	case e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_EQUAL:
		ret = C.TestCond_Expression_equal
	case e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN:
		ret = C.TestCond_Expression_greaterthan
	case e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN:
		ret = C.TestCond_Expression_lessthan
	case e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_CONTAINS:
		ret = C.TestCond_Expression_contains
	case e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_PRESENT:
		ret = C.TestCond_Expression_present
	default:
		return nil, fmt.Errorf("unexpected TestCondExpression %v", testCondExpression)
	}

	return &ret, nil
}

func decodeTestCondExpression(testCondExpressionC *C.TestCond_Expression_t) (*e2sm_kpm_v2.TestCondExpression, error) {

	testCondExpression := e2sm_kpm_v2.TestCondExpression(int32(*testCondExpressionC))

	return &testCondExpression, nil
}
