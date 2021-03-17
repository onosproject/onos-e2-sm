// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "QCI.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeQci(qci *e2sm_kpm_v2.Qci) ([]byte, error) {
	qciCP, err := newQci(qci)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeQci() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_QCI, unsafe.Pointer(qciCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeQci() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeQci(qci *e2sm_kpm_v2.Qci) ([]byte, error) {
	qciCP, err := newQci(qci)
	if err != nil {
		return nil, fmt.Errorf("perEncodeQci() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_QCI, unsafe.Pointer(qciCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeQci() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeQci(bytes []byte) (*e2sm_kpm_v2.Qci, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_QCI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeQci((*C.QCI_t)(unsafePtr))
}

func perDecodeQci(bytes []byte) (*e2sm_kpm_v2.Qci, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_QCI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeQci((*C.QCI_t)(unsafePtr))
}

func newQci(qci *e2sm_kpm_v2.Qci) (*C.QCI_t, error) {

	qciC := C.long(qci.Value)

	return &qciC, nil
}

func decodeQci(qciC *C.QCI_t) (*e2sm_kpm_v2.Qci, error) {

	qci := e2sm_kpm_v2.Qci{
		Value: int32(*qciC),
	}

	return &qci, nil
}
