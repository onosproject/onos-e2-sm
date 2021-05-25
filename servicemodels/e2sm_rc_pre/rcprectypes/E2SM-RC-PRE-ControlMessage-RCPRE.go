// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlMessage-RCPRE.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreControlMessage(e2SmRcPreControlMessage *e2sm_rc_pre_v2.E2SmRcPreControlMessage) ([]byte, error) {
	e2SmRcPreControlMessageCP, err := newE2SmRcPreControlMessage(e2SmRcPreControlMessage)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlMessage_RCPRE, unsafe.Pointer(e2SmRcPreControlMessageCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlMessage(e2SmRcPreControlMessage *e2sm_rc_pre_v2.E2SmRcPreControlMessage) ([]byte, error) {
	e2SmRcPreControlMessageCP, err := newE2SmRcPreControlMessage(e2SmRcPreControlMessage)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessage() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlMessage_RCPRE, unsafe.Pointer(e2SmRcPreControlMessageCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlMessage() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlMessage(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlMessage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlMessage_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlMessage((*C.E2SM_RC_PRE_ControlMessage_RCPRE_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlMessage(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlMessage_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlMessage((*C.E2SM_RC_PRE_ControlMessage_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreControlMessage(e2SmRcPreControlMessage *e2sm_rc_pre_v2.E2SmRcPreControlMessage) (*C.E2SM_RC_PRE_ControlMessage_RCPRE_t, error) {

	var pr C.E2SM_RC_PRE_ControlMessage_RCPRE_PR
	choiceC := [8]byte{}

	switch choice := e2SmRcPreControlMessage.E2SmRcPreControlMessage.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreControlMessage_ControlMessage:
		pr = C.E2SM_RC_PRE_ControlMessage_RCPRE_PR_controlMessage

		im, err := newE2SmRcPreControlMessageFormat1(choice.ControlMessage)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreControlMessage() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreControlMessage() %T not yet implemented", choice)
	}

	e2SmRcPreControlMessageC := C.E2SM_RC_PRE_ControlMessage_RCPRE_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmRcPreControlMessageC, nil
}

func decodeE2SmRcPreControlMessage(e2SmRcPreControlMessageC *C.E2SM_RC_PRE_ControlMessage_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreControlMessage, error) {

	//This is Decoder part (OneOf)
	e2SmRcPreControlMessage := new(e2sm_rc_pre_v2.E2SmRcPreControlMessage)

	switch e2SmRcPreControlMessageC.present {
	case C.E2SM_RC_PRE_ControlMessage_RCPRE_PR_controlMessage:
		e2SmRcPreControlMessagestructC, err := decodeE2SmRcPreControlMessageFormat1Bytes(e2SmRcPreControlMessageC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreControlMessage() %s", err.Error())
		}
		e2SmRcPreControlMessage.E2SmRcPreControlMessage = &e2sm_rc_pre_v2.E2SmRcPreControlMessage_ControlMessage{
			ControlMessage: e2SmRcPreControlMessagestructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmRcPreControlMessage() %v not yet implemented", e2SmRcPreControlMessageC.present)
	}

	return e2SmRcPreControlMessage, nil
}
