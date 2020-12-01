// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-RANfunction-Description.h"
//#include "RIC-EventTriggerStyle-List.h"
//#include "RIC-ReportStyle-List.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_ies.E2SmKpmRanfunctionDescription) ([]byte, error) {
	e2SmKpmRanfunctionDescriptionCP := newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription)
	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_RANfunction_Description, unsafe.Pointer(e2SmKpmRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_ies.E2SmKpmRanfunctionDescription) ([]byte, error) {
	e2SmKpmRanfunctionDescriptionCP := newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription)

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_RANfunction_Description, unsafe.Pointer(e2SmKpmRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmRanfunctionDescription(bytes []byte) (*e2sm_kpm_ies.E2SmKpmRanfunctionDescription, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmRanfunctionDescription((*C.E2SM_KPM_RANfunction_Description_t)(unsafePtr)), nil
}

func PerDecodeE2SmKpmRanfunctionDescription(bytes []byte) (*e2sm_kpm_ies.E2SmKpmRanfunctionDescription, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmRanfunctionDescription((*C.E2SM_KPM_RANfunction_Description_t)(unsafePtr)), nil
}

func newE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescription *e2sm_kpm_ies.E2SmKpmRanfunctionDescription) *C.E2SM_KPM_RANfunction_Description_t {

	e2SmKpmRanfunctionDescriptionC := new(C.E2SM_KPM_RANfunction_Description_t)
	ricEventtriggerstyleListCP := new(C.struct_E2SM_KPM_RANfunction_Description__e2SM_KPM_RANfunction_Item__ric_EventTriggerStyle_List)
	e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_EventTriggerStyle_List = ricEventtriggerstyleListCP
	for _, ricEventTriggerStyleListItem := range e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.GetRicEventTriggerStyleList() {
		ricEventTriggerStyleListItemC := newRicEventTriggerStyleList(ricEventTriggerStyleListItem)
		C.asn_sequence_add(unsafe.Pointer(ricEventtriggerstyleListCP), unsafe.Pointer(ricEventTriggerStyleListItemC))
	}

	ricReportstyleListCP := new(C.struct_E2SM_KPM_RANfunction_Description__e2SM_KPM_RANfunction_Item__ric_ReportStyle_List)
	e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_ReportStyle_List = ricReportstyleListCP
	for _, ricReportStyleListItem := range e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.GetRicReportStyleList() {
		ricReportStyleListItemC := newRicReportStyleListItem(ricReportStyleListItem)
		C.asn_sequence_add(unsafe.Pointer(ricReportstyleListCP), unsafe.Pointer(ricReportStyleListItemC))
	}

	ranfunctionNameC := newRanfunctionName(e2SmKpmRanfunctionDescription.RanFunctionName)
	e2SmKpmRanfunctionDescriptionC.ranFunction_Name = *ranfunctionNameC

	return e2SmKpmRanfunctionDescriptionC
}

func decodeE2SmKpmRanfunctionDescription(e2SmKpmRanfunctionDescriptionC *C.E2SM_KPM_RANfunction_Description_t) *e2sm_kpm_ies.E2SmKpmRanfunctionDescription {

	ranfunctionName := decodeRanfunctionName(&e2SmKpmRanfunctionDescriptionC.ranFunction_Name)

	e2SmKpmRanfunctionDescription := e2sm_kpm_ies.E2SmKpmRanfunctionDescription{
		RanFunctionName: ranfunctionName,
		E2SmKpmRanfunctionItem: &e2sm_kpm_ies.E2SmKpmRanfunctionDescription_E2SmKpmRanfunctionItem001{
			RicEventTriggerStyleList: make([]*e2sm_kpm_ies.RicEventTriggerStyleList, 0),
			RicReportStyleList:       make([]*e2sm_kpm_ies.RicReportStyleList, 0),
		},
	}

	ieCountOne := int(e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_EventTriggerStyle_List.list.count)
	for i := 0; i < ieCountOne; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) * uintptr(i)
		ricEventTriggerStyleListItemC := *(**C.RIC_EventTriggerStyle_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) + offset))
		ricEventTriggerStyleListItem := decodeRicEventTriggerStyleListItem(ricEventTriggerStyleListItemC)
		e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.RicEventTriggerStyleList = append(e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.RicEventTriggerStyleList, ricEventTriggerStyleListItem)
	}

	ieCountTwo := int(e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_ReportStyle_List.list.count)
	for i := 0; i < ieCountTwo; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_ReportStyle_List.list.array)) * uintptr(i)
		ricReportStyleListItemC := *(**C.RIC_ReportStyle_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmKpmRanfunctionDescriptionC.e2SM_KPM_RANfunction_Item.ric_ReportStyle_List.list.array)) + offset))
		ricReportStyleListItem := decodeRicReportStyleListItem(ricReportStyleListItemC)
		e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.RicReportStyleList = append(e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem.RicReportStyleList, ricReportStyleListItem)
	}

	return &e2SmKpmRanfunctionDescription
}
