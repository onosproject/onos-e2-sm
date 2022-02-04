// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-EventTriggerDefinition-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition) ([]byte, error) {

	E2SmRcPreEventTriggerDefinitionCP, err := newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_RCPRE, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition) ([]byte, error) {

	E2SmRcPreEventTriggerDefinitionCP, err := newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_RCPRE, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreEventTriggerDefinition(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinition((*C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE_t)(unsafePtr))
}

func XerDecodeE2SmRcPreEventTriggerDefinition(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinition((*C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition) (*C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE_t, error) {
	var presentC C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE__eventDefinition_formats_PR
	choiceC := [8]byte{}

	switch choice := E2SmRcPreEventTriggerDefinition.E2SmRcPreEventTriggerDefinitionEventDefinitionFormats.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1:
		presentC = C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE__eventDefinition_formats_PR_eventDefinition_Format1

		im, err := newE2SmRcPreEventTriggerDefinitionFormat1(choice.EventDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinition() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinition() %T not yet implemented", choice)
	}

	eventDefinitionFormatsC := C.struct_E2SM_RC_PRE_EventTriggerDefinition_RCPRE__eventDefinition_formats{
		present: presentC,
		choice:  choiceC,
	}

	E2SmRcPreEventTriggerDefinitionC := C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE_t{
		eventDefinition_formats: eventDefinitionFormatsC,
	}

	return &E2SmRcPreEventTriggerDefinitionC, nil
}

func decodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinitionC *C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition, error) {
	E2SmRcPreEventTriggerDefinition := new(e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition)

	eventDefinitionFormatsC := E2SmRcPreEventTriggerDefinitionC.eventDefinition_formats
	switch eventDefinitionFormatsC.present {
	case C.E2SM_RC_PRE_EventTriggerDefinition_RCPRE__eventDefinition_formats_PR_eventDefinition_Format1:
		E2SmRcPreEventTriggerDefinitionFormat1, err := decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes(eventDefinitionFormatsC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %s", err.Error())
		}

		E2SmRcPreEventTriggerDefinition.E2SmRcPreEventTriggerDefinitionEventDefinitionFormats = &e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: E2SmRcPreEventTriggerDefinitionFormat1,
		}

	default:
		return nil, fmt.Errorf("decodeE2SmRcPreEventTriggerDefinition() not yet implemented")
	}

	return E2SmRcPreEventTriggerDefinition, nil
}
