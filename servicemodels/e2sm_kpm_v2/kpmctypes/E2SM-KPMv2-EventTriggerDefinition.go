// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-EventTriggerDefinition.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

//func XerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_v2.E2SmKpmEventTriggerDefinition) ([]byte, error) {
//	e2SmKpmEventTriggerDefinitionCP, err := newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeE2SmKpmEventTriggerDefinition() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_EventTriggerDefinition, unsafe.Pointer(e2SmKpmEventTriggerDefinitionCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeE2SmKpmEventTriggerDefinition() %s", err.Error())
//	}
//	return bytes, nil
//}

func PerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_v2.E2SmKpmEventTriggerDefinition) ([]byte, error) {
	e2SmKpmEventTriggerDefinitionCP, err := newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmEventTriggerDefinition() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_EventTriggerDefinition, unsafe.Pointer(e2SmKpmEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmEventTriggerDefinition() %s", err.Error())
	}
	return bytes, nil
}

//func XerDecodeE2SmKpmEventTriggerDefinition(bytes []byte) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinition, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_EventTriggerDefinition)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeE2SmKpmEventTriggerDefinition((*C.E2SM_KPMv2_EventTriggerDefinition_t)(unsafePtr))
//}

func PerDecodeE2SmKpmEventTriggerDefinition(bytes []byte) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinition((*C.E2SM_KPMv2_EventTriggerDefinition_t)(unsafePtr))
}

func newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_v2.E2SmKpmEventTriggerDefinition) (*C.E2SM_KPMv2_EventTriggerDefinition_t, error) {

	var pr C.E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition.(type) {
	case *e2sm_kpm_v2.E2SmKpmEventTriggerDefinition_EventDefinitionFormat1:
		pr = C.E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR_eventDefinition_Format1

		im, err := newE2SmKpmEventTriggerDefinitionFormat1(choice.EventDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmEventTriggerDefinition() %T not yet implemented", choice)
	}

	e2SmKpmEventTriggerDefinitionFormatC := C.struct_E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats{
		present: pr,
		choice:  choiceC,
	}

	e2SmKpmEventTriggerDefinitionC := C.E2SM_KPMv2_EventTriggerDefinition_t{
		eventDefinition_formats: e2SmKpmEventTriggerDefinitionFormatC,
	}

	return &e2SmKpmEventTriggerDefinitionC, nil
}

func decodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinitionC *C.E2SM_KPMv2_EventTriggerDefinition_t) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinition, error) {

	e2SmKpmEventTriggerDefinition := new(e2sm_kpm_v2.E2SmKpmEventTriggerDefinition)

	switch e2SmKpmEventTriggerDefinitionC.eventDefinition_formats.present {
	case C.E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR_eventDefinition_Format1:
		eventTriggerDefinitionFormat1, err := decodeE2SmKpmEventTriggerDefinitionFormat1Bytes(e2SmKpmEventTriggerDefinitionC.eventDefinition_formats.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
		}
		e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition = &e2sm_kpm_v2.E2SmKpmEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: eventTriggerDefinitionFormat1,
		}
	case C.E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2SmKpmEventTriggerDefinition() An empty EventTriggerDefinition-Format has been sent %v", e2SmKpmEventTriggerDefinitionC.eventDefinition_formats.present)
	default:
		return nil, fmt.Errorf("decodeE2SmKpmEventTriggerDefinition() %v not yet implemented", e2SmKpmEventTriggerDefinitionC.eventDefinition_formats.present)
	}

	return e2SmKpmEventTriggerDefinition, nil
}
