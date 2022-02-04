// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

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
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKPMnodeENbID(globalKPMnodeENbID *e2sm_kpm_ies.GlobalKpmnodeEnbId) ([]byte, error) {

	globalKPMnodeENbIDCP, err := newGlobalKPMnodeENbID(globalKPMnodeENbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodeENbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_eNB_ID, unsafe.Pointer(globalKPMnodeENbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodeENbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKPMnodeENbID(bytes []byte) (*e2sm_kpm_ies.GlobalKpmnodeEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalKPMnodeENbID() pointer decoded from XER is nil")
	}
	return decodeGlobalKPMnodeENbID((*C.GlobalKPMnode_eNB_ID_t)(unsafePtr))
}

func newGlobalKPMnodeENbID(globalENbID *e2sm_kpm_ies.GlobalKpmnodeEnbId) (*C.GlobalKPMnode_eNB_ID_t, error) {

	globalENbIDC, err := newGlobalENbID(globalENbID.GetGlobalENbId())
	if err != nil {
		return nil, fmt.Errorf("newGlobalENbID() error encoding ENbID (C struct was not created successfully) \n%v", err)
	}

	globalKPMnodeENbIDC := C.GlobalKPMnode_eNB_ID_t{
		global_eNB_ID: *globalENbIDC,
	}
	return &globalKPMnodeENbIDC, nil
}

func decodeGlobalKPMnodeENbID(globalENbIDC *C.GlobalKPMnode_eNB_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeEnbId, error) {

	globalENbID, err := decodeGlobalENbID(&globalENbIDC.global_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalENbID() error decoding ENbID_C %v", err)
	}

	return &e2sm_kpm_ies.GlobalKpmnodeEnbId{
		GlobalENbId: globalENbID,
	}, nil
}

func decodeGlobalKPMnodeENbIDBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeEnbId, error) {
	enbC := (*C.GlobalKPMnode_eNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKPMnodeENbID(enbC)
}
