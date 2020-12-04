// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalgNB-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGlobalgNbID(globalgNbID *e2sm_kpm_ies.GlobalgNbId) ([]byte, error) {
	globalgNbIDCP, err := newGlobalgNbID(globalgNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalgNB_ID, unsafe.Pointer(globalgNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbId() %s", err.Error())
	}
	return bytes, nil
}

func newGlobalgNbID(globalgNbID *e2sm_kpm_ies.GlobalgNbId) (*C.GlobalgNB_ID_t, error) {

	plmnIDC := newPlmnIdentity(globalgNbID.PlmnId)
	gnbIDC, err := newGnbIDChoice(globalgNbID.GnbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalgNbId() %s", err.Error())
	}

	globalgNbIDC := C.GlobalgNB_ID_t{
		plmn_id: *plmnIDC,
		gnb_id:  *gnbIDC,
	}

	return &globalgNbIDC, nil
}

func decodeGlobalgNbID(globalgNbIDC *C.GlobalgNB_ID_t) (*e2sm_kpm_ies.GlobalgNbId, error) {
	globalgNbID := new(e2sm_kpm_ies.GlobalgNbId)
	plmnID := decodePlmnIdentity(&globalgNbIDC.plmn_id)
	gnbID, err := decodeGnbIDChoice(&globalgNbIDC.gnb_id)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalgNbId() error in str-to-int64 conversion %v", err)
	}

	globalgNbID.PlmnId = plmnID
	globalgNbID.GnbId = gnbID

	return globalgNbID, nil
}
