// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Cell-Measurement-Object-Item.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeCellMeasurementObjectItem(cellMeasurementObjectItem *e2sm_kpm_v2.CellMeasurementObjectItem) ([]byte, error) {
	cellMeasurementObjectItemCP, err := newCellMeasurementObjectItem(cellMeasurementObjectItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellMeasurementObjectItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Cell_Measurement_Object_Item, unsafe.Pointer(cellMeasurementObjectItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellMeasurementObjectItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCellMeasurementObjectItem(cellMeasurementObjectItem *e2sm_kpm_v2.CellMeasurementObjectItem) ([]byte, error) {
	cellMeasurementObjectItemCP, err := newCellMeasurementObjectItem(cellMeasurementObjectItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellMeasurementObjectItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Cell_Measurement_Object_Item, unsafe.Pointer(cellMeasurementObjectItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellMeasurementObjectItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCellMeasurementObjectItem(bytes []byte) (*e2sm_kpm_v2.CellMeasurementObjectItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Cell_Measurement_Object_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCellMeasurementObjectItem((*C.Cell_Measurement_Object_Item_t)(unsafePtr))
}

func perDecodeCellMeasurementObjectItem(bytes []byte) (*e2sm_kpm_v2.CellMeasurementObjectItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Cell_Measurement_Object_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCellMeasurementObjectItem((*C.Cell_Measurement_Object_Item_t)(unsafePtr))
}

func newCellMeasurementObjectItem(cellMeasurementObjectItem *e2sm_kpm_v2.CellMeasurementObjectItem) (*C.Cell_Measurement_Object_Item_t, error) {

	cellObjectIDC, err := newCellObjectID(cellMeasurementObjectItem.CellObjectId)
	if err != nil {
		return nil, fmt.Errorf("newCellObjectId() %s", err.Error())
	}

	cellGlobalIDC, err := newCellGlobalID(cellMeasurementObjectItem.CellGlobalId)
	if err != nil {
		return nil, fmt.Errorf("newCellGlobalId() %s", err.Error())
	}

	cellMeasurementObjectItemC := C.Cell_Measurement_Object_Item_t{
		cell_object_ID: *cellObjectIDC,
		cell_global_ID: *cellGlobalIDC,
	}

	return &cellMeasurementObjectItemC, nil
}

func decodeCellMeasurementObjectItem(cellMeasurementObjectItemC *C.Cell_Measurement_Object_Item_t) (*e2sm_kpm_v2.CellMeasurementObjectItem, error) {

	cellObjectID, err := decodeCellObjectID(&cellMeasurementObjectItemC.cell_object_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeCellObjectId() %s", err.Error())
	}

	cellGlobalID, err := decodeCellGlobalID(&cellMeasurementObjectItemC.cell_global_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeCellGlobalId() %s", err.Error())
	}

	cellMeasurementObjectItem := e2sm_kpm_v2.CellMeasurementObjectItem{
		CellObjectId: cellObjectID,
		CellGlobalId: cellGlobalID,
	}

	return &cellMeasurementObjectItem, nil
}
