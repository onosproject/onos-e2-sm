// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-ReportStyle-List.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRicReportStyleItem(ricReportStyleItem *e2sm_kpm_ies.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleItem(ricReportStyleItem)

	bytes, err := encodeXer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicReportStyleItem(ricReportStyleItem *e2sm_kpm_ies.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleItem(ricReportStyleItem)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicReportStyleItem(bytes []byte) (*e2sm_kpm_ies.RicReportStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicReportStyleItem((*C.RIC_ReportStyle_List_t)(unsafePtr)), nil
}

func perDecodeRicReportStyleItem(bytes []byte) (*e2sm_kpm_ies.RicReportStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicReportStyleItem((*C.RIC_ReportStyle_List_t)(unsafePtr)), nil
}

func newRicReportStyleItem(ricReportStyleItem *e2sm_kpm_ies.RicReportStyleList) *C.RIC_ReportStyle_List_t {

	ricReportStyleTypeC := newRicStyleType(ricReportStyleItem.RicReportStyleType)
	ricReportStyleNameC := newRicStyleName(ricReportStyleItem.RicReportStyleName)
	ricReportStyleIndicationHeaderFormatTypeC := newRicFormatType(ricReportStyleItem.RicIndicationHeaderFormatType)
	ricReportStyleIndicationMessageFormatTypeC := newRicFormatType(ricReportStyleItem.RicIndicationMessageFormatType)

	ricReportStyleItemC := C.RIC_ReportStyle_List_t{
		ric_ReportStyle_Type:             *ricReportStyleTypeC,
		ric_ReportStyle_Name:             *ricReportStyleNameC,
		ric_IndicationHeaderFormat_Type:  *ricReportStyleIndicationHeaderFormatTypeC,
		ric_IndicationMessageFormat_Type: *ricReportStyleIndicationMessageFormatTypeC,
	}

	return &ricReportStyleItemC
}

func decodeRicReportStyleItem(ricReportStyleItemC *C.RIC_ReportStyle_List_t) *e2sm_kpm_ies.RicReportStyleList {

	ricReportStyleType := decodeRicStyleType(&ricReportStyleItemC.ric_ReportStyle_Type)
	ricReportStyleName := decodeRicStyleName(&ricReportStyleItemC.ric_ReportStyle_Name)
	ricReportStyleIndicationHeaderFormatType := decodeRicFormatType(&ricReportStyleItemC.ric_IndicationHeaderFormat_Type)
	ricReportStyleIndicationMessageFormatType := decodeRicFormatType(&ricReportStyleItemC.ric_IndicationMessageFormat_Type)

	ricReportStyleItem := e2sm_kpm_ies.RicReportStyleList{
		RicReportStyleType:             ricReportStyleType,
		RicReportStyleName:             ricReportStyleName,
		RicIndicationHeaderFormatType:  ricReportStyleIndicationHeaderFormatType,
		RicIndicationMessageFormatType: ricReportStyleIndicationMessageFormatType,
	}
	return &ricReportStyleItem
}
