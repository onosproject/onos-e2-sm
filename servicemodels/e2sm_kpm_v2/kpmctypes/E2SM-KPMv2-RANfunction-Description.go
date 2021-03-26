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
//#include "RIC-KPMNode-Item.h"
//#include "RIC-EventTriggerStyle-Item.h"
//#include "RIC-ReportStyle-Item.h"
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

	//fmt.Printf("We're inside newE2SmKpmRanfunctionDescription(), starting encoding...")
	//fmt.Printf("Encoding RanFunctionName: \n %v \n", e2SmKpmRanfunctionDescription.RanFunctionName)
	ranFunctionNameC, err := newRanfunctionName(e2SmKpmRanfunctionDescription.RanFunctionName)
	if err != nil {
		return nil, fmt.Errorf("newRanfunctionName() %s", err.Error())
	}
	//fmt.Printf("That's the C version of encoded RanFunctionName: \n %v \n", ranFunctionNameC)

	//fmt.Printf("Encoding RicKpmNodeList: \n %v \n", e2SmKpmRanfunctionDescription.GetRicKpmNodeList())
	ricKpmNodeListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_KPM_Node_List)
	for _, ricKpmNodeListItem := range e2SmKpmRanfunctionDescription.GetRicKpmNodeList() {
		//fmt.Printf("Encoding RicKpmnodeItem: \n %v \n", ricKpmNodeListItem)
		ricKpmNodeListItemC, err := newRicKpmnodeItem(ricKpmNodeListItem)
		//fmt.Printf("That's the C version of encoded RicKpmnodeItem: \n %v \n", ricKpmNodeListItemC)
		if err != nil {
			return nil, fmt.Errorf("newRicKpmnodeItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(ricKpmNodeListC), unsafe.Pointer(ricKpmNodeListItemC)); err != nil {
			return nil, err
		}
	}
	//fmt.Printf("That's the C version of encoded RicKpmNodeList: \n %v \n", ricKpmNodeListC)

	//fmt.Printf("Encoding RicEventTriggerStyleList: \n %v \n", e2SmKpmRanfunctionDescription.GetRicEventTriggerStyleList())
	ricEventTriggerStyleListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_EventTriggerStyle_List)
	for _, ricEventTriggerStyleListItem := range e2SmKpmRanfunctionDescription.GetRicEventTriggerStyleList() {
		//fmt.Printf("Encoding RicEventTriggerStyleListItem: \n %v \n", ricEventTriggerStyleListItem)
		ricEventTriggerStyleListItemC, err := newRicEventTriggerStyleItem(ricEventTriggerStyleListItem)
		//fmt.Printf("That's the C version of encoded RicEventTriggerStyleListItem: \n %v \n", ricEventTriggerStyleListItemC)
		if err != nil {
			return nil, fmt.Errorf("newRicEventTriggerStyleItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(ricEventTriggerStyleListC), unsafe.Pointer(ricEventTriggerStyleListItemC)); err != nil {
			return nil, err
		}
	}
	//fmt.Printf("That's the C version of encoded RicEventTriggerStyleList: \n %v \n", ricEventTriggerStyleListC)

	//fmt.Printf("Encoding RicReportStyleList: \n %v \n", e2SmKpmRanfunctionDescription.GetRicReportStyleList())
	ricReportStyleListC := new(C.struct_E2SM_KPMv2_RANfunction_Description__ric_ReportStyle_List)
	for _, ricReportStyleListItem := range e2SmKpmRanfunctionDescription.GetRicReportStyleList() {
		//fmt.Printf("Encoding RicReportStyleListItem: \n %v \n", ricReportStyleListItem)
		ricReportStyleListItemC, err := newRicReportStyleItem(ricReportStyleListItem)
		//fmt.Printf("That's the C version of encoded RicReportStyleListItem: \n %v \n", ricReportStyleListItemC)
		if err != nil {
			return nil, fmt.Errorf("newRicReportStyleItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(ricReportStyleListC), unsafe.Pointer(ricReportStyleListItemC)); err != nil {
			return nil, err
		}
	}
	//fmt.Printf("That's the C version of encoded RicReportStyleList: \n %v \n", ricReportStyleListC)

	//fmt.Printf("Composing final C-structure...")
	e2SmKpmRanfunctionDescriptionC := C.E2SM_KPMv2_RANfunction_Description_t{
		ranFunction_Name:           *ranFunctionNameC,
		ric_KPM_Node_List:          ricKpmNodeListC,
		ric_EventTriggerStyle_List: ricEventTriggerStyleListC,
		ric_ReportStyle_List:       ricReportStyleListC,
	}
	//fmt.Printf("This is final C-structure: \n %v \n", e2SmKpmRanfunctionDescriptionC)

	return &e2SmKpmRanfunctionDescriptionC, nil
}

func decodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescriptionC *C.E2SM_KPMv2_RANfunction_Description_t) (*e2sm_kpm_v2.E2SmKpmRanfunctionDescription, error) {

	//fmt.Printf("We're inside decodeE2SmKpmRanfunctionDescription(), starting decoding...")
	//fmt.Printf("Decoding RanFunctionName: \n %v \n", e2SmKpmRanfunctionDescriptionC.ranFunction_Name)
	ranFunctionName, err := decodeRanfunctionName(&e2SmKpmRanfunctionDescriptionC.ranFunction_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRanfunctionName() %s", err.Error())
	}
	//fmt.Printf("That's what was decoded from C: \n %v \n", ranFunctionName)

	e2SmKpmRanfunctionDescription := e2sm_kpm_v2.E2SmKpmRanfunctionDescription{
		RanFunctionName:          ranFunctionName,
		RicKpmNodeList:           make([]*e2sm_kpm_v2.RicKpmnodeItem, 0),
		RicEventTriggerStyleList: make([]*e2sm_kpm_v2.RicEventTriggerStyleItem, 0),
		RicReportStyleList:       make([]*e2sm_kpm_v2.RicReportStyleItem, 0),

	}

	//fmt.Printf("Decoding RicKpmNodeList: \n %v \n", e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List)
	//fmt.Printf("We've got %v items to decode \n", int(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.count))
	var ieCount int
	ieCount = int(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.array)) * uintptr(i)
		ieC := *(**C.RIC_KPMNode_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_KPM_Node_List.list.array)) + offset))
		//fmt.Printf("Decoding RicKpmNodeItem: \n %v \n", ieC)
		ie, err := decodeRicKpmnodeItem(ieC)
		//fmt.Printf("Decoded RicKpmNodeItem is: \n %v \n", ie)
		if err != nil {
			return nil, fmt.Errorf("decodeRicKpmnodeItem() %s", err.Error())
		}
		e2SmKpmRanfunctionDescription.RicKpmNodeList = append(e2SmKpmRanfunctionDescription.RicKpmNodeList, ie)
	}
	//fmt.Printf("Decoded RicKpmNodeList is: \n %v \n", e2SmKpmRanfunctionDescription.RicKpmNodeList)

	//fmt.Printf("Decoding RicEventTriggerStyleList: \n %v \n", e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List)
	//fmt.Printf("We've got %v items to decode \n", int(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.count))
	ieCount = int(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.array)) * uintptr(i)
		ieC := *(**C.RIC_EventTriggerStyle_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_EventTriggerStyle_List.list.array)) + offset))
		//fmt.Printf("Decoding RicEventTriggerStyleItem: \n %v \n", ieC)
		ie, err := decodeRicEventTriggerStyleItem(ieC)
		//fmt.Printf("Decoded RicEventTriggerStyleItem is: \n %v \n", ie)
		if err != nil {
			return nil, fmt.Errorf("decodeRicEventTriggerStyleItem() %s", err.Error())
		}
		e2SmKpmRanfunctionDescription.RicEventTriggerStyleList = append(e2SmKpmRanfunctionDescription.RicEventTriggerStyleList, ie)
	}
	//fmt.Printf("Decoded RicEventTriggerStyleList is: \n %v \n", e2SmKpmRanfunctionDescription.RicEventTriggerStyleList)

	//fmt.Printf("Decoding RicReportStyleList: \n %v \n", e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List)
	//fmt.Printf("We've got %v items to decode \n", int(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.count))
	ieCount = int(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.array)) * uintptr(i)
		ieC := *(**C.RIC_ReportStyle_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.ric_ReportStyle_List.list.array)) + offset))
		//fmt.Printf("Decoding RicReportStyleItem: \n %v \n", ieC)
		ie, err := decodeRicReportStyleItem(ieC)
		//fmt.Printf("Decoded RicReportStyleItem is: \n %v \n", ie)
		if err != nil {
			return nil, fmt.Errorf("decodeRicReportStyleItem() %s", err.Error())
		}
		e2SmKpmRanfunctionDescription.RicReportStyleList = append(e2SmKpmRanfunctionDescription.RicReportStyleList, ie)
	}
	//fmt.Printf("Decoded RicReportStyleList is: \n %v \n", e2SmKpmRanfunctionDescription.RicReportStyleList)

	//fmt.Printf("This is final decoded structure: \n %v \n", e2SmKpmRanfunctionDescription)

	return &e2SmKpmRanfunctionDescription, nil
}
