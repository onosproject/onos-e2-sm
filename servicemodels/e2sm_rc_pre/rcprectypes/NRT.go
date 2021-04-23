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
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeNRT(nrt *e2sm_rc_pre_v2.Nrt) ([]byte, error) {

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

func newNRT(nrt *e2sm_rc_pre_v2.Nrt) (*C.NRT_t, error) {

	cgi, err := newCellGlobalID(nrt.GetCgi())
	if err != nil {
		return nil, err
	}
	earfcn, err := newARFCN(nrt.GetDlArfcn())
	if err != nil {
		return nil, err
	}
	cellSize, err := newCellSize(nrt.GetCellSize())
	if err != nil {
		return nil, err
	}
	pci := newPCI(nrt.GetPci())

	nrtC := C.NRT_t{
		cgi:       *cgi,
		dl_ARFCN:  *earfcn,
		cell_Size: cellSize,
		pci:       *pci,
	}

	return &nrtC, nil
}

func decodeNRT(nrtC *C.NRT_t) (*e2sm_rc_pre_v2.Nrt, error) {
	nrt := new(e2sm_rc_pre_v2.Nrt)

	cgi, err := decodeCellGlobalID(&nrtC.cgi)
	if err != nil {
		return nil, err
	}
	nrt.Cgi = cgi

	arfcn, err := decodeARFCN(&nrtC.dl_ARFCN)
	if err != nil {
		return nil, err
	}
	nrt.DlArfcn = arfcn

	cellSize := decodeCellSize(&nrtC.cell_Size)
	nrt.CellSize = cellSize

	pci := decodePCI(&nrtC.pci)
	nrt.Pci = pci

	return nrt, nil
}
