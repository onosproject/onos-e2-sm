// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementTypeID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementTypeID(measurementTypeID *e2sm_kpm_v2.MeasurementTypeId) ([]byte, error) {
	measurementTypeIDCP, err := newMeasurementTypeID(measurementTypeID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementTypeId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementTypeID_KPMv2, unsafe.Pointer(measurementTypeIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementTypeId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementTypeID(measurementTypeID *e2sm_kpm_v2.MeasurementTypeId) ([]byte, error) {
	measurementTypeIDCP, err := newMeasurementTypeID(measurementTypeID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementTypeId() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementTypeID_KPMv2, unsafe.Pointer(measurementTypeIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementTypeId() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementTypeID(bytes []byte) (*e2sm_kpm_v2.MeasurementTypeId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementTypeID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementTypeID((*C.MeasurementTypeID_KPMv2_t)(unsafePtr))
}

func perDecodeMeasurementTypeID(bytes []byte) (*e2sm_kpm_v2.MeasurementTypeId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementTypeID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementTypeID((*C.MeasurementTypeID_KPMv2_t)(unsafePtr))
}

func newMeasurementTypeID(measurementTypeID *e2sm_kpm_v2.MeasurementTypeId) (*C.MeasurementTypeID_KPMv2_t, error) {

	measurementTypeIDC := C.long(measurementTypeID.Value)

	return &measurementTypeIDC, nil
}

func decodeMeasurementTypeID(measurementTypeIDC *C.MeasurementTypeID_KPMv2_t) (*e2sm_kpm_v2.MeasurementTypeId, error) {

	measurementTypeID := e2sm_kpm_v2.MeasurementTypeId{
		Value: int32(*measurementTypeIDC),
	}

	return &measurementTypeID, nil
}

func decodeMeasurementTypeIDBytes(array [8]byte) (*e2sm_kpm_v2.MeasurementTypeId, error) {
	measurementTypeIDC := C.long(binary.LittleEndian.Uint64(array[0:]))
	return decodeMeasurementTypeID(&measurementTypeIDC)
}
