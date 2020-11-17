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

func xerEncodeGlobalgNbId(globalgNbId *e2sm_kpm_ies.GlobalgNbId) ([]byte, error) {
	globalgNbIdCP, err := newGlobalgNbId(globalgNbId)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbId() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalgNB_ID, unsafe.Pointer(globalgNbIdCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalgNbId() %s", err.Error())
	}
	return bytes, nil
}

func newGlobalgNbId(globalgNbId *e2sm_kpm_ies.GlobalgNbId) (*C.GlobalgNB_ID_t, error) {

	plmnIDC := newPlmnIdentity(globalgNbId.PlmnId)
	gnbIDC, err := newGnbIDChoice(globalgNbId.GnbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalgNbId() %s", err.Error())
	}

	globalgNbIdC := C.GlobalgNB_ID_t{
		plmn_id: *plmnIDC,
		gnb_id: *gnbIDC,
	}

	return &globalgNbIdC, nil
}

func decodeGlobalgNbId(globalgNbIdC *C.GlobalgNB_ID_t) (*e2sm_kpm_ies.GlobalgNbId, error) {
	globalgNbId := new(e2sm_kpm_ies.GlobalgNbId)
	plmnID := decodePlmnIdentity(&globalgNbIdC.plmn_id)
	gnbID, err := decodeGnbIDChoice(&globalgNbIdC.gnb_id)
	if err == nil {
		return nil, fmt.Errorf("decodeGlobalgNbId() error in str-to-int64 convertion %T", err)
	}

	globalgNbId.PlmnId = plmnID
	globalgNbId.GnbId = gnbID

	return globalgNbId, nil
}
