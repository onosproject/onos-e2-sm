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
	"strconv"
	"unsafe"
)

func xerEncodeGnbCuUpId(gnbCuUpId *e2sm_kpm_ies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIdCP := newGnbCuUpId(gnbCuUpId)

	bytes, err := encodeXer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIdCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newGnbCuUpId(gnbCuUpId *e2sm_kpm_ies.GnbCuUpId) (*C.GNB_CU_UP_ID_t) {

	// TODO: Check whether basis is decimal. Also consider putting out basis as a parameter
	gnbCuUpIdC := newInteger(strconv.FormatInt(gnbCuUpId.Value, 10))
	return gnbCuUpIdC
}

func decodeGnbCuUpId(gnbCuUpIdC *C.GNB_CU_UP_ID_t) (*e2sm_kpm_ies.GnbCuUpId, error) {
	gnbCuUpId := new(e2sm_kpm_ies.GnbCuUpId)
	resultStr := decodeInteger(gnbCuUpIdC)
	resultInt, err := strconv.ParseInt(resultStr, 10, 64)
	if err == nil {
		return nil, fmt.Errorf("decodeGnbCuUpId error in str-to-int64 convertion %T", err)
	}
	gnbCuUpId.Value = resultInt

	return gnbCuUpId, nil
}
