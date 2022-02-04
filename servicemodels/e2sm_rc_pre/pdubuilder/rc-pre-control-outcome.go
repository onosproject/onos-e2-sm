// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreControlOutcome(RANparameterID int32) (*e2sm_rc_pre_v2.E2SmRcPreControlOutcome, error) {

	e2smRcPreOutcomeFormat1 := e2sm_rc_pre_v2.E2SmRcPreControlOutcomeFormat1{
		OutcomeElementList: make([]*e2sm_rc_pre_v2.RanparameterItem, 0),
	}
	outcomeElementList := &e2sm_rc_pre_v2.RanparameterItem{
		RanParameterId: &e2sm_rc_pre_v2.RanparameterId{
			Value: RANparameterID,
		},
	}
	e2smRcPreOutcomeFormat1.OutcomeElementList = append(e2smRcPreOutcomeFormat1.OutcomeElementList, outcomeElementList)
	e2smRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreControlOutcome{
		E2SmRcPreControlOutcome: &e2sm_rc_pre_v2.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			ControlOutcomeFormat1: &e2smRcPreOutcomeFormat1,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}
