// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-MHO-ControlHeader-Format1.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1 *e2sm_mho.E2SmMhoControlHeaderFormat1) ([]byte, error) {
	e2SmMhoControlHeaderFormat1CP, err := newE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_MHO_ControlHeader_Format1, unsafe.Pointer(e2SmMhoControlHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmMhoControlHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1 *e2sm_mho.E2SmMhoControlHeaderFormat1) ([]byte, error) {
	e2SmMhoControlHeaderFormat1CP, err := newE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_MHO_ControlHeader_Format1, unsafe.Pointer(e2SmMhoControlHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmMhoControlHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmMhoControlHeaderFormat1(bytes []byte) (*e2sm_mho.E2SmMhoControlHeaderFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_MHO_ControlHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmMhoControlHeaderFormat1((*C.E2SM_MHO_ControlHeader_Format1_t)(unsafePtr))
}

func perDecodeE2SmMhoControlHeaderFormat1(bytes []byte) (*e2sm_mho.E2SmMhoControlHeaderFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_MHO_ControlHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmMhoControlHeaderFormat1((*C.E2SM_MHO_ControlHeader_Format1_t)(unsafePtr))
}

func newE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1 *e2sm_mho.E2SmMhoControlHeaderFormat1) (*C.E2SM_MHO_ControlHeader_Format1_t, error) {

	var err error
	e2SmMhoControlHeaderFormat1C := C.E2SM_MHO_ControlHeader_Format1_t{}

	rcCommandC, err := newMhoCommand(e2SmMhoControlHeaderFormat1.RcCommand)
	if err != nil {
		return nil, fmt.Errorf("newMhoCommand() %s", err.Error())
	}

	ricControlMessagePriorityC, err := newRicControlMessagePriority(e2SmMhoControlHeaderFormat1.RicControlMessagePriority)
	if err != nil {
		return nil, fmt.Errorf("newRicControlMessagePriority() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2SmMhoControlHeaderFormat1C.rc_command = rcCommandC
	e2SmMhoControlHeaderFormat1C.ric_Control_Message_Priority = ricControlMessagePriorityC

	return &e2SmMhoControlHeaderFormat1C, nil
}

func decodeE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1C *C.E2SM_MHO_ControlHeader_Format1_t) (*e2sm_mho.E2SmMhoControlHeaderFormat1, error) {

	var err error
	e2SmMhoControlHeaderFormat1 := e2sm_mho.E2SmMhoControlHeaderFormat1{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//RcCommand: rcCommand,
		//RicControlMessagePriority: ricControlMessagePriority,

	}

	e2SmMhoControlHeaderFormat1.RcCommand, err = decodeMhoCommand(e2SmMhoControlHeaderFormat1C.rc_command)
	if err != nil {
		return nil, fmt.Errorf("decodeMhoCommand() %s", err.Error())
	}

	e2SmMhoControlHeaderFormat1.RicControlMessagePriority, err = decodeRicControlMessagePriority(e2SmMhoControlHeaderFormat1C.ric_Control_Message_Priority)
	if err != nil {
		return nil, fmt.Errorf("decodeRicControlMessagePriority() %s", err.Error())
	}

	return &e2SmMhoControlHeaderFormat1, nil
}

func decodeE2SmMhoControlHeaderFormat1Bytes(array [8]byte) (*e2sm_mho.E2SmMhoControlHeaderFormat1, error) {
	e2SmMhoControlHeaderFormat1C := (*C.E2SM_MHO_ControlHeader_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmMhoControlHeaderFormat1(e2SmMhoControlHeaderFormat1C)
}
