// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementCondUEidItem.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementCondUeIDItem(measurementCondUeIDItem *e2sm_kpm_v2.MeasurementCondUeidItem) ([]byte, error) {
	measurementCondUeIDItemCP, err := newMeasurementCondUeIDItem(measurementCondUeIDItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondUeIDItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementCondUEidItem, unsafe.Pointer(measurementCondUeIDItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementCondUeIDItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementCondUeIDItem(measurementCondUeIDItem *e2sm_kpm_v2.MeasurementCondUeidItem) ([]byte, error) {
	measurementCondUeIDItemCP, err := newMeasurementCondUeIDItem(measurementCondUeIDItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondUeIDItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementCondUEidItem, unsafe.Pointer(measurementCondUeIDItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementCondUeIDItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementCondUeIDItem(bytes []byte) (*e2sm_kpm_v2.MeasurementCondUeidItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementCondUEidItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementCondUeIDItem((*C.MeasurementCondUEidItem_t)(unsafePtr))
}

func perDecodeMeasurementCondUeIDItem(bytes []byte) (*e2sm_kpm_v2.MeasurementCondUeidItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementCondUEidItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementCondUeIDItem((*C.MeasurementCondUEidItem_t)(unsafePtr))
}

func newMeasurementCondUeIDItem(measurementCondUeIDItem *e2sm_kpm_v2.MeasurementCondUeidItem) (*C.MeasurementCondUEidItem_t, error) {

	measTypeC, err := newMeasurementType(measurementCondUeIDItem.MeasType)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementType() %s", err.Error())
	}

	matchingCondC, err := newMatchingCondList(measurementCondUeIDItem.MatchingCond)
	if err != nil {
		return nil, fmt.Errorf("newMatchingCondList() %s", err.Error())
	}

	matchingUeIDListC, err := newMatchingUeIDList(measurementCondUeIDItem.MatchingUeidList)
	if err != nil {
		return nil, fmt.Errorf("newMatchingUeIDList() %s", err.Error())
	}

	measurementCondUeIDItemC := C.MeasurementCondUEidItem_t{
		measType:         *measTypeC,
		matchingCond:     *matchingCondC,
		matchingUEidList: matchingUeIDListC,
	}

	return &measurementCondUeIDItemC, nil
}

func decodeMeasurementCondUeIDItem(measurementCondUeIDItemC *C.MeasurementCondUEidItem_t) (*e2sm_kpm_v2.MeasurementCondUeidItem, error) {

	measType, err := decodeMeasurementType(&measurementCondUeIDItemC.measType)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementType() %s", err.Error())
	}

	matchingCond, err := decodeMatchingCondList(&measurementCondUeIDItemC.matchingCond)
	if err != nil {
		return nil, fmt.Errorf("decodeMatchingCondList() %s", err.Error())
	}

	matchingUeIDList, err := decodeMatchingUeIDList(measurementCondUeIDItemC.matchingUEidList)
	if err != nil {
		return nil, fmt.Errorf("decodeMatchingUeidList() %s", err.Error())
	}

	measurementCondUeIDItem := e2sm_kpm_v2.MeasurementCondUeidItem{
		MeasType:         measType,
		MatchingCond:     matchingCond,
		MatchingUeidList: matchingUeIDList,
	}

	return &measurementCondUeIDItem, nil
}
