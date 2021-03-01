// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRT.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodeNRT(nrt *e2sm_rc_pre_ies.Nrt) ([]byte, error) {

	nrtCP, err := newNRT(nrt)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRT() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRT, unsafe.Pointer(nrtCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRT() %s", err.Error())
	}
	return bytes, nil
}

func newNRT(nrt *e2sm_rc_pre_ies.Nrt) (*C.NRT_t, error) {

	cgi, _ := newCellGlobalID(nrt.Cgi)
	earfcn := newEARFCN(nrt.DlEarfcn)
	cellSize, _ := newCellSize(nrt.CellSize)
	pci := newPCI(nrt.Pci)
	nrIndex := C.long(nrt.NrIndex)

	nrtC := C.NRT_t{
		nrIndex:   nrIndex,
		cgi:       *cgi,
		dl_EARFCN: *earfcn,
		cell_Size: cellSize,
		pci:       *pci,
	}

	return &nrtC, nil
}

func decodeNRT(nrtC *C.NRT_t) (*e2sm_rc_pre_ies.Nrt, error) {
	nrt := new(e2sm_rc_pre_ies.Nrt)
	nrt.NrIndex = int32(nrtC.nrIndex)

	cgi, _ := decodeCellGlobalID(&nrtC.cgi)
	nrt.Cgi = cgi

	earfcn := decodeEARFCN(&nrtC.dl_EARFCN)
	nrt.DlEarfcn = earfcn

	cellSize := decodeCellSize(&nrtC.cell_Size)
	nrt.CellSize = cellSize

	pci := decodePCI(&nrtC.pci)
	nrt.Pci = pci

	return nrt, nil
}
