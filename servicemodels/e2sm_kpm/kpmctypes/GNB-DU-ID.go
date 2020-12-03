// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-DU-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGnbDuID(gnbDuID *e2sm_kpm_ies.GnbDuId) ([]byte, error) {
	gnbDuIDCP := newGnbDuID(gnbDuID)

	bytes, err := encodeXer(&C.asn_DEF_GNB_DU_ID, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbDuId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbDuID(gnbDuID *e2sm_kpm_ies.GnbDuId) ([]byte, error) {
	gnbDuIDCP := newGnbDuID(gnbDuID)

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_DU_ID, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbDuId() %s", err.Error())
	}
	return bytes, nil
}

func perDecodeGnbDuID(bytes []byte) (*e2sm_kpm_ies.GnbDuId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGnbDuID((*C.GNB_DU_ID_t)(unsafePtr)), nil
}

func newGnbDuID(gnbDuID *e2sm_kpm_ies.GnbDuId) *C.GNB_DU_ID_t {
	return newInteger(gnbDuID.Value)
}

func decodeGnbDuID(gnbDuIDC *C.GNB_DU_ID_t) *e2sm_kpm_ies.GnbDuId {
	return &e2sm_kpm_ies.GnbDuId{
		Value: decodeInteger(gnbDuIDC),
	}
}

func freeGnbDuID(gnbDuIDC *C.GNB_DU_ID_t) {
	freeInteger(gnbDuIDC)
}
