// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-ControlMessage-Format1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1 *e2sm_mho.E2SmMhoControlMessageFormat1) ([]byte, error) {
	e2SmMhoControlMessageFormat1CP, err := newE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_ControlMessage_Format1, unsafe.Pointer(e2SmMhoControlMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1 *e2sm_mho.E2SmMhoControlMessageFormat1) ([]byte, error) {
	e2SmMhoControlMessageFormat1CP, err := newE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlMessageFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_ControlMessage_Format1, unsafe.Pointer(e2SmMhoControlMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoControlMessageFormat1(bytes []byte) (*e2sm_mho.E2SmMhoControlMessageFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_ControlMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoControlMessageFormat1((*C.E2SM_MHO_ControlMessage_Format1_t)(unsafePtr))
}

func perDecodeE2SmMhoControlMessageFormat1(bytes []byte) (*e2sm_mho.E2SmMhoControlMessageFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_ControlMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoControlMessageFormat1((*C.E2SM_MHO_ControlMessage_Format1_t)(unsafePtr))
}

func newE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1 *e2sm_mho.E2SmMhoControlMessageFormat1) (*C.E2SM_MHO_ControlMessage_Format1_t, error) {

	var err error
	e2SmMhoControlMessageFormat1C := C.E2SM_MHO_ControlMessage_Format1_t{}

	servingCgiC, err := newCellGlobalID(e2SmMhoControlMessageFormat1.ServingCgi)
	if err != nil {
		return nil, fmt.Errorf("newCellGlobalID() %s", err.Error())
	}

	uedIDC, err := newUeIdentity(e2SmMhoControlMessageFormat1.UedId)
	if err != nil {
		return nil, fmt.Errorf("newUeIdentity() %s", err.Error())
	}

	targetCgiC, err := newCellGlobalID(e2SmMhoControlMessageFormat1.TargetCgi)
	if err != nil {
		return nil, fmt.Errorf("newCellGlobalID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoControlMessageFormat1C.serving_cgi = servingCgiC
	e2SmMhoControlMessageFormat1C.uedID = uedIdC
	e2SmMhoControlMessageFormat1C.target_cgi = targetCgiC

	return &e2SmMhoControlMessageFormat1C, nil
}

func decodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1C *C.E2SM_MHO_ControlMessage_Format1_t) (*e2sm_mho.E2SmMhoControlMessageFormat1, error) {

	var err error
	e2SmMhoControlMessageFormat1 := e2sm_mho.E2SmMhoControlMessageFormat1{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ServingCgi: servingCgi,
		//UedId: uedId,
		//TargetCgi: targetCgi,

	}

	e2SmMhoControlMessageFormat1.ServingCgi, err = decodeCellGlobalID(e2SmMhoControlMessageFormat1C.serving_cgi)
	if err != nil {
		return nil, fmt.Errorf("decodeCellGlobalID() %s", err.Error())
	}

	e2SmMhoControlMessageFormat1.UedId, err = decodeUeIdentity(e2SmMhoControlMessageFormat1C.uedID)
	if err != nil {
		return nil, fmt.Errorf("decodeUeIdentity() %s", err.Error())
	}

	e2SmMhoControlMessageFormat1.TargetCgi, err = decodeCellGlobalID(e2SmMhoControlMessageFormat1C.target_cgi)
	if err != nil {
		return nil, fmt.Errorf("decodeCellGlobalID() %s", err.Error())
	}

	return &e2SmMhoControlMessageFormat1, nil
}

func decodeE2SmMhoControlMessageFormat1Bytes(array [8]byte) (*e2sm_mho.E2SmMhoControlMessageFormat1, error) {
	e2SmMhoControlMessageFormat1C := (*C.E2SM_MHO_ControlMessage_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1C)
}
