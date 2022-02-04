// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationMessage-Format2.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2 *e2sm_mho.E2SmMhoIndicationMessageFormat2) ([]byte, error) {
	e2SmMhoIndicationMessageFormat2CP, err := newE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessageFormat2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationMessage_Format2, unsafe.Pointer(e2SmMhoIndicationMessageFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessageFormat2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2 *e2sm_mho.E2SmMhoIndicationMessageFormat2) ([]byte, error) {
	e2SmMhoIndicationMessageFormat2CP, err := newE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessageFormat2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationMessage_Format2, unsafe.Pointer(e2SmMhoIndicationMessageFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessageFormat2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoIndicationMessageFormat2(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationMessage_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoIndicationMessageFormat2((*C.E2SM_MHO_IndicationMessage_Format2_t)(unsafePtr))
}

func perDecodeE2SmMhoIndicationMessageFormat2(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationMessage_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationMessageFormat2((*C.E2SM_MHO_IndicationMessage_Format2_t)(unsafePtr))
}

func newE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2 *e2sm_mho.E2SmMhoIndicationMessageFormat2) (*C.E2SM_MHO_IndicationMessage_Format2_t, error) {

	var err error
	e2SmMhoIndicationMessageFormat2C := C.E2SM_MHO_IndicationMessage_Format2_t{}

	ueIDC, err := newUeIdentity(e2SmMhoIndicationMessageFormat2.UeId)
	if err != nil {
		return nil, fmt.Errorf("newUeIdentity() %s", err.Error())
	}

	rrcStatusC, err := newRrcstatus(&e2SmMhoIndicationMessageFormat2.RrcStatus)
	if err != nil {
		return nil, fmt.Errorf("newRrcstatus() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoIndicationMessageFormat2C.ueID = *ueIDC
	e2SmMhoIndicationMessageFormat2C.rrcStatus = *rrcStatusC

	return &e2SmMhoIndicationMessageFormat2C, nil
}

func decodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2C *C.E2SM_MHO_IndicationMessage_Format2_t) (*e2sm_mho.E2SmMhoIndicationMessageFormat2, error) {

	var err error
	e2SmMhoIndicationMessageFormat2 := e2sm_mho.E2SmMhoIndicationMessageFormat2{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//UeId: ueId,
		//RrcStatus: rrcStatus,

	}

	e2SmMhoIndicationMessageFormat2.UeId, err = decodeUeIdentity(&e2SmMhoIndicationMessageFormat2C.ueID)
	if err != nil {
		return nil, fmt.Errorf("decodeUeIdentity() %s", err.Error())
	}

	var t *e2sm_mho.Rrcstatus
	t, err = decodeRrcstatus(&e2SmMhoIndicationMessageFormat2C.rrcStatus)
	if err != nil {
		return nil, fmt.Errorf("decodeRrcstatus() %s", err.Error())
	}
	e2SmMhoIndicationMessageFormat2.RrcStatus = *t

	return &e2SmMhoIndicationMessageFormat2, nil
}

func decodeE2SmMhoIndicationMessageFormat2Bytes(array [8]byte) (*e2sm_mho.E2SmMhoIndicationMessageFormat2, error) {
	e2SmMhoIndicationMessageFormat2C := (*C.E2SM_MHO_IndicationMessage_Format2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2C)
}
