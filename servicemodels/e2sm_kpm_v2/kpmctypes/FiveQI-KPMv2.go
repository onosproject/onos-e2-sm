// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "FiveQI-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeFiveQi(fiveQi *e2sm_kpm_v2.FiveQi) ([]byte, error) {
	fiveQiCP, err := newFiveQi(fiveQi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeFiveQi() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_FiveQI_KPMv2, unsafe.Pointer(fiveQiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeFiveQi() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeFiveQi(fiveQi *e2sm_kpm_v2.FiveQi) ([]byte, error) {
	fiveQiCP, err := newFiveQi(fiveQi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeFiveQi() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_FiveQI_KPMv2, unsafe.Pointer(fiveQiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeFiveQi() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeFiveQi(bytes []byte) (*e2sm_kpm_v2.FiveQi, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_FiveQI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeFiveQi((*C.FiveQI_KPMv2_t)(unsafePtr))
}

func perDecodeFiveQi(bytes []byte) (*e2sm_kpm_v2.FiveQi, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_FiveQI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeFiveQi((*C.FiveQI_KPMv2_t)(unsafePtr))
}

func newFiveQi(fiveQi *e2sm_kpm_v2.FiveQi) (*C.FiveQI_KPMv2_t, error) {

	fiveQiC := C.long(fiveQi.Value)

	return &fiveQiC, nil
}

func decodeFiveQi(fiveQiC *C.FiveQI_KPMv2_t) (*e2sm_kpm_v2.FiveQi, error) {

	fiveQi := e2sm_kpm_v2.FiveQi{
		Value: int32(*fiveQiC),
	}

	return &fiveQi, nil
}
