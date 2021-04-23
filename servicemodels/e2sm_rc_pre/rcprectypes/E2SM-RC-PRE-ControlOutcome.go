// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlOutcome.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome *e2sm_rc_pre_v2.E2SmRcPreControlOutcome) ([]byte, error) {
	e2SmRcPreControlOutcomeCP, err := newE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcome() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlOutcome, unsafe.Pointer(e2SmRcPreControlOutcomeCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcome() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome *e2sm_rc_pre_v2.E2SmRcPreControlOutcome) ([]byte, error) {
	e2SmRcPreControlOutcomeCP, err := newE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcome() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlOutcome, unsafe.Pointer(e2SmRcPreControlOutcomeCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlOutcome() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlOutcome(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlOutcome, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlOutcome)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlOutcome((*C.E2SM_RC_PRE_ControlOutcome_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlOutcome(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlOutcome, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlOutcome)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlOutcome((*C.E2SM_RC_PRE_ControlOutcome_t)(unsafePtr))
}

func newE2SmRcPreControlOutcome(e2SmRcPreControlOutcome *e2sm_rc_pre_v2.E2SmRcPreControlOutcome) (*C.E2SM_RC_PRE_ControlOutcome_t, error) {

	var pr C.E2SM_RC_PRE_ControlOutcome_PR
	choiceC := [8]byte{}

	switch choice := e2SmRcPreControlOutcome.E2SmRcPreControlOutcome.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreControlOutcome_ControlOutcomeFormat1:
		pr = C.E2SM_RC_PRE_ControlOutcome_PR_controlOutcome_Format1

		im, err := newE2SmRcPreControlOutcomeFormat1(choice.ControlOutcomeFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreControlOutcome() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreControlOutcome() %T not yet implemented", choice)
	}

	e2SmRcPreControlOutcomeC := C.E2SM_RC_PRE_ControlOutcome_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmRcPreControlOutcomeC, nil
}

func decodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcomeC *C.E2SM_RC_PRE_ControlOutcome_t) (*e2sm_rc_pre_v2.E2SmRcPreControlOutcome, error) {

	//This is Decoder part (OneOf)
	e2SmRcPreControlOutcome := new(e2sm_rc_pre_v2.E2SmRcPreControlOutcome)

	switch e2SmRcPreControlOutcomeC.present {
	case C.E2SM_RC_PRE_ControlOutcome_PR_controlOutcome_Format1:
		e2SmRcPreControlOutcomestructC, err := decodeE2SmRcPreControlOutcomeFormat1Bytes(e2SmRcPreControlOutcomeC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreControlOutcome() %s", err.Error())
		}
		e2SmRcPreControlOutcome.E2SmRcPreControlOutcome = &e2sm_rc_pre_v2.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			ControlOutcomeFormat1: e2SmRcPreControlOutcomestructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmRcPreControlOutcome() %v not yet implemented", e2SmRcPreControlOutcomeC.present)
	}

	return e2SmRcPreControlOutcome, nil
}
