// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-eNB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKpmnodeEnbID(globalKpmnodeEnbID *e2sm_kpm_v2.GlobalKpmnodeEnbId) ([]byte, error) {
	globalKpmnodeEnbIDCP, err := newGlobalKpmnodeEnbID(globalKpmnodeEnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_eNB_ID, unsafe.Pointer(globalKpmnodeEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeEnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalKpmnodeEnbID(globalKpmnodeEnbID *e2sm_kpm_v2.GlobalKpmnodeEnbId) ([]byte, error) {
	globalKpmnodeEnbIDCP, err := newGlobalKpmnodeEnbID(globalKpmnodeEnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeEnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalKPMnode_eNB_ID, unsafe.Pointer(globalKpmnodeEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKpmnodeEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalKpmnodeEnbID((*C.GlobalKPMnode_eNB_ID_t)(unsafePtr))
}

func perDecodeGlobalKpmnodeEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeEnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalKPMnode_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalKpmnodeEnbID((*C.GlobalKPMnode_eNB_ID_t)(unsafePtr))
}

func newGlobalKpmnodeEnbID(globalKpmnodeEnbID *e2sm_kpm_v2.GlobalKpmnodeEnbId) (*C.GlobalKPMnode_eNB_ID_t, error) {

	globalENbIDC, err := newGlobalEnbID(globalKpmnodeEnbID.GlobalENbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalEnbID() %s", err.Error())
	}

	globalKpmnodeEnbIDC := C.GlobalKPMnode_eNB_ID_t{
		global_eNB_ID: *globalENbIDC,
	}

	return &globalKpmnodeEnbIDC, nil
}

func decodeGlobalKpmnodeEnbID(globalKpmnodeEnbIDC *C.GlobalKPMnode_eNB_ID_t) (*e2sm_kpm_v2.GlobalKpmnodeEnbId, error) {

	globalENbID, err := decodeGlobalEnbID(&globalKpmnodeEnbIDC.global_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalEnbID() %s", err.Error())
	}

	globalKpmnodeEnbID := e2sm_kpm_v2.GlobalKpmnodeEnbId{
		GlobalENbId: globalENbID,
	}

	return &globalKpmnodeEnbID, nil
}

func decodeGlobalKpmnodeEnbIDBytes(array [8]byte) (*e2sm_kpm_v2.GlobalKpmnodeEnbId, error) {
	globalKpmnodeEnbIDC := (*C.GlobalKPMnode_eNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmnodeEnbID(globalKpmnodeEnbIDC)
}
