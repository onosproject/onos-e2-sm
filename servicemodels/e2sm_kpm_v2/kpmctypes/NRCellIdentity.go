// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCellIdentity.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeNrcellIdentity(nrcellIdentity *e2sm_kpm_v2.NrcellIdentity) ([]byte, error) {
	nrcellIdentityCP, err := newNrcellIdentity(nrcellIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcellIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrcellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNrcellIdentity(nrcellIdentity *e2sm_kpm_v2.NrcellIdentity) ([]byte, error) {
	nrcellIdentityCP, err := newNrcellIdentity(nrcellIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcellIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrcellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNrcellIdentity(bytes []byte) (*e2sm_kpm_v2.NrcellIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NRCellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNrcellIdentity((*C.NRCellIdentity_t)(unsafePtr))
}

func perDecodeNrcellIdentity(bytes []byte) (*e2sm_kpm_v2.NrcellIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NRCellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNrcellIdentity((*C.NRCellIdentity_t)(unsafePtr))
}

func newNrcellIdentity(nrcellIdentity *e2sm_kpm_v2.NrcellIdentity) (*C.NRCellIdentity_t, error) {

	nrcellIdentityC, err := newBitString(nrcellIdentity.Value)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	return nrcellIdentityC, nil
}

func decodeNrcellIdentity(nrcellIdentityC *C.NRCellIdentity_t) (*e2sm_kpm_v2.NrcellIdentity, error) {

	nrcellIdentity := new(e2sm_kpm_v2.NrcellIdentity)
	nrcellIdentityValue, err := decodeBitString(nrcellIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}
	nrcellIdentity.Value = nrcellIdentityValue

	return nrcellIdentity, nil
}

func decodeNrcellIdentityBytes(array [8]byte) (*e2sm_kpm_v2.NrcellIdentity, error) {
	nrcellIdentityC := (*C.NRCellIdentity_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeNrcellIdentity(nrcellIdentityC)
}
