// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MatchingCondList-KPMv2.h"
//#include "MatchingCondItem-KPMv2.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMatchingCondList(matchingCondList *e2sm_kpm_v2.MatchingCondList) ([]byte, error) {
	matchingCondListCP, err := newMatchingCondList(matchingCondList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingCondList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MatchingCondList_KPMv2, unsafe.Pointer(matchingCondListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingCondList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMatchingCondList(matchingCondList *e2sm_kpm_v2.MatchingCondList) ([]byte, error) {
	matchingCondListCP, err := newMatchingCondList(matchingCondList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingCondList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MatchingCondList_KPMv2, unsafe.Pointer(matchingCondListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingCondList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMatchingCondList(bytes []byte) (*e2sm_kpm_v2.MatchingCondList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MatchingCondList_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMatchingCondList((*C.MatchingCondList_KPMv2_t)(unsafePtr))
}

func perDecodeMatchingCondList(bytes []byte) (*e2sm_kpm_v2.MatchingCondList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MatchingCondList_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMatchingCondList((*C.MatchingCondList_KPMv2_t)(unsafePtr))
}

func newMatchingCondList(matchingCondList *e2sm_kpm_v2.MatchingCondList) (*C.MatchingCondList_KPMv2_t, error) {

	matchingCondListC := new(C.MatchingCondList_KPMv2_t)
	for _, valueItem := range matchingCondList.GetValue() {
		itemC, err := newMatchingCondItem(valueItem)
		if err != nil {
			return nil, fmt.Errorf("newMatchingCondItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(matchingCondListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return matchingCondListC, nil
}

func decodeMatchingCondList(matchingCondListC *C.MatchingCondList_KPMv2_t) (*e2sm_kpm_v2.MatchingCondList, error) {

	matchingCondList := new(e2sm_kpm_v2.MatchingCondList)

	ieCount := int(matchingCondListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(matchingCondListC.list.array)) * uintptr(i)
		ieC := *(**C.MatchingCondItem_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(matchingCondListC.list.array)) + offset))
		ie, err := decodeMatchingCondItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMatchingCondItem() %s", err.Error())
		}
		matchingCondList.Value = append(matchingCondList.Value, ie)
	}

	return matchingCondList, nil
}
