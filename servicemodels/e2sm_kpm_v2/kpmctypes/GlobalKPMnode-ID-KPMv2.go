// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-ID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalKpmnodeID(globalKpmnodeID *e2sm_kpm_v2.GlobalKpmnodeId) ([]byte, error) {
	globalKpmnodeIDCP, err := newGlobalKpmnodeID(globalKpmnodeID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_ID_KPMv2, unsafe.Pointer(globalKpmnodeIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmnodeID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalKpmnodeID(globalKpmnodeID *e2sm_kpm_v2.GlobalKpmnodeId) ([]byte, error) {
	globalKpmnodeIDCP, err := newGlobalKpmnodeID(globalKpmnodeID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalKPMnode_ID_KPMv2, unsafe.Pointer(globalKpmnodeIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalKpmnodeID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalKpmnodeID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalKPMnode_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalKpmnodeID((*C.GlobalKPMnode_ID_KPMv2_t)(unsafePtr))
}

func perDecodeGlobalKpmnodeID(bytes []byte) (*e2sm_kpm_v2.GlobalKpmnodeId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalKPMnode_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalKpmnodeID((*C.GlobalKPMnode_ID_KPMv2_t)(unsafePtr))
}

func newGlobalKpmnodeID(globalKpmnodeID *e2sm_kpm_v2.GlobalKpmnodeId) (*C.GlobalKPMnode_ID_KPMv2_t, error) {

	var pr C.GlobalKPMnode_ID_KPMv2_PR
	choiceC := [8]byte{}

	switch choice := globalKpmnodeID.GlobalKpmnodeId.(type) {
	case *e2sm_kpm_v2.GlobalKpmnodeId_GNb:
		pr = C.GlobalKPMnode_ID_KPMv2_PR_gNB

		im, err := newGlobalKpmnodeGnbID(choice.GNb)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmnodeGnbID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.GlobalKpmnodeId_EnGNb:
		pr = C.GlobalKPMnode_ID_KPMv2_PR_en_gNB

		im, err := newGlobalKpmnodeEnGnbID(choice.EnGNb)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmnodeEnGnbID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.GlobalKpmnodeId_NgENb:
		pr = C.GlobalKPMnode_ID_KPMv2_PR_ng_eNB

		im, err := newGlobalKpmnodeNgEnbID(choice.NgENb)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmnodeNgEnbID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.GlobalKpmnodeId_ENb:
		pr = C.GlobalKPMnode_ID_KPMv2_PR_eNB

		im, err := newGlobalKpmnodeEnbID(choice.ENb)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmnodeEnbID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newGlobalKpmnodeID() %T not yet implemented", choice)
	}

	globalKpmnodeIDC := C.GlobalKPMnode_ID_KPMv2_t{
		present: pr,
		choice:  choiceC,
	}

	return &globalKpmnodeIDC, nil
}

func decodeGlobalKpmnodeID(globalKpmnodeIDC *C.GlobalKPMnode_ID_KPMv2_t) (*e2sm_kpm_v2.GlobalKpmnodeId, error) {

	//This is Decoder part (OneOf)
	globalKpmnodeID := new(e2sm_kpm_v2.GlobalKpmnodeId)

	switch globalKpmnodeIDC.present {
	case C.GlobalKPMnode_ID_KPMv2_PR_gNB:
		globalKpmnodeIDstructC, err := decodeGlobalKpmnodeGnbIDBytes(globalKpmnodeIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmnodeGnbIDBytes() %s", err.Error())
		}
		globalKpmnodeID.GlobalKpmnodeId = &e2sm_kpm_v2.GlobalKpmnodeId_GNb{
			GNb: globalKpmnodeIDstructC,
		}
	case C.GlobalKPMnode_ID_KPMv2_PR_en_gNB:
		globalKpmnodeIDstructC, err := decodeGlobalKpmnodeEnGnbIDBytes(globalKpmnodeIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmnodeEnGnbIDBytes() %s", err.Error())
		}
		globalKpmnodeID.GlobalKpmnodeId = &e2sm_kpm_v2.GlobalKpmnodeId_EnGNb{
			EnGNb: globalKpmnodeIDstructC,
		}
	case C.GlobalKPMnode_ID_KPMv2_PR_ng_eNB:
		globalKpmnodeIDstructC, err := decodeGlobalKpmnodeNgEnbIDBytes(globalKpmnodeIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmnodeNgEnbIDBytes() %s", err.Error())
		}
		globalKpmnodeID.GlobalKpmnodeId = &e2sm_kpm_v2.GlobalKpmnodeId_NgENb{
			NgENb: globalKpmnodeIDstructC,
		}
	case C.GlobalKPMnode_ID_KPMv2_PR_eNB:
		globalKpmnodeIDstructC, err := decodeGlobalKpmnodeEnbIDBytes(globalKpmnodeIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmnodeEnbIDBytes() %s", err.Error())
		}
		globalKpmnodeID.GlobalKpmnodeId = &e2sm_kpm_v2.GlobalKpmnodeId_ENb{
			ENb: globalKpmnodeIDstructC,
		}
	case C.GlobalKPMnode_ID_KPMv2_PR_NOTHING:
		return nil, fmt.Errorf("decodeGlobalKpmnodeID() An empty GlobalKPMnodeID has been sent %v", globalKpmnodeIDC.present)
	default:
		return nil, fmt.Errorf("decodeGlobalKpmnodeID() %v not yet implemented", globalKpmnodeIDC.present)
	}

	return globalKpmnodeID, nil
}
