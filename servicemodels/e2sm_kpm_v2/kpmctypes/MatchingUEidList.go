// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MatchingUEidList.h"
//#include "MatchingUEidItem.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMatchingUeIDList(matchingUeIDList *e2sm_kpm_v2.MatchingUeidList) ([]byte, error) {
	matchingUeIDListCP, err := newMatchingUeIDList(matchingUeIDList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingUeIDList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MatchingUEidList, unsafe.Pointer(matchingUeIDListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingUeIDList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMatchingUeIDList(matchingUeIDList *e2sm_kpm_v2.MatchingUeidList) ([]byte, error) {
	matchingUeIDListCP, err := newMatchingUeIDList(matchingUeIDList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingUeIDList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MatchingUEidList, unsafe.Pointer(matchingUeIDListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingUeIDList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMatchingUeIDList(bytes []byte) (*e2sm_kpm_v2.MatchingUeidList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MatchingUEidList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMatchingUeIDList((*C.MatchingUEidList_t)(unsafePtr))
}

func perDecodeMatchingUeIDList(bytes []byte) (*e2sm_kpm_v2.MatchingUeidList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MatchingUEidList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMatchingUeIDList((*C.MatchingUEidList_t)(unsafePtr))
}

func newMatchingUeIDList(matchingUeIDList *e2sm_kpm_v2.MatchingUeidList) (*C.MatchingUEidList_t, error) {

	matchingUeIDListC := new(C.MatchingUEidList_t)
	for _, valueItem := range matchingUeIDList.GetValue() {
		itemC, err := newMatchingUeIDItem(valueItem)
		if err != nil {
			return nil, fmt.Errorf("newMatchingUeIDItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(matchingUeIDListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return matchingUeIDListC, nil
}

func decodeMatchingUeIDList(matchingUeIDListC *C.MatchingUEidList_t) (*e2sm_kpm_v2.MatchingUeidList, error) {

	matchingUeIDList := new(e2sm_kpm_v2.MatchingUeidList)

	ieCount := int(matchingUeIDListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(matchingUeIDListC.list.array)) * uintptr(i)
		ieC := *(**C.MatchingUEidItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(matchingUeIDListC.list.array)) + offset))
		ie, err := decodeMatchingUeIDItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMatchingUeIDItem() %s", err.Error())
		}
		matchingUeIDList.Value = append(matchingUeIDList.Value, ie)
	}

	return matchingUeIDList, nil
}
