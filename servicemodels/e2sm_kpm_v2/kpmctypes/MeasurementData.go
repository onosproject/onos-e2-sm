// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementData.h"
//#include "MeasurementRecord.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeMeasurementData(measurementData *e2sm_kpm_v2.MeasurementData) ([]byte, error) {
	measurementDataCP, err := newMeasurementData(measurementData)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementData() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementData, unsafe.Pointer(measurementDataCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementData() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementData(measurementData *e2sm_kpm_v2.MeasurementData) ([]byte, error) {
	measurementDataCP, err := newMeasurementData(measurementData)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementData() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementData, unsafe.Pointer(measurementDataCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementData() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementData(bytes []byte) (*e2sm_kpm_v2.MeasurementData, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementData)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementData((*C.MeasurementData_t)(unsafePtr))
}

func perDecodeMeasurementData(bytes []byte) (*e2sm_kpm_v2.MeasurementData, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementData)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementData((*C.MeasurementData_t)(unsafePtr))
}

func newMeasurementData(measurementData *e2sm_kpm_v2.MeasurementData) (*C.MeasurementData_t, error) {

	measurementDataC := new(C.MeasurementData_t)
	for _, item := range measurementData.GetValue() {
		itemC, err := newMeasurementRecord(item)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementRecord() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementDataC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return measurementDataC, nil
}

func decodeMeasurementData(measurementDataC *C.MeasurementData_t) (*e2sm_kpm_v2.MeasurementData, error) {

	measurementData := new(e2sm_kpm_v2.MeasurementData)
	var ieCount = int(measurementDataC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementDataC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementRecord_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementDataC.list.array)) + offset))
		ie, err := decodeMeasurementRecord(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementRecord() %s", err.Error())
		}
		measurementData.Value = append(measurementData.Value, ie)
	}

	return measurementData, nil
}
