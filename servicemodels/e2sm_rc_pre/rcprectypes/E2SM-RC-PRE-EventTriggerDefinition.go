// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-EventTriggerDefinition.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition) ([]byte, error) {

	E2SmRcPreEventTriggerDefinitionCP, err := newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition) ([]byte, error) {

	E2SmRcPreEventTriggerDefinitionCP, err := newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreEventTriggerDefinition(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinition((*C.E2SM_RC_PRE_EventTriggerDefinition_t)(unsafePtr))
}

func XerDecodeE2SmRcPreEventTriggerDefinition(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinition((*C.E2SM_RC_PRE_EventTriggerDefinition_t)(unsafePtr))
}

func newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition) (*C.E2SM_RC_PRE_EventTriggerDefinition_t, error) {
	var present C.E2SM_RC_PRE_EventTriggerDefinition_PR
	choiceC := [8]byte{}

	switch choice := E2SmRcPreEventTriggerDefinition.E2SmRcPreEventTriggerDefinition.(type) {
	case *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1:
		present = C.E2SM_RC_PRE_EventTriggerDefinition_PR_eventDefinition_Format1

		im, err := newE2SmRcPreEventTriggerDefinitionFormat1(choice.EventDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinition() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinition() %T not yet implemented", choice)
	}

	E2SmRcPreEventTriggerDefinitionC := C.E2SM_RC_PRE_EventTriggerDefinition_t{
		present: present,
		choice:  choiceC,
	}

	return &E2SmRcPreEventTriggerDefinitionC, nil
}

func decodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinitionC *C.E2SM_RC_PRE_EventTriggerDefinition_t) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition, error) {
	E2SmRcPreEventTriggerDefinition := new(e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition)

	switch E2SmRcPreEventTriggerDefinitionC.present {
	case C.E2SM_RC_PRE_EventTriggerDefinition_PR_eventDefinition_Format1:
		E2SmRcPreEventTriggerDefinitionFormat1, err := decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes(E2SmRcPreEventTriggerDefinitionC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %s", err.Error())
		}

		E2SmRcPreEventTriggerDefinition.E2SmRcPreEventTriggerDefinition = &e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: E2SmRcPreEventTriggerDefinitionFormat1,
		}

	default:
		return nil, fmt.Errorf("decodeE2SmRcPreEventTriggerDefinition() %v not yet implemented", E2SmRcPreEventTriggerDefinitionC.present)
	}

	return E2SmRcPreEventTriggerDefinition, nil
}
