// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMv2node-en-gNB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbID *e2sm_kpm_v2.GlobalKpmnodeEnGnbId) ([]byte, error) {
	globalKpmnodeEnGnbIDCP, err := newGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeEnGnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMv2node_en_gNB_ID, unsafe.Pointer(globalKpmnodeEnGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeEnGnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbID *e2sm_kpm_v2.GlobalKpmnodeEnGnbId) ([]byte, error) {
	globalKpmnodeEnGnbIDCP, err := newGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeEnGnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalKPMv2node_en_gNB_ID, unsafe.Pointer(globalKpmnodeEnGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeEnGnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKpmnodeEnGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeEnGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMv2node_en_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalKpmnodeEnGnbID((*C.GlobalKPMv2node_en_gNB_ID_t)(unsafePtr))
}

func perDecodeGlobalKpmnodeEnGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeEnGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalKPMv2node_en_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalKpmnodeEnGnbID((*C.GlobalKPMv2node_en_gNB_ID_t)(unsafePtr))
}

func newGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbID *e2sm_kpm_v2.GlobalKpmnodeEnGnbId) (*C.GlobalKPMv2node_en_gNB_ID_t, error) {

	globalGNbIDC, err := newGlobalenGnbID(globalKpmnodeEnGnbID.GlobalGNbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalenGnbID() %s", err.Error())
	}

	globalKpmnodeEnGnbIDC := C.GlobalKPMv2node_en_gNB_ID_t{
		global_gNB_ID: *globalGNbIDC,
		//gNB_CU_UP_ID:  gNbCuUpIDC,
		//gNB_DU_ID:     gNbDuIDC,
	}

	if globalKpmnodeEnGnbID.GNbCuUpId != nil {
		globalKpmnodeEnGnbIDC.gNB_CU_UP_ID, err = newGnbCuUpID(globalKpmnodeEnGnbID.GNbCuUpId)
		if err != nil {
			return nil, fmt.Errorf("newGnbCuUpID() %s", err.Error())
		}
	}

	if globalKpmnodeEnGnbID.GNbDuId != nil {
		globalKpmnodeEnGnbIDC.gNB_DU_ID, err = newGnbDuID(globalKpmnodeEnGnbID.GNbDuId)
		if err != nil {
			return nil, fmt.Errorf("newGnbDuID() %s", err.Error())
		}
	}

	return &globalKpmnodeEnGnbIDC, nil
}

func decodeGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbIDC *C.GlobalKPMv2node_en_gNB_ID_t) (*e2sm_kpm_v2.GlobalKpmnodeEnGnbId, error) {

	globalGNbID, err := decodeGlobalenGnbID(&globalKpmnodeEnGnbIDC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalenGnbID() %s", err.Error())
	}

	globalKpmnodeEnGnbID := e2sm_kpm_v2.GlobalKpmnodeEnGnbId{
		GlobalGNbId: globalGNbID,
		//GNbCuUpId:   gNbCuUpID,
		//GNbDuId:     gNbDuID,
	}

	if globalKpmnodeEnGnbIDC.gNB_CU_UP_ID != nil {
		globalKpmnodeEnGnbID.GNbCuUpId, err = decodeGnbCuUpID(globalKpmnodeEnGnbIDC.gNB_CU_UP_ID)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbCuUpID() %s", err.Error())
		}
	}

	if globalKpmnodeEnGnbIDC.gNB_DU_ID != nil {
		globalKpmnodeEnGnbID.GNbDuId, err = decodeGnbDuID(globalKpmnodeEnGnbIDC.gNB_DU_ID)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbDuID() %s", err.Error())
		}
	}

	return &globalKpmnodeEnGnbID, nil
}

func decodeGlobalKpmnodeEnGnbIDBytes(array [8]byte) (*e2sm_kpm_v2.GlobalKpmnodeEnGnbId, error) {
	globalKpmnodeEnGnbIDC := (*C.GlobalKPMv2node_en_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmnodeEnGnbID(globalKpmnodeEnGnbIDC)
}
