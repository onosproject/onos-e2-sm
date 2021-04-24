// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-RANfunction-Description.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
//#include ".h" //ToDo - include correct .h file for corresponding C-struct of "Repeated" field or other anonymous structure defined in .h file
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) ([]byte, error) {
	e2SmMhoRanfunctionDescriptionCP, err := newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_RANfunction_Description, unsafe.Pointer(e2SmMhoRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) ([]byte, error) {
	e2SmMhoRanfunctionDescriptionCP, err := newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_RANfunction_Description, unsafe.Pointer(e2SmMhoRanfunctionDescriptionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoRanfunctionDescription() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoRanfunctionDescription(bytes []byte) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoRanfunctionDescription((*C.E2SM_MHO_RANfunction_Description_t)(unsafePtr))
}

func perDecodeE2SmMhoRanfunctionDescription(bytes []byte) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_RANfunction_Description)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoRanfunctionDescription((*C.E2SM_MHO_RANfunction_Description_t)(unsafePtr))
}

func newE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription *e2sm_mho.E2SmMhoRanfunctionDescription) (*C.E2SM_MHO_RANfunction_Description_t, error) {

	var err error
	e2SmMhoRanfunctionDescriptionC := C.E2SM_MHO_RANfunction_Description_t{}

	ricEventTriggerStyleListC := new(C.struct_E2SM_MHO_RANfunction_Description__E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item) //ToDo - verify correctness of the variable's name
	for _, ricEventTriggerStyleListItem := range e2SmMhoRanfunctionDescription.GetRicEventTriggerStyleList() {                               //ToDo - Verify if GetSmth() function is called correctly
		ricEventTriggerStyleListItemC, err := newRicEventTriggerStyleList(ricEventTriggerStyleListItem)
		if err != nil {
			return nil, fmt.Errorf("newRicEventTriggerStyleList() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(ricEventTriggerStyleListC), unsafe.Pointer(ricEventTriggerStyleListItemC)); err != nil {
			return nil, err
		}
	}

	ricReportStyleListC := new(C.struct_E2SM_MHO_RANfunction_Description__E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item) //ToDo - verify correctness of the variable's name
	for _, ricReportStyleListItem := range e2SmMhoRanfunctionDescription.GetRicReportStyleList() {                                     //ToDo - Verify if GetSmth() function is called correctly
		ricReportStyleListItemC, err := newRicReportStyleList(ricReportStyleListItem)
		if err != nil {
			return nil, fmt.Errorf("newRicReportStyleList() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(ricReportStyleListC), unsafe.Pointer(ricReportStyleListItemC)); err != nil {
			return nil, err
		}
	}

	ranFunctionNameC, err := newRanfunctionName(e2SmMhoRanfunctionDescription.RanFunctionName)
	if err != nil {
		return nil, fmt.Errorf("newRanfunctionName() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoRanfunctionDescriptionC.ranFunction_Name = ranFunctionNameC
	e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item = ricEventTriggerStyleListC
	e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item = ricReportStyleListC

	return &e2SmMhoRanfunctionDescriptionC, nil
}

func decodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescriptionC *C.E2SM_MHO_RANfunction_Description_t) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {

	var err error
	e2SmMhoRanfunctionDescription := e2sm_mho.E2SmMhoRanfunctionDescription{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//RanFunctionName: ranFunctionName,
		RicEventTriggerStyleList: make([]*e2sm_mho.RicEventTriggerStyleList, 0), //ToDo - Check if protobuf structure is implemented correctly (mainly naming)
		RicReportStyleList:       make([]*e2sm_mho.RicReportStyleList, 0),       //ToDo - Check if protobuf structure is implemented correctly (mainly naming)

	}

	e2SmMhoRanfunctionDescription.RanFunctionName, err = decodeRanfunctionName(e2SmMhoRanfunctionDescriptionC.ranFunction_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRanfunctionName() %s", err.Error())
	}

	var ieCount int

	ieCount = int(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.array)) * uintptr(i)
		ieC := *(**C.RicEventTriggerStyleList_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.array)) + offset))
		ie, err := decodeRicEventTriggerStyleList(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicEventTriggerStyleList() %s", err.Error())
		}
		e2SmMhoRanfunctionDescription.RicEventTriggerStyleList = append(e2SmMhoRanfunctionDescription.RicEventTriggerStyleList, ie)
	}

	ieCount = int(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.array)) * uintptr(i)
		ieC := *(**C.RicReportStyleList_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmMhoRanfunctionDescriptionC.E2SM_MHO_RANfunction_Description__e2SM_MHO_RANfunction_Item.list.array)) + offset))
		ie, err := decodeRicReportStyleList(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicReportStyleList() %s", err.Error())
		}
		e2SmMhoRanfunctionDescription.RicReportStyleList = append(e2SmMhoRanfunctionDescription.RicReportStyleList, ie)
	}

	return &e2SmMhoRanfunctionDescription, nil
}

func decodeE2SmMhoRanfunctionDescriptionBytes(array [8]byte) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {
	e2SmMhoRanfunctionDescriptionC := (*C.E2SM_MHO_RANfunction_Description_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescriptionC)
}
