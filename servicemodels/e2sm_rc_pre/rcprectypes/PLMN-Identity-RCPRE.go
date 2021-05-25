// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PLMN-Identity-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodePlmnIdentity(plmnID *e2sm_rc_pre_v2.PlmnIdentity) ([]byte, error) {

	plmnIDCP := newPlmnIdentity(plmnID)

	bytes, err := encodeXer(&C.asn_DEF_PLMN_Identity_RCPRE, unsafe.Pointer(plmnIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func newPlmnIdentity(plmnID *e2sm_rc_pre_v2.PlmnIdentity) *C.PLMN_Identity_RCPRE_t {

	plmnIDC := newOctetString(string(plmnID.GetValue()))

	return plmnIDC
}

func decodePlmnIdentity(plmnIDC *C.PLMN_Identity_RCPRE_t) *e2sm_rc_pre_v2.PlmnIdentity {
	plmnID := new(e2sm_rc_pre_v2.PlmnIdentity)

	plmnIDO := decodeOctetString(plmnIDC)
	plmnID.Value = []byte(plmnIDO)

	return plmnID
}
