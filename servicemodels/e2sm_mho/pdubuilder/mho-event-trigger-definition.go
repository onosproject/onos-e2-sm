// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

func CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod int32) (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC,
		ReportingPeriodMs: rtPeriod,
	}

	E2SmMhoPdu := e2sm_mho.E2SmMhoEventTriggerDefinition{
		E2SmMhoEventTriggerDefinition: &e2sm_mho.E2SmMhoEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: eventDefinitionFormat1,
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport() (*e2sm_mho.E2SmMhoEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT,
		ReportingPeriodMs: 0,
	}

	E2SmMhoPdu := e2sm_mho.E2SmMhoEventTriggerDefinition{
		E2SmMhoEventTriggerDefinition: &e2sm_mho.E2SmMhoEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: eventDefinitionFormat1,
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}
