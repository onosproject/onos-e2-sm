// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementInfoList.h"
//#include "MeasurementInfoItem.h"
//#include "Cell-Measurement-Object-Item.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementInfoList(measurementInfoList *e2sm_kpm_v2.MeasurementInfoList) ([]byte, error) {
	measurementInfoListCP, err := newMeasurementInfoList(measurementInfoList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementInfoList, unsafe.Pointer(measurementInfoListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementInfoList(measurementInfoList *e2sm_kpm_v2.MeasurementInfoList) ([]byte, error) {
	measurementInfoListCP, err := newMeasurementInfoList(measurementInfoList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementInfoList, unsafe.Pointer(measurementInfoListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementInfoList(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementInfoList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementInfoList((*C.MeasurementInfoList_t)(unsafePtr))
}

func perDecodeMeasurementInfoList(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementInfoList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementInfoList((*C.MeasurementInfoList_t)(unsafePtr))
}

func newMeasurementInfoList(measurementInfoList *e2sm_kpm_v2.MeasurementInfoList) (*C.MeasurementInfoList_t, error) {

	measurementInfoListC := new(C.MeasurementInfoList_t)
	for _, item := range measurementInfoList.GetValue() {
		measurementInfoItemC, err := newMeasurementInfoItem(item)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementInfoItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementInfoListC), unsafe.Pointer(measurementInfoItemC)); err != nil {
			return nil, err
		}
	}

	return measurementInfoListC, nil
}

func decodeMeasurementInfoList(measurementInfoListC *C.MeasurementInfoList_t) (*e2sm_kpm_v2.MeasurementInfoList, error) {

	measurementInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}

	ieCount := int(measurementInfoListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementInfoListC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementInfoItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementInfoListC.list.array)) + offset))
		ie, err := decodeMeasurementInfoItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementInfoItem() %s", err.Error())
		}
		measurementInfoList.Value = append(measurementInfoList.Value, ie)
	}

	return &measurementInfoList, nil
}

