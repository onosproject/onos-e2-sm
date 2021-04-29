// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-ControlHeader.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func XerEncodeE2SmMhoControlHeader(e2SmMhoControlHeader *e2sm_mho.E2SmMhoControlHeader) ([]byte, error) {
	e2SmMhoControlHeaderCP, err := newE2SmMhoControlHeader(e2SmMhoControlHeader)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_ControlHeader, unsafe.Pointer(e2SmMhoControlHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmMhoControlHeader(e2SmMhoControlHeader *e2sm_mho.E2SmMhoControlHeader) ([]byte, error) {
	e2SmMhoControlHeaderCP, err := newE2SmMhoControlHeader(e2SmMhoControlHeader)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_ControlHeader, unsafe.Pointer(e2SmMhoControlHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmMhoControlHeader(bytes []byte) (*e2sm_mho.E2SmMhoControlHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_ControlHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoControlHeader((*C.E2SM_MHO_ControlHeader_t)(unsafePtr))
}

func PerDecodeE2SmMhoControlHeader(bytes []byte) (*e2sm_mho.E2SmMhoControlHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_ControlHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoControlHeader((*C.E2SM_MHO_ControlHeader_t)(unsafePtr))
}

func newE2SmMhoControlHeader(e2SmMhoControlHeader *e2sm_mho.E2SmMhoControlHeader) (*C.E2SM_MHO_ControlHeader_t, error) {

	var pr C.E2SM_MHO_ControlHeader_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}               //ToDo - Check if number of bytes is sufficient

	switch choice := e2SmMhoControlHeader.E2SmMhoControlHeader.(type) {
	case *e2sm_mho.E2SmMhoControlHeader_ControlHeaderFormat1:
		pr = C.E2SM_MHO_ControlHeader_PR_controlHeader_Format1 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoControlHeaderFormat1(choice.ControlHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoControlHeaderFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoControlHeader() %T not yet implemented", choice)
	}

	e2SmMhoControlHeaderC := C.E2SM_MHO_ControlHeader_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmMhoControlHeaderC, nil
}

func decodeE2SmMhoControlHeader(e2SmMhoControlHeaderC *C.E2SM_MHO_ControlHeader_t) (*e2sm_mho.E2SmMhoControlHeader, error) {

	e2SmMhoControlHeader := new(e2sm_mho.E2SmMhoControlHeader)

	switch e2SmMhoControlHeaderC.present {
	case C.E2SM_MHO_ControlHeader_PR_controlHeader_Format1:
		e2SmMhoControlHeaderstructC, err := decodeE2SmMhoControlHeaderFormat1Bytes(e2SmMhoControlHeaderC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoControlHeaderFormat1Bytes() %s", err.Error())
		}
		e2SmMhoControlHeader.E2SmMhoControlHeader = &e2sm_mho.E2SmMhoControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: e2SmMhoControlHeaderstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoControlHeader() %v not yet implemented", e2SmMhoControlHeaderC.present)
	}

	return e2SmMhoControlHeader, nil
}

func decodeE2SmMhoControlHeaderBytes(array [8]byte) (*e2sm_mho.E2SmMhoControlHeader, error) {
	e2SmMhoControlHeaderC := (*C.E2SM_MHO_ControlHeader_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoControlHeader(e2SmMhoControlHeaderC)
}
