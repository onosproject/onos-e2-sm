// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreControlOutcome(elementList []*e2smrcprev2.RanparameterItem) (*e2smrcprev2.E2SmRcPreControlOutcome, error) {

	e2smRcPrePdu := e2smrcprev2.E2SmRcPreControlOutcome{
		E2SmRcPreControlOutcome: &e2smrcprev2.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			ControlOutcomeFormat1: &e2smrcprev2.E2SmRcPreControlOutcomeFormat1{
				OutcomeElementList: elementList,
			},
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

//CreateE2SmRcPreControlOutcomeEmpty is used just to generate signature, in case it is needed
func CreateE2SmRcPreControlOutcomeEmpty() (*e2smrcprev2.E2SmRcPreControlOutcome, error) {

	e2smRcPrePdu := e2smrcprev2.E2SmRcPreControlOutcome{
		E2SmRcPreControlOutcome: &e2smrcprev2.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			//ControlOutcomeFormat1: nil,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateRanParameterItem(ranParameterID int32) *e2smrcprev2.RanparameterItem {

	return &e2smrcprev2.RanparameterItem{
		RanParameterId: &e2smrcprev2.RanparameterId{
			Value: ranParameterID,
		},
	}
}
