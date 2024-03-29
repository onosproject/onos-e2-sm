// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Trigger-ConditionIE-Item.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

// ToDo: Solve "Cannot convert rtPeriodIeC (type _Ctype_long) to type unsafe.Pointer
func xerEncodeTriggerConditionIeItem(triggerConditionIeItem *e2sm_kpm_ies.TriggerConditionIeItem) ([]byte, error) {
	triggerConditionIeItemCP, err := newTriggerConditionIeItem(triggerConditionIeItem)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_Trigger_ConditionIE_Item, unsafe.Pointer(triggerConditionIeItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTriggerConditionIeItem(triggerConditionIeItem *e2sm_kpm_ies.TriggerConditionIeItem) ([]byte, error) {
	triggerConditionIeItemCP, err := newTriggerConditionIeItem(triggerConditionIeItem)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Trigger_ConditionIE_Item, unsafe.Pointer(triggerConditionIeItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTriggerConditionIeItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTriggerConditionIeItem(bytes []byte) (*e2sm_kpm_ies.TriggerConditionIeItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Trigger_ConditionIE_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTriggerConditionIeItem((*C.Trigger_ConditionIE_Item_t)(unsafePtr)), nil
}

func perDecodeTriggerConditionIeItem(bytes []byte) (*e2sm_kpm_ies.TriggerConditionIeItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Trigger_ConditionIE_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTriggerConditionIeItem((*C.Trigger_ConditionIE_Item_t)(unsafePtr)), nil
}

func newTriggerConditionIeItem(triggerConditionIeItem *e2sm_kpm_ies.TriggerConditionIeItem) (*C.Trigger_ConditionIE_Item_t, error) {

	rtPeriodIeC, err := newRtPeriodIe(triggerConditionIeItem.GetReportPeriodIe())
	if err != nil {
		return nil, err
	}

	triggerConditionIeItemC := C.Trigger_ConditionIE_Item_t{
		report_Period_IE: rtPeriodIeC,
	}
	return &triggerConditionIeItemC, nil
}

func decodeTriggerConditionIeItem(triggerConditionIeItemC *C.Trigger_ConditionIE_Item_t) *e2sm_kpm_ies.TriggerConditionIeItem {
	rtPeriodIe := decodeRtPeriodIe(&triggerConditionIeItemC.report_Period_IE)
	triggerConditionIeItem := &e2sm_kpm_ies.TriggerConditionIeItem{
		ReportPeriodIe: rtPeriodIe,
	}

	return triggerConditionIeItem
}
