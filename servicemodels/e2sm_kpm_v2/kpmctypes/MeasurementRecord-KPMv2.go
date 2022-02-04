// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementRecord-KPMv2.h"
//#include "MeasurementRecordItem-KPMv2.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementRecord(measurementRecord *e2sm_kpm_v2.MeasurementRecord) ([]byte, error) {
	measurementRecordCP, err := newMeasurementRecord(measurementRecord)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementRecord() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementRecord_KPMv2, unsafe.Pointer(measurementRecordCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementRecord() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementRecord(measurementRecord *e2sm_kpm_v2.MeasurementRecord) ([]byte, error) {
	measurementRecordCP, err := newMeasurementRecord(measurementRecord)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementRecord() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementRecord_KPMv2, unsafe.Pointer(measurementRecordCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementRecord() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementRecord(bytes []byte) (*e2sm_kpm_v2.MeasurementRecord, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementRecord_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementRecord((*C.MeasurementRecord_KPMv2_t)(unsafePtr))
}

func perDecodeMeasurementRecord(bytes []byte) (*e2sm_kpm_v2.MeasurementRecord, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementRecord_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementRecord((*C.MeasurementRecord_KPMv2_t)(unsafePtr))
}

func newMeasurementRecord(measurementRecord *e2sm_kpm_v2.MeasurementRecord) (*C.MeasurementRecord_KPMv2_t, error) {

	measurementRecordC := new(C.MeasurementRecord_KPMv2_t)
	for _, item := range measurementRecord.GetValue() {
		itemC, err := newMeasurementRecordItem(item)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementRecordItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementRecordC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return measurementRecordC, nil
}

func decodeMeasurementRecord(measurementRecordC *C.MeasurementRecord_KPMv2_t) (*e2sm_kpm_v2.MeasurementRecord, error) {

	measurementRecord := new(e2sm_kpm_v2.MeasurementRecord)

	ieCount := int(measurementRecordC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementRecordC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementRecordItem_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementRecordC.list.array)) + offset))
		ie, err := decodeMeasurementRecordItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementRecordItem() %s", err.Error())
		}
		measurementRecord.Value = append(measurementRecord.Value, ie)
	}

	return measurementRecord, nil
}
