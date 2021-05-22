// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-gNB-ID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalKpmnodeGnbID(globalKpmnodeGnbID *e2sm_kpm_v2.GlobalKpmnodeGnbId) ([]byte, error) {
	globalKpmnodeGnbIDCP, err := newGlobalKpmnodeGnbID(globalKpmnodeGnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeGnNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_gNB_ID_KPMv2, unsafe.Pointer(globalKpmnodeGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeGnNbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalKpmnodeGnbID(globalKpmnodeGnbID *e2sm_kpm_v2.GlobalKpmnodeGnbId) ([]byte, error) {
	globalKpmnodeGnbIDCP, err := newGlobalKpmnodeGnbID(globalKpmnodeGnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeGnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalKPMnode_gNB_ID_KPMv2, unsafe.Pointer(globalKpmnodeGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeGnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKpmnodeGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_gNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalKpmnodeGnbID((*C.GlobalKPMnode_gNB_ID_KPMv2_t)(unsafePtr))
}

func perDecodeGlobalKpmnodeGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalKPMnode_gNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalKpmnodeGnbID((*C.GlobalKPMnode_gNB_ID_KPMv2_t)(unsafePtr))
}

func newGlobalKpmnodeGnbID(globalKpmnodeGnbID *e2sm_kpm_v2.GlobalKpmnodeGnbId) (*C.GlobalKPMnode_gNB_ID_KPMv2_t, error) {

	globalGnbIDC, err := newGlobalgNbID(globalKpmnodeGnbID.GlobalGNbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalgNbID() %s", err.Error())
	}

	globalKpmnodeGnbIDC := C.GlobalKPMnode_gNB_ID_KPMv2_t{
		global_gNB_ID: *globalGnbIDC,
		//gNB_CU_UP_ID:  gNbCuUpIDC,
		//gNB_DU_ID:     gNbDuIDC,
	}

	if globalKpmnodeGnbID.GNbCuUpId != nil {
		globalKpmnodeGnbIDC.gNB_CU_UP_ID, err = newGnbCuUpID(globalKpmnodeGnbID.GNbCuUpId)
		if err != nil {
			return nil, fmt.Errorf("newGnbCuUpID() %s", err.Error())
		}
	}

	if globalKpmnodeGnbID.GNbDuId != nil {
		globalKpmnodeGnbIDC.gNB_DU_ID, err = newGnbDuID(globalKpmnodeGnbID.GNbDuId)
		if err != nil {
			return nil, fmt.Errorf("newGnbDuID() %s", err.Error())
		}
	}

	return &globalKpmnodeGnbIDC, nil
}

func decodeGlobalKpmnodeGnbID(globalKpmnodeGnbIDC *C.GlobalKPMnode_gNB_ID_KPMv2_t) (*e2sm_kpm_v2.GlobalKpmnodeGnbId, error) {

	globalGnbID, err := decodeGlobalgNbID(&globalKpmnodeGnbIDC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalgNbID() %s", err.Error())
	}

	globalKpmnodeNgEnbID := e2sm_kpm_v2.GlobalKpmnodeGnbId{
		GlobalGNbId: globalGnbID,
		//GNbCuUpId:   gNbCuUpID,
		//GNbDuId:     gNbDuID,
	}

	if globalKpmnodeGnbIDC.gNB_CU_UP_ID != nil {
		globalKpmnodeNgEnbID.GNbCuUpId, err = decodeGnbCuUpID(globalKpmnodeGnbIDC.gNB_CU_UP_ID)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbCuUpID() %s", err.Error())
		}
	}

	if globalKpmnodeGnbIDC.gNB_DU_ID != nil {
		globalKpmnodeNgEnbID.GNbDuId, err = decodeGnbDuID(globalKpmnodeGnbIDC.gNB_DU_ID)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbDuID() %s", err.Error())
		}
	}
	return &globalKpmnodeNgEnbID, nil
}

func decodeGlobalKpmnodeGnbIDBytes(array [8]byte) (*e2sm_kpm_v2.GlobalKpmnodeGnbId, error) {
	globalKpmnodeGnbIDC := (*C.GlobalKPMnode_gNB_ID_KPMv2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmnodeGnbID(globalKpmnodeGnbIDC)
}
