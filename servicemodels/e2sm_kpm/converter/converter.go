// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package converter

import (
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	e2smkpmies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

// The `main` package and the `kpmctypes` cannot be referenced by outside applications,
// so this package is to create a facade to them

// RanFuncDescriptionObjectToASN1 -
func RanFuncDescriptionObjectToASN1(protoObj *e2smkpmies.E2SmKpmRanfunctionDescription) ([]byte, error) {
	perBytes, err := kpmctypes.PerEncodeE2SmKpmRanfunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

// EventTriggerDefinitionObjectToASN1 -
func EventTriggerDefinitionObjectToASN1(protoObj *e2smkpmies.E2SmKpmEventTriggerDefinition) ([]byte, error) {
	perBytes, err := kpmctypes.PerEncodeE2SmKpmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

// ActionDefinitionObjectToASN1 -
func ActionDefinitionObjectToASN1(protoObj *e2smkpmies.E2SmKpmActionDefinition) ([]byte, error) {
	perBytes, err := kpmctypes.PerEncodeE2SmKpmActionDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

// IndicationHeaderObjectToASN1 -
func IndicationHeaderObjectToASN1(protoObj *e2smkpmies.E2SmKpmIndicationHeader) ([]byte, error) {
	perBytes, err := kpmctypes.PerEncodeE2SmKpmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding PerEncodeE2SmKpmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

// IndicationMessageObjectToASN1 -
func IndicationMessageObjectToASN1(protoObj *e2smkpmies.E2SmKpmIndicationMessage) ([]byte, error) {
	perBytes, err := kpmctypes.PerEncodeE2SmKpmIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding PerEncodeE2SmKpmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}
