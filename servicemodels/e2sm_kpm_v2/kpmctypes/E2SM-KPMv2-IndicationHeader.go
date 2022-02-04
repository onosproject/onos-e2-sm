// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-IndicationHeader.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func XerEncodeE2SmKpmIndicationHeader(e2SmKpmIndicationHeader *e2sm_kpm_v2.E2SmKpmIndicationHeader) ([]byte, error) {
	e2SmKpmIndicationHeaderCP, err := newE2SmKpmIndicationHeader(e2SmKpmIndicationHeader)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_IndicationHeader, unsafe.Pointer(e2SmKpmIndicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmIndicationHeader(e2SmKpmIndicationHeader *e2sm_kpm_v2.E2SmKpmIndicationHeader) ([]byte, error) {
	e2SmKpmIndicationHeaderCP, err := newE2SmKpmIndicationHeader(e2SmKpmIndicationHeader)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_IndicationHeader, unsafe.Pointer(e2SmKpmIndicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmIndicationHeader(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmIndicationHeader((*C.E2SM_KPMv2_IndicationHeader_t)(unsafePtr))
}

func PerDecodeE2SmKpmIndicationHeader(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_IndicationHeader)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationHeader((*C.E2SM_KPMv2_IndicationHeader_t)(unsafePtr))
}

func newE2SmKpmIndicationHeader(e2SmKpmIndicationHeader *e2sm_kpm_v2.E2SmKpmIndicationHeader) (*C.E2SM_KPMv2_IndicationHeader_t, error) {

	var pr C.E2SM_KPMv2_IndicationHeader__indicationHeader_formats_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmIndicationHeader.E2SmKpmIndicationHeader.(type) {
	case *e2sm_kpm_v2.E2SmKpmIndicationHeader_IndicationHeaderFormat1:
		pr = C.E2SM_KPMv2_IndicationHeader__indicationHeader_formats_PR_indicationHeader_Format1

		im, err := newE2SmKpmIndicationHeaderFormat1(choice.IndicationHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmIndicationHeaderFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmIndicationHeader() %T not yet implemented", choice)
	}

	e2SmKpmEventTriggerDefinitionFormatC := C.struct_E2SM_KPMv2_IndicationHeader__indicationHeader_formats{
		present: pr,
		choice:  choiceC,
	}

	e2SmKpmIndicationHeaderC := C.E2SM_KPMv2_IndicationHeader_t{
		indicationHeader_formats: e2SmKpmEventTriggerDefinitionFormatC,
	}

	return &e2SmKpmIndicationHeaderC, nil
}

func decodeE2SmKpmIndicationHeader(e2SmKpmIndicationHeaderC *C.E2SM_KPMv2_IndicationHeader_t) (*e2sm_kpm_v2.E2SmKpmIndicationHeader, error) {

	e2SmKpmIndicationHeader := new(e2sm_kpm_v2.E2SmKpmIndicationHeader)

	switch e2SmKpmIndicationHeaderC.indicationHeader_formats.present {
	case C.E2SM_KPMv2_IndicationHeader__indicationHeader_formats_PR_indicationHeader_Format1:
		indicationHeaderFormat1, err := decodeE2SmKpmIndicationHeaderFormat1Bytes(e2SmKpmIndicationHeaderC.indicationHeader_formats.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationHeaderFormat1Bytes() %s", err.Error())
		}
		e2SmKpmIndicationHeader.E2SmKpmIndicationHeader = &e2sm_kpm_v2.E2SmKpmIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: indicationHeaderFormat1,
		}
	case C.E2SM_KPMv2_IndicationHeader__indicationHeader_formats_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2SmKpmIndicationHeader() An empty IndicationHeader-Format has been sent %v", e2SmKpmIndicationHeaderC.indicationHeader_formats.present)
	default:
		return nil, fmt.Errorf("decodeE2SmKpmIndicationHeader() %v not yet implemented", e2SmKpmIndicationHeaderC.indicationHeader_formats.present)
	}

	return e2SmKpmIndicationHeader, nil
}
