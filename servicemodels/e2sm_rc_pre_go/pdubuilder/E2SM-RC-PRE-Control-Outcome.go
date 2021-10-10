// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
)

func CreateE2SmRcPreControlOutcome(elementList []*e2sm_rc_pre_go.RanparameterItem) (*e2sm_rc_pre_go.E2SmRcPreControlOutcome, error) {

	e2smRcPrePdu := e2sm_rc_pre_go.E2SmRcPreControlOutcome{
		E2SmRcPreControlOutcome: &e2sm_rc_pre_go.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			ControlOutcomeFormat1: &e2sm_rc_pre_go.E2SmRcPreControlOutcomeFormat1{
				OutcomeElementList: elementList,
			},
		},
	}

	//if err := e2smRcPrePdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &e2smRcPrePdu, nil
}

//CreateE2SmRcPreControlOutcomeEmpty is used just to generate signature, in case it is needed
func CreateE2SmRcPreControlOutcomeEmpty() (*e2sm_rc_pre_go.E2SmRcPreControlOutcome, error) {

	e2smRcPrePdu := e2sm_rc_pre_go.E2SmRcPreControlOutcome{
		E2SmRcPreControlOutcome: &e2sm_rc_pre_go.E2SmRcPreControlOutcome_ControlOutcomeFormat1{
			//ControlOutcomeFormat1: nil,
		},
	}

	//if err := e2smRcPrePdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &e2smRcPrePdu, nil
}

func CreateRanParameterItem(ranParameterID int32) *e2sm_rc_pre_go.RanparameterItem {

	return &e2sm_rc_pre_go.RanparameterItem{
		RanParameterId: &e2sm_rc_pre_go.RanparameterId{
			Value: ranParameterID,
		},
	}
}
