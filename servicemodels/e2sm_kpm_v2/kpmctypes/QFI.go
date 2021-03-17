// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "QFI.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeQfi(qfi *e2sm_kpm_v2.Qfi) ([]byte, error) {
	qfiCP, err := newQfi(qfi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeQfi() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_QFI, unsafe.Pointer(qfiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeQfi() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeQfi(qfi *e2sm_kpm_v2.Qfi) ([]byte, error) {
	qfiCP, err := newQfi(qfi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeQfi() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_QFI, unsafe.Pointer(qfiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeQfi() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeQfi(bytes []byte) (*e2sm_kpm_v2.Qfi, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_QFI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeQfi((*C.QFI_t)(unsafePtr))
}

func perDecodeQfi(bytes []byte) (*e2sm_kpm_v2.Qfi, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_QFI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeQfi((*C.QFI_t)(unsafePtr))
}

func newQfi(qfi *e2sm_kpm_v2.Qfi) (*C.QFI_t, error) {

	qfiC := C.long(qfi.Value)

	return &qfiC, nil
}

func decodeQfi(qfiC *C.QFI_t) (*e2sm_kpm_v2.Qfi, error) {

	qfi := e2sm_kpm_v2.Qfi{
		Value: int32(*qfiC),
	}

	return &qfi, nil
}
