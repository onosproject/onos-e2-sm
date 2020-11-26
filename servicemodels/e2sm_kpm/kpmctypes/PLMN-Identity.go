// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PLMN-Identity.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodePlmnIdentity(plmnID *e2sm_kpm_ies.PlmnIdentity) ([]byte, error) {

	plmnIDCP := newPlmnIdentity(plmnID)

	bytes, err := encodeXer(&C.asn_DEF_PLMN_Identity, unsafe.Pointer(plmnIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func newPlmnIdentity(plmnID *e2sm_kpm_ies.PlmnIdentity) *C.PLMN_Identity_t {

	plmnIDC := newOctetString(string(plmnID.Value))

	return plmnIDC
}

func decodePlmnIdentity(plmnIDC *C.PLMN_Identity_t) *e2sm_kpm_ies.PlmnIdentity {
	plmnID := new(e2sm_kpm_ies.PlmnIdentity)

	plmnIDO := decodeOctetString(plmnIDC)
	plmnID.Value = []byte(plmnIDO)

	return plmnID
}
