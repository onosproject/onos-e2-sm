// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
)

func CreateE2SmKpmEventTriggerDefinition(rtPeriod int64) (*e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition, error) {
	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_kpm_v2_go.EventTriggerDefinitionFormats{
			E2SmKpmEventTriggerDefinition: &e2sm_kpm_v2_go.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinitionFormat1{
					ReportingPeriod: rtPeriod,
				},
			},
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}
