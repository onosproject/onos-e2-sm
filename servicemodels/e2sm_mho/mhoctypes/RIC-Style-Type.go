// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Type.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRicStyleType(ricStyleType *e2sm_mho.RicStyleType) ([]byte, error) {
	ricStyleTypeCP, err := newRicStyleType(ricStyleType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Type, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleType(ricStyleType *e2sm_mho.RicStyleType) ([]byte, error) {
	ricStyleTypeCP, err := newRicStyleType(ricStyleType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Type, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleType(bytes []byte) (*e2sm_mho.RicStyleType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_t)(unsafePtr))
}

func perDecodeRicStyleType(bytes []byte) (*e2sm_mho.RicStyleType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_t)(unsafePtr))
}

func newRicStyleType(ricStyleType *e2sm_mho.RicStyleType) (*C.RIC_Style_Type_t, error) {

	ricStyleTypeC := C.long(ricStyleType.Value)

	return &ricStyleTypeC, nil
}

func decodeRicStyleType(ricStyleTypeC *C.RIC_Style_Type_t) (*e2sm_mho.RicStyleType, error) {

	ricStyleType := e2sm_mho.RicStyleType{
		Value: int32(*ricStyleTypeC),
	}

	return &ricStyleType, nil
}

//func decodeRicStyleTypeBytes(array [8]byte) (*e2sm_mho.RicStyleType, error) {
//	ricStyleTypeC := (*C.RIC_Style_Type_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
//
//	return decodeRicStyleType(ricStyleTypeC)
//}
