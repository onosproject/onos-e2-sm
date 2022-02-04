// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "SNSSAI-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeSnssai(snssai *e2sm_kpm_v2.Snssai) ([]byte, error) {
	snssaiCP, err := newSnssai(snssai)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSnssai() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_SNSSAI_KPMv2, unsafe.Pointer(snssaiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeSnssai() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeSnssai(snssai *e2sm_kpm_v2.Snssai) ([]byte, error) {
	snssaiCP, err := newSnssai(snssai)
	if err != nil {
		return nil, fmt.Errorf("perEncodeSnssai() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_SNSSAI_KPMv2, unsafe.Pointer(snssaiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeSnssai() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeSnssai(bytes []byte) (*e2sm_kpm_v2.Snssai, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_SNSSAI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeSnssai((*C.SNSSAI_KPMv2_t)(unsafePtr))
}

func perDecodeSnssai(bytes []byte) (*e2sm_kpm_v2.Snssai, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_SNSSAI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeSnssai((*C.SNSSAI_KPMv2_t)(unsafePtr))
}

func newSnssai(snssai *e2sm_kpm_v2.Snssai) (*C.SNSSAI_KPMv2_t, error) {

	sStC, err := newOctetString(string(snssai.SSt))
	if err != nil {
		return nil, err
	}

	snssaiC := C.SNSSAI_KPMv2_t{
		sST: *sStC,
		//sD:  sDC,
	}

	if snssai.SD != nil {
		snssaiC.sD, err = newOctetString(string(snssai.SD))
		if err != nil {
			return nil, err
		}
	}

	return &snssaiC, nil
}

func decodeSnssai(snssaiC *C.SNSSAI_KPMv2_t) (*e2sm_kpm_v2.Snssai, error) {

	sSt, err := decodeOctetString(&snssaiC.sST)
	if err != nil {
		return nil, err
	}

	snssai := e2sm_kpm_v2.Snssai{
		SSt: []byte(sSt),
		//SD:  []byte(sD),
	}

	if snssaiC.sD != nil {
		sd, err := decodeOctetString(snssaiC.sD)
		if err != nil {
			return nil, err
		}
		snssai.SD = []byte(sd)
	}

	return &snssai, nil
}
