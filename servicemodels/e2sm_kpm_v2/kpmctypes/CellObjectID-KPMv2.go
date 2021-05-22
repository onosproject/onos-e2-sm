// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CellObjectID-KPMv2.h"
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

	bytes, err := encodeXer(&C.asn_DEF_CellObjectID_KPMv2, unsafe.Pointer(cellObjectIDCP))
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

	bytes, err := encodePerBuffer(&C.asn_DEF_CellObjectID_KPMv2, unsafe.Pointer(cellObjectIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellObjectID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCellObjectID(bytes []byte) (*e2sm_kpm_v2.CellObjectId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CellObjectID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCellObjectID((*C.CellObjectID_KPMv2_t)(unsafePtr))
}

func perDecodeCellObjectID(bytes []byte) (*e2sm_kpm_v2.CellObjectId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CellObjectID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCellObjectID((*C.CellObjectID_KPMv2_t)(unsafePtr))
}

func newCellObjectID(cellObjectID *e2sm_kpm_v2.CellObjectId) (*C.CellObjectID_KPMv2_t, error) {

	cellObjectIDC, err := newPrintableString(cellObjectID.Value)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	return cellObjectIDC, nil
}

func decodeCellObjectID(cellObjectIDC *C.CellObjectID_KPMv2_t) (*e2sm_kpm_v2.CellObjectId, error) {

	cellObjectID := new(e2sm_kpm_v2.CellObjectId)
	cellObjectIDValue, err := decodePrintableString(cellObjectIDC)
	if err != nil {
		return nil, fmt.Errorf("decodeString() %s", err.Error())
	}
	cellObjectID.Value = cellObjectIDValue

	return cellObjectID, nil
}
