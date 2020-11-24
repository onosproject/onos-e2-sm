// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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

func xerEncodeGnbCuUpId(gnbCuUpId *e2sm_kpm_ies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIdCP := newGnbCuUpId(gnbCuUpId)

	bytes, err := encodeXer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIdCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbCuUpId() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbCuUpId(gnbCuUpId *e2sm_kpm_ies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIdCP := newGnbCuUpId(gnbCuUpId)

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIdCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbCuUpId() %s", err.Error())
	}
	fmt.Printf("perEncodeGnbCuUpId -- bytes -- %v\n", bytes)
	return bytes, nil
}

func perDecodeGnbCuUpId(bytes []byte) (*e2sm_kpm_ies.GnbCuUpId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	fmt.Printf("perDecodeGnbCuUpId -- unsafePtr -- %v\n", unsafePtr)
	return decodeGnbCuUpId((*C.GNB_CU_UP_ID_t)(unsafePtr))
}

func newGnbCuUpId(gnbCuUpId *e2sm_kpm_ies.GnbCuUpId) (*C.GNB_CU_UP_ID_t) {

	fmt.Printf("newGnbCuUpId -- GnbCuUpId -- %v\n", gnbCuUpId)
	gnbCuUpIdC := newInteger(gnbCuUpId.Value)
	fmt.Printf("newGnbCuUpId -- GnbCuUpIdC -- %v\n", gnbCuUpIdC)
	return gnbCuUpIdC
}

func decodeGnbCuUpId(gnbCuUpIdC *C.GNB_CU_UP_ID_t) (*e2sm_kpm_ies.GnbCuUpId, error) {
	gnbCuUpId := new(e2sm_kpm_ies.GnbCuUpId)
	fmt.Printf("decodeGnbCuUpId -- GnbCuUpIdC -- %v\n", gnbCuUpIdC)
	resultStr := decodeInteger(gnbCuUpIdC)
	gnbCuUpId.Value = resultStr
	fmt.Printf("decodeGnbCuUpId -- gnbCuUpId -- %v\n", gnbCuUpId)

	return gnbCuUpId, nil
}
