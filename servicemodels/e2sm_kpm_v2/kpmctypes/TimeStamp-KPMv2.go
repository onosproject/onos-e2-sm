// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TimeStamp-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeTimeStamp(timeStamp *e2sm_kpm_v2.TimeStamp) ([]byte, error) {
	timeStampCP, err := newTimeStamp(timeStamp)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTimeStamp() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TimeStamp_KPMv2, unsafe.Pointer(timeStampCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTimeStamp() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTimeStamp(timeStamp *e2sm_kpm_v2.TimeStamp) ([]byte, error) {
	timeStampCP, err := newTimeStamp(timeStamp)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTimeStamp() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TimeStamp_KPMv2, unsafe.Pointer(timeStampCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTimeStamp() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTimeStamp(bytes []byte) (*e2sm_kpm_v2.TimeStamp, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TimeStamp_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTimeStamp((*C.TimeStamp_KPMv2_t)(unsafePtr))
}

func perDecodeTimeStamp(bytes []byte) (*e2sm_kpm_v2.TimeStamp, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TimeStamp_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTimeStamp((*C.TimeStamp_KPMv2_t)(unsafePtr))
}

func newTimeStamp(timeStamp *e2sm_kpm_v2.TimeStamp) (*C.TimeStamp_KPMv2_t, error) {

	timeStampC, err := newOctetString(string(timeStamp.Value))
	if err != nil {
		return nil, err
	}

	return timeStampC, nil
}

func decodeTimeStamp(timeStampC *C.TimeStamp_KPMv2_t) (*e2sm_kpm_v2.TimeStamp, error) {

	res, err := decodeOctetString(timeStampC)
	if err != nil {
		return nil, err
	}

	timeStamp := e2sm_kpm_v2.TimeStamp{
		Value: []byte(res),
	}

	return &timeStamp, nil
}
