// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Name.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRicStyleName(ricStyleName *e2sm_mho.RicStyleName) ([]byte, error) {
	ricStyleNameCP, err := newRicStyleName(ricStyleName)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleName() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(ricStyleNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleName(ricStyleName *e2sm_mho.RicStyleName) ([]byte, error) {
	ricStyleNameCP, err := newRicStyleName(ricStyleName)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleName() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(ricStyleNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleName(bytes []byte) (*e2sm_mho.RicStyleName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr))
}

func perDecodeRicStyleName(bytes []byte) (*e2sm_mho.RicStyleName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr))
}

func newRicStyleName(ricStyleName *e2sm_mho.RicStyleName) (*C.RIC_Style_Name_t, error) {

	ricStyleNameC, err := newPrintableString(ricStyleName.Value)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	return &ricStyleNameC, nil
}

func decodeRicStyleName(ricStyleNameC *C.RIC_Style_Name_t) (*e2sm_mho.RicStyleName, error) {

	ricStyleName := new(e2sm_mho.RicStyleName)
	ricStyleNameValue, err := decodePrintableString(ricStyleNameC)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}
	ricStyleName.Value = ricStyleNameValue

	return &ricStyleName, nil
}

func decodeRicStyleNameBytes(array [8]byte) (*e2sm_mho.RicStyleName, error) {
	ricStyleNameC := (*C.RIC_Style_Name_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicStyleName(ricStyleNameC)
}
