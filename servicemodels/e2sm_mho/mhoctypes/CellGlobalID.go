// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CellGlobalID.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeCellGlobalID(cellGlobalID *e2sm_mho.CellGlobalId) ([]byte, error) {
	cellGlobalIDCP, err := newCellGlobalID(cellGlobalID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellGlobalID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_CellGlobalID, unsafe.Pointer(cellGlobalIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCellGlobalID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCellGlobalID(cellGlobalID *e2sm_mho.CellGlobalId) ([]byte, error) {
	cellGlobalIDCP, err := newCellGlobalID(cellGlobalID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellGlobalID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CellGlobalID, unsafe.Pointer(cellGlobalIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCellGlobalID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCellGlobalID(bytes []byte) (*e2sm_mho.CellGlobalId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CellGlobalID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCellGlobalID((*C.CellGlobalID_t)(unsafePtr))
}

func perDecodeCellGlobalID(bytes []byte) (*e2sm_mho.CellGlobalId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CellGlobalID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCellGlobalID((*C.CellGlobalID_t)(unsafePtr))
}

func newCellGlobalID(cellGlobalID *e2sm_mho.CellGlobalId) (*C.CellGlobalID_t, error) {

	var pr C.CellGlobalID_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}     //ToDo - Check if number of bytes is sufficient

	switch choice := cellGlobalID.CellGlobalId.(type) {
	case *e2sm_mho.CellGlobalId_NrCgi:
		pr = C.CellGlobalID_PR_nr_CGI //ToDo - Check if it's correct PR's name

		im, err := newNrcgi(choice.NrCgi)
		if err != nil {
			return nil, fmt.Errorf("newNrcgi() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_mho.CellGlobalId_EUtraCgi:
		pr = C.CellGlobalID_PR_eUTRA_CGI //ToDo - Check if it's correct PR's name

		im, err := newEutracgi(choice.EUtraCgi)
		if err != nil {
			return nil, fmt.Errorf("newEutracgi() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newCellGlobalID() %T not yet implemented", choice)
	}

	cellGlobalIDC := C.CellGlobalID_t{
		present: pr,
		choice:  choiceC,
	}

	return &cellGlobalIDC, nil
}

func decodeCellGlobalID(cellGlobalIDC *C.CellGlobalID_t) (*e2sm_mho.CellGlobalId, error) {

	cellGlobalID := new(e2sm_mho.CellGlobalId)

	switch cellGlobalIDC.present {
	case C.CellGlobalID_PR_nr_CGI:
		cellGlobalIDstructC, err := decodeNrcgiBytes(cellGlobalIDC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeNrcgiBytes() %s", err.Error())
		}
		cellGlobalID.NrCgi = &e2sm_mho.CellGlobalId_NrCgi{
			NrCgi: cellGlobalIDstructC,
		}
	case C.CellGlobalID_PR_eUTRA_CGI:
		cellGlobalIDstructC, err := decodeEutracgiBytes(cellGlobalIDC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeEutracgiBytes() %s", err.Error())
		}
		cellGlobalID.EUtraCgi = &e2sm_mho.CellGlobalId_EUtraCgi{
			EUtraCgi: cellGlobalIDstructC,
		}
	default:
		return nil, fmt.Errorf("decodeCellGlobalID() %v not yet implemented", cellGlobalIDC.present)
	}

	return &cellGlobalID, nil
}

func decodeCellGlobalIDBytes(array [8]byte) (*e2sm_mho.CellGlobalId, error) {
	cellGlobalIDC := (*C.CellGlobalID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeCellGlobalID(cellGlobalIDC)
}
