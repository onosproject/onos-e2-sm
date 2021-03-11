// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationMessage.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func XerEncodeE2SmKpmIndicationMessage(e2SmKpmIndicationMessage *e2sm_kpm_v2.E2SmKpmIndicationMessage) ([]byte, error) {
	e2SmKpmIndicationMessageCP, err := newE2SmKpmIndicationMessage(e2SmKpmIndicationMessage)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationMessage, unsafe.Pointer(e2SmKpmIndicationMessageCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmIndicationMessage(e2SmKpmIndicationMessage *e2sm_kpm_v2.E2SmKpmIndicationMessage) ([]byte, error) {
	e2SmKpmIndicationMessageCP, err := newE2SmKpmIndicationMessage(e2SmKpmIndicationMessage)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_IndicationMessage, unsafe.Pointer(e2SmKpmIndicationMessageCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmIndicationMessage(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_IndicationMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmIndicationMessage((*C.E2SM_KPM_IndicationMessage_t)(unsafePtr))
}

func PerDecodeE2SmKpmIndicationMessage(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_IndicationMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationMessage((*C.E2SM_KPM_IndicationMessage_t)(unsafePtr))
}

func newE2SmKpmIndicationMessage(e2SmKpmIndicationMessage *e2sm_kpm_v2.E2SmKpmIndicationMessage) (*C.E2SM_KPM_IndicationMessage_t, error) {

	var pr C.E2SM_KPM_IndicationMessage__indicationMessage_formats_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmIndicationMessage.E2SmKpmIndicationMessage.(type) {
	case *e2sm_kpm_v2.E2SmKpmIndicationMessage_IndicationMessageFormat1:
		pr = C.E2SM_KPM_IndicationMessage__indicationMessage_formats_PR_indicationMessage_Format1

		im, err := newE2SmKpmIndicationMessageFormat1(choice.IndicationMessageFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmIndicationMessageFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmIndicationMessage() %T not yet implemented", choice)
	}

	e2SmKpmIndicationMessageFormatC := C.struct_E2SM_KPM_IndicationMessage__indicationMessage_formats{
		present: pr,
		choice:  choiceC,
	}

	e2SmKpmIndicationMessageC := C.E2SM_KPM_IndicationMessage_t{
		indicationMessage_formats: e2SmKpmIndicationMessageFormatC,
	}

	return &e2SmKpmIndicationMessageC, nil
}

func decodeE2SmKpmIndicationMessage(e2SmKpmIndicationMessageC *C.E2SM_KPM_IndicationMessage_t) (*e2sm_kpm_v2.E2SmKpmIndicationMessage, error) {

	e2SmKpmIndicationMessage := new(e2sm_kpm_v2.E2SmKpmIndicationMessage)

	switch e2SmKpmIndicationMessageC.indicationMessage_formats.present {
	case C.E2SM_KPM_IndicationMessage__indicationMessage_formats_PR_indicationMessage_Format1:
		indicationMessageFormat1, err := decodeE2SmKpmIndicationMessageFormat1Bytes(e2SmKpmIndicationMessageC.indicationMessage_formats.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationMessageFormat1Bytes() %s", err.Error())
		}
		e2SmKpmIndicationMessage.E2SmKpmIndicationMessage = &e2sm_kpm_v2.E2SmKpmIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: indicationMessageFormat1,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmKpmIndicationMessage() %v not yet implemented", e2SmKpmIndicationMessageC.indicationMessage_formats.present)
	}

	return e2SmKpmIndicationMessage, nil
}
