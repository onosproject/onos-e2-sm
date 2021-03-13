// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-ReportStyle-Item.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRicReportStyleItem(ricReportStyleItem *e2sm_kpm_v2.RicReportStyleItem) ([]byte, error) {
	ricReportStyleItemCP, err := newRicReportStyleItem(ricReportStyleItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_ReportStyle_Item, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicReportStyleItem(ricReportStyleItem *e2sm_kpm_v2.RicReportStyleItem) ([]byte, error) {
	ricReportStyleItemCP, err := newRicReportStyleItem(ricReportStyleItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_ReportStyle_Item, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicReportStyleItem(bytes []byte) (*e2sm_kpm_v2.RicReportStyleItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_ReportStyle_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicReportStyleItem((*C.RIC_ReportStyle_Item_t)(unsafePtr))
}

//func perDecodeRicReportStyleItem(bytes []byte) (*e2sm_kpm_v2.RicReportStyleItem, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_ReportStyle_Item)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRicReportStyleItem((*C.RIC_ReportStyle_Item_t)(unsafePtr))
//}

func newRicReportStyleItem(ricReportStyleItem *e2sm_kpm_v2.RicReportStyleItem) (*C.RIC_ReportStyle_Item_t, error) {

	ricReportStyleTypeC, err := newRicStyleType(ricReportStyleItem.RicReportStyleType)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleType() %s", err.Error())
	}

	ricReportStyleNameC, err := newRicStyleName(ricReportStyleItem.RicReportStyleName)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleName() %s", err.Error())
	}

	ricActionFormatTypeC, err := newRicFormatType(ricReportStyleItem.RicActionFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	measInfoActionListC, err := newMeasurementInfoActionList(ricReportStyleItem.MeasInfoActionList)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementInfoActionList() %s", err.Error())
	}

	ricIndicationHeaderFormatTypeC, err := newRicFormatType(ricReportStyleItem.RicIndicationHeaderFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	ricIndicationMessageFormatTypeC, err := newRicFormatType(ricReportStyleItem.RicIndicationMessageFormatType)
	if err != nil {
		return nil, fmt.Errorf("newRicFormatType() %s", err.Error())
	}

	ricReportStyleItemC := C.RIC_ReportStyle_Item_t{
		ric_ReportStyle_Type:             *ricReportStyleTypeC,
		ric_ReportStyle_Name:             *ricReportStyleNameC,
		ric_ActionFormat_Type:            *ricActionFormatTypeC,
		measInfo_Action_List:             *measInfoActionListC,
		ric_IndicationHeaderFormat_Type:  *ricIndicationHeaderFormatTypeC,
		ric_IndicationMessageFormat_Type: *ricIndicationMessageFormatTypeC,
	}

	return &ricReportStyleItemC, nil
}

func decodeRicReportStyleItem(ricReportStyleItemC *C.RIC_ReportStyle_Item_t) (*e2sm_kpm_v2.RicReportStyleItem, error) {

	ricReportStyleType, err := decodeRicStyleType(&ricReportStyleItemC.ric_ReportStyle_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleType() %s", err.Error())
	}

	ricReportStyleName, err := decodeRicStyleName(&ricReportStyleItemC.ric_ReportStyle_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleName() %s", err.Error())
	}

	ricActionFormatType, err := decodeRicFormatType(&ricReportStyleItemC.ric_ActionFormat_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	measInfoActionList, err := decodeMeasurementInfoActionList(&ricReportStyleItemC.measInfo_Action_List)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementInfoActionList() %s", err.Error())
	}

	ricIndicationHeaderFormatType, err := decodeRicFormatType(&ricReportStyleItemC.ric_IndicationHeaderFormat_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	ricIndicationMessageFormatType, err := decodeRicFormatType(&ricReportStyleItemC.ric_IndicationMessageFormat_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicFormatType() %s", err.Error())
	}

	ricReportStyleItem := e2sm_kpm_v2.RicReportStyleItem{
		RicReportStyleType:             ricReportStyleType,
		RicReportStyleName:             ricReportStyleName,
		RicActionFormatType:            ricActionFormatType,
		MeasInfoActionList:             measInfoActionList,
		RicIndicationHeaderFormatType:  ricIndicationHeaderFormatType,
		RicIndicationMessageFormatType: ricIndicationMessageFormatType,
	}

	return &ricReportStyleItem, nil
}
