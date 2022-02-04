// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-CU-UP-ID.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGnbCuUpID(gnbCuUpID *e2sm_kpm_ies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIDCP := newGnbCuUpID(gnbCuUpID)
	defer freeGnbCuUpID(gnbCuUpIDCP)
	bytes, err := encodeXer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbCuUpId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbCuUpID(gnbCuUpID *e2sm_kpm_ies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIDCP := newGnbCuUpID(gnbCuUpID)

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbCuUpId() %s", err.Error())
	}
	return bytes, nil
}

func perDecodeGnbCuUpID(bytes []byte) (*e2sm_kpm_ies.GnbCuUpId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGnbCuUpID((*C.GNB_CU_UP_ID_t)(unsafePtr)), nil
}

func newGnbCuUpID(gnbCuUpID *e2sm_kpm_ies.GnbCuUpId) *C.GNB_CU_UP_ID_t {

	return newInteger(gnbCuUpID.Value)
}

func decodeGnbCuUpID(gnbCuUpIDC *C.GNB_CU_UP_ID_t) *e2sm_kpm_ies.GnbCuUpId {
	return &e2sm_kpm_ies.GnbCuUpId{
		Value: decodeInteger(gnbCuUpIDC),
	}
}

func freeGnbCuUpID(gnbCuUpIDC *C.GNB_CU_UP_ID_t) {
	freeInteger(gnbCuUpIDC)
}
