// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationHeader.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoIndicationHeader(e2SmMhoIndicationHeader *e2sm_mho.E2SmMhoIndicationHeader) ([]byte, error) {
	e2SmMhoIndicationHeaderCP, err := newE2SmMhoIndicationHeader(e2SmMhoIndicationHeader)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationHeader, unsafe.Pointer(e2SmMhoIndicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoIndicationHeader(e2SmMhoIndicationHeader *e2sm_mho.E2SmMhoIndicationHeader) ([]byte, error) {
	e2SmMhoIndicationHeaderCP, err := newE2SmMhoIndicationHeader(e2SmMhoIndicationHeader)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationHeader, unsafe.Pointer(e2SmMhoIndicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoIndicationHeader(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoIndicationHeader((*C.E2SM_MHO_IndicationHeader_t)(unsafePtr))
}

func perDecodeE2SmMhoIndicationHeader(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationHeader((*C.E2SM_MHO_IndicationHeader_t)(unsafePtr))
}

func newE2SmMhoIndicationHeader(e2SmMhoIndicationHeader *e2sm_mho.E2SmMhoIndicationHeader) (*C.E2SM_MHO_IndicationHeader_t, error) {

	var pr C.E2SM_MHO_IndicationHeader_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}                  //ToDo - Check if number of bytes is sufficient

	switch choice := e2SmMhoIndicationHeader.E2SmMhoIndicationHeader.(type) {
	case *e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1:
		pr = C.E2SM_MHO_IndicationHeader_PR_indicationHeader_Format1 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoIndicationHeaderFormat1(choice.IndicationHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoIndicationHeaderFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoIndicationHeader() %T not yet implemented", choice)
	}

	e2SmMhoIndicationHeaderC := C.E2SM_MHO_IndicationHeader_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmMhoIndicationHeaderC, nil
}

func decodeE2SmMhoIndicationHeader(e2SmMhoIndicationHeaderC *C.E2SM_MHO_IndicationHeader_t) (*e2sm_mho.E2SmMhoIndicationHeader, error) {

	e2SmMhoIndicationHeader := new(e2sm_mho.E2SmMhoIndicationHeader)

	switch e2SmMhoIndicationHeaderC.present {
	case C.E2SM_MHO_IndicationHeader_PR_indicationHeader_Format1:
		e2SmMhoIndicationHeaderstructC, err := decodeE2SmMhoIndicationHeaderFormat1Bytes(e2SmMhoIndicationHeaderC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoIndicationHeaderFormat1Bytes() %s", err.Error())
		}
		e2SmMhoIndicationHeader.E2SmMhoIndicationHeader = &e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: e2SmMhoIndicationHeaderstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoIndicationHeader() %v not yet implemented", e2SmMhoIndicationHeaderC.present)
	}

	return e2SmMhoIndicationHeader, nil
}

func decodeE2SmMhoIndicationHeaderBytes(array [8]byte) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	e2SmMhoIndicationHeaderC := (*C.E2SM_MHO_IndicationHeader_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoIndicationHeader(e2SmMhoIndicationHeaderC)
}
