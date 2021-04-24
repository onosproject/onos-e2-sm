// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Control-Message-Priority.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRicControlMessagePriority(ricControlMessagePriority *e2sm_mho.RicControlMessagePriority) ([]byte, error) {
	ricControlMessagePriorityCP, err := newRicControlMessagePriority(ricControlMessagePriority)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicControlMessagePriority() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_Control_Message_Priority, unsafe.Pointer(ricControlMessagePriorityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicControlMessagePriority() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicControlMessagePriority(ricControlMessagePriority *e2sm_mho.RicControlMessagePriority) ([]byte, error) {
	ricControlMessagePriorityCP, err := newRicControlMessagePriority(ricControlMessagePriority)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicControlMessagePriority() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Control_Message_Priority, unsafe.Pointer(ricControlMessagePriorityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicControlMessagePriority() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicControlMessagePriority(bytes []byte) (*e2sm_mho.RicControlMessagePriority, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Control_Message_Priority)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicControlMessagePriority((*C.RIC_Control_Message_Priority_t)(unsafePtr))
}

func perDecodeRicControlMessagePriority(bytes []byte) (*e2sm_mho.RicControlMessagePriority, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Control_Message_Priority)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicControlMessagePriority((*C.RIC_Control_Message_Priority_t)(unsafePtr))
}

func newRicControlMessagePriority(ricControlMessagePriority *e2sm_mho.RicControlMessagePriority) (*C.RIC_Control_Message_Priority_t, error) {

	ricControlMessagePriorityC := C.long(ricControlMessagePriority.Value)

	return &ricControlMessagePriorityC, nil
}

func decodeRicControlMessagePriority(ricControlMessagePriorityC *C.RIC_Control_Message_Priority_t) (*e2sm_mho.RicControlMessagePriority, error) {

	ricControlMessagePriority := e2sm_mho.RicControlMessagePriority{
		Value: int32(ricControlMessagePriorityC),
	}

	return &ricControlMessagePriority, nil
}

func decodeRicControlMessagePriorityBytes(array [8]byte) (*e2sm_mho.RicControlMessagePriority, error) {
	ricControlMessagePriorityC := (*C.RIC_Control_Message_Priority_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicControlMessagePriority(ricControlMessagePriorityC)
}
