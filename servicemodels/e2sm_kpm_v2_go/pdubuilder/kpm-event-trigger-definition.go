// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmEventTriggerDefinition(rtPeriod int64) (*e2smkpmv2.E2SmKpmEventTriggerDefinition, error) {
	e2SmKpmPdu := e2smkpmv2.E2SmKpmEventTriggerDefinition{
		EventDefinitionFormats: &e2smkpmv2.EventTriggerDefinitionFormats{
			E2SmKpmEventTriggerDefinition: &e2smkpmv2.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2smkpmv2.E2SmKpmEventTriggerDefinitionFormat1{
					ReportingPeriod: rtPeriod,
				},
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmEventTriggerDefinition(): error validating E2SmKpmPDU %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}
