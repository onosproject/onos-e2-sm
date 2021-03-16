// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-ActionDefinition-Format2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2) ([]byte, error) {
	e2SmKpmActionDefinitionFormat2CP, err := newE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_ActionDefinition_Format2, unsafe.Pointer(e2SmKpmActionDefinitionFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinitionFormat2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2) ([]byte, error) {
	e2SmKpmActionDefinitionFormat2CP, err := newE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_ActionDefinition_Format2, unsafe.Pointer(e2SmKpmActionDefinitionFormat2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinitionFormat2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2SmKpmActionDefinitionFormat2(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_ActionDefinition_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat2((*C.E2SM_KPM_ActionDefinition_Format2_t)(unsafePtr))
}

func perDecodeE2SmKpmActionDefinitionFormat2(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_ActionDefinition_Format2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmActionDefinitionFormat2((*C.E2SM_KPM_ActionDefinition_Format2_t)(unsafePtr))
}

func newE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2 *e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2) (*C.E2SM_KPM_ActionDefinition_Format2_t, error) {

	ueIDC, err := newUeIdentity(e2SmKpmActionDefinitionFormat2.UeId)
	if err != nil {
		return nil, fmt.Errorf("newUeIdentity() %s", err.Error())
	}

	subscriptInfoC, err := newE2SmKpmActionDefinitionFormat1(e2SmKpmActionDefinitionFormat2.SubscriptInfo)
	if err != nil {
		return nil, fmt.Errorf("newE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat2C := C.E2SM_KPM_ActionDefinition_Format2_t{
		ueID:          *ueIDC,
		subscriptInfo: *subscriptInfoC,
	}

	return &e2SmKpmActionDefinitionFormat2C, nil
}

func decodeE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2C *C.E2SM_KPM_ActionDefinition_Format2_t) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2, error) {

	ueID, err := decodeUeIdentity(&e2SmKpmActionDefinitionFormat2C.ueID)
	if err != nil {
		return nil, fmt.Errorf("decodeUeIdentity() %s", err.Error())
	}

	subscriptInfo, err := decodeE2SmKpmActionDefinitionFormat1(&e2SmKpmActionDefinitionFormat2C.subscriptInfo)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmKpmActionDefinitionFormat1() %s", err.Error())
	}

	e2SmKpmActionDefinitionFormat2 := e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2{
		UeId:          ueID,
		SubscriptInfo: subscriptInfo,
	}

	return &e2SmKpmActionDefinitionFormat2, nil
}

func decodeE2SmKpmActionDefinitionFormat2Bytes(array [8]byte) (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat2, error) {
	e2SmKpmActionDefinitionFormat2C := (*C.E2SM_KPM_ActionDefinition_Format2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmActionDefinitionFormat2(e2SmKpmActionDefinitionFormat2C)
}
