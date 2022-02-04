// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Format-Type-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeRicFormatType(ricFormatType *e2sm_kpm_v2.RicFormatType) ([]byte, error) {
	ricFormatTypeCP, err := newRicFormatType(ricFormatType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicFormatType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_Format_Type_KPMv2, unsafe.Pointer(ricFormatTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicFormatType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicFormatType(ricFormatType *e2sm_kpm_v2.RicFormatType) ([]byte, error) {
	ricFormatTypeCP, err := newRicFormatType(ricFormatType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicFormatType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Format_Type_KPMv2, unsafe.Pointer(ricFormatTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicFormatType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicFormatType(bytes []byte) (*e2sm_kpm_v2.RicFormatType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Format_Type_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicFormatType((*C.RIC_Format_Type_KPMv2_t)(unsafePtr))
}

func perDecodeRicFormatType(bytes []byte) (*e2sm_kpm_v2.RicFormatType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Format_Type_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicFormatType((*C.RIC_Format_Type_KPMv2_t)(unsafePtr))
}

func newRicFormatType(ricFormatType *e2sm_kpm_v2.RicFormatType) (*C.RIC_Format_Type_KPMv2_t, error) {

	ricFormatTypeC := C.long(ricFormatType.Value)

	return &ricFormatTypeC, nil
}

func decodeRicFormatType(ricFormatTypeC *C.RIC_Format_Type_KPMv2_t) (*e2sm_kpm_v2.RicFormatType, error) {

	ricFormatType := e2sm_kpm_v2.RicFormatType{
		Value: int32(*ricFormatTypeC),
	}

	return &ricFormatType, nil
}
