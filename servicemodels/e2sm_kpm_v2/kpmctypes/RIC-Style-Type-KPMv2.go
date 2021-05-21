// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Type-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeRicStyleType(ricStyleType *e2sm_kpm_v2.RicStyleType) ([]byte, error) {
	ricStyleTypeCP, err := newRicStyleType(ricStyleType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Type_KPMv2, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleType(ricStyleType *e2sm_kpm_v2.RicStyleType) ([]byte, error) {
	ricStyleTypeCP, err := newRicStyleType(ricStyleType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Type_KPMv2, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleType(bytes []byte) (*e2sm_kpm_v2.RicStyleType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Type_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_KPMv2_t)(unsafePtr))
}

func perDecodeRicStyleType(bytes []byte) (*e2sm_kpm_v2.RicStyleType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Type_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_KPMv2_t)(unsafePtr))
}

func newRicStyleType(ricStyleType *e2sm_kpm_v2.RicStyleType) (*C.RIC_Style_Type_KPMv2_t, error) {

	ricStyleTypeC := C.long(ricStyleType.Value)

	return &ricStyleTypeC, nil
}

func decodeRicStyleType(ricStyleTypeC *C.RIC_Style_Type_KPMv2_t) (*e2sm_kpm_v2.RicStyleType, error) {

	ricStyleType := e2sm_kpm_v2.RicStyleType{
		Value: int32(*ricStyleTypeC),
	}

	return &ricStyleType, nil
}
