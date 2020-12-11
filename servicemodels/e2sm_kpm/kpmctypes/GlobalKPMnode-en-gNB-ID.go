// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-en-gNB-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKPMnodeEnGNbID(globalKPMnodeENbID *e2sm_kpm_ies.GlobalKpmnodeEnGnbId) ([]byte, error) {

	globalKPMnodeENbIDCP, err := newGlobalKPMnodeEnGNbID(globalKPMnodeENbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodeEnGNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_en_gNB_ID, unsafe.Pointer(globalKPMnodeENbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodeEnGNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKPMnodeEnGNbID(bytes []byte) (*e2sm_kpm_ies.GlobalKpmnodeEnGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_en_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalKPMnodeEnGNbID() pointer decoded from XER is nil")
	}
	return decodeGlobalKPMnodeEnGNbID((*C.GlobalKPMnode_en_gNB_ID_t)(unsafePtr))
}

func newGlobalKPMnodeEnGNbID(globalenGNbID *e2sm_kpm_ies.GlobalKpmnodeEnGnbId) (*C.GlobalKPMnode_en_gNB_ID_t, error) {

	globalenGNbIDC, err := newGlobalenGNbID(globalenGNbID.GetGlobalGNbId())
	if err != nil {
		return nil, fmt.Errorf("newGlobalKPMnodeEnGNbID() error encoding ENGNbID (C struct was not created successfully) \n%v", err)
	}

	globalKPMnodeENbIDC := C.GlobalKPMnode_en_gNB_ID_t{
		global_gNB_ID: *globalenGNbIDC,
	}
	return &globalKPMnodeENbIDC, nil
}

func decodeGlobalKPMnodeEnGNbID(globalenGNbIDC *C.GlobalKPMnode_en_gNB_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeEnGnbId, error) {

	globalenGNbID, err := decodeGlobalenGNbID(&globalenGNbIDC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalKPMnodeEnGNbID() error decoding ENGNbID_C %v", err)
	}

	return &e2sm_kpm_ies.GlobalKpmnodeEnGnbId{
		GlobalGNbId: globalenGNbID,
	}, nil
}

//func decodeGlobalKPMnodeEnGNbIDBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeEnGnbId, error) {
//	enbC := (*C.GlobalKPMnode_en_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
//
//	return decodeGlobalKPMnodeEnGNbID(enbC)
//}
