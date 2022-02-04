// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-ActionDefinition-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1) ([]byte, error) {
	e2SmKpmActionDefinitionFormat1CP, err := newE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format1, unsafe.Pointer(e2SmKpmActionDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1) ([]byte, error) {
	e2SmKpmActionDefinitionFormat1CP, err := newE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format1, unsafe.Pointer(e2SmKpmActionDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmActionDefinitionFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat1((*C.E2SM_KPMv2_ActionDefinition_Format1_t)(unsafePtr))
}

func perDecodeE2SmKpmActionDefinitionFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat1((*C.E2SM_KPMv2_ActionDefinition_Format1_t)(unsafePtr))
}

func newE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1) (*C.E2SM_KPMv2_ActionDefinition_Format1_t, error) {

	cellObjIDC, err := newCellObjectID(e2SmKpmActionDefinitionFormat1.CellObjId)
	if err != nil {
		return nil, fmt.Errorf("newCellObjectID() %s", err.Error())
	}

	measInfoListC, err := newMeasurementInfoList(e2SmKpmActionDefinitionFormat1.MeasInfoList)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementInfoList() %s", err.Error())
	}

	granulPeriodC, err := newGranularityPeriod(e2SmKpmActionDefinitionFormat1.GranulPeriod)
	if err != nil {
		return nil, fmt.Errorf("newGranularityPeriod() %s", err.Error())
	}

	subscriptIDC, err := newSubscriptionID(e2SmKpmActionDefinitionFormat1.SubscriptId)
	if err != nil {
		return nil, fmt.Errorf("newSubscriptionID() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat1C := C.E2SM_KPMv2_ActionDefinition_Format1_t{
		cellObjID:    *cellObjIDC,
		measInfoList: *measInfoListC,
		granulPeriod: *granulPeriodC,
		subscriptID:  *subscriptIDC,
	}

	return &e2SmKpmActionDefinitionFormat1C, nil
}

func decodeE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1C *C.E2SM_KPMv2_ActionDefinition_Format1_t) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1, error) {

	cellObjID, err := decodeCellObjectID(&e2SmKpmActionDefinitionFormat1C.cellObjID)
	if err != nil {
		return nil, fmt.Errorf("decodeCellObjectId() %s", err.Error())
	}

	measInfoList, err := decodeMeasurementInfoList(&e2SmKpmActionDefinitionFormat1C.measInfoList)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementInfoList() %s", err.Error())
	}

	granulPeriod, err := decodeGranularityPeriod(&e2SmKpmActionDefinitionFormat1C.granulPeriod)
	if err != nil {
		return nil, fmt.Errorf("decodeGranularityPeriod() %s", err.Error())
	}

	subscriptID, err := decodeSubscriptionID(&e2SmKpmActionDefinitionFormat1C.subscriptID)
	if err != nil {
		return nil, fmt.Errorf("decodeSubscriptionId() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat1 := e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1{
		CellObjId:    cellObjID,
		MeasInfoList: measInfoList,
		GranulPeriod: granulPeriod,
		SubscriptId:  subscriptID,
	}

	return &e2SmKpmActionDefinitionFormat1, nil
}

func decodeE2SmKpmActionDefinitionFormat1Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1, error) {
	e2SmKpmActionDefinitionFormat1C := (*C.E2SM_KPMv2_ActionDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat1C)
}
