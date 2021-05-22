// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementInfo-Action-List-KPMv2.h"
//#include "MeasurementInfo-Action-Item-KPMv2.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementInfoActionList(measurementInfoActionList *e2sm_kpm_v2.MeasurementInfoActionList) ([]byte, error) {
	measurementInfoActionListCP, err := newMeasurementInfoActionList(measurementInfoActionList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoActionList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementInfo_Action_List_KPMv2, unsafe.Pointer(measurementInfoActionListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoActionList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementInfoActionList(measurementInfoActionList *e2sm_kpm_v2.MeasurementInfoActionList) ([]byte, error) {
	measurementInfoActionListCP, err := newMeasurementInfoActionList(measurementInfoActionList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoActionList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementInfo_Action_List_KPMv2, unsafe.Pointer(measurementInfoActionListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoActionList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementInfoActionList(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoActionList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementInfo_Action_List_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementInfoActionList((*C.MeasurementInfo_Action_List_KPMv2_t)(unsafePtr))
}

func perDecodeMeasurementInfoActionList(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoActionList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementInfo_Action_List_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementInfoActionList((*C.MeasurementInfo_Action_List_KPMv2_t)(unsafePtr))
}

func newMeasurementInfoActionList(measurementInfoActionList *e2sm_kpm_v2.MeasurementInfoActionList) (*C.MeasurementInfo_Action_List_KPMv2_t, error) {

	measurementInfoActionListC := new(C.MeasurementInfo_Action_List_KPMv2_t)
	for _, item := range measurementInfoActionList.GetValue() {
		itemC, err := newMeasurementInfoActionItem(item)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementInfoActionItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(measurementInfoActionListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return measurementInfoActionListC, nil
}

func decodeMeasurementInfoActionList(measurementInfoActionListC *C.MeasurementInfo_Action_List_KPMv2_t) (*e2sm_kpm_v2.MeasurementInfoActionList, error) {

	measurementInfoActionList := new(e2sm_kpm_v2.MeasurementInfoActionList)

	var ieCount = int(measurementInfoActionListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(measurementInfoActionListC.list.array)) * uintptr(i)
		ieC := *(**C.MeasurementInfo_Action_Item_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(measurementInfoActionListC.list.array)) + offset))
		ie, err := decodeMeasurementInfoActionItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementInfoActionItem() %s", err.Error())
		}
		measurementInfoActionList.Value = append(measurementInfoActionList.Value, ie)
	}

	return measurementInfoActionList, nil
}
