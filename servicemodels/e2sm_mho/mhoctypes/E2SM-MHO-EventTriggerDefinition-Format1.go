// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-EventTriggerDefinition-Format1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1 *e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmMhoEventTriggerDefinitionFormat1CP, err := newE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmMhoEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1 *e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmMhoEventTriggerDefinitionFormat1CP, err := newE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmMhoEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoEventTriggerDefinitionFormat1((*C.E2SM_MHO_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func perDecodeE2SmMhoEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoEventTriggerDefinitionFormat1((*C.E2SM_MHO_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func newE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1 *e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1) (*C.E2SM_MHO_EventTriggerDefinition_Format1_t, error) {

	var err error
	e2SmMhoEventTriggerDefinitionFormat1C := C.E2SM_MHO_EventTriggerDefinition_Format1_t{}

    triggerTypeC, err := newMhoTriggerType(&e2SmMhoEventTriggerDefinitionFormat1.TriggerType)
    if err != nil {
        return nil, fmt.Errorf("newMhoTriggerType() %s", err.Error())
    }

    //instance is optional
    if e2SmMhoEventTriggerDefinitionFormat1.ReportingPeriodMs != -1 {

        reportingPeriodMsC := C.long(e2SmMhoEventTriggerDefinitionFormat1.ReportingPeriodMs)
        *e2SmMhoEventTriggerDefinitionFormat1C.reportingPeriod_ms = reportingPeriodMsC
    }
    //ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
    e2SmMhoEventTriggerDefinitionFormat1C.triggerType = *triggerTypeC

	return &e2SmMhoEventTriggerDefinitionFormat1C, nil
}

func decodeE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1C *C.E2SM_MHO_EventTriggerDefinition_Format1_t) (*e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1, error) {

	var err error
	e2SmMhoEventTriggerDefinitionFormat1 := e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//TriggerType: triggerType,
		//ReportingPeriodMs: reportingPeriodMs,

	}

	var t *e2sm_mho.MhoTriggerType
	t, err = decodeMhoTriggerType(&e2SmMhoEventTriggerDefinitionFormat1C.triggerType)
	if err != nil {
		return nil, fmt.Errorf("decodeMhoTriggerType() %s", err.Error())
	}
	e2SmMhoEventTriggerDefinitionFormat1.TriggerType = *t

	//instance is optional
	if e2SmMhoEventTriggerDefinitionFormat1C.reportingPeriod_ms != nil {

		e2SmMhoEventTriggerDefinitionFormat1.ReportingPeriodMs = int32(*e2SmMhoEventTriggerDefinitionFormat1C.reportingPeriod_ms)
	}

	return &e2SmMhoEventTriggerDefinitionFormat1, nil
}

func decodeE2SmMhoEventTriggerDefinitionFormat1Bytes(array [8]byte) (*e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1, error) {
	e2SmMhoEventTriggerDefinitionFormat1C := (*C.E2SM_MHO_EventTriggerDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoEventTriggerDefinitionFormat1(e2SmMhoEventTriggerDefinitionFormat1C)
}
