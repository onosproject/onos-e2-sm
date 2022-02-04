// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-gNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKpmNodeGnbID(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) ([]byte, error) {
	gnbCP, err := newGlobalKpmNodeGnbID(gnb)
	defer freeGlobalKpmNodeGnbID(gnbCP)

	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeGnbId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_gNB_ID, unsafe.Pointer(gnbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeGnbId() %s", err.Error())
	}
	return bytes, nil
}

func newGlobalKpmNodeGnbID(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) (*C.GlobalKPMnode_gNB_ID_t, error) {

	globalGnbID, err := newGlobalgNbID(gnb.GlobalGNbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalKpmNodeGnbId() %s", err.Error())
	}
	gnbCuUpID := newGnbCuUpID(gnb.GNbCuUpId)
	gnbDuID := newGnbDuID(gnb.GNbDuId)

	gnbC := C.GlobalKPMnode_gNB_ID_t{
		global_gNB_ID: *globalGnbID,
		gNB_CU_UP_ID:  gnbCuUpID,
		gNB_DU_ID:     gnbDuID,
	}

	return &gnbC, nil
}

func decodeGlobalKpmNodeGnbID(gnbC *C.GlobalKPMnode_gNB_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeGnbId, error) {

	globalGnbID, err := decodeGlobalgNbID(&gnbC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalKpmNodeGnbId() %s", err.Error())
	}

	gnb := e2sm_kpm_ies.GlobalKpmnodeGnbId{
		GlobalGNbId: globalGnbID,
	}
	if gnbC.gNB_CU_UP_ID != nil {
		gnb.GNbCuUpId = decodeGnbCuUpID(gnbC.gNB_CU_UP_ID)
	}
	if gnbC.gNB_DU_ID != nil {
		gnb.GNbDuId = decodeGnbDuID(gnbC.gNB_DU_ID)
	}
	return &gnb, nil
}

func decodeGlobalKpmNodeGnbIDBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeGnbId, error) {
	gnbC := (*C.GlobalKPMnode_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmNodeGnbID(gnbC)
}

func freeGlobalKpmNodeGnbID(ptr *C.GlobalKPMnode_gNB_ID_t) {
	freeGnbCuUpID(ptr.gNB_CU_UP_ID)
	freeGnbDuID(ptr.gNB_DU_ID)
}
