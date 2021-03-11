// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationMessage-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1) ([]byte, error) {
	e2SmKpmIndicationMessageFormat1CP, err := newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationMessage_Format1, unsafe.Pointer(e2SmKpmIndicationMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1) ([]byte, error) {
	e2SmKpmIndicationMessageFormat1CP, err := newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_IndicationMessage_Format1, unsafe.Pointer(e2SmKpmIndicationMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmIndicationMessageFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_IndicationMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmIndicationMessageFormat1((*C.E2SM_KPM_IndicationMessage_Format1_t)(unsafePtr))
}

func perDecodeE2SmKpmIndicationMessageFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_IndicationMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationMessageFormat1((*C.E2SM_KPM_IndicationMessage_Format1_t)(unsafePtr))
}

func newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1 *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1) (*C.E2SM_KPM_IndicationMessage_Format1_t, error) {

	subscriptIDC, err := newSubscriptionID(e2SmKpmIndicationMessageFormat1.SubscriptId)
	if err != nil {
		return nil, fmt.Errorf("newSubscriptionID() %s", err.Error())
	}

	cellObjIDC, err := newCellObjectID(e2SmKpmIndicationMessageFormat1.CellObjId)
	if err != nil {
		return nil, fmt.Errorf("newCellObjectId() %s", err.Error())
	}

	granulPeriodC, err := newGranularityPeriod(e2SmKpmIndicationMessageFormat1.GranulPeriod)
	if err != nil {
		return nil, fmt.Errorf("newGranularityPeriod() %s", err.Error())
	}

	//measInfoListC, err := newMeasurementInfoList(e2SmKpmIndicationMessageFormat1.MeasInfoList)
	//if err != nil {
	//	return nil, fmt.Errorf("newMeasurementInfoList() %s", err.Error())
	//}

	//measDataC, err := newMeasurementData(e2SmKpmIndicationMessageFormat1.MeasData)
	//if err != nil {
	//	return nil, fmt.Errorf("newMeasurementData() %s", err.Error())
	//}

	e2SmKpmIndicationMessageFormat1C := C.E2SM_KPM_IndicationMessage_Format1_t{
		subscriptID:  *subscriptIDC,
		cellObjID:    cellObjIDC,
		granulPeriod: granulPeriodC,
		//measInfoList: measInfoListC,
		//measData:     *measDataC,
	}

	return &e2SmKpmIndicationMessageFormat1C, nil
}

func decodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1C *C.E2SM_KPM_IndicationMessage_Format1_t) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1, error) {

	subscriptID, err := decodeSubscriptionID(&e2SmKpmIndicationMessageFormat1C.subscriptID)
	if err != nil {
		return nil, fmt.Errorf("decodeSubscriptionId() %s", err.Error())
	}

	cellObjID, err := decodeCellObjectID(e2SmKpmIndicationMessageFormat1C.cellObjID)
	if err != nil {
		return nil, fmt.Errorf("decodeCellObjectId() %s", err.Error())
	}

	granulPeriod, err := decodeGranularityPeriod(e2SmKpmIndicationMessageFormat1C.granulPeriod)
	if err != nil {
		return nil, fmt.Errorf("decodeGranularityPeriod() %s", err.Error())
	}

	//measInfoList, err := decodeMeasurementInfoList(e2SmKpmIndicationMessageFormat1C.measInfoList)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeMeasurementInfoList() %s", err.Error())
	//}

	//measData, err := decodeMeasurementData(&e2SmKpmIndicationMessageFormat1C.measData)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeMeasurementData() %s", err.Error())
	//}

	e2SmKpmIndicationMessageFormat1 := e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1{
		SubscriptId:  subscriptID,
		CellObjId:    cellObjID,
		GranulPeriod: granulPeriod,
		//MeasInfoList: measInfoList,
		//MeasData:     measData,
	}

	return &e2SmKpmIndicationMessageFormat1, nil
}

func decodeE2SmKpmIndicationMessageFormat1Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1, error) {
	e2SmKpmIndicationMessageFormat1C := (*C.E2SM_KPM_IndicationMessage_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1C)
}
