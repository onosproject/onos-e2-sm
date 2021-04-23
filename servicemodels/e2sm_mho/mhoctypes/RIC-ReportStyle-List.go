// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-ReportStyle-List.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRicReportStyleList(ricReportStyleList *e2sm_mho.RicReportStyleList) ([]byte, error) {
	ricReportStyleListCP, err := newRicReportStyleList(ricReportStyleList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicReportStyleList(ricReportStyleList *e2sm_mho.RicReportStyleList) ([]byte, error) {
	ricReportStyleListCP, err := newRicReportStyleList(ricReportStyleList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicReportStyleList(bytes []byte) (*e2sm_mho.RicReportStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicReportStyleList((*C.RIC_ReportStyle_List_t)(unsafePtr))
}

func perDecodeRicReportStyleList(bytes []byte) (*e2sm_mho.RicReportStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicReportStyleList((*C.RIC_ReportStyle_List_t)(unsafePtr))
}

func newRicReportStyleList(ricReportStyleList *e2sm_mho.RicReportStyleList) (*C.RIC_ReportStyle_List_t, error) {

	var err error
	ricReportStyleListC := C.RIC_ReportStyle_List_t{}

	ricReportStyleTypeC, err := newRicStyleType(ricReportStyleList.RicReportStyleType)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleType() %s", err.Error())
	}

	ricReportStyleNameC, err := newRicStyleName(ricReportStyleList.RicReportStyleName)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleName() %s", err.Error())
	}

	ricIndicationHeaderFormatTypeC, err := newRicFormatType(ricReportStyleList.RicIndicationHeaderFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	ricIndicationMessageFormatTypeC, err := newRicFormatType(ricReportStyleList.RicIndicationMessageFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	ricReportStyleListC.RIC_ReportStyle_List = ricReportStyleTypeC
	ricReportStyleListC.RIC_ReportStyle_List = ricReportStyleNameC
	ricReportStyleListC.RIC_ReportStyle_List = ricIndicationHeaderFormatTypeC
	ricReportStyleListC.RIC_ReportStyle_List = ricIndicationMessageFormatTypeC

	return &ricReportStyleListC, nil
}

func decodeRicReportStyleList(ricReportStyleListC *C.RIC_ReportStyle_List_t) (*e2sm_mho.RicReportStyleList, error) {

	var err error
	ricReportStyleList := e2sm_mho.RicReportStyleList{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//RicReportStyleType: ricReportStyleType,
		//RicReportStyleName: ricReportStyleName,
		//RicIndicationHeaderFormatType: ricIndicationHeaderFormatType,
		//RicIndicationMessageFormatType: ricIndicationMessageFormatType,

	}

	ricReportStyleList.RicReportStyleType, err = decodeRicStyleType(ricReportStyleListC.RIC_ReportStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleType() %s", err.Error())
	}

	ricReportStyleList.RicReportStyleName, err = decodeRicStyleName(ricReportStyleListC.RIC_ReportStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleName() %s", err.Error())
	}

	ricReportStyleList.RicIndicationHeaderFormatType, err = decodeRicFormatType(ricReportStyleListC.RIC_ReportStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	ricReportStyleList.RicIndicationMessageFormatType, err = decodeRicFormatType(ricReportStyleListC.RIC_ReportStyle_List)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	return &ricReportStyleList, nil
}

func decodeRicReportStyleListBytes(array [8]byte) (*e2sm_mho.RicReportStyleList, error) {
	ricReportStyleListC := (*C.RIC_ReportStyle_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicReportStyleList(ricReportStyleListC)
}
