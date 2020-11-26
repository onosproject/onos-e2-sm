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
	ricEventTriggerStyleItemCP := newRicEventTriggerStyleItem(ricEventTriggerStyleItem)

	bytes, err := encodeXer(&C.asn_DEF_RIC_EventTriggerStyle_List, unsafe.Pointer(ricEventTriggerStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_ies.RicEventTriggerStyleList) ([]byte, error) {
	ricEventTriggerStyleItemCP := newRicEventTriggerStyleItem(ricEventTriggerStyleItem)

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
	return decodeRicEventTriggerStyleItem((*C.RIC_EventTriggerStyle_List_t)(unsafePtr)), nil
}

func perDecodeRicEventTriggerStyleItem(bytes []byte) (*e2sm_kpm_ies.RicEventTriggerStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicEventTriggerStyleItem((*C.RIC_EventTriggerStyle_List_t)(unsafePtr)), nil
}

func newRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_ies.RicEventTriggerStyleList) *C.RIC_EventTriggerStyle_List_t {

	ricEventTriggerStyleType := newRicStyleType(ricEventTriggerStyleItem.RicEventTriggerStyleType)
	ricEventTriggerStyleName := newRicStyleName(ricEventTriggerStyleItem.RicEventTriggerStyleName)
	ricEventTriggerFormatType := newRicFormatType(ricEventTriggerStyleItem.RicEventTriggerFormatType)

	ricEventTriggerStyleItemC := C.RIC_EventTriggerStyle_List_t{
		ric_EventTriggerStyle_Type:  *ricEventTriggerStyleType,
		ric_EventTriggerStyle_Name:  *ricEventTriggerStyleName,
		ric_EventTriggerFormat_Type: *ricEventTriggerFormatType,
	}

	return &ricEventTriggerStyleItemC
}

func decodeRicEventTriggerStyleItem(ricEventTriggerStyleItemC *C.RIC_EventTriggerStyle_List_t) *e2sm_kpm_ies.RicEventTriggerStyleList {

	ricStyleType := decodeRicStyleType((*C.RIC_Style_Type_t)(*ricEventTriggerStyleItemC.ric_EventTriggerStyle_Type))
	ricStyleName := decodeRicStyleName((*C.RIC_Style_Name_t)(*ricEventTriggerStyleItemC.ric_EventTriggerStyle_Name))
	ricFormatType := decodeRicFormatType((*C.RIC_Format_Type_t)(*ricEventTriggerStyleItemC.ric_EventTriggerFormat_Type))

	ricEventTriggerStyleItem := e2sm_kpm_ies.RicEventTriggerStyleList{
		RicEventTriggerStyleType:  ricStyleType,
		RicEventTriggerStyleName:  ricStyleName,
		RicEventTriggerFormatType: ricFormatType,
	}
	return &ricEventTriggerStyleItem
}
