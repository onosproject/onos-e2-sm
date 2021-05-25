// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCellIdentity-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeNRCellIdentity(nrCellIdentity *e2sm_rc_pre_v2.NrcellIdentity) ([]byte, error) {
	nrCellIdentityCP := newNRCellIdentity(nrCellIdentity)

	bytes, err := encodeXer(&C.asn_DEF_NRCellIdentity_RCPRE, unsafe.Pointer(nrCellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRCellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNRCellIdentity(nrCellIdentity *e2sm_rc_pre_v2.NrcellIdentity) ([]byte, error) {
	nrCellIdentityCP := newNRCellIdentity(nrCellIdentity)

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCellIdentity_RCPRE, unsafe.Pointer(nrCellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNRCellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func newNRCellIdentity(nrCellIdentity *e2sm_rc_pre_v2.NrcellIdentity) *C.NRCellIdentity_RCPRE_t {
	nrCellIdentityC := newBitString(nrCellIdentity.Value)

	return nrCellIdentityC
}

func decodeNRCellIdentity(nrCellIdentityC *C.NRCellIdentity_RCPRE_t) (*e2sm_rc_pre_v2.NrcellIdentity, error) {
	nrCellIdentity := new(e2sm_rc_pre_v2.NrcellIdentity)

	nrCellIdentityBs, err := decodeBitString(nrCellIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeNRCellIdentity() %s", err.Error())
	}
	nrCellIdentity.Value = nrCellIdentityBs

	return nrCellIdentity, nil
}
