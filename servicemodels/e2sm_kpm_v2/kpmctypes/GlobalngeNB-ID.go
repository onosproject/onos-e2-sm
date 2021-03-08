// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalngeNB-ID.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalngeNbID(globalngeNbID *e2sm_kpm_v2.GlobalngeNbId) ([]byte, error) {
	globalngeNbIDCP, err := newGlobalngeNbID(globalngeNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalngeNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalngeNB_ID, unsafe.Pointer(globalngeNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalngeNbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalngeNbID(globalngeNbID *e2sm_kpm_v2.GlobalngeNbId) ([]byte, error) {
	globalngeNbIDCP, err := newGlobalngeNbID(globalngeNbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalngeNbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalngeNB_ID, unsafe.Pointer(globalngeNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalngeNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalngeNbID(bytes []byte) (*e2sm_kpm_v2.GlobalngeNbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalngeNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalngeNbID((*C.GlobalngeNB_ID_t)(unsafePtr))
}

func perDecodeGlobalngeNbID(bytes []byte) (*e2sm_kpm_v2.GlobalngeNbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalngeNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalngeNbID((*C.GlobalngeNB_ID_t)(unsafePtr))
}

func newGlobalngeNbID(globalngeNbID *e2sm_kpm_v2.GlobalngeNbId) (*C.GlobalngeNB_ID_t, error) {

	plmnIDC, err := newPlmnIdentity(globalngeNbID.PlmnId)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	enbIDC, err := newEnbIDChoice(globalngeNbID.EnbId)
	if err != nil {
		return nil, fmt.Errorf("newEnbIDChoice() %s", err.Error())
	}

	shortMacroENbIDC, err := newBitString(globalngeNbID.ShortMacroENbId)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	longMacroENbIDC, err := newBitString(globalngeNbID.LongMacroENbId)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	globalngeNbIDC := C.GlobalngeNB_ID_t{
		plmn_id:            *plmnIDC,
		enb_id:             *enbIDC,
		short_Macro_eNB_ID: *shortMacroENbIDC,
		long_Macro_eNB_ID:  *longMacroENbIDC,
	}

	return &globalngeNbIDC, nil
}

func decodeGlobalngeNbID(globalngeNbIDC *C.GlobalngeNB_ID_t) (*e2sm_kpm_v2.GlobalngeNbId, error) {

	plmnID, err := decodePlmnIdentity(&globalngeNbIDC.plmn_id)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	enbID, err := decodeEnbIDChoice(&globalngeNbIDC.enb_id)
	if err != nil {
		return nil, fmt.Errorf("decodeEnbIdChoice() %s", err.Error())
	}

	shortMacroENbID, err := decodeBitString(&globalngeNbIDC.short_Macro_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}

	longMacroENbID, err := decodeBitString(&globalngeNbIDC.long_Macro_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}

	globalngeNbID := e2sm_kpm_v2.GlobalngeNbId{
		PlmnId:          plmnID,
		EnbId:           enbID,
		ShortMacroENbId: shortMacroENbID,
		LongMacroENbId:  longMacroENbID,
	}

	return &globalngeNbID, nil
}
