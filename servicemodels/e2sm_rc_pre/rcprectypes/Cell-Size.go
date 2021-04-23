// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Cell-Size.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

//ToDo: Solve "Cannot convert cellSizeC (type _Ctype_long) to type unsafe.Pointer"
func xerEncodeCellSize(cellSize e2sm_rc_pre_v2.CellSize) ([]byte, error) {
	cellSizeC, err := newCellSize(cellSize)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_Cell_Size, unsafe.Pointer(&cellSizeC))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellSize() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCellSize(cellSize e2sm_rc_pre_v2.CellSize) ([]byte, error) {
	cellSizeC, err := newCellSize(cellSize)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Cell_Size, unsafe.Pointer(&cellSizeC))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellSize() %s", err.Error())
	}
	return bytes, nil
}

//ToDo: Decide what to return instead of nil, which cannot be returned since return value is not a pointer anymore
func xerDecodeCellSize(bytes []byte) (e2sm_rc_pre_v2.CellSize, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Cell_Size)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCellSize((*C.Cell_Size_t)(unsafePtr)), nil
}

func perDecodeCellSize(bytes []byte) (e2sm_rc_pre_v2.CellSize, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Cell_Size)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCellSize((*C.Cell_Size_t)(unsafePtr)), nil
}

func newCellSize(cellSize e2sm_rc_pre_v2.CellSize) (C.Cell_Size_t, error) {
	var ret C.Cell_Size_t
	switch cellSize {
	case e2sm_rc_pre_v2.CellSize_CELL_SIZE_FEMTO:
		ret = C.Cell_Size_femto
	case e2sm_rc_pre_v2.CellSize_CELL_SIZE_ENTERPRISE:
		ret = C.Cell_Size_enterprise
	case e2sm_rc_pre_v2.CellSize_CELL_SIZE_OUTDOOR_SMALL:
		ret = C.Cell_Size_outdoorSmall
	case e2sm_rc_pre_v2.CellSize_CELL_SIZE_MACRO:
		ret = C.Cell_Size_macro
	default:
		return 0, fmt.Errorf("unexpected CellSize %v", cellSize)
	}
	return ret, nil
}

func decodeCellSize(cellSizeC *C.Cell_Size_t) e2sm_rc_pre_v2.CellSize {

	return e2sm_rc_pre_v2.CellSize(int32(*cellSizeC))
}
