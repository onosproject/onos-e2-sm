// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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

func xerEncodeGlobalKpmNodeGnbId(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) ([]byte, error) {
	gnbCP, err := newGlobalKpmNodeGnbId(gnb)
	defer freeGlobalKpmNodeGnbId(gnbCP)

	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeGnbId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_gNB_ID, unsafe.Pointer(gnbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeGnbId() %s", err.Error())
	}
	return bytes, nil
}

func newGlobalKpmNodeGnbId(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) (*C.GlobalKPMnode_gNB_ID_t, error) {

	globalGnbId, err := newGlobalgNbId(gnb.GlobalGNbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalKpmNodeGnbId() %s", err.Error())
	}
	gnbCuUpId := newGnbCuUpId(gnb.GNbCuUpId)
	gnbDuId := newGnbDuId(gnb.GNbDuId)

	gnbC := C.GlobalKPMnode_gNB_ID_t{
		global_gNB_ID: *globalGnbId,
		gNB_CU_UP_ID: gnbCuUpId,
		gNB_DU_ID: gnbDuId,
	}

	return &gnbC, nil
}

func decodeGlobalKpmNodeGnbId(gnbC *C.GlobalKPMnode_gNB_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeGnbId, error) {

	globalGnbId, err := decodeGlobalgNbId(&gnbC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalKpmNodeGnbId() %s", err.Error())
	}

	gnb := e2sm_kpm_ies.GlobalKpmnodeGnbId{
		GlobalGNbId:          globalGnbId,
		GNbCuUpId:            decodeGnbCuUpId(gnbC.gNB_CU_UP_ID),
		GNbDuId:              decodeGnbDuId(gnbC.gNB_DU_ID),
	}

	return &gnb, nil
}

func decodeGlobalKpmNodeGnbIdBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeGnbId, error) {
	gnbC := (*C.GlobalKPMnode_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmNodeGnbId(gnbC)
}

func freeGlobalKpmNodeGnbId(ptr *C.GlobalKPMnode_gNB_ID_t) {
	freeGnbCuUpId(ptr.gNB_CU_UP_ID)
	freeGnbDuId(ptr.gNB_DU_ID)
}
