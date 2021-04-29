// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-EventTriggerDefinition.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinition *e2sm_mho.E2SmMhoEventTriggerDefinition) ([]byte, error) {
	e2SmMhoEventTriggerDefinitionCP, err := newE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoEventTriggerDefinition() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_EventTriggerDefinition, unsafe.Pointer(e2SmMhoEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoEventTriggerDefinition() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinition *e2sm_mho.E2SmMhoEventTriggerDefinition) ([]byte, error) {
	e2SmMhoEventTriggerDefinitionCP, err := newE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinition)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoEventTriggerDefinition() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_EventTriggerDefinition, unsafe.Pointer(e2SmMhoEventTriggerDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoEventTriggerDefinition() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoEventTriggerDefinition(bytes []byte) (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoEventTriggerDefinition((*C.E2SM_MHO_EventTriggerDefinition_t)(unsafePtr))
}

func perDecodeE2SmMhoEventTriggerDefinition(bytes []byte) (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_EventTriggerDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoEventTriggerDefinition((*C.E2SM_MHO_EventTriggerDefinition_t)(unsafePtr))
}

func newE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinition *e2sm_mho.E2SmMhoEventTriggerDefinition) (*C.E2SM_MHO_EventTriggerDefinition_t, error) {

    var pr C.E2SM_MHO_EventTriggerDefinition__eventDefinition_formats_PR //ToDo - verify correctness of the name
    choiceC := [8]byte{}                                                 //ToDo - Check if number of bytes is sufficient

	switch choice := e2SmMhoEventTriggerDefinition.E2SmMhoEventTriggerDefinition.(type) {
	case *e2sm_mho.E2SmMhoEventTriggerDefinition_EventDefinitionFormat1:
        pr = C.E2SM_MHO_EventTriggerDefinition__eventDefinition_formats_PR_eventDefinition_Format1 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoEventTriggerDefinitionFormat1(choice.EventDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoEventTriggerDefinitionFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoEventTriggerDefinition() %T not yet implemented", choice)
	}

    e2SmMhoEventTriggerDefinitionC := C.E2SM_MHO_EventTriggerDefinition_t{}
    e2SmMhoEventTriggerDefinitionC.eventDefinition_formats.present = pr
    e2SmMhoEventTriggerDefinitionC.eventDefinition_formats.choice = choiceC

	return &e2SmMhoEventTriggerDefinitionC, nil
}

func decodeE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinitionC *C.E2SM_MHO_EventTriggerDefinition_t) (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {

	e2SmMhoEventTriggerDefinition := new(e2sm_mho.E2SmMhoEventTriggerDefinition)

    switch e2SmMhoEventTriggerDefinitionC.eventDefinition_formats.present {
    case C.E2SM_MHO_EventTriggerDefinition__eventDefinition_formats_PR_eventDefinition_Format1:
        e2SmMhoEventTriggerDefinitionstructC, err := decodeE2SmMhoEventTriggerDefinitionFormat1Bytes(e2SmMhoEventTriggerDefinitionC.eventDefinition_formats.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoEventTriggerDefinitionFormat1Bytes() %s", err.Error())
		}
        e2SmMhoEventTriggerDefinition.E2SmMhoEventTriggerDefinition = &e2sm_mho.E2SmMhoEventTriggerDefinition_EventDefinitionFormat1{
            EventDefinitionFormat1: e2SmMhoEventTriggerDefinitionstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoEventTriggerDefinition() %v not yet implemented", e2SmMhoEventTriggerDefinitionC.eventDefinition_formats.present)
	}

    return e2SmMhoEventTriggerDefinition, nil
}

func decodeE2SmMhoEventTriggerDefinitionBytes(array [8]byte) (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {
	e2SmMhoEventTriggerDefinitionC := (*C.E2SM_MHO_EventTriggerDefinition_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoEventTriggerDefinition(e2SmMhoEventTriggerDefinitionC)
}
