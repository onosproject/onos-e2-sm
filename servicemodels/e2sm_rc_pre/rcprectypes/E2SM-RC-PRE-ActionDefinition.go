// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ActionDefinition.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func XerEncodeE2SmRcPreActionDefinition(actionDefinition *e2sm_rc_pre_ies.E2SmRcPreActionDefinition) ([]byte, error) {

	actionDefinitionCP := newE2SmRcPreActionDefinition(actionDefinition)

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ActionDefinition, unsafe.Pointer(actionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreActionDefinition(actionDefinition *e2sm_rc_pre_ies.E2SmRcPreActionDefinition) ([]byte, error) {
	actionDefinitionCP := newE2SmRcPreActionDefinition(actionDefinition)

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ActionDefinition, unsafe.Pointer(actionDefinitionCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreActionDefinition() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreActionDefinition(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreActionDefinition, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreActionDefinition((*C.E2SM_RC_PRE_ActionDefinition_t)(unsafePtr)), nil
}

func PerDecodeE2SmRcPreActionDefinition(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreActionDefinition, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ActionDefinition)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreActionDefinition((*C.E2SM_RC_PRE_ActionDefinition_t)(unsafePtr)), nil
}

func newE2SmRcPreActionDefinition(actionDefinition *e2sm_rc_pre_ies.E2SmRcPreActionDefinition) *C.E2SM_RC_PRE_ActionDefinition_t {

	ricStyleTypeC := newRicStyleType(actionDefinition.GetRicStyleType())

	actionDefinitionC := C.E2SM_RC_PRE_ActionDefinition_t{
		ric_Style_Type: *ricStyleTypeC,
	}

	return &actionDefinitionC
}

func decodeE2SmRcPreActionDefinition(actionDefinitionC *C.E2SM_RC_PRE_ActionDefinition_t) *e2sm_rc_pre_ies.E2SmRcPreActionDefinition {

	ricStyleType := decodeRicStyleType(&actionDefinitionC.ric_Style_Type)

	return &e2sm_rc_pre_ies.E2SmRcPreActionDefinition{
		RicStyleType: ricStyleType,
	}
}
