// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-RANfunction-Description.h"
//#include "RIC-EventTriggerStyle-List.h"
//#include "RIC-ReportStyle-List.h"
import "C"
import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"unsafe"
)

func XerEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) ([]byte, error) {
	e2SmMhoRanfunctionDescriptionCP := newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)
	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_RANfunction_Description, unsafe.Pointer(e2SmMhoRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) ([]byte, error) {
	e2SmMhoRanfunctionDescriptionCP := newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_RANfunction_Description, unsafe.Pointer(e2SmMhoRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmMhoRanfunctionDescription(bytes []byte) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoRanfunctionDescription((*C.E2SM_MHO_RANfunction_Description_t)(unsafePtr)), nil
}

func PerDecodeE2SmMhoRanfunctionDescription(bytes []byte) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoRanfunctionDescription((*C.E2SM_MHO_RANfunction_Description_t)(unsafePtr)), nil
}

func newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) *C.E2SM_MHO_RANfunction_Description_t {

	e2SmMhoRanfunctionDescriptionC := new(C.E2SM_MHO_RANfunction_Description_t)

	if e2SmMhoRanfunctionDescription.GetRicEventTriggerStyleList() != nil {
		ricEventtriggerstyleListCP := new(C.struct_E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item__ric_EventTriggerStyle_List)
		e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_EventTriggerStyle_List = ricEventtriggerstyleListCP
		for _, ricEventTriggerStyleListItem := range e2SmMhoRanfunctionDescription.GetRicEventTriggerStyleList() {
			ricEventTriggerStyleListItemC := newRicEventTriggerStyleList(ricEventTriggerStyleListItem)
			C.asn_sequence_add(unsafe.Pointer(ricEventtriggerstyleListCP), unsafe.Pointer(ricEventTriggerStyleListItemC))
		}
	}

	if e2SmMhoRanfunctionDescription.GetRicReportStyleList() != nil {
		ricReportstyleListCP := new(C.struct_E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item__ric_ReportStyle_List)
		e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_ReportStyle_List = ricReportstyleListCP
		for _, ricReportStyleListItem := range e2SmMhoRanfunctionDescription.GetRicReportStyleList() {
			ricReportStyleListItemC := newRicReportStyleListItem(ricReportStyleListItem)
			C.asn_sequence_add(unsafe.Pointer(ricReportstyleListCP), unsafe.Pointer(ricReportStyleListItemC))
		}
	}

	ranfunctionNameC := newRanfunctionName(e2SmMhoRanfunctionDescription.RanFunctionName)
	e2SmMhoRanfunctionDescriptionC.ranFunction_Name = *ranfunctionNameC

	return e2SmMhoRanfunctionDescriptionC
}

func decodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescriptionC *C.E2SM_MHO_RANfunction_Description_t) *e2sm_mho.E2SmMhoRanfunctionDescription {

	ranfunctionName := decodeRanfunctionName(&e2SmMhoRanfunctionDescriptionC.ranFunction_Name)

	e2SmMhoRanfunctionDescription := e2sm_mho.E2SmMhoRanfunctionDescription{
		RanFunctionName: ranfunctionName,
		RicEventTriggerStyleList: make([]*e2sm_mho.RicEventTriggerStyleList, 0),
		RicReportStyleList:       make([]*e2sm_mho.RicReportStyleList, 0),
	}

	if e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_EventTriggerStyle_List != nil {
		ieCountOne := int(e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_EventTriggerStyle_List.list.count)
		for i := 0; i < ieCountOne; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(*e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) * uintptr(i)
			ricEventTriggerStyleListItemC := *(**C.RIC_EventTriggerStyle_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) + offset))
			ricEventTriggerStyleListItem := decodeRicEventTriggerStyleListItem(ricEventTriggerStyleListItemC)
			e2SmMhoRanfunctionDescription.RicEventTriggerStyleList = append(e2SmMhoRanfunctionDescription.RicEventTriggerStyleList, ricEventTriggerStyleListItem)
		}
	}

	if e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_ReportStyle_List != nil {
		ieCountTwo := int(e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_ReportStyle_List.list.count)
		for i := 0; i < ieCountTwo; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(*e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_ReportStyle_List.list.array)) * uintptr(i)
			ricReportStyleListItemC := *(**C.RIC_ReportStyle_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.e2SM_MHO_RANfunction_Item.ric_ReportStyle_List.list.array)) + offset))
			ricReportStyleListItem := decodeRicReportStyleListItem(ricReportStyleListItemC)
			e2SmMhoRanfunctionDescription.RicReportStyleList = append(e2SmMhoRanfunctionDescription.RicReportStyleList, ricReportStyleListItem)
		}
	}

	return &e2SmMhoRanfunctionDescription
}
