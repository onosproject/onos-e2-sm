// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-ng-eNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKPMnodengeNBID(globalKpmnodeNgEnbID *e2sm_kpm_ies.GlobalKpmnodeNgEnbId) ([]byte, error) {

	globalKpmnodeNgEnbIDCP, err := newGlobalKPMnodengeNBID(globalKpmnodeNgEnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodengeNBID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_ng_eNB_ID, unsafe.Pointer(globalKpmnodeNgEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKPMnodengeNBID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKPMnodengeNBID(bytes []byte) (*e2sm_kpm_ies.GlobalKpmnodeNgEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_ng_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalKPMnodengeNBID() pointer decoded from XER is nil")
	}
	return decodeGlobalKPMnodengeNBID((*C.GlobalKPMnode_ng_eNB_ID_t)(unsafePtr))
}

func newGlobalKPMnodengeNBID(globalKpmnodeNgEnbID *e2sm_kpm_ies.GlobalKpmnodeNgEnbId) (*C.GlobalKPMnode_ng_eNB_ID_t, error) {

	globalngeNBIDC, err := newGlobalNgeNBID(globalKpmnodeNgEnbID.GetGlobalNgENbId())
	if err != nil {
		return nil, fmt.Errorf("newGlobalKPMnodengeNBID() %s", err.Error())
	}

	globalKpmnodeNgEnbIDC := C.GlobalKPMnode_ng_eNB_ID_t{
		global_ng_eNB_ID: *globalngeNBIDC,
	}

	return &globalKpmnodeNgEnbIDC, nil
}

func decodeGlobalKPMnodengeNBID(globalKpmnodeNgEnbIDC *C.GlobalKPMnode_ng_eNB_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeNgEnbId, error) {

	globalngeNBID, err := decodeGlobalNgeNBID(&globalKpmnodeNgEnbIDC.global_ng_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalKPMnodengeNBID() %s", err.Error())
	}

	return &e2sm_kpm_ies.GlobalKpmnodeNgEnbId{
		GlobalNgENbId: globalngeNBID,
	}, nil
}

func decodeGlobalKPMnodengeNBIDBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeNgEnbId, error) {
	enbC := (*C.GlobalKPMnode_ng_eNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKPMnodengeNBID(enbC)
}
