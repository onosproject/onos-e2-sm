// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MatchingCondItem.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMatchingCondItem(matchingCondItem *e2sm_kpm_v2.MatchingCondItem) ([]byte, error) {
	matchingCondItemCP, err := newMatchingCondItem(matchingCondItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingCondItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MatchingCondItem, unsafe.Pointer(matchingCondItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingCondItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMatchingCondItem(matchingCondItem *e2sm_kpm_v2.MatchingCondItem) ([]byte, error) {
	matchingCondItemCP, err := newMatchingCondItem(matchingCondItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingCondItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MatchingCondItem, unsafe.Pointer(matchingCondItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingCondItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMatchingCondItem(bytes []byte) (*e2sm_kpm_v2.MatchingCondItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MatchingCondItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMatchingCondItem((*C.MatchingCondItem_t)(unsafePtr))
}

func perDecodeMatchingCondItem(bytes []byte) (*e2sm_kpm_v2.MatchingCondItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MatchingCondItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMatchingCondItem((*C.MatchingCondItem_t)(unsafePtr))
}

func newMatchingCondItem(matchingCondItem *e2sm_kpm_v2.MatchingCondItem) (*C.MatchingCondItem_t, error) {

	var pr C.MatchingCondItem_PR
	choiceC := [8]byte{}

	switch choice := matchingCondItem.MatchingCondItem.(type) {
	case *e2sm_kpm_v2.MatchingCondItem_MeasLabel:
		pr = C.MatchingCondItem_PR_measLabel

		im, err := newMeasurementLabel(choice.MeasLabel)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementLabel() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.MatchingCondItem_TestCondInfo:
		pr = C.MatchingCondItem_PR_testCondInfo

		im, err := newTestCondInfo(choice.TestCondInfo)
		if err != nil {
			return nil, fmt.Errorf("newTestCondInfo() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newMatchingCondItem() %T not yet implemented", choice)
	}

	matchingCondItemC := C.MatchingCondItem_t{
		present: pr,
		choice:  choiceC,
	}

	return &matchingCondItemC, nil
}

func decodeMatchingCondItem(matchingCondItemC *C.MatchingCondItem_t) (*e2sm_kpm_v2.MatchingCondItem, error) {

	matchingCondItem := new(e2sm_kpm_v2.MatchingCondItem)

	switch matchingCondItemC.present {
	case C.MatchingCondItem_PR_measLabel:
		matchingCondItemstructC, err := decodeMeasurementLabelBytes(matchingCondItemC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementLabelBytes() %s", err.Error())
		}
		matchingCondItem.MatchingCondItem = &e2sm_kpm_v2.MatchingCondItem_MeasLabel{
			MeasLabel: matchingCondItemstructC,
		}
	case C.MatchingCondItem_PR_testCondInfo:
		matchingCondItemstructC, err := decodeTestCondInfoBytes(matchingCondItemC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeTestCondInfoBytes() %s", err.Error())
		}
		matchingCondItem.MatchingCondItem = &e2sm_kpm_v2.MatchingCondItem_TestCondInfo{
			TestCondInfo: matchingCondItemstructC,
		}
	case C.MatchingCondItem_PR_NOTHING:
		return nil, fmt.Errorf("decodeMatchingCondItem() An empty MatchingCondItem has been sent %v", matchingCondItemC.present)
	default:
		return nil, fmt.Errorf("decodeMatchingCondItem() %v not yet implemented", matchingCondItemC.present)
	}

	return matchingCondItem, nil
}
