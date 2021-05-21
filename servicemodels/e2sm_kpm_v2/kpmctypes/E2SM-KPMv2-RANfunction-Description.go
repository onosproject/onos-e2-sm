// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-RANfunction-Description.h"
//#include "RIC-KPMNode-Item-KPMv2.h"
//#include "RIC-EventTriggerStyle-Item-KPMv2.h"
//#include "RIC-ReportStyle-Item-KPMv2.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func XerEncodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_v2.E2SmKpmRanfunctionDescription) ([]byte, error) {
	e2SmKpmRanfunctionDescriptionCP, err := newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_RANfunction_Description, unsafe.Pointer(e2SmKpmRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_v2.E2SmKpmRanfunctionDescription) ([]byte, error) {
	e2SmKpmRanfunctionDescriptionCP, err := newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_RANfunction_Description, unsafe.Pointer(e2SmKpmRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmRanfunctionDescription(bytes []byte) (*e2sm_kpm_v2.E2SmKpmRanfunctionDescription, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmRanfunctionDescription((*C.E2SM_KPMv2_RANfunction_Description_t)(unsafePtr))
}

func PerDecodeE2SmKpmRanfunctionDescription(bytes []byte) (*e2sm_kpm_v2.E2SmKpmRanfunctionDescription, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmRanfunctionDescription((*C.E2SM_KPMv2_RANfunction_Description_t)(unsafePtr))
}

func newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_v2.E2SmKpmRanfunctionDescription) (*C.E2SM_KPMv2_RANfunction_Description_t, error) {

	ranFunctionNameC, err := newRanfunctionName(e2SmKpmRanfunctionDescription.RanFunctionName)
	if err != nil {
		return nil, fmt.Errorf("newRanfunctionName() %s", err.Error())
	}

	e2SmKpmRanfunctionDescriptionC := C.E2SM_KPMv2_RANfunction_Description_t{
		ranFunction_Name:           *ranFunctionNameC,
		//ric_KPM_Node_List:          ricKpmNodeListC,
		//ric_EventTriggerStyle_List: ricEventTriggerStyleListC,
		//ric_ReportStyle_List:       ricReportStyleListC,
	}

	if e2SmKpmRanfunctionDescription.RicKpmNodeList != nil {
		ricKpmNodeListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_KPM_Node_List)
		for _, ricKpmNodeListItem := range e2SmKpmRanfunctionDescription.GetRicKpmNodeList() {
			ricKpmNodeListItemC, err := newRicKpmnodeItem(ricKpmNodeListItem)
			if err != nil {
				return nil, fmt.Errorf("newRicKpmnodeItem() %s", err.Error())
			}
			if _, err = C.asn_sequence_add(unsafe.Pointer(ricKpmNodeListC), unsafe.Pointer(ricKpmNodeListItemC)); err != nil {
				return nil, err
			}
		}
		e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List = ricKpmNodeListC
	}

	if e2SmKpmRanfunctionDescription.RicEventTriggerStyleList != nil {
		ricEventTriggerStyleListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_EventTriggerStyle_List)
		for _, ricEventTriggerStyleListItem := range e2SmKpmRanfunctionDescription.GetRicEventTriggerStyleList() {
			ricEventTriggerStyleListItemC, err := newRicEventTriggerStyleItem(ricEventTriggerStyleListItem)
			if err != nil {
				return nil, fmt.Errorf("newRicEventTriggerStyleItem() %s", err.Error())
			}
			if _, err = C.asn_sequence_add(unsafe.Pointer(ricEventTriggerStyleListC), unsafe.Pointer(ricEventTriggerStyleListItemC)); err != nil {
				return nil, err
			}
		}
		e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List = ricEventTriggerStyleListC
	}

	if e2SmKpmRanfunctionDescription.RicReportStyleList != nil {
		ricReportStyleListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_ReportStyle_List)
		for _, ricReportStyleListItem := range e2SmKpmRanfunctionDescription.GetRicReportStyleList() {
			ricReportStyleListItemC, err := newRicReportStyleItem(ricReportStyleListItem)
			if err != nil {
				return nil, fmt.Errorf("newRicReportStyleItem() %s", err.Error())
			}
			if _, err = C.asn_sequence_add(unsafe.Pointer(ricReportStyleListC), unsafe.Pointer(ricReportStyleListItemC)); err != nil {
				return nil, err
			}
		}
		e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List = ricReportStyleListC
	}

	return &e2SmKpmRanfunctionDescriptionC, nil
}

func decodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescriptionC *C.E2SM_KPMv2_RANfunction_Description_t) (*e2sm_kpm_v2.E2SmKpmRanfunctionDescription, error) {

	ranFunctionName, err := decodeRanfunctionName(&e2SmKpmRanfunctionDescriptionC.ranFunction_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRanfunctionName() %s", err.Error())
	}

	e2SmKpmRanfunctionDescription := e2sm_kpm_v2.E2SmKpmRanfunctionDescription{
		RanFunctionName:          ranFunctionName,
		RicKpmNodeList:           make([]*e2sm_kpm_v2.RicKpmnodeItem, 0),
		RicEventTriggerStyleList: make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0),
		RicReportStyleList:       make([]*e2sm_kpm_v2.RicReportStyleItem, 0),
	}

	// instance is optional
	if e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List != nil {
		ieCount := int(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.count)
		for i := 0; i < ieCount; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.array)) * uintptr(i)
			ieC := *(**C.RIC_KPMNode_Item_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.array)) + offset))
			ie, err := decodeRicKpmnodeItem(ieC)
			if err != nil {
				return nil, fmt.Errorf("decodeRicKpmnodeItem() %s", err.Error())
			}
			e2SmKpmRanfunctionDescription.RicKpmNodeList = append(e2SmKpmRanfunctionDescription.RicKpmNodeList, ie)
		}
	}

	//instance is optional
	if e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List != nil {
		ieCount := int(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.count)
		for i := 0; i < ieCount; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.array)) * uintptr(i)
			ieC := *(**C.RIC_EventTriggerStyle_Item_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.array)) + offset))
			ie, err := decodeRicEventTriggerStyleItem(ieC)
			if err != nil {
				return nil, fmt.Errorf("decodeRicEventTriggerStyleItem() %s", err.Error())
			}
			e2SmKpmRanfunctionDescription.RicEventTriggerStyleList = append(e2SmKpmRanfunctionDescription.RicEventTriggerStyleList, ie)
		}
	}

	//instance is optional
	if e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List != nil {
		ieCount := int(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.count)
		for i := 0; i < ieCount; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.array)) * uintptr(i)
			ieC := *(**C.RIC_ReportStyle_Item_KPMv2_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.array)) + offset))
			ie, err := decodeRicReportStyleItem(ieC)
			if err != nil {
				return nil, fmt.Errorf("decodeRicReportStyleItem() %s", err.Error())
			}
			e2SmKpmRanfunctionDescription.RicReportStyleList = append(e2SmKpmRanfunctionDescription.RicReportStyleList, ie)
		}
	}

	return &e2SmKpmRanfunctionDescription, nil
}
