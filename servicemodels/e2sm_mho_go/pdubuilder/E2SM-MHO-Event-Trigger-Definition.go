// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMhoEventTriggerDefinition(triggerType e2sm_mho_go.MhoTriggerType) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType: triggerType,
		//ReportingPeriodMs: &rtPeriod,
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoEventTriggerDefinition(): error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod int32) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_mho_go.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC,
		ReportingPeriodMs: &rtPeriod,
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoEventTriggerDefinitionPeriodic(): error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport() (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType: e2sm_mho_go.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT,
		//ReportingPeriodMs: 0,
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport(): error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}
