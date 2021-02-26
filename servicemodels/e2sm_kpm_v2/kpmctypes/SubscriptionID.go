// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "SubscriptionID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeSubscriptionID(subscriptionId *e2sm_kpm_v2.SubscriptionId) ([]byte, error) {
	subscriptionIdCP, err := newSubscriptionID(subscriptionId)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSubscriptionId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_SubscriptionID, unsafe.Pointer(subscriptionIdCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSubscriptionId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeSubscriptionID(subscriptionId *e2sm_kpm_v2.SubscriptionId) ([]byte, error) {
	subscriptionIdCP, err := newSubscriptionID(subscriptionId)
	if err != nil {
		return nil, fmt.Errorf("perEncodeSubscriptionId() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_SubscriptionID, unsafe.Pointer(subscriptionIdCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeSubscriptionId() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeSubscriptionID(bytes []byte) (*e2sm_kpm_v2.SubscriptionId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_SubscriptionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeSubscriptionID((*C.SubscriptionID_t)(unsafePtr))
}

func perDecodeSubscriptionID(bytes []byte) (*e2sm_kpm_v2.SubscriptionId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_SubscriptionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeSubscriptionID((*C.SubscriptionID_t)(unsafePtr))
}

func newSubscriptionID(subscriptionId *e2sm_kpm_v2.SubscriptionId) (*C.SubscriptionID_t, error) {

	subscriptionIdC, err := newInteger(subscriptionId.Value)
	if err != nil {
		return nil, fmt.Errorf("newInt64() %s", err.Error())
	}

	return &subscriptionIdC, nil
}

func decodeSubscriptionID(subscriptionIdC *C.SubscriptionID_t) (*e2sm_kpm_v2.SubscriptionId, error) {

	subscriptionId := new(e2sm_kpm_v2.SubscriptionId)
	subscriptionIdValue, err := decodeInteger(subscriptionIdC)
	if err != nil {
		return nil, fmt.Errorf("decodeInt64() %s", err.Error())
	}
	subscriptionId.Value = subscriptionIdValue

	return subscriptionId, nil
}

func decodeSubscriptionIdBytes(array [8]byte) (*e2sm_kpm_v2.SubscriptionId, error) {
	subscriptionIdC := (*C.SubscriptionID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeSubscriptionID(subscriptionIdC)
}
