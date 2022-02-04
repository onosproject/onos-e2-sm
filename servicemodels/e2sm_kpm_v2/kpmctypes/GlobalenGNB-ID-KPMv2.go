// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalenGNB-ID-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalenGnbID(globalenGnbID *e2sm_kpm_v2.GlobalenGnbId) ([]byte, error) {
	globalenGnbIDCP, err := newGlobalenGnbID(globalenGnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalenGNB_ID_KPMv2, unsafe.Pointer(globalenGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalenGnbID(globalenGnbID *e2sm_kpm_v2.GlobalenGnbId) ([]byte, error) {
	globalenGnbIDCP, err := newGlobalenGnbID(globalenGnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalenGnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalenGNB_ID_KPMv2, unsafe.Pointer(globalenGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalenGnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalenGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalenGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalenGNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalenGnbID((*C.GlobalenGNB_ID_KPMv2_t)(unsafePtr))
}

func perDecodeGlobalenGnbID(bytes []byte) (*e2sm_kpm_v2.GlobalenGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalenGNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalenGnbID((*C.GlobalenGNB_ID_KPMv2_t)(unsafePtr))
}

func newGlobalenGnbID(globalenGnbID *e2sm_kpm_v2.GlobalenGnbId) (*C.GlobalenGNB_ID_KPMv2_t, error) {

	pLmnIdentityC, err := newPlmnIdentity(globalenGnbID.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	gNbIDC, err := newEngnbID(globalenGnbID.GNbId)
	if err != nil {
		return nil, fmt.Errorf("newEngnbId() %s", err.Error())
	}

	globalenGnbIDC := C.GlobalenGNB_ID_KPMv2_t{
		pLMN_Identity: *pLmnIdentityC,
		gNB_ID:        *gNbIDC,
	}

	return &globalenGnbIDC, nil
}

func decodeGlobalenGnbID(globalenGnbIDC *C.GlobalenGNB_ID_KPMv2_t) (*e2sm_kpm_v2.GlobalenGnbId, error) {

	pLmnIdentity, err := decodePlmnIdentity(&globalenGnbIDC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	gNbID, err := decodeEngnbID(&globalenGnbIDC.gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeEngnbId() %s", err.Error())
	}

	globalenGnbID := e2sm_kpm_v2.GlobalenGnbId{
		PLmnIdentity: pLmnIdentity,
		GNbId:        gNbID,
	}

	return &globalenGnbID, nil
}
