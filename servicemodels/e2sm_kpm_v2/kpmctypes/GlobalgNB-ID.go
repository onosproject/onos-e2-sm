// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalgNB-ID.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalgNbID(globalgNbID *e2sm_kpm_v2.GlobalgNbId) ([]byte, error) {
	globalgNbIDCP, err := newGlobalgNbID(globalgNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalgNB_ID, unsafe.Pointer(globalgNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalgNbID(globalgNbID *e2sm_kpm_v2.GlobalgNbId) ([]byte, error) {
	globalgNbIDCP, err := newGlobalgNbID(globalgNbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalgNbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalgNB_ID, unsafe.Pointer(globalgNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalgNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalgNbID(bytes []byte) (*e2sm_kpm_v2.GlobalgNbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalgNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalgNbID((*C.GlobalgNB_ID_t)(unsafePtr))
}

func perDecodeGlobalgNbID(bytes []byte) (*e2sm_kpm_v2.GlobalgNbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalgNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalgNbID((*C.GlobalgNB_ID_t)(unsafePtr))
}

func newGlobalgNbID(globalgNbID *e2sm_kpm_v2.GlobalgNbId) (*C.GlobalgNB_ID_t, error) {

	plmnIDC, err := newPlmnIdentity(globalgNbID.PlmnId)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	gnbIDC, err := newGnbIDChoice(globalgNbID.GnbId)
	if err != nil {
		return nil, fmt.Errorf("newGnbIdChoice() %s", err.Error())
	}

	globalgNbIDC := C.GlobalgNB_ID_t{
		plmn_id: *plmnIDC,
		gnb_id:  *gnbIDC,
	}

	return &globalgNbIDC, nil
}

func decodeGlobalgNbID(globalgNbIDC *C.GlobalgNB_ID_t) (*e2sm_kpm_v2.GlobalgNbId, error) {

	plmnID, err := decodePlmnIdentity(&globalgNbIDC.plmn_id)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	gnbID, err := decodeGnbIDChoice(&globalgNbIDC.gnb_id)
	if err != nil {
		return nil, fmt.Errorf("decodeGnbIdChoice() %s", err.Error())
	}

	globalgNbID := e2sm_kpm_v2.GlobalgNbId{
		PlmnId: plmnID,
		GnbId:  gnbID,
	}

	return &globalgNbID, nil
}
