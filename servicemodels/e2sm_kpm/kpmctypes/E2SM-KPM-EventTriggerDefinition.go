// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-EventTriggerDefinition.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func PerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_ies.E2SmKpmEventTriggerDefinition) ([]byte, error) {

	e2SmKpmEventTriggerDefinitionCP, err := newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition, unsafe.Pointer(e2SmKpmEventTriggerDefinitionCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_ies.E2SmKpmEventTriggerDefinition) ([]byte, error) {

	e2SmKpmEventTriggerDefinitionCP, err := newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition, unsafe.Pointer(e2SmKpmEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmKpmEventTriggerDefinition(bytes []byte) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinition((*C.E2SM_KPM_EventTriggerDefinition_t)(unsafePtr))
}

func XerDecodeE2SmKpmEventTriggerDefinition(bytes []byte) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinition((*C.E2SM_KPM_EventTriggerDefinition_t)(unsafePtr))
}

func newE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition *e2sm_kpm_ies.E2SmKpmEventTriggerDefinition) (*C.E2SM_KPM_EventTriggerDefinition_t, error) {
	var present C.E2SM_KPM_EventTriggerDefinition_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition.(type) {
	case *e2sm_kpm_ies.E2SmKpmEventTriggerDefinition_EventDefinitionFormat1:
		present = C.E2SM_KPM_EventTriggerDefinition_PR_eventDefinition_Format1

		im, err := newE2SmKpmEventTriggerDefinitionFormat1(choice.EventDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmEventTriggerDefinition() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmEventTriggerDefinition() %T not yet implemented", choice)
	}

	e2SmKpmEventTriggerDefinitionC := C.E2SM_KPM_EventTriggerDefinition_t{
		present: present,
		choice:  choiceC,
	}

	return &e2SmKpmEventTriggerDefinitionC, nil
}

func decodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinitionC *C.E2SM_KPM_EventTriggerDefinition_t) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinition, error) {
	e2SmKpmEventTriggerDefinition := new(e2sm_kpm_ies.E2SmKpmEventTriggerDefinition)

	switch e2SmKpmEventTriggerDefinitionC.present {
	case C.E2SM_KPM_EventTriggerDefinition_PR_eventDefinition_Format1:
		e2SmKpmEventTriggerDefinitionFormat1, err := decodeE2SmKpmEventTriggerDefinitionFormat1Bytes(e2SmKpmEventTriggerDefinitionC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationMessage() %s", err.Error())
		}

		e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition = &e2sm_kpm_ies.E2SmKpmEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: e2SmKpmEventTriggerDefinitionFormat1,
		}

	default:
		return nil, fmt.Errorf("decodeE2SmKpmEventTriggerDefinition() %v not yet implemented", e2SmKpmEventTriggerDefinitionC.present)
	}

	return e2SmKpmEventTriggerDefinition, nil
}
