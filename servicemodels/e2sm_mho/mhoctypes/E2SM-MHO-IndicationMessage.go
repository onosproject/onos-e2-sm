// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-IndicationMessage.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func XerEncodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessage *e2sm_mho.E2SmMhoIndicationMessage) ([]byte, error) {
	e2SmMhoIndicationMessageCP, err := newE2SmMhoIndicationMessage(e2SmMhoIndicationMessage)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_IndicationMessage, unsafe.Pointer(e2SmMhoIndicationMessageCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessage *e2sm_mho.E2SmMhoIndicationMessage) ([]byte, error) {
	e2SmMhoIndicationMessageCP, err := newE2SmMhoIndicationMessage(e2SmMhoIndicationMessage)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessage() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_IndicationMessage, unsafe.Pointer(e2SmMhoIndicationMessageCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmMhoIndicationMessage(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_IndicationMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoIndicationMessage((*C.E2SM_MHO_IndicationMessage_t)(unsafePtr))
}

func PerDecodeE2SmMhoIndicationMessage(bytes []byte) (*e2sm_mho.E2SmMhoIndicationMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_IndicationMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoIndicationMessage((*C.E2SM_MHO_IndicationMessage_t)(unsafePtr))
}

func newE2SmMhoIndicationMessage(e2SmMhoIndicationMessage *e2sm_mho.E2SmMhoIndicationMessage) (*C.E2SM_MHO_IndicationMessage_t, error) {

	var pr C.E2SM_MHO_IndicationMessage_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}                   //ToDo - Check if number of bytes is sufficient

	switch choice := e2SmMhoIndicationMessage.E2SmMhoIndicationMessage.(type) {
	case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1:
		pr = C.E2SM_MHO_IndicationMessage_PR_indicationMessage_Format1 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoIndicationMessageFormat1(choice.IndicationMessageFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoIndicationMessageFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2:
		pr = C.E2SM_MHO_IndicationMessage_PR_indicationMessage_Format2 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoIndicationMessageFormat2(choice.IndicationMessageFormat2)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoIndicationMessageFormat2() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoIndicationMessage() %T not yet implemented", choice)
	}

	e2SmMhoIndicationMessageC := C.E2SM_MHO_IndicationMessage_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmMhoIndicationMessageC, nil
}

func decodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessageC *C.E2SM_MHO_IndicationMessage_t) (*e2sm_mho.E2SmMhoIndicationMessage, error) {

	e2SmMhoIndicationMessage := new(e2sm_mho.E2SmMhoIndicationMessage)

	switch e2SmMhoIndicationMessageC.present {
	case C.E2SM_MHO_IndicationMessage_PR_indicationMessage_Format1:
		e2SmMhoIndicationMessagestructC, err := decodeE2SmMhoIndicationMessageFormat1Bytes(e2SmMhoIndicationMessageC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoIndicationMessageFormat1Bytes() %s", err.Error())
		}
		e2SmMhoIndicationMessage.E2SmMhoIndicationMessage = &e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: e2SmMhoIndicationMessagestructC,
		}
	case C.E2SM_MHO_IndicationMessage_PR_indicationMessage_Format2:
		e2SmMhoIndicationMessagestructC, err := decodeE2SmMhoIndicationMessageFormat2Bytes(e2SmMhoIndicationMessageC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoIndicationMessageFormat2Bytes() %s", err.Error())
		}
		e2SmMhoIndicationMessage.E2SmMhoIndicationMessage = &e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: e2SmMhoIndicationMessagestructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoIndicationMessage() %v not yet implemented", e2SmMhoIndicationMessageC.present)
	}

	return e2SmMhoIndicationMessage, nil
}

//func decodeE2SmMhoIndicationMessageBytes(array [8]byte) (*e2sm_mho.E2SmMhoIndicationMessage, error) {
//	e2SmMhoIndicationMessageC := (*C.E2SM_MHO_IndicationMessage_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
//
//	return decodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessageC)
//}
