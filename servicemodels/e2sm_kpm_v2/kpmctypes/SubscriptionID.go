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
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeSubscriptionID(subscriptionID *e2sm_kpm_v2.SubscriptionId) ([]byte, error) {
	subscriptionIDCP, err := newSubscriptionID(subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSubscriptionId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_SubscriptionID, unsafe.Pointer(subscriptionIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSubscriptionId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeSubscriptionID(subscriptionID *e2sm_kpm_v2.SubscriptionId) ([]byte, error) {
	subscriptionIDCP, err := newSubscriptionID(subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeSubscriptionId() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_SubscriptionID, unsafe.Pointer(subscriptionIDCP))
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

func newSubscriptionID(subscriptionID *e2sm_kpm_v2.SubscriptionId) (*C.SubscriptionID_t, error) {

	//subscriptionIDC, err := newInteger(subscriptionID.Value)
	//if err != nil {
	//	return nil, fmt.Errorf("newInt64() %s", err.Error())
	//}

	res := C.ulong(subscriptionID.Value)

	return &res, nil
}

func decodeSubscriptionID(subscriptionIDC *C.SubscriptionID_t) (*e2sm_kpm_v2.SubscriptionId, error) {

	subscriptionID := new(e2sm_kpm_v2.SubscriptionId)
	//subscriptionIDValue, err := decodeInteger(subscriptionIDC)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeInteger() %s", err.Error())
	//}
	subscriptionID.Value = int64(*subscriptionIDC)

	return subscriptionID, nil
}
