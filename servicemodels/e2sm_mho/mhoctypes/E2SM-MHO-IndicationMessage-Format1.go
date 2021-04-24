// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationMessage-Format1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1 *e2sm_mho.E2SmMhoIndicationMessageFormat1) ([]byte, error) {
	e2SmMhoIndicationMessageFormat1CP, err := newE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationMessage_Format1, unsafe.Pointer(e2SmMhoIndicationMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1 *e2sm_mho.E2SmMhoIndicationMessageFormat1) ([]byte, error) {
	e2SmMhoIndicationMessageFormat1CP, err := newE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationMessage_Format1, unsafe.Pointer(e2SmMhoIndicationMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoIndicationMessageFormat1(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoIndicationMessageFormat1((*C.E2SM_MHO_IndicationMessage_Format1_t)(unsafePtr))
}

func perDecodeE2SmMhoIndicationMessageFormat1(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationMessage_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationMessageFormat1((*C.E2SM_MHO_IndicationMessage_Format1_t)(unsafePtr))
}

func newE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1 *e2sm_mho.E2SmMhoIndicationMessageFormat1) (*C.E2SM_MHO_IndicationMessage_Format1_t, error) {

	var err error
	e2SmMhoIndicationMessageFormat1C := C.E2SM_MHO_IndicationMessage_Format1_t{}

	ueIDC, err := newUeIdentity(e2SmMhoIndicationMessageFormat1.UeId)
	if err != nil {
		return nil, fmt.Errorf("newUeIdentity() %s", err.Error())
	}

	rsrpC, err := newRsrp(e2SmMhoIndicationMessageFormat1.Rsrp)
	if err != nil {
		return nil, fmt.Errorf("newRsrp() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoIndicationMessageFormat1C.ueID = ueIdC
	e2SmMhoIndicationMessageFormat1C.rsrp = rsrpC

	return &e2SmMhoIndicationMessageFormat1C, nil
}

func decodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1C *C.E2SM_MHO_IndicationMessage_Format1_t) (*e2sm_mho.E2SmMhoIndicationMessageFormat1, error) {

	var err error
	e2SmMhoIndicationMessageFormat1 := e2sm_mho.E2SmMhoIndicationMessageFormat1{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//UeId: ueId,
		//Rsrp: rsrp,

	}

	e2SmMhoIndicationMessageFormat1.UeId, err = decodeUeIdentity(e2SmMhoIndicationMessageFormat1C.ueID)
	if err != nil {
		return nil, fmt.Errorf("decodeUeIdentity() %s", err.Error())
	}

	e2SmMhoIndicationMessageFormat1.Rsrp, err = decodeRsrp(e2SmMhoIndicationMessageFormat1C.rsrp)
	if err != nil {
		return nil, fmt.Errorf("decodeRsrp() %s", err.Error())
	}

	return &e2SmMhoIndicationMessageFormat1, nil
}

func decodeE2SmMhoIndicationMessageFormat1Bytes(array [8]byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat1, error) {
	e2SmMhoIndicationMessageFormat1C := (*C.E2SM_MHO_IndicationMessage_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1C)
}
