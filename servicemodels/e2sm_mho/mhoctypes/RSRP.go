// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RSRP.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRsrp(rsrp *e2sm_mho.Rsrp) ([]byte, error) {
	rsrpCP, err := newRsrp(rsrp)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRsrp() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RSRP, unsafe.Pointer(rsrpCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRsrp() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRsrp(rsrp *e2sm_mho.Rsrp) ([]byte, error) {
	rsrpCP, err := newRsrp(rsrp)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRsrp() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RSRP, unsafe.Pointer(rsrpCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRsrp() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRsrp(bytes []byte) (*e2sm_mho.Rsrp, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RSRP)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRsrp((*C.RSRP_t)(unsafePtr))
}

func perDecodeRsrp(bytes []byte) (*e2sm_mho.Rsrp, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RSRP)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRsrp((*C.RSRP_t)(unsafePtr))
}

func newRsrp(rsrp *e2sm_mho.Rsrp) (*C.RSRP_t, error) {

	rsrpC := C.long(rsrp.Value)

	return &rsrpC, nil
}

func decodeRsrp(rsrpC *C.RSRP_t) (*e2sm_mho.Rsrp, error) {

	rsrp := e2sm_mho.Rsrp{
		Value: int32(*rsrpC),
	}

	return &rsrp, nil
}

func decodeRsrpBytes(array [8]byte) (*e2sm_mho.Rsrp, error) {
	rsrpC := (*C.RSRP_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRsrp(rsrpC)
}
