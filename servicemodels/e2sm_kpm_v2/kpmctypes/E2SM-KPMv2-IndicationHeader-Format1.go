// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPMv2-IndicationHeader-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1 *e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1) ([]byte, error) {
	e2SmKpmIndicationHeaderFormat1CP, err := newE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPMv2_IndicationHeader_Format1, unsafe.Pointer(e2SmKpmIndicationHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1 *e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1) ([]byte, error) {
	e2SmKpmIndicationHeaderFormat1CP, err := newE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPMv2_IndicationHeader_Format1, unsafe.Pointer(e2SmKpmIndicationHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmIndicationHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmIndicationHeaderFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPMv2_IndicationHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmIndicationHeaderFormat1((*C.E2SM_KPMv2_IndicationHeader_Format1_t)(unsafePtr))
}

func perDecodeE2SmKpmIndicationHeaderFormat1(bytes []byte) (*e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPMv2_IndicationHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmIndicationHeaderFormat1((*C.E2SM_KPMv2_IndicationHeader_Format1_t)(unsafePtr))
}

func newE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1 *e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1) (*C.E2SM_KPMv2_IndicationHeader_Format1_t, error) {

	colletStartTimeC, err := newTimeStamp(e2SmKpmIndicationHeaderFormat1.ColletStartTime)
	if err != nil {
		return nil, fmt.Errorf("newTimeStamp() %s", err.Error())
	}

	e2SmKpmIndicationHeaderFormat1C := C.E2SM_KPMv2_IndicationHeader_Format1_t{
		colletStartTime: *colletStartTimeC,
		//fileFormatversion: fileFormatversionC,
		//senderName:        senderNameC,
		//senderType:        senderTypeC,
		//vendorName:        vendorNameC,
		//kpmNodeID:         kpmNodeIDC,
	}

	if e2SmKpmIndicationHeaderFormat1.FileFormatversion != "" {
		e2SmKpmIndicationHeaderFormat1C.fileFormatversion, err = newPrintableString(e2SmKpmIndicationHeaderFormat1.FileFormatversion)
		if err != nil {
			return nil, fmt.Errorf("newPrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1.SenderName != "" {
		e2SmKpmIndicationHeaderFormat1C.senderName, err = newPrintableString(e2SmKpmIndicationHeaderFormat1.SenderName)
		if err != nil {
			return nil, fmt.Errorf("newPrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1.SenderType != "" {
		e2SmKpmIndicationHeaderFormat1C.senderType, err = newPrintableString(e2SmKpmIndicationHeaderFormat1.SenderType)
		if err != nil {
			return nil, fmt.Errorf("newPrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1.VendorName != "" {
		e2SmKpmIndicationHeaderFormat1C.vendorName, err = newPrintableString(e2SmKpmIndicationHeaderFormat1.VendorName)
		if err != nil {
			return nil, fmt.Errorf("newPrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1.KpmNodeId != nil {
		e2SmKpmIndicationHeaderFormat1C.kpmNodeID, err = newGlobalKpmnodeID(e2SmKpmIndicationHeaderFormat1.KpmNodeId)
		if err != nil {
			return nil, fmt.Errorf("newGlobalKpmnodeID() %s", err.Error())
		}
	}

	return &e2SmKpmIndicationHeaderFormat1C, nil
}

func decodeE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1C *C.E2SM_KPMv2_IndicationHeader_Format1_t) (*e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1, error) {

	colletStartTime, err := decodeTimeStamp(&e2SmKpmIndicationHeaderFormat1C.colletStartTime)
	if err != nil {
		return nil, fmt.Errorf("decodeTimeStamp() %s", err.Error())
	}

	e2SmKpmIndicationHeaderFormat1 := e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1{
		ColletStartTime: colletStartTime,
		//FileFormatversion: fileFormatversion,
		//SenderName:        senderName,
		//SenderType:        senderType,
		//VendorName:        vendorName,
		//KpmNodeId:         kpmNodeID,
	}

	if e2SmKpmIndicationHeaderFormat1C.fileFormatversion != nil {
		e2SmKpmIndicationHeaderFormat1.FileFormatversion, err = decodePrintableString(e2SmKpmIndicationHeaderFormat1C.fileFormatversion)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1C.senderName != nil {
		e2SmKpmIndicationHeaderFormat1.SenderName, err = decodePrintableString(e2SmKpmIndicationHeaderFormat1C.senderName)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1C.senderType != nil {
		e2SmKpmIndicationHeaderFormat1.SenderType, err = decodePrintableString(e2SmKpmIndicationHeaderFormat1C.senderType)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1C.vendorName != nil {
		e2SmKpmIndicationHeaderFormat1.VendorName, err = decodePrintableString(e2SmKpmIndicationHeaderFormat1C.vendorName)
		if err != nil {
			return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
		}
	}
	if e2SmKpmIndicationHeaderFormat1C.kpmNodeID != nil {
		e2SmKpmIndicationHeaderFormat1.KpmNodeId, err = decodeGlobalKpmnodeID(e2SmKpmIndicationHeaderFormat1C.kpmNodeID)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalKpmnodeId() %s", err.Error())
		}
	}

	return &e2SmKpmIndicationHeaderFormat1, nil
}

func decodeE2SmKpmIndicationHeaderFormat1Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmIndicationHeaderFormat1, error) {
	e2SmKpmIndicationHeaderFormat1C := (*C.E2SM_KPMv2_IndicationHeader_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmIndicationHeaderFormat1(e2SmKpmIndicationHeaderFormat1C)
}
