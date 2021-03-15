// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementTypeName.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementTypeName(measurementTypeName *e2sm_kpm_v2.MeasurementTypeName) ([]byte, error) {
	measurementTypeNameCP, err := newMeasurementTypeName(measurementTypeName)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementTypeName() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementTypeName, unsafe.Pointer(measurementTypeNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementTypeName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementTypeName(measurementTypeName *e2sm_kpm_v2.MeasurementTypeName) ([]byte, error) {
	measurementTypeNameCP, err := newMeasurementTypeName(measurementTypeName)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementTypeName() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementTypeName, unsafe.Pointer(measurementTypeNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementTypeName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementTypeName(bytes []byte) (*e2sm_kpm_v2.MeasurementTypeName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementTypeName)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementTypeName((*C.MeasurementTypeName_t)(unsafePtr))
}

func perDecodeMeasurementTypeName(bytes []byte) (*e2sm_kpm_v2.MeasurementTypeName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementTypeName)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementTypeName((*C.MeasurementTypeName_t)(unsafePtr))
}

func newMeasurementTypeName(measurementTypeName *e2sm_kpm_v2.MeasurementTypeName) (*C.MeasurementTypeName_t, error) {

	measurementTypeNameC, err := newPrintableString(measurementTypeName.Value)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	return measurementTypeNameC, nil
}

func decodeMeasurementTypeName(measurementTypeNameC *C.MeasurementTypeName_t) (*e2sm_kpm_v2.MeasurementTypeName, error) {

	measurementTypeName := new(e2sm_kpm_v2.MeasurementTypeName)
	measurementTypeNameValue, err := decodePrintableString(measurementTypeNameC)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}
	measurementTypeName.Value = measurementTypeNameValue

	return measurementTypeName, nil
}

func decodeMeasurementTypeNameBytes(array [16]byte) (*e2sm_kpm_v2.MeasurementTypeName, error) {
	measurementTypeNameC := C.MeasurementTypeName_t{
		buf:  (*C.uchar)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8])))),
		size: C.ulong(binary.LittleEndian.Uint64(array[8:16])),
	}

	return decodeMeasurementTypeName(&measurementTypeNameC)
}
