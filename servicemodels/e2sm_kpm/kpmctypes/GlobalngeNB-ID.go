// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalngeNB-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalNgeNBID(globalNgeNB *e2sm_kpm_ies.GlobalngeNbId) ([]byte, error) {

	globalNgeNBCP, err := newGlobalNgeNBID(globalNgeNB)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalNgeNBID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalngeNB_ID, unsafe.Pointer(globalNgeNBCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalNgeNBID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalNgeNBID(bytes []byte) (*e2sm_kpm_ies.GlobalngeNbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalngeNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeGlobalNgeNBID() pointer decoded from XER is nil")
	}
	return decodeGlobalNgeNBID((*C.GlobalngeNB_ID_t)(unsafePtr))
}

func newGlobalNgeNBID(ngeNBID *e2sm_kpm_ies.GlobalngeNbId) (*C.GlobalngeNB_ID_t, error) {

	plmnIDC := newPlmnIdentity(ngeNBID.GetPlmnId())
	eNbIDC, err := newENbIDChoice(ngeNBID.GetEnbId())
	if err != nil {
		return nil, fmt.Errorf("newNgeNBID() %s", err.Error())
	}

	ngeNBIDC := C.GlobalngeNB_ID_t{
		plmn_id: *plmnIDC,
		enb_id:  *eNbIDC,
	}

	return &ngeNBIDC, nil
}

func decodeGlobalNgeNBID(ngeNBIDC *C.GlobalngeNB_ID_t) (*e2sm_kpm_ies.GlobalngeNbId, error) {

	plmnID := decodePlmnIdentity(&ngeNBIDC.plmn_id)
	eNbID, err := decodeENbIDChoice(&ngeNBIDC.enb_id)
	if err != nil {
		return nil, fmt.Errorf("decodeNgeNBID() %s", err.Error())
	}

	return &e2sm_kpm_ies.GlobalngeNbId{
		PlmnId: plmnID,
		EnbId:  eNbID,
	}, nil
}
