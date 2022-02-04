// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationMessage-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func PerEncodeE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage) ([]byte, error) {

	E2SmRcPreIndicationMsgCP, err := newE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage_RCPRE, unsafe.Pointer(E2SmRcPreIndicationMsgCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XerEncodeE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage) ([]byte, error) {

	E2SmRcPreIndicationMsgCP, err := newE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage_RCPRE, unsafe.Pointer(E2SmRcPreIndicationMsgCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreIndicationMessage(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_IndicationMessage_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationMessage((*C.E2SM_RC_PRE_IndicationMessage_RCPRE_t)(unsafePtr))
}

func XerDecodeE2SmRcPreIndicationMessage(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_IndicationMessage_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationMessage((*C.E2SM_RC_PRE_IndicationMessage_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreIndicationMessage(e2SmRcPreIndicationMsg *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage) (*C.E2SM_RC_PRE_IndicationMessage_RCPRE_t, error) {
	var present C.E2SM_RC_PRE_IndicationMessage_RCPRE_PR
	choiceC := [8]byte{}

	switch choice := e2SmRcPreIndicationMsg.E2SmRcPreIndicationMessage.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1:
		present = C.E2SM_RC_PRE_IndicationMessage_RCPRE_PR_indicationMessage_Format1

		im, err := newE2SmRcPreIndicationMessageFormat1(choice)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreIndicationMessage() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_RicStyleType:
		present = C.E2SM_RC_PRE_IndicationMessage_RCPRE_PR_ric_Style_Type

		im := newRicStyleType(choice.RicStyleType)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreIndicationMessage() %T not yet implemented", choice)
	}

	e2SmRcPreIndicationMsgC := C.E2SM_RC_PRE_IndicationMessage_RCPRE_t{
		present: present,
		choice:  choiceC,
	}

	return &e2SmRcPreIndicationMsgC, nil
}

func decodeE2SmRcPreIndicationMessage(e2SmRcPreIndicationMsgC *C.E2SM_RC_PRE_IndicationMessage_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage, error) {
	e2SmRcPreIndicationMsg := new(e2sm_rc_pre_v2.E2SmRcPreIndicationMessage)

	switch e2SmRcPreIndicationMsgC.present {
	case C.E2SM_RC_PRE_IndicationMessage_RCPRE_PR_indicationMessage_Format1:
		E2SmRcPreIndicationMsgFormat1, err := decodeE2SmRcPreIndicationMessageFormat1Bytes(e2SmRcPreIndicationMsgC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %s", err.Error())
		}

		e2SmRcPreIndicationMsg.E2SmRcPreIndicationMessage = E2SmRcPreIndicationMsgFormat1
	case C.E2SM_RC_PRE_IndicationMessage_RCPRE_PR_ric_Style_Type:
		ricStyleType := decodeRicStyleTypeBytes(e2SmRcPreIndicationMsgC.choice)

		e2SmRcPreIndicationMsg.E2SmRcPreIndicationMessage = &e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_RicStyleType{
			RicStyleType: ricStyleType,
		}

	default:
		return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %v not yet implemented", e2SmRcPreIndicationMsgC.present)
	}

	return e2SmRcPreIndicationMsg, nil
}
