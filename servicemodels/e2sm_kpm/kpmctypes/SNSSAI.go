// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "SNSSAI.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeSnssai(snssai *e2sm_kpm_ies.Snssai) ([]byte, error) {

	snssaiCP := newSnssai(snssai)

	bytes, err := encodeXer(&C.asn_DEF_SNSSAI, unsafe.Pointer(snssaiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSnssai() %s", err.Error())
	}
	return bytes, nil
}

func newSnssai(snssai *e2sm_kpm_ies.Snssai) *C.SNSSAI_t {

	sst := newOctetString(string(snssai.SSt))
	sd := newOctetString(string(snssai.SD))

	snssaiC := C.SNSSAI_t{
		sST: *sst,
		sD:  sd,
	}

	return &snssaiC
}

func decodeSnssai(snssaiC *C.SNSSAI_t) *e2sm_kpm_ies.Snssai {
	snssai := new(e2sm_kpm_ies.Snssai)

	sst := decodeOctetString(&snssaiC.sST)
	snssai.SSt = []byte(sst)

	sd := decodeOctetString(snssaiC.sD)
	snssai.SD = []byte(sd)

	return snssai
}
