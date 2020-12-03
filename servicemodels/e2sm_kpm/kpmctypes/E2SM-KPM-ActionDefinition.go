// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-ActionDefinition.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func XerEncodeE2SmKpmActionDefinition(actionDefinition *e2sm_kpm_ies.E2SmKpmActionDefinition) ([]byte, error) {

	actionDefinitionCP := newE2SmKpmActionDefinition(actionDefinition)

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_ActionDefinition, unsafe.Pointer(actionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmKpmActionDefinition(actionDefinition *e2sm_kpm_ies.E2SmKpmActionDefinition) ([]byte, error) {
	actionDefinitionCP := newE2SmKpmActionDefinition(actionDefinition)

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_KPM_ActionDefinition, unsafe.Pointer(actionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmKpmActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmKpmActionDefinition(bytes []byte) (*e2sm_kpm_ies.E2SmKpmActionDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_KPM_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmKpmActionDefinition((*C.E2SM_KPM_ActionDefinition_t)(unsafePtr)), nil
}

func PerDecodeE2SmKpmActionDefinition(bytes []byte) (*e2sm_kpm_ies.E2SmKpmActionDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_KPM_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmKpmActionDefinition((*C.E2SM_KPM_ActionDefinition_t)(unsafePtr)), nil
}

func newE2SmKpmActionDefinition(actionDefinition *e2sm_kpm_ies.E2SmKpmActionDefinition) *C.E2SM_KPM_ActionDefinition_t {

	ricStyleTypeC := newRicStyleType(actionDefinition.GetRicStyleType())

	actionDefinitionC := C.E2SM_KPM_ActionDefinition_t{
		ric_Style_Type: *ricStyleTypeC,
	}

	return &actionDefinitionC
}

func decodeE2SmKpmActionDefinition(actionDefinitionC *C.E2SM_KPM_ActionDefinition_t) *e2sm_kpm_ies.E2SmKpmActionDefinition {

	ricStyleType := decodeRicStyleType(&actionDefinitionC.ric_Style_Type)

	return &e2sm_kpm_ies.E2SmKpmActionDefinition{
		RicStyleType: ricStyleType,
	}
}
