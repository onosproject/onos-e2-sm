// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-EventTriggerDefinition-Format1.h"
//#include "Trigger-ConditionIE-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmKpmEventTriggerDefinitionFormat1CP, err := newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmKpmEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1) ([]byte, error) {
	e2SmKpmEventTriggerDefinitionFormat1CP, err := newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1, unsafe.Pointer(e2SmKpmEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinitionFormat1((*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func perDecodeE2SmKpmEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmEventTriggerDefinitionFormat1((*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func newE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1 *e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1) (*C.E2SM_KPM_EventTriggerDefinition_Format1_t, error) {
	policyTestListC := new(C.struct_E2SM_KPM_EventTriggerDefinition_Format1__policyTest_List)
	format1 := C.E2SM_KPM_EventTriggerDefinition_Format1_t{
		policyTest_List: policyTestListC,
	}
	for _, policyTestItem := range e2SmKpmEventTriggerDefinitionFormat1.GetPolicyTestList() {
		policyTestItemC, err := newTriggerConditionIeItem(policyTestItem)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmEventTriggerDefinitionFormat1() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(policyTestListC), unsafe.Pointer(policyTestItemC)); err != nil {
			return nil, err
		}
	}

	return &format1, nil
}

func decodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1C *C.E2SM_KPM_EventTriggerDefinition_Format1_t) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1, error) {
	e2SmKpmEventTriggerDefinitionFormat1 := &e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1{
		PolicyTestList: make([]*e2sm_kpm_ies.TriggerConditionIeItem, 0),
	}

	ieCount := int(e2SmKpmEventTriggerDefinitionFormat1C.policyTest_List.list.count)
	//fmt.Printf("RanFunctionListC %T List %T %v Array %T %v Deref %v\n", rflC, rflC.list, rflC.list, rflC.list.array, *rflC.list.array, *(rflC.list.array))
	for i := 0; i < ieCount; i++ {
		//ToDo: policyTest_List is defined as a POINTER!! Could cause a problem there
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmKpmEventTriggerDefinitionFormat1C.policyTest_List.list.array)) * uintptr(i)
		policyTestItemC := *(**C.Trigger_ConditionIE_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmEventTriggerDefinitionFormat1C.policyTest_List.list.array)) + offset))
		//fmt.Printf("Value %T %p %v\n", rfiIeC, rfiIeC, rfiIeC)
		policyTestItem := decodeTriggerConditionIeItem(policyTestItemC)

		e2SmKpmEventTriggerDefinitionFormat1.PolicyTestList = append(e2SmKpmEventTriggerDefinitionFormat1.PolicyTestList, policyTestItem)
	}

	return e2SmKpmEventTriggerDefinitionFormat1, nil
}

func decodeE2SmKpmEventTriggerDefinitionFormat1Bytes(array [8]byte) (*e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1, error) {
	e2SmKpmEventTriggerDefinitionFormat1C := (*C.E2SM_KPM_EventTriggerDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmKpmEventTriggerDefinitionFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
