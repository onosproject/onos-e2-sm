// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCellIdentity.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeNRCellIdentity(nrCellIdentity *e2sm_kpm_ies.NrcellIdentity) ([]byte, error) {
	nrCellIdentityCP := newNRCellIdentity(nrCellIdentity)

	bytes, err := encodeXer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrCellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRCellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNRCellIdentity(nrCellIdentity *e2sm_kpm_ies.NrcellIdentity) ([]byte, error) {
	nrCellIdentityCP := newNRCellIdentity(nrCellIdentity)

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrCellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNRCellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func newNRCellIdentity(nrCellIdentity *e2sm_kpm_ies.NrcellIdentity) *C.NRCellIdentity_t {
	nrCellIdentityC := newBitString(nrCellIdentity.Value)

	return nrCellIdentityC
}

func decodeNRCellIdentity(nrCellIdentityC *C.NRCellIdentity_t) (*e2sm_kpm_ies.NrcellIdentity, error) {
	nrCellIdentity := new(e2sm_kpm_ies.NrcellIdentity)

	nrCellIdentityBs, err := decodeBitString(nrCellIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeNRCellIdentity() %s", err.Error())
	}
	nrCellIdentity.Value = nrCellIdentityBs

	return nrCellIdentity, nil
}
