// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalENB-ID.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGlobalEnbID(globalEnbID *e2sm_kpm_v2.GlobalEnbId) ([]byte, error) {
	globalEnbIDCP, err := newGlobalEnbID(globalEnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalENB_ID, unsafe.Pointer(globalEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalEnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalEnbID(globalEnbID *e2sm_kpm_v2.GlobalEnbId) ([]byte, error) {
	globalEnbIDCP, err := newGlobalEnbID(globalEnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalEnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalENB_ID, unsafe.Pointer(globalEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalEnbID((*C.GlobalENB_ID_t)(unsafePtr))
}

func perDecodeGlobalEnbID(bytes []byte) (*e2sm_kpm_v2.GlobalEnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalEnbID((*C.GlobalENB_ID_t)(unsafePtr))
}

func newGlobalEnbID(globalEnbID *e2sm_kpm_v2.GlobalEnbId) (*C.GlobalENB_ID_t, error) {

	pLmnIdentityC, err := newPlmnIdentity(globalEnbID.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	eNbIDC, err := newEnbID(globalEnbID.ENbId)
	if err != nil {
		return nil, fmt.Errorf("newEnbId() %s", err.Error())
	}

	globalEnbIDC := C.GlobalENB_ID_t{
		pLMN_Identity: *pLmnIdentityC,
		eNB_ID:        *eNbIDC,
	}

	return &globalEnbIDC, nil
}

func decodeGlobalEnbID(globalEnbIDC *C.GlobalENB_ID_t) (*e2sm_kpm_v2.GlobalEnbId, error) {

	pLmnIdentity, err := decodePlmnIdentity(&globalEnbIDC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	eNbID, err := decodeEnbID(&globalEnbIDC.eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeEnbID() %s", err.Error())
	}

	globalEnbID := e2sm_kpm_v2.GlobalEnbId{
		PLmnIdentity: pLmnIdentity,
		ENbId:        eNbID,
	}

	return &globalEnbID, nil
}
