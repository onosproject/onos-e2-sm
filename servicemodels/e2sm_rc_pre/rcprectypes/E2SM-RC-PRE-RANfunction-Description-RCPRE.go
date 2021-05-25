// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-RANfunction-Description-RCPRE.h"
//#include "RIC-EventTriggerStyle-List-RCPRE.h"
//#include "RIC-ReportStyle-List-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescription *e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription) ([]byte, error) {
	e2SmRcPreRanfunctionDescriptionCP := newE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescription)
	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_RANfunction_Description_RCPRE, unsafe.Pointer(e2SmRcPreRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescription *e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription) ([]byte, error) {
	e2SmRcPreRanfunctionDescriptionCP := newE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescription)

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_RANfunction_Description_RCPRE, unsafe.Pointer(e2SmRcPreRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreRanfunctionDescription(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_RANfunction_Description_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreRanfunctionDescription((*C.E2SM_RC_PRE_RANfunction_Description_RCPRE_t)(unsafePtr)), nil
}

func PerDecodeE2SmRcPreRanfunctionDescription(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_RANfunction_Description_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreRanfunctionDescription((*C.E2SM_RC_PRE_RANfunction_Description_RCPRE_t)(unsafePtr)), nil
}

func newE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescription *e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription) *C.E2SM_RC_PRE_RANfunction_Description_RCPRE_t {

	e2SmRcPreRanfunctionDescriptionC := new(C.E2SM_RC_PRE_RANfunction_Description_RCPRE_t)

	if e2SmRcPreRanfunctionDescription.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList() != nil {
		ricEventtriggerstyleListCP := new(C.struct_E2SM_RC_PRE_RANfunction_Description_RCPRE__e2SM_RC_PRE_RANfunction_Item__ric_EventTriggerStyle_List)
		e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_EventTriggerStyle_List = ricEventtriggerstyleListCP
		for _, ricEventTriggerStyleListItem := range e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.GetRicEventTriggerStyleList() {
			ricEventTriggerStyleListItemC := newRicEventTriggerStyleList(ricEventTriggerStyleListItem)
			C.asn_sequence_add(unsafe.Pointer(ricEventtriggerstyleListCP), unsafe.Pointer(ricEventTriggerStyleListItemC))
		}
	}

	if e2SmRcPreRanfunctionDescription.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList() != nil {
		ricReportstyleListCP := new(C.struct_E2SM_RC_PRE_RANfunction_Description_RCPRE__e2SM_RC_PRE_RANfunction_Item__ric_ReportStyle_List)
		e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_ReportStyle_List = ricReportstyleListCP
		for _, ricReportStyleListItem := range e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.GetRicReportStyleList() {
			ricReportStyleListItemC := newRicReportStyleListItem(ricReportStyleListItem)
			C.asn_sequence_add(unsafe.Pointer(ricReportstyleListCP), unsafe.Pointer(ricReportStyleListItemC))
		}
	}

	ranfunctionNameC := newRanfunctionName(e2SmRcPreRanfunctionDescription.RanFunctionName)
	e2SmRcPreRanfunctionDescriptionC.ranFunction_Name = *ranfunctionNameC

	return e2SmRcPreRanfunctionDescriptionC
}

func decodeE2SmRcPreRanfunctionDescription(e2SmRcPreRanfunctionDescriptionC *C.E2SM_RC_PRE_RANfunction_Description_RCPRE_t) *e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription {

	ranfunctionName := decodeRanfunctionName(&e2SmRcPreRanfunctionDescriptionC.ranFunction_Name)

	e2SmRcPreRanfunctionDescription := e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription{
		RanFunctionName: ranfunctionName,
		E2SmRcPreRanfunctionItem: &e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription_E2SmRcPreRanfunctionItem001{
			RicEventTriggerStyleList: make([]*e2sm_rc_pre_v2.RicEventTriggerStyleList, 0),
			RicReportStyleList:       make([]*e2sm_rc_pre_v2.RicReportStyleList, 0),
		},
	}

	if e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_EventTriggerStyle_List != nil {
		ieCountOne := int(e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_EventTriggerStyle_List.list.count)
		for i := 0; i < ieCountOne; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(*e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) * uintptr(i)
			ricEventTriggerStyleListItemC := *(**C.RIC_EventTriggerStyle_List_RCPRE_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_EventTriggerStyle_List.list.array)) + offset))
			ricEventTriggerStyleListItem := decodeRicEventTriggerStyleListItem(ricEventTriggerStyleListItemC)
			e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.RicEventTriggerStyleList = append(e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.RicEventTriggerStyleList, ricEventTriggerStyleListItem)
		}
	}

	if e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_ReportStyle_List != nil {
		ieCountTwo := int(e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_ReportStyle_List.list.count)
		for i := 0; i < ieCountTwo; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(*e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_ReportStyle_List.list.array)) * uintptr(i)
			ricReportStyleListItemC := *(**C.RIC_ReportStyle_List_RCPRE_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmRcPreRanfunctionDescriptionC.e2SM_RC_PRE_RANfunction_Item.ric_ReportStyle_List.list.array)) + offset))
			ricReportStyleListItem := decodeRicReportStyleListItem(ricReportStyleListItemC)
			e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.RicReportStyleList = append(e2SmRcPreRanfunctionDescription.E2SmRcPreRanfunctionItem.RicReportStyleList, ricReportStyleListItem)
		}
	}

	return &e2SmRcPreRanfunctionDescription
}
