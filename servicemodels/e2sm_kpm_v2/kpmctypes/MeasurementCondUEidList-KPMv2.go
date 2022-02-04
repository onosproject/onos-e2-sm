// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementCondUEidList-KPMv2.h"
//#include "MeasurementCondUEidItem-KPMv2.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementCondUeIDList(measurementCondUeIDList *e2sm_kpm_v2.MeasurementCondUeidList) ([]byte, error) {
	measurementCondUeIDListCP, err := newMeasurementCondUeIDList(measurementCondUeIDList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondUeIDList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementCondUEidList_KPMv2, unsafe.Pointer(measurementCondUeIDListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondUeIDList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementCondUeIDList(measurementCondUeIDList *e2sm_kpm_v2.MeasurementCondUeidList) ([]byte, error) {
	measurementCondUeIDListCP, err := newMeasurementCondUeIDList(measurementCondUeIDList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondUeIDList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementCondUEidList_KPMv2, unsafe.Pointer(measurementCondUeIDListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondUeIDList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementCondUeIDList(bytes []byte) (*e2sm_kpm_v2.MeasurementCondUeidList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementCondUEidList_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementCondUeIDList((*C.MeasurementCondUEidList_KPMv2_t)(unsafePtr))
}

func perDecodeMeasurementCondUeIDList(bytes []byte) (*e2sm_kpm_v2.MeasurementCondUeidList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementCondUEidList_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementCondUeIDList((*C.MeasurementCondUEidList_KPMv2_t)(unsafePtr))
}

func newMeasurementCondUeIDList(measurementCondUeIDList *e2sm_kpm_v2.MeasurementCondUeidList) (*C.MeasurementCondUEidList_KPMv2_t, error) {

	measurementCondUeIDListC := new(C.MeasurementCondUEidList_KPMv2_t)
	for _, item := range measurementCondUeIDList.GetValue() {
		itemC, err := newMeasurementCondUeIDItem(item)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementCondUeidItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementCondUeIDListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return measurementCondUeIDListC, nil
}

func decodeMeasurementCondUeIDList(measurementCondUeIDListC *C.MeasurementCondUEidList_KPMv2_t) (*e2sm_kpm_v2.MeasurementCondUeidList, error) {

	measurementCondUeIDList := new(e2sm_kpm_v2.MeasurementCondUeidList)

	ieCount := int(measurementCondUeIDListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementCondUeIDListC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementCondUEidItem_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementCondUeIDListC.list.array)) + offset))
		ie, err := decodeMeasurementCondUeIDItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementCondUeIDItem() %s", err.Error())
		}
		measurementCondUeIDList.Value = append(measurementCondUeIDList.Value, ie)
	}

	return measurementCondUeIDList, nil
}
