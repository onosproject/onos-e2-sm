// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalENB-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalENbID(globalENbID *e2sm_kpm_ies.GlobalEnbId) ([]byte, error) {

	globalENbIDCP, err := newGlobalENbID(globalENbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalENbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalENB_ID, unsafe.Pointer(globalENbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalENbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalENbID(bytes []byte) (*e2sm_kpm_ies.GlobalEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalENbID() pointer decoded from XER is nil")
	}
	return decodeGlobalENbID((*C.GlobalENB_ID_t)(unsafePtr))
}

func newGlobalENbID(globalENbID *e2sm_kpm_ies.GlobalEnbId) (*C.GlobalENB_ID_t, error) {

	plmnIDC := newPlmnIdentity(globalENbID.GetPLmnIdentity())
	enbIDC, err := newENbID(globalENbID.GetENbId())
	if err != nil {
		return nil, fmt.Errorf("newGlobalENbID() error encoding ENbID (C struct was not created successfully) \n%v", err)
	}

	globalENbIDC := C.GlobalENB_ID_t{
		pLMN_Identity: *plmnIDC,
		eNB_ID:        *enbIDC,
	}

	return &globalENbIDC, nil
}

func decodeGlobalENbID(globalENbIDC *C.GlobalENB_ID_t) (*e2sm_kpm_ies.GlobalEnbId, error) {

	plmnID := decodePlmnIdentity(&globalENbIDC.pLMN_Identity)
	enbID, err := decodeENbID(&globalENbIDC.eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalENbID() error decoding ENbID_C %v", err)
	}

	return &e2sm_kpm_ies.GlobalEnbId{
		PLmnIdentity: plmnID,
		ENbId:        enbID,
	}, nil
}
