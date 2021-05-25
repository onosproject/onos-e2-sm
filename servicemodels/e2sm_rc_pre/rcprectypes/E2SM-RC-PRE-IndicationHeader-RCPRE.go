// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationHeader-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func PerEncodeE2SmRcPreIndicationHeader(indicationHeader *e2sm_rc_pre_v2.E2SmRcPreIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmRcPreIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_IndicationHeader_RCPRE, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func XerEncodeE2SmRcPreIndicationHeader(indicationHeader *e2sm_rc_pre_v2.E2SmRcPreIndicationHeader) ([]byte, error) {
	indicationHeaderCP, err := newE2SmRcPreIndicationHeader(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationHeader_RCPRE, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreIndicationHeader(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeader, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_IndicationHeader_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationHeader((*C.E2SM_RC_PRE_IndicationHeader_RCPRE_t)(unsafePtr))
}

func XerDecodeE2SmRcPreIndicationHeader(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeader, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_IndicationHeader_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationHeader((*C.E2SM_RC_PRE_IndicationHeader_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreIndicationHeader(indicationHeader *e2sm_rc_pre_v2.E2SmRcPreIndicationHeader) (*C.E2SM_RC_PRE_IndicationHeader_RCPRE_t, error) {
	var pr C.E2SM_RC_PRE_IndicationHeader_RCPRE_PR
	choiceC := [8]byte{}

	switch choice := indicationHeader.E2SmRcPreIndicationHeader.(type) {
	case *e2sm_rc_pre_v2.E2SmRcPreIndicationHeader_IndicationHeaderFormat1:
		pr = C.E2SM_RC_PRE_IndicationHeader_RCPRE_PR_indicationHeader_Format1

		im, err := newE2SmRcPreIndicationHeaderFormat1(choice.IndicationHeaderFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreIndicationHeader() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmRcPreIndicationHeader() %T not yet implemented", choice)
	}

	indicationHeaderC := C.E2SM_RC_PRE_IndicationHeader_RCPRE_t{
		present: pr,
		choice:  choiceC,
	}

	return &indicationHeaderC, nil
}

func decodeE2SmRcPreIndicationHeader(indicationHeaderC *C.E2SM_RC_PRE_IndicationHeader_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeader, error) {
	indicationHeader := new(e2sm_rc_pre_v2.E2SmRcPreIndicationHeader)

	switch indicationHeaderC.present {
	case C.E2SM_RC_PRE_IndicationHeader_RCPRE_PR_indicationHeader_Format1:
		indicationHeaderFormat1, err := decodeE2SmRcPreIndicationHeaderFormat1Bytes(indicationHeaderC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationHeader() %s", err.Error())
		}

		indicationHeader.E2SmRcPreIndicationHeader = &e2sm_rc_pre_v2.E2SmRcPreIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: indicationHeaderFormat1,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmRcPreIndicationHeader() %v not yet implemented", indicationHeaderC.present)
	}

	return indicationHeader, nil
}
