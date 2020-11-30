// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-EventTriggerStyle-List.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_ies.RicEventTriggerStyleList) ([]byte, error) {
	ricEventTriggerStyleItemCP := newRicEventTriggerStyleList(ricEventTriggerStyleItem)

	bytes, err := encodeXer(&C.asn_DEF_RIC_EventTriggerStyle_List, unsafe.Pointer(ricEventTriggerStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_ies.RicEventTriggerStyleList) ([]byte, error) {
	ricEventTriggerStyleItemCP := newRicEventTriggerStyleList(ricEventTriggerStyleItem)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_EventTriggerStyle_List, unsafe.Pointer(ricEventTriggerStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicEventTriggerStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicEventTriggerStyleItem(bytes []byte) (*e2sm_kpm_ies.RicEventTriggerStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicEventTriggerStyleListItem((*C.RIC_EventTriggerStyle_List_t)(unsafePtr)), nil
}

func perDecodeRicEventTriggerStyleItem(bytes []byte) (*e2sm_kpm_ies.RicEventTriggerStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicEventTriggerStyleListItem((*C.RIC_EventTriggerStyle_List_t)(unsafePtr)), nil
}

func newRicEventTriggerStyleList(ricEventTriggerStyleList *e2sm_kpm_ies.RicEventTriggerStyleList) *C.RIC_EventTriggerStyle_List_t {

	ricEventTriggerStyleTypeC := newRicStyleType(ricEventTriggerStyleList.RicEventTriggerStyleType)
	ricEventTriggerStyleNameC := newRicStyleName(ricEventTriggerStyleList.RicEventTriggerStyleName)
	ricEventTriggerFormatTypeC := newRicFormatType(ricEventTriggerStyleList.RicEventTriggerFormatType)

	ricEventTriggerStyleListC := C.RIC_EventTriggerStyle_List_t{
		ric_EventTriggerStyle_Type:  *ricEventTriggerStyleTypeC,
		ric_EventTriggerStyle_Name:  *ricEventTriggerStyleNameC,
		ric_EventTriggerFormat_Type: *ricEventTriggerFormatTypeC,
	}

	return &ricEventTriggerStyleListC
}

func decodeRicEventTriggerStyleListItem(ricEventTriggerStyleListC *C.RIC_EventTriggerStyle_List_t) *e2sm_kpm_ies.RicEventTriggerStyleList {

	ricStyleType := decodeRicStyleType(&ricEventTriggerStyleListC.ric_EventTriggerStyle_Type)
	ricStyleName := decodeRicStyleName(&ricEventTriggerStyleListC.ric_EventTriggerStyle_Name)
	ricFormatType := decodeRicFormatType(&ricEventTriggerStyleListC.ric_EventTriggerFormat_Type)

	ricEventTriggerStyleList := e2sm_kpm_ies.RicEventTriggerStyleList{
		RicEventTriggerStyleType:  ricStyleType,
		RicEventTriggerStyleName:  ricStyleName,
		RicEventTriggerFormatType: ricFormatType,
	}
	return &ricEventTriggerStyleList
}
