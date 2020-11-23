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

func xerEncodeGnbDuId(gnbDuId *e2sm_kpm_ies.GnbDuId) ([]byte, error) {
	gnbDuIdCP := newGnbDuId(gnbDuId)

	bytes, err := encodeXer(&C.asn_DEF_GNB_DU_ID, unsafe.Pointer(gnbDuIdCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbDuId() %s", err.Error())
	}
	return bytes, nil
}

func newGnbDuId(gnbDuId *e2sm_kpm_ies.GnbDuId) (*C.GNB_DU_ID_t) {
	fmt.Printf("newGnbDuId -- gnbDuId -- %v\n", gnbDuId)
	gnbDuIdC := newInteger(gnbDuId.Value)
	fmt.Printf("newGnbDuId -- gnbDuIdC -- %v\n", gnbDuIdC)
	return gnbDuIdC
}

func decodeGnbDuId(gnbDuIdC *C.GNB_DU_ID_t) (*e2sm_kpm_ies.GnbDuId, error) {
	gnbDuId := new(e2sm_kpm_ies.GnbDuId)
	fmt.Printf("decodeGnbDuId -- gnbDuIdC -- %v\n", gnbDuIdC)
	resultStr := decodeInteger(gnbDuIdC)
	gnbDuId.Value = resultStr
	fmt.Printf("decodeGnbDuId -- gnbDuId -- %v\n", gnbDuId)

	return gnbDuId, nil
}
