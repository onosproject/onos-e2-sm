// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CellObjectID.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeCellObjectID(cellObjectID *e2sm_kpm_v2.CellObjectId) ([]byte, error) {
	cellObjectIDCP, err := newCellObjectID(cellObjectID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellObjectID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_CellObjectID, unsafe.Pointer(cellObjectIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellObjectID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCellObjectID(cellObjectID *e2sm_kpm_v2.CellObjectId) ([]byte, error) {
	cellObjectIDCP, err := newCellObjectID(cellObjectID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellObjectID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CellObjectID, unsafe.Pointer(cellObjectIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellObjectID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCellObjectID(bytes []byte) (*e2sm_kpm_v2.CellObjectId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CellObjectID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCellObjectID((*C.CellObjectID_t)(unsafePtr))
}

func perDecodeCellObjectID(bytes []byte) (*e2sm_kpm_v2.CellObjectId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CellObjectID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCellObjectID((*C.CellObjectID_t)(unsafePtr))
}

func newCellObjectID(cellObjectID *e2sm_kpm_v2.CellObjectId) (*C.CellObjectID_t, error) {

	//fmt.Printf("We're inside newCellObjectID(), starting encoding...")
	//fmt.Printf("Encoding PrintableString: \n %v \n", cellObjectID.Value)
	cellObjectIDC, err := newPrintableString(cellObjectID.Value)
	//fmt.Printf("That's the C version of encoded PrintableString: \n %v \n", cellObjectIDC)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	return cellObjectIDC, nil
}

func decodeCellObjectID(cellObjectIDC *C.CellObjectID_t) (*e2sm_kpm_v2.CellObjectId, error) {

	//fmt.Printf("We're inside decodeCellMeasurementObjectItem(), starting decoding...")
	cellObjectID := new(e2sm_kpm_v2.CellObjectId)
	//fmt.Printf("Decoding PrintableString: \n %v \n", cellObjectIDC)
	cellObjectIDValue, err := decodePrintableString(cellObjectIDC)
	//fmt.Printf("That's what was decoded from C: \n %v \n", cellObjectIDValue)
	if err != nil {
		return nil, fmt.Errorf("decodeString() %s", err.Error())
	}
	cellObjectID.Value = cellObjectIDValue

	//fmt.Printf("This is final decoded structure: \n %v \n", cellObjectID)
	return cellObjectID, nil
}
