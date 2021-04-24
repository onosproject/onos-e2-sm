// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCellIdentity.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeNrcellIDentity(nrcellIDentity *e2sm_mho.NrcellIdentity) ([]byte, error) {
	nrcellIDentityCP, err := newNrcellIDentity(nrcellIDentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcellIDentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrcellIDentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcellIDentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNrcellIDentity(nrcellIDentity *e2sm_mho.NrcellIdentity) ([]byte, error) {
	nrcellIDentityCP, err := newNrcellIDentity(nrcellIDentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcellIDentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCellIdentity, unsafe.Pointer(nrcellIDentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcellIDentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNrcellIDentity(bytes []byte) (*e2sm_mho.NrcellIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NRCellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNrcellIDentity((*C.NRCellIdentity_t)(unsafePtr))
}

func perDecodeNrcellIDentity(bytes []byte) (*e2sm_mho.NrcellIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NRCellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNrcellIDentity((*C.NRCellIdentity_t)(unsafePtr))
}

func newNrcellIDentity(nrcellIDentity *e2sm_mho.NrcellIdentity) (*C.NRCellIdentity_t, error) {

	nrcellIDentityC, err := newBitString(nrcellIDentity.Value)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	return nrcellIDentityC, nil
}

func decodeNrcellIDentity(nrcellIDentityC *C.NRCellIdentity_t) (*e2sm_mho.NrcellIdentity, error) {

	nrcellIDentity := new(e2sm_mho.NrcellIdentity)
	nrcellIDentityValue, err := decodeBitString(nrcellIDentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}
	nrcellIDentity.Value = nrcellIDentityValue

	return nrcellIDentity, nil
}

func decodeNrcellIDentityBytes(array [8]byte) (*e2sm_mho.NrcellIdentity, error) {
	nrcellIDentityC := (*C.NRCellIdentity_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeNrcellIDentity(nrcellIDentityC)
}
