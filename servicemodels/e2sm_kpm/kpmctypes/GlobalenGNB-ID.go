// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalenGNB-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalenGNbID(globalenGNbID *e2sm_kpm_ies.GlobalenGnbId) ([]byte, error) {

	globalenGNbIDCP, err := newGlobalenGNbID(globalenGNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalenGNB_ID, unsafe.Pointer(globalenGNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalenGNbID(bytes []byte) (*e2sm_kpm_ies.GlobalenGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalenGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalenGNbID() pointer decoded from XER is nil")
	}
	return decodeGlobalenGNbID((*C.GlobalenGNB_ID_t)(unsafePtr))
}

func newGlobalenGNbID(globalENbID *e2sm_kpm_ies.GlobalenGnbId) (*C.GlobalenGNB_ID_t, error) {

	plmnIDC := newPlmnIdentity(globalENbID.GetPLmnIdentity())
	enGNbIDC, err := newENGNbID(globalENbID.GetGNbId())
	if err != nil {
		return nil, fmt.Errorf("newGlobalenGNbID() error encoding ENGNbID (C struct was not created successfully) \n%v", err)
	}

	globalenGNbIDC := C.GlobalenGNB_ID_t{
		pLMN_Identity: *plmnIDC,
		gNB_ID:        *enGNbIDC,
	}

	return &globalenGNbIDC, nil
}

func decodeGlobalenGNbID(globalenGNbIDC *C.GlobalenGNB_ID_t) (*e2sm_kpm_ies.GlobalenGnbId, error) {

	plmnID := decodePlmnIdentity(&globalenGNbIDC.pLMN_Identity)
	enGNbID, err := decodeENGNbID(&globalenGNbIDC.gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalenGNbID() error decoding ENGNbID_C %v", err)
	}

	return &e2sm_kpm_ies.GlobalenGnbId{
		PLmnIdentity: plmnID,
		GNbId:        enGNbID,
	}, nil
}
