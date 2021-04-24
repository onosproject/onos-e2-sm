// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EUTRACellIdentity.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeEutracellIDentity(eutracellIDentity *e2sm_mho.EutracellIdentity) ([]byte, error) {
	eutracellIDentityCP, err := newEutracellIDentity(eutracellIDentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracellIDentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_EUTRACellIdentity, unsafe.Pointer(eutracellIDentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracellIDentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEutracellIDentity(eutracellIDentity *e2sm_mho.EutracellIdentity) ([]byte, error) {
	eutracellIDentityCP, err := newEutracellIDentity(eutracellIDentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracellIDentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_EUTRACellIdentity, unsafe.Pointer(eutracellIDentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracellIDentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEutracellIDentity(bytes []byte) (*e2sm_mho.EutracellIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_EUTRACellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEutracellIDentity((*C.EUTRACellIdentity_t)(unsafePtr))
}

func perDecodeEutracellIDentity(bytes []byte) (*e2sm_mho.EutracellIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_EUTRACellIdentity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEutracellIDentity((*C.EUTRACellIdentity_t)(unsafePtr))
}

func newEutracellIDentity(eutracellIDentity *e2sm_mho.EutracellIdentity) (*C.EUTRACellIdentity_t, error) {

	eutracellIDentityC, err := newBitString(eutracellIDentity.Value)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	return &eutracellIDentityC, nil
}

func decodeEutracellIDentity(eutracellIDentityC *C.EUTRACellIdentity_t) (*e2sm_mho.EutracellIdentity, error) {

	eutracellIDentity := new(e2sm_mho.EutracellIdentity)
	eutracellIDentityValue, err := decodeBitString(eutracellIDentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}
	eutracellIDentity.Value = eutracellIDentityValue

	return &eutracellIDentity, nil
}

func decodeEutracellIDentityBytes(array [8]byte) (*e2sm_mho.EutracellIdentity, error) {
	eutracellIDentityC := (*C.EUTRACellIdentity_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeEutracellIDentity(eutracellIDentityC)
}
