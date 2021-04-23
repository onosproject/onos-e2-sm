// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-ControlMessage.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoControlMessage(e2SmMhoControlMessage *e2sm_mho.E2SmMhoControlMessage) ([]byte, error) {
	e2SmMhoControlMessageCP, err := newE2SmMhoControlMessage(e2SmMhoControlMessage)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_ControlMessage, unsafe.Pointer(e2SmMhoControlMessageCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlMessage() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoControlMessage(e2SmMhoControlMessage *e2sm_mho.E2SmMhoControlMessage) ([]byte, error) {
	e2SmMhoControlMessageCP, err := newE2SmMhoControlMessage(e2SmMhoControlMessage)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlMessage() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_ControlMessage, unsafe.Pointer(e2SmMhoControlMessageCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlMessage() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoControlMessage(bytes []byte) (*e2sm_mho.E2SmMhoControlMessage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_ControlMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoControlMessage((*C.E2SM_MHO_ControlMessage_t)(unsafePtr))
}

func perDecodeE2SmMhoControlMessage(bytes []byte) (*e2sm_mho.E2SmMhoControlMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_ControlMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoControlMessage((*C.E2SM_MHO_ControlMessage_t)(unsafePtr))
}

func newE2SmMhoControlMessage(e2SmMhoControlMessage *e2sm_mho.E2SmMhoControlMessage) (*C.E2SM_MHO_ControlMessage_t, error) {

	var pr C.E2SM_MHO_ControlMessage_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}                //ToDo - Check if number of bytes is sufficient

	switch choice := e2SmMhoControlMessage.E2SmMhoControlMessage.(type) {
	case *e2sm_mho.E2SmMhoControlMessage_ControlMessageFormat1:
		pr = C.E2SM_MHO_ControlMessage_PR_controlMessage_Format1 //ToDo - Check if it's correct PR's name

		im, err := newE2SmMhoControlMessageFormat1(choice.ControlMessageFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmMhoControlMessageFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmMhoControlMessage() %T not yet implemented", choice)
	}

	e2SmMhoControlMessageC := C.E2SM_MHO_ControlMessage_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmMhoControlMessageC, nil
}

func decodeE2SmMhoControlMessage(e2SmMhoControlMessageC *C.E2SM_MHO_ControlMessage_t) (*e2sm_mho.E2SmMhoControlMessage, error) {

	e2SmMhoControlMessage := new(e2sm_mho.E2SmMhoControlMessage)

	switch e2SmMhoControlMessageC.present {
	case C.E2SM_MHO_ControlMessage_PR_controlMessage_Format1:
		e2SmMhoControlMessagestructC, err := decodeE2SmMhoControlMessageFormat1Bytes(e2SmMhoControlMessageC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmMhoControlMessageFormat1Bytes() %s", err.Error())
		}
		e2SmMhoControlMessage.ControlMessageFormat1 = &e2sm_mho.E2SmMhoControlMessage_ControlMessageFormat1{
			ControlMessageFormat1: e2SmMhoControlMessagestructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmMhoControlMessage() %v not yet implemented", e2SmMhoControlMessageC.present)
	}

	return &e2SmMhoControlMessage, nil
}

func decodeE2SmMhoControlMessageBytes(array [8]byte) (*e2sm_mho.E2SmMhoControlMessage, error) {
	e2SmMhoControlMessageC := (*C.E2SM_MHO_ControlMessage_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoControlMessage(e2SmMhoControlMessageC)
}
