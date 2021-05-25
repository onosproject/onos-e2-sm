// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlMessage-Format1-RCPRE.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1) ([]byte, error) {
	e2SmRcPreControlMessageFormat1CP, err := newE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlMessage_Format1_RCPRE, unsafe.Pointer(e2SmRcPreControlMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1) ([]byte, error) {
	e2SmRcPreControlMessageFormat1CP, err := newE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlMessage_Format1_RCPRE, unsafe.Pointer(e2SmRcPreControlMessageFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlMessageFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlMessage_Format1_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlMessageFormat1((*C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlMessageFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlMessage_Format1_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlMessageFormat1((*C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1 *e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1) (*C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t, error) {

	parameterTypeC, err := newRanparameterDefItem(e2SmRcPreControlMessageFormat1.ParameterType)
	if err != nil {
		return nil, fmt.Errorf("newE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	parameterValC, err := newRanparameterValue(e2SmRcPreControlMessageFormat1.ParameterVal)
	if err != nil {
		return nil, fmt.Errorf("newE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	e2SmRcPreControlMessageFormat1C := C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t{
		parameterType: *parameterTypeC,
		parameterVal:  *parameterValC,
	}

	return &e2SmRcPreControlMessageFormat1C, nil
}

func decodeE2SmRcPreControlMessageFormat1(e2SmRcPreControlMessageFormat1C *C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1, error) {

	parameterType, err := decodeRanparameterDefItem(&e2SmRcPreControlMessageFormat1C.parameterType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	parameterVal, err := decodeRanparameterValue(&e2SmRcPreControlMessageFormat1C.parameterVal)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreControlMessageFormat1() %s", err.Error())
	}

	e2SmRcPreControlMessageFormat1 := e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1{
		ParameterType: parameterType,
		ParameterVal:  parameterVal,
	}

	return &e2SmRcPreControlMessageFormat1, nil
}

func decodeE2SmRcPreControlMessageFormat1Bytes(array [8]byte) (*e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1, error) {
	controlMessageFormat1C := (*C.E2SM_RC_PRE_ControlMessage_Format1_RCPRE_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmRcPreControlMessageFormat1(controlMessageFormat1C)
}
