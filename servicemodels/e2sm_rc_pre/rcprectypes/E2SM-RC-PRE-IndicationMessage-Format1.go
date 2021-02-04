// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationMessage-Format1.h"
//#include "NRT.h"
//#include "PCI-Range.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodeE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1 *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1) ([]byte, error) {

	E2SmRcPreIndicationMsgFormat1CP, err := newE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage_Format1, unsafe.Pointer(E2SmRcPreIndicationMsgFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func newE2SmRcPreIndicationMessageFormat1(E2SmRcPreIndicationMsgFormat1 *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1) (*C.E2SM_RC_PRE_IndicationMessage_Format1_t, error) {
	neighborsListC := new(C.struct_E2SM_RC_PRE_IndicationMessage_Format1__neighbors)
	pciPoolListC := new(C.struct_E2SM_RC_PRE_IndicationMessage_Format1__pci_Pool)
	cgiC, _ := newCellGlobalID(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetCgi())
	earfcnC := newEARFCN(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetDlEarfcn())
	cellSizeC, _ := newCellSize(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetCellSize())
	pciC := newPCI(E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetPci())

	format1 := C.E2SM_RC_PRE_IndicationMessage_Format1_t{
		cgi:       *cgiC,
		dl_EARFCN: *earfcnC,
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

	for _, pciPool := range E2SmRcPreIndicationMsgFormat1.IndicationMessageFormat1.GetPciPool() {
		pciPoolC, err := newPciRange(pciPool)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreIndicationMessageFormat1() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(pciPoolListC), unsafe.Pointer(pciPoolC)); err != nil {
			return nil, err
		}
	}
	format1.pci_Pool = *pciPoolListC

	return &format1, nil
}

func decodeE2SmRcPreIndicationMessageFormat1(e2SmIindicationMsgFormat1C *C.E2SM_RC_PRE_IndicationMessage_Format1_t) (*e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1, error) {

	cgi, _ := decodeCellGlobalID(&e2SmIindicationMsgFormat1C.cgi)
	earfcn := decodeEARFCN(&e2SmIindicationMsgFormat1C.dl_EARFCN)
	cellSize := decodeCellSize(&e2SmIindicationMsgFormat1C.cell_Size)
	pci := decodePCI(&e2SmIindicationMsgFormat1C.pci)

	e2SmIindicationMsgFormat1 := &e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_ies.E2SmRcPreIndicationMessageFormat1{
			Cgi:       cgi,
			DlEarfcn:  earfcn,
			CellSize:  cellSize,
			PciPool:   make([]*e2sm_rc_pre_ies.PciRange, 0),
			Pci:       pci,
			Neighbors: make([]*e2sm_rc_pre_ies.Nrt, 0),
		},
	}

	nbrCount := int(e2SmIindicationMsgFormat1C.neighbors.list.count)
	for i := 0; i < nbrCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmIindicationMsgFormat1C.neighbors.list.array)) * uintptr(i)
		neighborC := *(**C.NRT_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmIindicationMsgFormat1C.neighbors.list.array)) + offset))
		neighbor, err := decodeNRT(neighborC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
		}
		e2SmIindicationMsgFormat1.IndicationMessageFormat1.Neighbors = append(e2SmIindicationMsgFormat1.IndicationMessageFormat1.Neighbors, neighbor)
	}

	pciPoolCount := int(e2SmIindicationMsgFormat1C.pci_Pool.list.count)
	for i := 0; i < pciPoolCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmIindicationMsgFormat1C.pci_Pool.list.array)) * uintptr(i)
		pciPoolC := *(**C.PCI_Range_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmIindicationMsgFormat1C.pci_Pool.list.array)) + offset))
		pciPool, err := decodePciRange(pciPoolC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessageFormat1() %s", err.Error())
		}
		e2SmIindicationMsgFormat1.IndicationMessageFormat1.PciPool = append(e2SmIindicationMsgFormat1.IndicationMessageFormat1.PciPool, pciPool)
	}

	return e2SmIindicationMsgFormat1, nil
}

func decodeE2SmRcPreIndicationMessageFormat1Bytes(array [8]byte) (*e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1, error) {
	e2SmRcPreIndicationMsgFormat1C := (*C.E2SM_RC_PRE_IndicationMessage_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmRcPreIndicationMessageFormat1(e2SmRcPreIndicationMsgFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessageFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
