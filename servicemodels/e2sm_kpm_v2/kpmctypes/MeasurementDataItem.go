// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementDataItem.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementDataItem(measurementDataItem *e2sm_kpm_v2.MeasurementDataItem) ([]byte, error) {
	measurementDataItemCP, err := newMeasurementDataItem(measurementDataItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementDataItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementDataItem, unsafe.Pointer(measurementDataItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementDataItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementDataItem(measurementDataItem *e2sm_kpm_v2.MeasurementDataItem) ([]byte, error) {
	measurementDataItemCP, err := newMeasurementDataItem(measurementDataItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementDataItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementDataItem, unsafe.Pointer(measurementDataItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementDataItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementDataItem(bytes []byte) (*e2sm_kpm_v2.MeasurementDataItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementDataItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementDataItem((*C.MeasurementDataItem_t)(unsafePtr))
}

func perDecodeMeasurementDataItem(bytes []byte) (*e2sm_kpm_v2.MeasurementDataItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementDataItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementDataItem((*C.MeasurementDataItem_t)(unsafePtr))
}

func newMeasurementDataItem(measurementDataItem *e2sm_kpm_v2.MeasurementDataItem) (*C.MeasurementDataItem_t, error) {

	measRecordC, err := newMeasurementRecord(measurementDataItem.MeasRecord)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementRecord() %s", err.Error())
	}

	var incompleteFlagC C.MeasurementDataItem__incompleteFlag_t
	switch measurementDataItem.IncompleteFlag {
	case e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE:
		incompleteFlagC = C.MeasurementDataItem__incompleteFlag_true
	default:
		return nil, fmt.Errorf("unexpected MeasurementDataItem IncompleteFlag %v", measurementDataItem.IncompleteFlag)
	}

	measurementDataItemC := C.MeasurementDataItem_t{
		measRecord:     *measRecordC,
		incompleteFlag: &incompleteFlagC,
	}

	return &measurementDataItemC, nil
}

func decodeMeasurementDataItem(measurementDataItemC *C.MeasurementDataItem_t) (*e2sm_kpm_v2.MeasurementDataItem, error) {

	measRecord, err := decodeMeasurementRecord(&measurementDataItemC.measRecord)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementRecord() %s", err.Error())
	}

	measurementDataItem := e2sm_kpm_v2.MeasurementDataItem{
		MeasRecord:     measRecord,
		IncompleteFlag: e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE,
	}

	return &measurementDataItem, nil
}
