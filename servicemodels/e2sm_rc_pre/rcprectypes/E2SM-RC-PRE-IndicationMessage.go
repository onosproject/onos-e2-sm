// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationMessage.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func PerEncodeE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage) ([]byte, error) {

	E2SmRcPreIndicationMsgCP, err := newE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage, unsafe.Pointer(E2SmRcPreIndicationMsgCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XerEncodeE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage) ([]byte, error) {

	E2SmRcPreIndicationMsgCP, err := newE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationMessage, unsafe.Pointer(E2SmRcPreIndicationMsgCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmRcPreIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreIndicationMessage(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreIndicationMessage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_IndicationMessage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationMessage((*C.E2SM_RC_PRE_IndicationMessage_t)(unsafePtr))
}

func newE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsg *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage) (*C.E2SM_RC_PRE_IndicationMessage_t, error) {
	var present C.E2SM_RC_PRE_IndicationMessage_PR
	choiceC := [8]byte{}

	switch choice := E2SmRcPreIndicationMsg.E2SmRcPreIndicationMessage.(type) {
	case *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1:
		present = C.E2SM_RC_PRE_IndicationMessage_PR_indicationMessage_Format1

		im, err := newE2SmRcPreIndicationMessageFormat1(choice)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreIndicationMessage() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	//case *e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_RicStyleType:
	//TODO: Implement RicStyleType
	default:
		return nil, fmt.Errorf("newE2SmRcPreIndicationMessage() %T not yet implemented", choice)
	}

	E2SmRcPreIndicationMsgC := C.E2SM_RC_PRE_IndicationMessage_t{
		present: present,
		//TODO: Implement RicStyleType
		//ric_Style_Type: nil,
		choice: choiceC,
	}

	return &E2SmRcPreIndicationMsgC, nil
}

func decodeE2SmRcPreIndicationMessage(E2SmRcPreIndicationMsgC *C.E2SM_RC_PRE_IndicationMessage_t) (*e2sm_rc_pre_ies.E2SmRcPreIndicationMessage, error) {
	E2SmRcPreIndicationMsg := new(e2sm_rc_pre_ies.E2SmRcPreIndicationMessage)

	switch E2SmRcPreIndicationMsgC.present {
	case C.E2SM_RC_PRE_IndicationMessage_PR_indicationMessage_Format1:
		E2SmRcPreIndicationMsgFormat1, err := decodeE2SmRcPreIndicationMessageFormat1Bytes(E2SmRcPreIndicationMsgC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %s", err.Error())
		}

		E2SmRcPreIndicationMsg.E2SmRcPreIndicationMessage = E2SmRcPreIndicationMsgFormat1

	default:
		return nil, fmt.Errorf("decodeE2SmRcPreIndicationMessage() %v not yet implemented", E2SmRcPreIndicationMsgC.present)
	}

	return E2SmRcPreIndicationMsg, nil
}
