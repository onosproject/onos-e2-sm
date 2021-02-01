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
//#include "Trigger-ConditionIE-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1) ([]byte, error) {
	E2SmRcPreEventTriggerDefinitionFormat1CP, err := newE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1) ([]byte, error) {
	E2SmRcPreEventTriggerDefinitionFormat1CP, err := newE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1, unsafe.Pointer(E2SmRcPreEventTriggerDefinitionFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmRcPreEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinitionFormat1((*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func perDecodeE2SmRcPreEventTriggerDefinitionFormat1(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_EventTriggerDefinition_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreEventTriggerDefinitionFormat1((*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafePtr))
}

func newE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1 *e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1) (*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t, error) {
	policyTestListC := new(C.struct_E2SM_RC_PRE_EventTriggerDefinition_Format1__policyTest_List)
	format1 := C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t{
		policyTest_List: policyTestListC,
	}
	for _, policyTestItem := range E2SmRcPreEventTriggerDefinitionFormat1.GetPolicyTestList() {
		policyTestItemC, err := newTriggerConditionIeItem(policyTestItem)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreEventTriggerDefinitionFormat1() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(policyTestListC), unsafe.Pointer(policyTestItemC)); err != nil {
			return nil, err
		}
	}

	return &format1, nil
}

func decodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1C *C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	E2SmRcPreEventTriggerDefinitionFormat1 := &e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1{
		PolicyTestList: make([]*e2sm_rc_pre_ies.TriggerConditionIeItem, 0),
	}

	ieCount := int(E2SmRcPreEventTriggerDefinitionFormat1C.policyTest_List.list.count)
	//fmt.Printf("RanFunctionListC %T List %T %v Array %T %v Deref %v\n", rflC, rflC.list, rflC.list, rflC.list.array, *rflC.list.array, *(rflC.list.array))
	for i := 0; i < ieCount; i++ {
		//ToDo: policyTest_List is defined as a POINTER!! Could cause a problem there
		offset := unsafe.Sizeof(unsafe.Pointer(*E2SmRcPreEventTriggerDefinitionFormat1C.policyTest_List.list.array)) * uintptr(i)
		policyTestItemC := *(**C.Trigger_ConditionIE_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SmRcPreEventTriggerDefinitionFormat1C.policyTest_List.list.array)) + offset))
		//fmt.Printf("Value %T %p %v\n", rfiIeC, rfiIeC, rfiIeC)
		policyTestItem := decodeTriggerConditionIeItem(policyTestItemC)

		E2SmRcPreEventTriggerDefinitionFormat1.PolicyTestList = append(E2SmRcPreEventTriggerDefinitionFormat1.PolicyTestList, policyTestItem)
	}

	return E2SmRcPreEventTriggerDefinitionFormat1, nil
}

func decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes(array [8]byte) (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1, error) {
	E2SmRcPreEventTriggerDefinitionFormat1C := (*C.E2SM_RC_PRE_EventTriggerDefinition_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreEventTriggerDefinitionFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
