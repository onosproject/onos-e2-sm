// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationMessage-Format1-RCPRE.h"
//#include "NRT-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1 *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1) ([]byte, error) {

	E2SmRcPreIndicationMsgFormat1CP, err := newE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage_Format1_RCPRE, unsafe.Pointer(E2SmRcPreIndicationMsgFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func newE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1 *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1) (*C.E2SM_RC_PRE_IndicationMessage_Format1_RCPRE_t, error) {
	neighborsListC := new(C.struct_E2SM_RC_PRE_IndicationMessage_Format1_RCPRE__neighbors)
	arfcnC, err := newARFCN(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetDlArfcn())
	if err != nil {
		return nil, err
	}
	cellSizeC, err := newCellSize(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetCellSize())
	if err != nil {
		return nil, err
	}
	pciC := newPCI(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetPci())

	format1 := C.E2SM_RC_PRE_IndicationMessage_Format1_RCPRE_t{
		dl_ARFCN:  *arfcnC,
		cell_Size: cellSizeC,
		pci:       *pciC,
	}

	for _, neighbor := range E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetNeighbors() {
		neighborC, err := newNRT(neighbor)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreIndicationMessageFormat1() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(neighborsListC), unsafe.Pointer(neighborC)); err != nil {
			return nil, err
		}
	}
	format1.neighbors = *neighborsListC

	return &format1, nil
}

func decodeE2SmRcPreIndicationMessageFormat1(e2SmIindicationMsgFormat1C *C.E2SM_RC_PRE_IndicationMessage_Format1_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1, error) {

	arfcn, err := decodeARFCN(&e2SmIindicationMsgFormat1C.dl_ARFCN)
	if err != nil {
		return nil, err
	}
	cellSize := decodeCellSize(&e2SmIindicationMsgFormat1C.cell_Size)
	pci := decodePCI(&e2SmIindicationMsgFormat1C.pci)

	e2SmIindicationMsgFormat1 := &e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_v2.E2SmRcPreIndicationMessageFormat1{
			DlArfcn:   arfcn,
			CellSize:  cellSize,
			Pci:       pci,
			Neighbors: make([]*e2sm_rc_pre_v2.Nrt, 0),
		},
	}

	nbrCount := int(e2SmIindicationMsgFormat1C.neighbors.list.count)
	for i := 0; i < nbrCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmIindicationMsgFormat1C.neighbors.list.array)) * uintptr(i)
		neighborC := *(**C.NRT_RCPRE_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmIindicationMsgFormat1C.neighbors.list.array)) + offset))
		neighbor, err := decodeNRT(neighborC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
		}
		e2SmIindicationMsgFormat1.IndicationMessageFormat1.Neighbors = append(e2SmIindicationMsgFormat1.IndicationMessageFormat1.Neighbors, neighbor)
	}

	return e2SmIindicationMsgFormat1, nil
}

func decodeE2SmRcPreIndicationMessageFormat1Bytes(array [8]byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1, error) {
	e2SmRcPreIndicationMsgFormat1C := (*C.E2SM_RC_PRE_IndicationMessage_Format1_RCPRE_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmRcPreIndicationMessageFormat1(e2SmRcPreIndicationMsgFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessageFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
