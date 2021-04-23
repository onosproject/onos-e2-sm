// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlHeader-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1) ([]byte, error) {
	e2SmRcPreControlHeaderFormat1CP, err := newE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlHeader_Format1, unsafe.Pointer(e2SmRcPreControlHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1) ([]byte, error) {
	e2SmRcPreControlHeaderFormat1CP, err := newE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlHeader_Format1, unsafe.Pointer(e2SmRcPreControlHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlHeaderFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlHeaderFormat1((*C.E2SM_RC_PRE_ControlHeader_Format1_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlHeaderFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlHeader_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlHeaderFormat1((*C.E2SM_RC_PRE_ControlHeader_Format1_t)(unsafePtr))
}

func newE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1) (*C.E2SM_RC_PRE_ControlHeader_Format1_t, error) {
	rcCommandC, err := newRcPreCommand(&e2SmRcPreControlHeaderFormat1.RcCommand)
	if err != nil {
		return nil, fmt.Errorf("newE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}

	e2SmRcPreControlHeaderFormat1C := C.E2SM_RC_PRE_ControlHeader_Format1_t{
		//cgi:                          cgi,
		rc_command: *rcCommandC,
		//ric_Control_Message_Priority: ricControlMessagePriorityC,
	}

	if e2SmRcPreControlHeaderFormat1.GetCgi() != nil {
		cgi, err := newCellGlobalID(e2SmRcPreControlHeaderFormat1.Cgi)
		if err != nil {
			return nil, err
		}
		e2SmRcPreControlHeaderFormat1C.cgi = cgi
	}

	if e2SmRcPreControlHeaderFormat1.GetRicControlMessagePriority() != nil {
		ricControlMessagePriorityC, err := newRicControlMessagePriority(e2SmRcPreControlHeaderFormat1.RicControlMessagePriority)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreControlHeaderFormat1() %s", err.Error())
		}
		e2SmRcPreControlHeaderFormat1C.ric_Control_Message_Priority = ricControlMessagePriorityC
	}

	return &e2SmRcPreControlHeaderFormat1C, nil
}

func decodeE2SmRcPreControlHeaderFormat1(e2SmRcPreControlHeaderFormat1C *C.E2SM_RC_PRE_ControlHeader_Format1_t) (*e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1, error) {

	rcCommand, err := decodeRcPreCommand(&e2SmRcPreControlHeaderFormat1C.rc_command)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
	}

	e2SmRcPreControlHeaderFormat1 := e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1{
		RcCommand: *rcCommand,
	}

	if e2SmRcPreControlHeaderFormat1C.cgi != nil { // Is optional
		cgi, err := decodeCellGlobalID(e2SmRcPreControlHeaderFormat1C.cgi)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationHeaderFormat1() %s", err.Error())
		}
		e2SmRcPreControlHeaderFormat1.Cgi = cgi
	}

	if e2SmRcPreControlHeaderFormat1C.ric_Control_Message_Priority != nil {
		ricControlMessagePriority, err := decodeRicControlMessagePriority(e2SmRcPreControlHeaderFormat1C.ric_Control_Message_Priority)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreControlHeaderFormat1() %s", err.Error())
		}
		e2SmRcPreControlHeaderFormat1.RicControlMessagePriority = ricControlMessagePriority
	}

	return &e2SmRcPreControlHeaderFormat1, nil
}

func decodeE2SmRcPreControlHeaderFormat1Bytes(array [8]byte) (*e2sm_rc_pre_v2.E2SmRcPreControlHeaderFormat1, error) {
	controlHeaderFormat1C := (*C.E2SM_RC_PRE_ControlHeader_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmRcPreControlHeaderFormat1(controlHeaderFormat1C)
}
