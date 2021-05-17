// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-EventTriggerDefinition-Format1.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1) ([]byte, error) {
	E2SmRcPreEventTriggerDefinitionFormat1CP, err := newE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1) ([]byte, error) {
	E2SmRcPreEventTriggerDefinitionFormat1CP, err := newE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinitionFormat1((*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func PerDecodeE2SmRcPreEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinitionFormat1((*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func newE2SmRcPreEventTriggerDefinitionFormat1(e2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1) (*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t, error) {

	triggerTypeC, err := newRcPreTriggerType(&e2SmRcPreEventTriggerDefinitionFormat1.TriggerType)
	if err != nil {
		return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}

	e2SmRcPreEventTriggerDefinitionFormat1C := C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t{
		triggerType: *triggerTypeC,
		//reportingPeriod_ms: &reportingPeriodMsC,
	}

	if e2SmRcPreEventTriggerDefinitionFormat1.GetReportingPeriodMs() != 0 {
		reportingPeriodMsC := C.ulong(e2SmRcPreEventTriggerDefinitionFormat1.ReportingPeriodMs)
		e2SmRcPreEventTriggerDefinitionFormat1C.reportingPeriod_ms = &reportingPeriodMsC
	}

	return &e2SmRcPreEventTriggerDefinitionFormat1C, nil
}

func decodeE2SmRcPreEventTriggerDefinitionFormat1(e2SmRcPreEventTriggerDefinitionFormat1C *C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1, error) {

	triggerType, err := decodeRcPreTriggerType(&e2SmRcPreEventTriggerDefinitionFormat1C.triggerType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
	}

	e2SmRcPreEventTriggerDefinitionFormat1 := e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1{
		TriggerType: *triggerType,
		//ReportingPeriodMs: reportingPeriodMs,
	}

	if e2SmRcPreEventTriggerDefinitionFormat1C.reportingPeriod_ms != nil {
		reportingPeriodMs := uint32(*e2SmRcPreEventTriggerDefinitionFormat1C.reportingPeriod_ms)
		e2SmRcPreEventTriggerDefinitionFormat1.ReportingPeriodMs = reportingPeriodMs
	}

	return &e2SmRcPreEventTriggerDefinitionFormat1, nil
}

func decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes(array [8]byte) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	e2SmRcPreEventTriggerDefinitionFormat1C := (*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmRcPreEventTriggerDefinitionFormat1(e2SmRcPreEventTriggerDefinitionFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
