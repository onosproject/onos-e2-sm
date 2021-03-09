// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-ActionDefinition.h"
//#include "E2SM-KPM-ActionDefinition-Format1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func XerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition *e2sm_kpm_v2.E2SmKpmActionDefinition) ([]byte, error) {
	e2SmKpmActionDefinitionCP, err := newE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinition() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_ActionDefinition, unsafe.Pointer(e2SmKpmActionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition *e2sm_kpm_v2.E2SmKpmActionDefinition) ([]byte, error) {
	e2SmKpmActionDefinitionCP, err := newE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinition() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_ActionDefinition, unsafe.Pointer(e2SmKpmActionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2SmKpmActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmActionDefinition(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmActionDefinition((*C.E2SM_KPM_ActionDefinition_t)(unsafePtr))
}

func PerDecodeE2SmKpmActionDefinition(bytes []byte) (*e2sm_kpm_v2.E2SmKpmActionDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmActionDefinition((*C.E2SM_KPM_ActionDefinition_t)(unsafePtr))
}

func newE2SmKpmActionDefinition(e2SmKpmActionDefinition *e2sm_kpm_v2.E2SmKpmActionDefinition) (*C.E2SM_KPM_ActionDefinition_t, error) {

	ricStyleTypeC, err := newRicStyleType(e2SmKpmActionDefinition.RicStyleType)
	if err != nil {
		return nil, fmt.Errorf("newRicStyleType() %s", err.Error())
	}

	var pr C.E2SM_KPM_ActionDefinition__actionDefinition_formats_PR
	choiceC := [8]byte{}

	switch choice := e2SmKpmActionDefinition.E2SmKpmActionDefinition.(type) {
	case *e2sm_kpm_v2.E2SmKpmActionDefinition_ActionDefinitionFormat1:
		pr = C.E2SM_KPM_ActionDefinition__actionDefinition_formats_PR_actionDefinition_Format1

		im, err := newE2SmKpmActionDefinitionFormat1(choice.ActionDefinitionFormat1)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmActionDefinitionFormat1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2SmKpmActionDefinition() %T not yet implemented", choice)
	}

	e2SmKpmActionDefinitionFormatC := C.struct_E2SM_KPM_ActionDefinition__actionDefinition_formats{
		present: pr,
		choice:  choiceC,
	}

	e2SmKpmActionDefinitionC := C.E2SM_KPM_ActionDefinition_t{
		ric_Style_Type:           *ricStyleTypeC,
		actionDefinition_formats: e2SmKpmActionDefinitionFormatC,
	}

	return &e2SmKpmActionDefinitionC, nil
}

func decodeE2SmKpmActionDefinition(e2SmKpmActionDefinitionC *C.E2SM_KPM_ActionDefinition_t) (*e2sm_kpm_v2.E2SmKpmActionDefinition, error) {

	e2SmKpmActionDefinition := new(e2sm_kpm_v2.E2SmKpmActionDefinition)

	ricStyleType, err := decodeRicStyleType(&e2SmKpmActionDefinitionC.ric_Style_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRicStyleType() %s", err.Error())
	}

	e2SmKpmActionDefinition.RicStyleType = ricStyleType

	switch e2SmKpmActionDefinitionC.actionDefinition_formats.present {
	case C.E2SM_KPM_ActionDefinition__actionDefinition_formats_PR_actionDefinition_Format1:
		actionDefinitionFormat1, err := decodeE2SmKpmActionDefinitionFormat1Bytes(e2SmKpmActionDefinitionC.actionDefinition_formats.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		e2SmKpmActionDefinition.E2SmKpmActionDefinition = &e2sm_kpm_v2.E2SmKpmActionDefinition_ActionDefinitionFormat1{
			ActionDefinitionFormat1: actionDefinitionFormat1,
		}
	default:
		return nil, fmt.Errorf("decodeE2SmKpmActionDefinition() %v not yet implemented", e2SmKpmActionDefinitionC.actionDefinition_formats.present)
	}

	return e2SmKpmActionDefinition, nil
}
