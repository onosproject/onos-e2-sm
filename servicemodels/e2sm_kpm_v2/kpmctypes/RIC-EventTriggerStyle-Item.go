// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-EventTriggerStyle-Item.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_v2.RicEventTriggerStyleItem) ([]byte, error) {
	ricEventTriggerStyleItemCP, err := newRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_EventTriggerStyle_Item, unsafe.Pointer(ricEventTriggerStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_v2.RicEventTriggerStyleItem) ([]byte, error) {
	ricEventTriggerStyleItemCP, err := newRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicEventTriggerStyleItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_EventTriggerStyle_Item, unsafe.Pointer(ricEventTriggerStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicEventTriggerStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicEventTriggerStyleItem(bytes []byte) (*e2sm_kpm_v2.RicEventTriggerStyleItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_EventTriggerStyle_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicEventTriggerStyleItem((*C.RIC_EventTriggerStyle_Item_t)(unsafePtr))
}

func perDecodeRicEventTriggerStyleItem(bytes []byte) (*e2sm_kpm_v2.RicEventTriggerStyleItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_EventTriggerStyle_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicEventTriggerStyleItem((*C.RIC_EventTriggerStyle_Item_t)(unsafePtr))
}

func newRicEventTriggerStyleItem(ricEventTriggerStyleItem *e2sm_kpm_v2.RicEventTriggerStyleItem) (*C.RIC_EventTriggerStyle_Item_t, error) {

	ricEventTriggerStyleTypeC, err := newRicStyleType(ricEventTriggerStyleItem.RicEventTriggerStyleType)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleType() %s", err.Error())
	}

	ricEventTriggerStyleNameC, err := newRicStyleName(ricEventTriggerStyleItem.RicEventTriggerStyleName)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleName() %s", err.Error())
	}

	ricEventTriggerFormatTypeC, err := newRicFormatType(ricEventTriggerStyleItem.RicEventTriggerFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	ricEventTriggerStyleItemC := C.RIC_EventTriggerStyle_Item_t{
		ric_EventTriggerStyle_Type:  *ricEventTriggerStyleTypeC,
		ric_EventTriggerStyle_Name:  *ricEventTriggerStyleNameC,
		ric_EventTriggerFormat_Type: *ricEventTriggerFormatTypeC,
	}

	return &ricEventTriggerStyleItemC, nil
}

func decodeRicEventTriggerStyleItem(ricEventTriggerStyleItemC *C.RIC_EventTriggerStyle_Item_t) (*e2sm_kpm_v2.RicEventTriggerStyleItem, error) {

	ricEventTriggerStyleType, err := decodeRicStyleType(&ricEventTriggerStyleItemC.ric_EventTriggerStyle_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleType() %s", err.Error())
	}

	ricEventTriggerStyleName, err := decodeRicStyleName(&ricEventTriggerStyleItemC.ric_EventTriggerStyle_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleName() %s", err.Error())
	}

	ricEventTriggerFormatType, err := decodeRicFormatType(&ricEventTriggerStyleItemC.ric_EventTriggerFormat_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	ricEventTriggerStyleItem := e2sm_kpm_v2.RicEventTriggerStyleItem{
		RicEventTriggerStyleType:  ricEventTriggerStyleType,
		RicEventTriggerStyleName:  ricEventTriggerStyleName,
		RicEventTriggerFormatType: ricEventTriggerFormatType,
	}

	return &ricEventTriggerStyleItem, nil
}
