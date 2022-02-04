// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GranularityPeriod-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGranularityPeriod(granularityPeriod *e2sm_kpm_v2.GranularityPeriod) ([]byte, error) {
	granularityPeriodCP, err := newGranularityPeriod(granularityPeriod)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGranularityPeriod() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GranularityPeriod_KPMv2, unsafe.Pointer(granularityPeriodCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGranularityPeriod() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGranularityPeriod(granularityPeriod *e2sm_kpm_v2.GranularityPeriod) ([]byte, error) {
	granularityPeriodCP, err := newGranularityPeriod(granularityPeriod)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGranularityPeriod() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GranularityPeriod_KPMv2, unsafe.Pointer(granularityPeriodCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGranularityPeriod() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGranularityPeriod(bytes []byte) (*e2sm_kpm_v2.GranularityPeriod, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GranularityPeriod_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGranularityPeriod((*C.GranularityPeriod_KPMv2_t)(unsafePtr))
}

func perDecodeGranularityPeriod(bytes []byte) (*e2sm_kpm_v2.GranularityPeriod, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GranularityPeriod_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGranularityPeriod((*C.GranularityPeriod_KPMv2_t)(unsafePtr))
}

func newGranularityPeriod(granularityPeriod *e2sm_kpm_v2.GranularityPeriod) (*C.GranularityPeriod_KPMv2_t, error) {

	granularityPeriodC := C.ulong(granularityPeriod.Value)

	return &granularityPeriodC, nil
}

func decodeGranularityPeriod(granularityPeriodC *C.GranularityPeriod_KPMv2_t) (*e2sm_kpm_v2.GranularityPeriod, error) {

	granularityPeriod := e2sm_kpm_v2.GranularityPeriod{
		Value: uint32(*granularityPeriodC),
	}

	return &granularityPeriod, nil
}
