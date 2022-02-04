// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationHeader-Format1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func XerEncodeE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1 *e2sm_mho.E2SmMhoIndicationHeaderFormat1) ([]byte, error) {
	e2SmMhoIndicationHeaderFormat1CP, err := newE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationHeader_Format1, unsafe.Pointer(e2SmMhoIndicationHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1 *e2sm_mho.E2SmMhoIndicationHeaderFormat1) ([]byte, error) {
	e2SmMhoIndicationHeaderFormat1CP, err := newE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationHeader_Format1, unsafe.Pointer(e2SmMhoIndicationHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmMhoIndicationHeaderFormat1(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeaderFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoIndicationHeaderFormat1((*C.E2SM_MHO_IndicationHeader_Format1_t)(unsafePtr))
}

func PerDecodeE2SmMhoIndicationHeaderFormat1(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeaderFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationHeaderFormat1((*C.E2SM_MHO_IndicationHeader_Format1_t)(unsafePtr))
}

func newE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1 *e2sm_mho.E2SmMhoIndicationHeaderFormat1) (*C.E2SM_MHO_IndicationHeader_Format1_t, error) {

	var err error
	e2SmMhoIndicationHeaderFormat1C := C.E2SM_MHO_IndicationHeader_Format1_t{}

	cgiC, err := newCellGlobalID(e2SmMhoIndicationHeaderFormat1.Cgi)
	if err != nil {
		return nil, fmt.Errorf("newCellGlobalID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoIndicationHeaderFormat1C.cgi = *cgiC

	return &e2SmMhoIndicationHeaderFormat1C, nil
}

func decodeE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1C *C.E2SM_MHO_IndicationHeader_Format1_t) (*e2sm_mho.E2SmMhoIndicationHeaderFormat1, error) {

	var err error
	e2SmMhoIndicationHeaderFormat1 := e2sm_mho.E2SmMhoIndicationHeaderFormat1{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//Cgi: cgi,

	}

	e2SmMhoIndicationHeaderFormat1.Cgi, err = decodeCellGlobalID(&e2SmMhoIndicationHeaderFormat1C.cgi)
	if err != nil {
		return nil, fmt.Errorf("decodeCellGlobalID() %s", err.Error())
	}

	return &e2SmMhoIndicationHeaderFormat1, nil
}

func decodeE2SmMhoIndicationHeaderFormat1Bytes(array [8]byte) (*e2sm_mho.E2SmMhoIndicationHeaderFormat1, error) {
	e2SmMhoIndicationHeaderFormat1C := (*C.E2SM_MHO_IndicationHeader_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoIndicationHeaderFormat1(e2SmMhoIndicationHeaderFormat1C)
}
