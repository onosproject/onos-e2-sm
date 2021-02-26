// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-EventTriggerDefinition-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmKpmEventTriggerDefinitionFormat1CP, err := newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmKpmEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmKpmEventTriggerDefinitionFormat1CP, err := newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmKpmEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinitionFormat1((*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func perDecodeE2SmKpmEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinitionFormat1((*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1) (*C.E2SM_KPM_EventTriggerDefinition_Format1_t, error) {

	reportingPeriodC := C.long(e2SmKpmEventTriggerDefinitionFormat1.ReportingPeriod)
	e2SmKpmEventTriggerDefinitionFormat1C := C.E2SM_KPM_EventTriggerDefinition_Format1_t{
		reportingPeriod: reportingPeriodC,
	}

	return &e2SmKpmEventTriggerDefinitionFormat1C, nil
}

func decodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1C *C.E2SM_KPM_EventTriggerDefinition_Format1_t) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1, error) {

	reportingPeriod := int32(e2SmKpmEventTriggerDefinitionFormat1C.reportingPeriod)
	e2SmKpmEventTriggerDefinitionFormat1 := e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1{
		ReportingPeriod: reportingPeriod,
	}

	return &e2SmKpmEventTriggerDefinitionFormat1, nil
}

func decodeE2SmKpmEventTriggerDefinitionFormat1Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1, error) {
	e2SmKpmEventTriggerDefinitionFormat1C := (*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1C)
}
