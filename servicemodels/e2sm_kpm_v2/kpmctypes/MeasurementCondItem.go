// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementCondItem.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementCondItem(measurementCondItem *e2sm_kpm_v2.MeasurementCondItem) ([]byte, error) {
	measurementCondItemCP, err := newMeasurementCondItem(measurementCondItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementCondItem, unsafe.Pointer(measurementCondItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementCondItem(measurementCondItem *e2sm_kpm_v2.MeasurementCondItem) ([]byte, error) {
	measurementCondItemCP, err := newMeasurementCondItem(measurementCondItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementCondItem, unsafe.Pointer(measurementCondItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementCondItem(bytes []byte) (*e2sm_kpm_v2.MeasurementCondItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementCondItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementCondItem((*C.MeasurementCondItem_t)(unsafePtr))
}

func perDecodeMeasurementCondItem(bytes []byte) (*e2sm_kpm_v2.MeasurementCondItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementCondItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementCondItem((*C.MeasurementCondItem_t)(unsafePtr))
}

func newMeasurementCondItem(measurementCondItem *e2sm_kpm_v2.MeasurementCondItem) (*C.MeasurementCondItem_t, error) {

	measTypeC, err := newMeasurementType(measurementCondItem.MeasType)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementType() %s", err.Error())
	}

	matchingCondC, err := newMatchingCondList(measurementCondItem.MatchingCond)
	if err != nil {
		return nil, fmt.Errorf("newMatchingCondList() %s", err.Error())
	}

	measurementCondItemC := C.MeasurementCondItem_t{
		measType:     *measTypeC,
		matchingCond: *matchingCondC,
	}

	return &measurementCondItemC, nil
}

func decodeMeasurementCondItem(measurementCondItemC *C.MeasurementCondItem_t) (*e2sm_kpm_v2.MeasurementCondItem, error) {

	measType, err := decodeMeasurementType(&measurementCondItemC.measType)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementType() %s", err.Error())
	}

	matchingCond, err := decodeMatchingCondList(&measurementCondItemC.matchingCond)
	if err != nil {
		return nil, fmt.Errorf("decodeMatchingCondList() %s", err.Error())
	}

	measurementCondItem := e2sm_kpm_v2.MeasurementCondItem{
		MeasType:     measType,
		MatchingCond: matchingCond,
	}

	return &measurementCondItem, nil
}
