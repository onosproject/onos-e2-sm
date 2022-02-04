// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-ng-eNB-ID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbID *e2sm_kpm_v2.GlobalKpmnodeNgEnbId) ([]byte, error) {
	globalKpmnodeNgEnbIDCP, err := newGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeNgEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_ng_eNB_ID_KPMv2, unsafe.Pointer(globalKpmnodeNgEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeNgEnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbID *e2sm_kpm_v2.GlobalKpmnodeNgEnbId) ([]byte, error) {
	globalKpmnodeNgEnbIDCP, err := newGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeNgEnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalKPMnode_ng_eNB_ID_KPMv2, unsafe.Pointer(globalKpmnodeNgEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeNgEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKpmnodeNgEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeNgEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_ng_eNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalKpmnodeNgEnbID((*C.GlobalKPMnode_ng_eNB_ID_KPMv2_t)(unsafePtr))
}

func perDecodeGlobalKpmnodeNgEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeNgEnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalKPMnode_ng_eNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalKpmnodeNgEnbID((*C.GlobalKPMnode_ng_eNB_ID_KPMv2_t)(unsafePtr))
}

func newGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbID *e2sm_kpm_v2.GlobalKpmnodeNgEnbId) (*C.GlobalKPMnode_ng_eNB_ID_KPMv2_t, error) {

	globalNgENbIDC, err := newGlobalngeNbID(globalKpmnodeNgEnbID.GlobalNgENbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalngeNbID() %s", err.Error())
	}

	globalKpmnodeNgEnbIDC := C.GlobalKPMnode_ng_eNB_ID_KPMv2_t{
		global_ng_eNB_ID: *globalNgENbIDC,
		//gNB_DU_ID:        gNbDuIDC,
	}

	if globalKpmnodeNgEnbID.GNbDuId != nil {
		globalKpmnodeNgEnbIDC.gNB_DU_ID, err = newGnbDuID(globalKpmnodeNgEnbID.GNbDuId)
		if err != nil {
			return nil, fmt.Errorf("newGnbDuID() %s", err.Error())
		}
	}

	return &globalKpmnodeNgEnbIDC, nil
}

func decodeGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbIDC *C.GlobalKPMnode_ng_eNB_ID_KPMv2_t) (*e2sm_kpm_v2.GlobalKpmnodeNgEnbId, error) {

	globalNgENbID, err := decodeGlobalngeNbID(&globalKpmnodeNgEnbIDC.global_ng_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalngeNbID() %s", err.Error())
	}

	globalKpmnodeNgEnbID := e2sm_kpm_v2.GlobalKpmnodeNgEnbId{
		GlobalNgENbId: globalNgENbID,
		//GNbDuId:       gNbDuID,
	}

	if globalKpmnodeNgEnbIDC.gNB_DU_ID != nil {
		globalKpmnodeNgEnbID.GNbDuId, err = decodeGnbDuID(globalKpmnodeNgEnbIDC.gNB_DU_ID)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbDuID() %s", err.Error())
		}
	}

	return &globalKpmnodeNgEnbID, nil
}

func decodeGlobalKpmnodeNgEnbIDBytes(array [8]byte) (*e2sm_kpm_v2.GlobalKpmnodeNgEnbId, error) {
	globalKpmnodeNgEnbIDC := (*C.GlobalKPMnode_ng_eNB_ID_KPMv2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmnodeNgEnbID(globalKpmnodeNgEnbIDC)
}
