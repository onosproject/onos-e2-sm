// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MatchingUEidItem-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMatchingUeIDItem(matchingUeIDItem *e2sm_kpm_v2.MatchingUeidItem) ([]byte, error) {
	matchingUeIDItemCP, err := newMatchingUeIDItem(matchingUeIDItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingUeIDItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MatchingUEidItem_KPMv2, unsafe.Pointer(matchingUeIDItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMatchingUeidItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMatchingUeIDItem(matchingUeIDItem *e2sm_kpm_v2.MatchingUeidItem) ([]byte, error) {
	matchingUeIDItemCP, err := newMatchingUeIDItem(matchingUeIDItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingUeidItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MatchingUEidItem_KPMv2, unsafe.Pointer(matchingUeIDItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMatchingUeidItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMatchingUeIDItem(bytes []byte) (*e2sm_kpm_v2.MatchingUeidItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MatchingUEidItem_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMatchingUeIDItem((*C.MatchingUEidItem_KPMv2_t)(unsafePtr))
}

func perDecodeMatchingUeIDItem(bytes []byte) (*e2sm_kpm_v2.MatchingUeidItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MatchingUEidItem_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMatchingUeIDItem((*C.MatchingUEidItem_KPMv2_t)(unsafePtr))
}

func newMatchingUeIDItem(matchingUeIDItem *e2sm_kpm_v2.MatchingUeidItem) (*C.MatchingUEidItem_KPMv2_t, error) {

	ueIDC, err := newUeIdentity(matchingUeIDItem.UeId)
	if err != nil {
		return nil, fmt.Errorf("newUeIdentity() %s", err.Error())
	}

	matchingUeIDItemC := C.MatchingUEidItem_KPMv2_t{
		ueID: *ueIDC,
	}

	return &matchingUeIDItemC, nil
}

func decodeMatchingUeIDItem(matchingUeIDItemC *C.MatchingUEidItem_KPMv2_t) (*e2sm_kpm_v2.MatchingUeidItem, error) {

	ueID, err := decodeUeIdentity(&matchingUeIDItemC.ueID)
	if err != nil {
		return nil, fmt.Errorf("decodeUeIdentity() %s", err.Error())
	}

	matchingUeIDItem := e2sm_kpm_v2.MatchingUeidItem{
		UeId: ueID,
	}

	return &matchingUeIDItem, nil
}
