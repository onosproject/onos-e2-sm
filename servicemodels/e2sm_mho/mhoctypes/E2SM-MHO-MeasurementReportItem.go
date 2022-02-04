// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-MeasurementReportItem.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem *e2sm_mho.E2SmMhoMeasurementReportItem) ([]byte, error) {
	e2SmMhoMeasurementReportItemCP, err := newE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoMeasurementReportItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_MeasurementReportItem, unsafe.Pointer(e2SmMhoMeasurementReportItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoMeasurementReportItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem *e2sm_mho.E2SmMhoMeasurementReportItem) ([]byte, error) {
	e2SmMhoMeasurementReportItemCP, err := newE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoMeasurementReportItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_MeasurementReportItem, unsafe.Pointer(e2SmMhoMeasurementReportItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoMeasurementReportItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoMeasurementReportItem(bytes []byte) (*e2sm_mho.E2SmMhoMeasurementReportItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_MeasurementReportItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoMeasurementReportItem((*C.E2SM_MHO_MeasurementReportItem_t)(unsafePtr))
}

func perDecodeE2SmMhoMeasurementReportItem(bytes []byte) (*e2sm_mho.E2SmMhoMeasurementReportItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_MeasurementReportItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoMeasurementReportItem((*C.E2SM_MHO_MeasurementReportItem_t)(unsafePtr))
}

func newE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem *e2sm_mho.E2SmMhoMeasurementReportItem) (*C.E2SM_MHO_MeasurementReportItem_t, error) {

	var err error
	e2SmMhoMeasurementReportItemC := C.E2SM_MHO_MeasurementReportItem_t{}

	cgiC, err := newCellGlobalID(e2SmMhoMeasurementReportItem.Cgi)
	if err != nil {
		return nil, fmt.Errorf("newCellGlobalID() %s", err.Error())
	}

	rsrpC, err := newRsrp(e2SmMhoMeasurementReportItem.Rsrp)
	if err != nil {
		return nil, fmt.Errorf("newRsrp() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoMeasurementReportItemC.cgi = *cgiC
	e2SmMhoMeasurementReportItemC.rsrp = *rsrpC

	return &e2SmMhoMeasurementReportItemC, nil
}

func decodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItemC *C.E2SM_MHO_MeasurementReportItem_t) (*e2sm_mho.E2SmMhoMeasurementReportItem, error) {

	var err error
	e2SmMhoMeasurementReportItem := e2sm_mho.E2SmMhoMeasurementReportItem{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//Cgi: cgi,
		//Rsrp: rsrp,

	}

	e2SmMhoMeasurementReportItem.Cgi, err = decodeCellGlobalID(&e2SmMhoMeasurementReportItemC.cgi)
	if err != nil {
		return nil, fmt.Errorf("decodeCellGlobalID() %s", err.Error())
	}

	e2SmMhoMeasurementReportItem.Rsrp, err = decodeRsrp(&e2SmMhoMeasurementReportItemC.rsrp)
	if err != nil {
		return nil, fmt.Errorf("decodeRsrp() %s", err.Error())
	}

	return &e2SmMhoMeasurementReportItem, nil
}

//func decodeE2SmMhoMeasurementReportItemBytes(array [8]byte) (*e2sm_mho.E2SmMhoMeasurementReportItem, error) {
//	e2SmMhoMeasurementReportItemC := (*C.E2SM_MHO_MeasurementReportItem_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
//
//	return decodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItemC)
//}
