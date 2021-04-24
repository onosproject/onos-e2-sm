// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-EventTriggerStyle-List.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRicEventTriggerStyleList(ricEventTriggerStyleList *e2sm_mho.RicEventTriggerStyleList) ([]byte, error) {
	ricEventTriggerStyleListCP, err := newRicEventTriggerStyleList(ricEventTriggerStyleList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_EventTriggerStyle_List, unsafe.Pointer(ricEventTriggerStyleListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicEventTriggerStyleList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicEventTriggerStyleList(ricEventTriggerStyleList *e2sm_mho.RicEventTriggerStyleList) ([]byte, error) {
	ricEventTriggerStyleListCP, err := newRicEventTriggerStyleList(ricEventTriggerStyleList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicEventTriggerStyleList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_EventTriggerStyle_List, unsafe.Pointer(ricEventTriggerStyleListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicEventTriggerStyleList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicEventTriggerStyleList(bytes []byte) (*e2sm_mho.RicEventTriggerStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicEventTriggerStyleList((*C.RIC_EventTriggerStyle_List_t)(unsafePtr))
}

func perDecodeRicEventTriggerStyleList(bytes []byte) (*e2sm_mho.RicEventTriggerStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicEventTriggerStyleList((*C.RIC_EventTriggerStyle_List_t)(unsafePtr))
}

func newRicEventTriggerStyleList(ricEventTriggerStyleList *e2sm_mho.RicEventTriggerStyleList) (*C.RIC_EventTriggerStyle_List_t, error) {

	var err error
	ricEventTriggerStyleListC := C.RIC_EventTriggerStyle_List_t{}

	ricEventTriggerStyleTypeC, err := newRicStyleType(ricEventTriggerStyleList.RicEventTriggerStyleType)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleType() %s", err.Error())
	}

	ricEventTriggerStyleNameC, err := newRicStyleName(ricEventTriggerStyleList.RicEventTriggerStyleName)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleName() %s", err.Error())
	}

	ricEventTriggerFormatTypeC, err := newRicFormatType(ricEventTriggerStyleList.RicEventTriggerFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	ricEventTriggerStyleListC.RIC_EventTriggerStyle_List = ricEventTriggerStyleTypeC
	ricEventTriggerStyleListC.RIC_EventTriggerStyle_List = ricEventTriggerStyleNameC
	ricEventTriggerStyleListC.RIC_EventTriggerStyle_List = ricEventTriggerFormatTypeC

	return &ricEventTriggerStyleListC, nil
}

func decodeRicEventTriggerStyleList(ricEventTriggerStyleListC *C.RIC_EventTriggerStyle_List_t) (*e2sm_mho.RicEventTriggerStyleList, error) {

	var err error
	ricEventTriggerStyleList := e2sm_mho.RicEventTriggerStyleList{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//RicEventTriggerStyleType: ricEventTriggerStyleType,
		//RicEventTriggerStyleName: ricEventTriggerStyleName,
		//RicEventTriggerFormatType: ricEventTriggerFormatType,

	}

	ricEventTriggerStyleList.RicEventTriggerStyleType, err = decodeRicStyleType(ricEventTriggerStyleListC.RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleType() %s", err.Error())
	}

	ricEventTriggerStyleList.RicEventTriggerStyleName, err = decodeRicStyleName(ricEventTriggerStyleListC.RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleName() %s", err.Error())
	}

	ricEventTriggerStyleList.RicEventTriggerFormatType, err = decodeRicFormatType(ricEventTriggerStyleListC.RIC_EventTriggerStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	return &ricEventTriggerStyleList, nil
}

func decodeRicEventTriggerStyleListBytes(array [8]byte) (*e2sm_mho.RicEventTriggerStyleList, error) {
	ricEventTriggerStyleListC := (*C.RIC_EventTriggerStyle_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicEventTriggerStyleList(ricEventTriggerStyleListC)
}
