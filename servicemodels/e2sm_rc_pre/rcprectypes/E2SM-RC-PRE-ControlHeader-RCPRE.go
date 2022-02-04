// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlHeader-RCPRE.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader *e2sm_rc_pre_v2.E2SmRcPreControlHeader) ([]byte, error) {
	e2SmRcPreControlHeaderCP, err := newE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlHeader_RCPRE, unsafe.Pointer(e2SmRcPreControlHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader *e2sm_rc_pre_v2.E2SmRcPreControlHeader) ([]byte, error) {
	e2SmRcPreControlHeaderCP, err := newE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlHeader_RCPRE, unsafe.Pointer(e2SmRcPreControlHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlHeader(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlHeader_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlHeader((*C.E2SM_RC_PRE_ControlHeader_RCPRE_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlHeader(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlHeader_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlHeader((*C.E2SM_RC_PRE_ControlHeader_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreControlHeader(e2SmRcPreControlHeader *e2sm_rc_pre_v2.E2SmRcPreControlHeader) (*C.E2SM_RC_PRE_ControlHeader_RCPRE_t, error) {

	var pr C.E2SM_RC_PRE_ControlHeader_RCPRE_PR
	choiceC := [8]byte{}

	switch choice := e2SmRcPreControlHeader.E2SmRcPreControlHeader.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreControlHeader_ControlHeaderFormat1:
		pr = C.E2SM_RC_PRE_ControlHeader_RCPRE_PR_controlHeader_Format1

		im, err := newE2SmRcPreControlHeaderFormat1(choice.ControlHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreControlHeader() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreControlHeader() %T not yet implemented", choice)
	}

	e2SmRcPreControlHeaderC := C.E2SM_RC_PRE_ControlHeader_RCPRE_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2SmRcPreControlHeaderC, nil
}

func decodeE2SmRcPreControlHeader(e2SmRcPreControlHeaderC *C.E2SM_RC_PRE_ControlHeader_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreControlHeader, error) {

	e2SmRcPreControlHeader := new(e2sm_rc_pre_v2.E2SmRcPreControlHeader)

	switch e2SmRcPreControlHeaderC.present {
	case C.E2SM_RC_PRE_ControlHeader_RCPRE_PR_controlHeader_Format1:
		e2SmRcPreControlHeaderstructC, err := decodeE2SmRcPreControlHeaderFormat1Bytes(e2SmRcPreControlHeaderC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreControlHeader() %s", err.Error())
		}
		e2SmRcPreControlHeader.E2SmRcPreControlHeader = &e2sm_rc_pre_v2.E2SmRcPreControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: e2SmRcPreControlHeaderstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmRcPreControlHeader() %v not yet implemented", e2SmRcPreControlHeaderC.present)
	}

	return e2SmRcPreControlHeader, nil
}
