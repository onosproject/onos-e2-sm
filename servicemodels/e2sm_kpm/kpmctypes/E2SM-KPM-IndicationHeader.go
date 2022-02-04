// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationHeader.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func PerEncodeE2SmKpmIndicationHeader(indicationHeader *e2sm_kpm_ies.E2SmKpmIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmKpmIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_IndicationHeader, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerEncodeE2SmKpmIndicationHeader(indicationHeader *e2sm_kpm_ies.E2SmKpmIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmKpmIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationHeader, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmKpmIndicationHeader(bytes []byte) (*e2sm_kpm_ies.E2SmKpmIndicationHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationHeader((*C.E2SM_KPM_IndicationHeader_t)(unsafePtr))
}

func XerDecodeE2SmKpmIndicationHeader(bytes []byte) (*e2sm_kpm_ies.E2SmKpmIndicationHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationHeader((*C.E2SM_KPM_IndicationHeader_t)(unsafePtr))
}

func newE2SmKpmIndicationHeader(indicationHeader *e2sm_kpm_ies.E2SmKpmIndicationHeader) (*C.E2SM_KPM_IndicationHeader_t, error) {
	var pr C.E2SM_KPM_IndicationHeader_PR
	choiceC := [8]byte{}

	switch choice := indicationHeader.E2SmKpmIndicationHeader.(type) {
	case *e2sm_kpm_ies.E2SmKpmIndicationHeader_IndicationHeaderFormat1:
		pr = C.E2SM_KPM_IndicationHeader_PR_indicationHeader_Format1

		im, err := newE2SmKpmIndicationHeaderFormat1(choice.IndicationHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmIndicationHeader() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmIndicationHeader() %T not yet implemented", choice)
	}

	indicationHeaderC := C.E2SM_KPM_IndicationHeader_t{
		present: pr,
		choice:  choiceC,
	}

	return &indicationHeaderC, nil
}

func decodeE2SmKpmIndicationHeader(indicationHeaderC *C.E2SM_KPM_IndicationHeader_t) (*e2sm_kpm_ies.E2SmKpmIndicationHeader, error) {
	indicationHeader := new(e2sm_kpm_ies.E2SmKpmIndicationHeader)

	switch indicationHeaderC.present {
	case C.E2SM_KPM_IndicationHeader_PR_indicationHeader_Format1:
		indicationHeaderFormat1, err := decodeE2SmKpmIndicationHeaderFormat1Bytes(indicationHeaderC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationHeader() %s", err.Error())
		}

		indicationHeader.E2SmKpmIndicationHeader = &e2sm_kpm_ies.E2SmKpmIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: indicationHeaderFormat1,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmKpmIndicationHeader() %v not yet implemented", indicationHeaderC.present)
	}

	return indicationHeader, nil
}
