// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod int64) (*e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition, error) {

	if rtPeriod > 4294967295 || rtPeriod < 0 {
		return nil, errors.NewInvalid("RT-Period should be within range 0 to 4294967295")
	}

	eventDefinitionFormat1 := &e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_rc_pre_go.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_PERIODIC,
		ReportingPeriodMs: &rtPeriod,
	}

	E2SmRcPrePdu := e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_rc_pre_go.EventTriggerDefinitionFormats{
			E2SmRcPreEventTriggerDefinitionEventDefinitionFormats: &e2sm_rc_pre_go.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateE2SmRcPreEventTriggerDefinitionUponChange() (*e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinitionFormat1{
		TriggerType: e2sm_rc_pre_go.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE,
	}

	E2SmRcPrePdu := e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_rc_pre_go.EventTriggerDefinitionFormats{
			E2SmRcPreEventTriggerDefinitionEventDefinitionFormats: &e2sm_rc_pre_go.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}
