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
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKpmNodeGnbId(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) ([]byte, error) {
	gnbCP, err := newGlobalKpmNodeGnbId(gnb)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_gNB_ID, unsafe.Pointer(gnbCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newGlobalKpmNodeGnbId(gnb *e2sm_kpm_ies.GlobalKpmnodeGnbId) (*C.GlobalKPMnode_gNB_ID_t, error) {

	globalGnbId, err := newGlobalgNbId(gnb.GlobalGNbId)
	if err != nil {
		return nil, err
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
	gnb := new(e2sm_kpm_ies.GlobalKpmnodeGnbId)

	globalGnbId, err := decodeGlobalgNbId(&gnbC.global_gNB_ID)
	if err != nil {
		return nil, err
	}
	gnb.GlobalGNbId = globalGnbId

	gnbCuUpId, err := decodeGnbCuUpId(gnbC.gNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	gnb.GNbCuUpId = gnbCuUpId

	gnbDuId, err := decodeGnbDuId(gnbC.gNB_DU_ID)
	if err != nil {
		return nil, err
	}
	gnb.GNbDuId = gnbDuId

	return gnb, nil
}

func decodeGlobalKpmNodeGnbIdBytes(array [8]byte) (*e2sm_kpm_ies.GlobalKpmnodeGnbId, error) {
	gnbC := (*C.GlobalKPMnode_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalKpmNodeGnbId(gnbC)
}

