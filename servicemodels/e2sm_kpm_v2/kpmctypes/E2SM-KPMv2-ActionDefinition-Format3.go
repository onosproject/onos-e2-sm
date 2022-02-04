// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-ActionDefinition-Format3.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3) ([]byte, error) {
	e2SmKpmActionDefinitionFormat3CP, err := newE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat3() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format3, unsafe.Pointer(e2SmKpmActionDefinitionFormat3CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat3() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3) ([]byte, error) {
	e2SmKpmActionDefinitionFormat3CP, err := newE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat3() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format3, unsafe.Pointer(e2SmKpmActionDefinitionFormat3CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat3() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmActionDefinitionFormat3(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat3((*C.E2SM_KPMv2_ActionDefinition_Format3_t)(unsafePtr))
}

func perDecodeE2SmKpmActionDefinitionFormat3(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_ActionDefinition_Format3)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat3((*C.E2SM_KPMv2_ActionDefinition_Format3_t)(unsafePtr))
}

func newE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3) (*C.E2SM_KPMv2_ActionDefinition_Format3_t, error) {

	cellObjIDC, err := newCellObjectID(e2SmKpmActionDefinitionFormat3.CellObjId)
	if err != nil {
		return nil, fmt.Errorf("newCellObjectId() %s", err.Error())
	}

	measCondListC, err := newMeasurementCondList(e2SmKpmActionDefinitionFormat3.MeasCondList)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementCondList() %s", err.Error())
	}

	granulPeriodC, err := newGranularityPeriod(e2SmKpmActionDefinitionFormat3.GranulPeriod)
	if err != nil {
		return nil, fmt.Errorf("newGranularityPeriod() %s", err.Error())
	}

	subscriptIDC, err := newSubscriptionID(e2SmKpmActionDefinitionFormat3.SubscriptId)
	if err != nil {
		return nil, fmt.Errorf("newSubscriptionId() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat3C := C.E2SM_KPMv2_ActionDefinition_Format3_t{
		//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
		cellObjID:    *cellObjIDC,
		measCondList: *measCondListC,
		granulPeriod: *granulPeriodC,
		subscriptID:  *subscriptIDC,
	}

	return &e2SmKpmActionDefinitionFormat3C, nil
}

func decodeE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3C *C.E2SM_KPMv2_ActionDefinition_Format3_t) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3, error) {

	cellObjID, err := decodeCellObjectID(&e2SmKpmActionDefinitionFormat3C.cellObjID)
	if err != nil {
		return nil, fmt.Errorf("decodeCellObjectId() %s", err.Error())
	}

	measCondList, err := decodeMeasurementCondList(&e2SmKpmActionDefinitionFormat3C.measCondList)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementCondList() %s", err.Error())
	}

	granulPeriod, err := decodeGranularityPeriod(&e2SmKpmActionDefinitionFormat3C.granulPeriod)
	if err != nil {
		return nil, fmt.Errorf("decodeGranularityPeriod() %s", err.Error())
	}

	subscriptID, err := decodeSubscriptionID(&e2SmKpmActionDefinitionFormat3C.subscriptID)
	if err != nil {
		return nil, fmt.Errorf("decodeSubscriptionId() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat3 := e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		CellObjId:    cellObjID,
		MeasCondList: measCondList,
		GranulPeriod: granulPeriod,
		SubscriptId:  subscriptID,
	}

	return &e2SmKpmActionDefinitionFormat3, nil
}

func decodeE2SmKpmActionDefinitionFormat3Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3, error) {
	e2SmKpmActionDefinitionFormat3C := (*C.E2SM_KPMv2_ActionDefinition_Format3_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmActionDefinitionFormat3(e2SmKpmActionDefinitionFormat3C)
}
