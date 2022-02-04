// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationHeader.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"unsafe"
)

func PerEncodeE2SmMhoIndicationHeader(indicationHeader *e2sm_mho.E2SmMhoIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmMhoIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationHeader, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerEncodeE2SmMhoIndicationHeader(indicationHeader *e2sm_mho.E2SmMhoIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmMhoIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationHeader, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmMhoIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmMhoIndicationHeader(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationHeader((*C.E2SM_MHO_IndicationHeader_t)(unsafePtr))
}

func XerDecodeE2SmMhoIndicationHeader(bytes []byte) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationHeader((*C.E2SM_MHO_IndicationHeader_t)(unsafePtr))
}

func newE2SmMhoIndicationHeader(indicationHeader *e2sm_mho.E2SmMhoIndicationHeader) (*C.E2SM_MHO_IndicationHeader_t, error) {
	var pr C.E2SM_MHO_IndicationHeader_PR
	choiceC := [8]byte{}

	switch choice := indicationHeader.E2SmMhoIndicationHeader.(type) {
	case *e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1:
		pr = C.E2SM_MHO_IndicationHeader_PR_indicationHeader_Format1

		im, err := newE2SmMhoIndicationHeaderFormat1(choice.IndicationHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoIndicationHeader() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoIndicationHeader() %T not yet implemented", choice)
	}

	indicationHeaderC := C.E2SM_MHO_IndicationHeader_t{
		present: pr,
		choice:  choiceC,
	}

	return &indicationHeaderC, nil
}

func decodeE2SmMhoIndicationHeader(indicationHeaderC *C.E2SM_MHO_IndicationHeader_t) (*e2sm_mho.E2SmMhoIndicationHeader, error) {
	indicationHeader := new(e2sm_mho.E2SmMhoIndicationHeader)

	switch indicationHeaderC.present {
	case C.E2SM_MHO_IndicationHeader_PR_indicationHeader_Format1:
		indicationHeaderFormat1, err := decodeE2SmMhoIndicationHeaderFormat1Bytes(indicationHeaderC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoIndicationHeader() %s", err.Error())
		}

		indicationHeader.E2SmMhoIndicationHeader = &e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: indicationHeaderFormat1,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoIndicationHeader() %v not yet implemented", indicationHeaderC.present)
	}

	return indicationHeader, nil
}
