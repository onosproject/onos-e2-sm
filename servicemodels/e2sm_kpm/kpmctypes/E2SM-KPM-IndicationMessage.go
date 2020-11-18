// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationMessage.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func PerEncodeE2SmKpmIndicationMessage(e2SmKpmIndicationMsg *e2sm_kpm_ies.E2SmKpmIndicationMessage) ([]byte, error) {

	e2SmKpmIndicationMsgCP, err := newE2SmKpmIndicationMessage(e2SmKpmIndicationMsg)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_IndicationMessage, unsafe.Pointer(e2SmKpmIndicationMsgCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func xerEncodeE2SmKpmIndicationMessage(e2SmKpmIndicationMsg *e2sm_kpm_ies.E2SmKpmIndicationMessage) ([]byte, error) {

	e2SmKpmIndicationMsgCP, err := newE2SmKpmIndicationMessage(e2SmKpmIndicationMsg)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationMessage, unsafe.Pointer(e2SmKpmIndicationMsgCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessage() %s", err.Error())
	}
	return bytes, nil
}

func newE2SmKpmIndicationMessage(e2SmKpmIndicationMsg *e2sm_kpm_ies.E2SmKpmIndicationMessage) (*C.E2SM_KPM_IndicationMessage_t, error) {
	var present C.E2SM_KPM_IndicationMessage_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmIndicationMsg.E2SmKpmIndicationMessage.(type) {
	case *e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1:
		present = C.E2SM_KPM_IndicationMessage_PR_indicationMessage_Format1

		im, err := newE2SmKpmIndicationMessageFormat1(choice)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmIndicationMessage() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	//case *e2sm_kpm_ies.E2SmKpmIndicationMessage_RicStyleType:
		//TODO: Implement RicStyleType
	default:
		return nil, fmt.Errorf("newE2SmKpmIndicationMessage() %T not yet implemented", choice)
	}

	e2SmKpmIndicationMsgC := C.E2SM_KPM_IndicationMessage_t{
		present: present,
		//TODO: Implement RicStyleType
		//ric_Style_Type: nil,
		choice: choiceC,
	}

	return &e2SmKpmIndicationMsgC, nil
}

func decodeE2SmKpmIndicationMessage(e2SmKpmIndicationMsgC *C.E2SM_KPM_IndicationMessage_t) (*e2sm_kpm_ies.E2SmKpmIndicationMessage, error) {
	e2SmKpmIndicationMsg := new(e2sm_kpm_ies.E2SmKpmIndicationMessage)

	switch e2SmKpmIndicationMsgC.present {
	case C.E2SM_KPM_IndicationMessage_PR_indicationMessage_Format1:
		e2SmKpmIndicationMsgFormat1, err := decodeE2SmKpmIndicationMessageFormat1Bytes(e2SmKpmIndicationMsgC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationMessage() %s", err.Error())
		}

		e2SmKpmIndicationMsg.E2SmKpmIndicationMessage = e2SmKpmIndicationMsgFormat1

	default:
		return nil, fmt.Errorf("decodeE2SmKpmIndicationMessage() %v not yet implemented", e2SmKpmIndicationMsgC.present)
	}

	return e2SmKpmIndicationMsg, nil
}
