// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementCondList.h"
//#include "MeasurementCondItem.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementCondList(measurementCondList *e2sm_kpm_v2.MeasurementCondList) ([]byte, error) {
	measurementCondListCP, err := newMeasurementCondList(measurementCondList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementCondList, unsafe.Pointer(measurementCondListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementCondList(measurementCondList *e2sm_kpm_v2.MeasurementCondList) ([]byte, error) {
	measurementCondListCP, err := newMeasurementCondList(measurementCondList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementCondList, unsafe.Pointer(measurementCondListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementCondList(bytes []byte) (*e2sm_kpm_v2.MeasurementCondList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementCondList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementCondList((*C.MeasurementCondList_t)(unsafePtr))
}

func perDecodeMeasurementCondList(bytes []byte) (*e2sm_kpm_v2.MeasurementCondList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementCondList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementCondList((*C.MeasurementCondList_t)(unsafePtr))
}

func newMeasurementCondList(measurementCondList *e2sm_kpm_v2.MeasurementCondList) (*C.MeasurementCondList_t, error) {

	measurementCondListC := new(C.MeasurementCondList_t)
	for _, valueItem := range measurementCondList.GetValue() {
		itemC, err := newMeasurementCondItem(valueItem)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementCondItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementCondListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return measurementCondListC, nil
}

func decodeMeasurementCondList(measurementCondListC *C.MeasurementCondList_t) (*e2sm_kpm_v2.MeasurementCondList, error) {

	measurementCondList := new(e2sm_kpm_v2.MeasurementCondList)

	ieCount := int(measurementCondListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementCondListC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementCondItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementCondListC.list.array)) + offset))
		ie, err := decodeMeasurementCondItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementCondItem() %s", err.Error())
		}
		measurementCondList.Value = append(measurementCondList.Value, ie)
	}

	return measurementCondList, nil
}
