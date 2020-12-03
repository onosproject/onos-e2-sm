// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalKPMnode-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalKpmNodeID(idGlobalKpmnodeID *e2sm_kpm_ies.GlobalKpmnodeId) ([]byte, error) {
	idGlobalKpmnodeIDCP, err := newGlobalKpmNodeID(idGlobalKpmnodeID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalKPMnode_ID, unsafe.Pointer(idGlobalKpmnodeIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalKpmNodeId() %s", err.Error())
	}
	return bytes, nil
}

func newGlobalKpmNodeID(idGlobalKpmnodeID *e2sm_kpm_ies.GlobalKpmnodeId) (*C.GlobalKPMnode_ID_t, error) {
	var pr C.GlobalKPMnode_ID_PR
	choiceC := [8]byte{}

	switch choice := idGlobalKpmnodeID.GlobalKpmnodeId.(type) {
	case *e2sm_kpm_ies.GlobalKpmnodeId_GNb:
		pr = C.GlobalKPMnode_ID_PR_gNB
		globalKpmNodeGnbIDC, err := newGlobalKpmNodeGnbID(choice.GNb)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmNodeId() %s", err.Error())
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalKpmNodeGnbIDC))))
	default:
		return nil, fmt.Errorf("newGlobalKpmNodeId() unexpected type %T not yet implemented", choice)
	}

	idGlobalKpmnodeIDC := C.GlobalKPMnode_ID_t{
		present: pr,
		choice:  choiceC,
	}

	return &idGlobalKpmnodeIDC, nil
}

func decodeGlobalKpmNodeID(idGlobalKpmnodeIDC *C.GlobalKPMnode_ID_t) (*e2sm_kpm_ies.GlobalKpmnodeId, error) {
	result := new(e2sm_kpm_ies.GlobalKpmnodeId)

	switch idGlobalKpmnodeIDC.present {
	case C.GlobalKPMnode_ID_PR_gNB:

		gnb, err := decodeGlobalKpmNodeGnbIDBytes(idGlobalKpmnodeIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmNodeId() %s", err.Error())
		}

		result.GlobalKpmnodeId = &e2sm_kpm_ies.GlobalKpmnodeId_GNb{
			GNb: gnb,
		}
	default:
		return nil, fmt.Errorf("decodeGlobalKpmNodeId() %v not yet implemented", idGlobalKpmnodeIDC.present)
	}

	return result, nil
}
