// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-IndicationMessage-Format2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2) ([]byte, error) {
	e2SmKpmIndicationMessageFormat2CP, err := newE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_IndicationMessage_Format2, unsafe.Pointer(e2SmKpmIndicationMessageFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2) ([]byte, error) {
	e2SmKpmIndicationMessageFormat2CP, err := newE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessageFormat2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_IndicationMessage_Format2, unsafe.Pointer(e2SmKpmIndicationMessageFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessageFormat2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmIndicationMessageFormat2(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_IndicationMessage_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmIndicationMessageFormat2((*C.E2SM_KPMv2_IndicationMessage_Format2_t)(unsafePtr))
}

func perDecodeE2SmKpmIndicationMessageFormat2(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_IndicationMessage_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationMessageFormat2((*C.E2SM_KPMv2_IndicationMessage_Format2_t)(unsafePtr))
}

func newE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2) (*C.E2SM_KPMv2_IndicationMessage_Format2_t, error) {

	subscriptIDC, err := newSubscriptionID(e2SmKpmIndicationMessageFormat2.SubscriptId)
	if err != nil {
		return nil, fmt.Errorf("newSubscriptionId() %s", err.Error())
	}

	measCondUeIDListC, err := newMeasurementCondUeIDList(e2SmKpmIndicationMessageFormat2.MeasCondUeidList)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementCondUeidList() %s", err.Error())
	}

	measDataC, err := newMeasurementData(e2SmKpmIndicationMessageFormat2.MeasData)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementData() %s", err.Error())
	}

	e2SmKpmIndicationMessageFormat2C := C.E2SM_KPMv2_IndicationMessage_Format2_t{
		subscriptID: *subscriptIDC,
		//cellObjID:        cellObjIDC,
		//granulPeriod:     granulPeriodC,
		measCondUEidList: *measCondUeIDListC,
		measData:         *measDataC,
	}

	if e2SmKpmIndicationMessageFormat2.CellObjId != nil {
		e2SmKpmIndicationMessageFormat2C.cellObjID, err = newCellObjectID(e2SmKpmIndicationMessageFormat2.CellObjId)
		if err != nil {
			return nil, fmt.Errorf("newCellObjectId() %s", err.Error())
		}
	}

	if e2SmKpmIndicationMessageFormat2.GranulPeriod != nil {
		e2SmKpmIndicationMessageFormat2C.granulPeriod, err = newGranularityPeriod(e2SmKpmIndicationMessageFormat2.GranulPeriod)
		if err != nil {
			return nil, fmt.Errorf("newGranularityPeriod() %s", err.Error())
		}
	}

	return &e2SmKpmIndicationMessageFormat2C, nil
}

func decodeE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2C *C.E2SM_KPMv2_IndicationMessage_Format2_t) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2, error) {

	subscriptID, err := decodeSubscriptionID(&e2SmKpmIndicationMessageFormat2C.subscriptID)
	if err != nil {
		return nil, fmt.Errorf("decodeSubscriptionId() %s", err.Error())
	}

	measCondUeIDList, err := decodeMeasurementCondUeIDList(&e2SmKpmIndicationMessageFormat2C.measCondUEidList)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementCondUeidList() %s", err.Error())
	}

	measData, err := decodeMeasurementData(&e2SmKpmIndicationMessageFormat2C.measData)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementData() %s", err.Error())
	}

	e2SmKpmIndicationMessageFormat2 := e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2{
		SubscriptId:      subscriptID,
		MeasCondUeidList: measCondUeIDList,
		MeasData:         measData,
	}

	if e2SmKpmIndicationMessageFormat2C.cellObjID != nil {
		e2SmKpmIndicationMessageFormat2.CellObjId, err = decodeCellObjectID(e2SmKpmIndicationMessageFormat2C.cellObjID)
		if err != nil {
			return nil, fmt.Errorf("decodeCellObjectId() %s", err.Error())
		}
	}

	if e2SmKpmIndicationMessageFormat2C.granulPeriod != nil {
		e2SmKpmIndicationMessageFormat2.GranulPeriod, err = decodeGranularityPeriod(e2SmKpmIndicationMessageFormat2C.granulPeriod)
		if err != nil {
			return nil, fmt.Errorf("decodeGranularityPeriod() %s", err.Error())
		}
	}

	return &e2SmKpmIndicationMessageFormat2, nil
}

func decodeE2SmKpmIndicationMessageFormat2Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2, error) {
	e2SmKpmIndicationMessageFormat2C := (*C.E2SM_KPMv2_IndicationMessage_Format2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmIndicationMessageFormat2(e2SmKpmIndicationMessageFormat2C)
}
