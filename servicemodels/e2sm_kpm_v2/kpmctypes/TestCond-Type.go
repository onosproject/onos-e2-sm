// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TestCond-Type.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeTestCondType(testCondType *e2sm_kpm_v2.TestCondType) ([]byte, error) {
	testCondTypeCP, err := newTestCondType(testCondType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TestCond_Type, unsafe.Pointer(testCondTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestCondType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTestCondType(testCondType *e2sm_kpm_v2.TestCondType) ([]byte, error) {
	testCondTypeCP, err := newTestCondType(testCondType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TestCond_Type, unsafe.Pointer(testCondTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTestCondType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestCondType(bytes []byte) (*e2sm_kpm_v2.TestCondType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TestCond_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestCondType((*C.TestCond_Type_t)(unsafePtr))
}

func perDecodeTestCondType(bytes []byte) (*e2sm_kpm_v2.TestCondType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TestCond_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestCondType((*C.TestCond_Type_t)(unsafePtr))
}

func newTestCondType(testCondType *e2sm_kpm_v2.TestCondType) (*C.TestCond_Type_t, error) {

	var pr C.TestCond_Type_PR
	choiceC := [8]byte{}

	switch choice := testCondType.TestCondType.(type) {
	case *e2sm_kpm_v2.TestCondType_GBr:
		pr = C.TestCond_Type_PR_gBR

		var gbrC C.TestCond_Type__gBR_t
		switch choice.GBr {
		case e2sm_kpm_v2.GBR_GBR_TRUE:
			gbrC = C.TestCond_Type__gBR_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType GBR %v", choice.GBr)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(gbrC))
	case *e2sm_kpm_v2.TestCondType_AMbr:
		pr = C.TestCond_Type_PR_aMBR

		var ambrC C.TestCond_Type__aMBR_t
		switch choice.AMbr {
		case e2sm_kpm_v2.AMBR_AMBR_TRUE:
			ambrC = C.TestCond_Type__aMBR_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType AMBR %v", choice.AMbr)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(ambrC))
	case *e2sm_kpm_v2.TestCondType_IsStat:
		pr = C.TestCond_Type_PR_isStat

		var isStatC C.TestCond_Type__isStat_t
		switch choice.IsStat {
		case e2sm_kpm_v2.ISSTAT_ISSTAT_TRUE:
			isStatC = C.TestCond_Type__aMBR_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType IsStat %v", choice.IsStat)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(isStatC))
	case *e2sm_kpm_v2.TestCondType_IsCatM:
		pr = C.TestCond_Type_PR_isCatM

		var isCatMC C.TestCond_Type__isCatM_t
		switch choice.IsCatM {
		case e2sm_kpm_v2.ISCATM_ISCATM_TRUE:
			isCatMC = C.TestCond_Type__isCatM_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType IsCatM %v", choice.IsCatM)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(isCatMC))
	case *e2sm_kpm_v2.TestCondType_RSrp:
		pr = C.TestCond_Type_PR_rSRP

		var rsrpC C.TestCond_Type__rSRP_t
		switch choice.RSrp {
		case e2sm_kpm_v2.RSRP_RSRP_TRUE:
			rsrpC = C.TestCond_Type__rSRP_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType RSRP %v", choice.RSrp)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(rsrpC))
	case *e2sm_kpm_v2.TestCondType_RSrq:
		pr = C.TestCond_Type_PR_rSRQ

		var rsrqC C.TestCond_Type__rSRQ_t
		switch choice.RSrq {
		case e2sm_kpm_v2.RSRQ_RSRQ_TRUE:
			rsrqC = C.TestCond_Type__rSRQ_true
		default:
			return nil, fmt.Errorf("unexpected TestCondType RSRP %v", choice.RSrq)
		}

		binary.LittleEndian.PutUint32(choiceC[0:], uint32(rsrqC))
	default:
		return nil, fmt.Errorf("newTestCondType() %T not yet implemented", choice)
	}

	testCondTypeC := C.TestCond_Type_t{
		present: pr,
		choice:  choiceC,
	}

	return &testCondTypeC, nil
}

func decodeTestCondType(testCondTypeC *C.TestCond_Type_t) (*e2sm_kpm_v2.TestCondType, error) {

	testCondType := new(e2sm_kpm_v2.TestCondType)

	switch testCondTypeC.present {
	case C.TestCond_Type_PR_gBR:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_GBr{
			GBr: e2sm_kpm_v2.GBR_GBR_TRUE,
		}
		//int32(binary.LittleEndian.Uint32(testCondTypeC.choice[:8]))
	case C.TestCond_Type_PR_aMBR:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
		}
	case C.TestCond_Type_PR_isStat:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_IsStat{
			IsStat: e2sm_kpm_v2.ISSTAT_ISSTAT_TRUE,
		}
	case C.TestCond_Type_PR_isCatM:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_IsCatM{
			IsCatM: e2sm_kpm_v2.ISCATM_ISCATM_TRUE,
		}
	case C.TestCond_Type_PR_rSRP:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_RSrp{
			RSrp: e2sm_kpm_v2.RSRP_RSRP_TRUE,
		}
	case C.TestCond_Type_PR_rSRQ:
		testCondType.TestCondType = &e2sm_kpm_v2.TestCondType_RSrq{
			RSrq: e2sm_kpm_v2.RSRQ_RSRQ_TRUE,
		}
	case C.TestCond_Type_PR_NOTHING:
		return nil, fmt.Errorf("decodeTestCondType() An empty TestCondType has been sent %v", testCondTypeC.present)
	default:
		return nil, fmt.Errorf("decodeTestCondType() %v not yet implemented", testCondTypeC.present)
	}

	return testCondType, nil
}
