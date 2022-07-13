//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcControlHeaderFormat1(ueID *e2smcommonies.Ueid, rst int32, rcaID int32) (*e2smrcv1.E2SmRcControlHeader, error) {

	ch := &e2smrcv1.E2SmRcControlHeader{
		RicControlHeaderFormats: &e2smrcv1.RicControlHeaderFormats{
			RicControlHeaderFormats: &e2smrcv1.RicControlHeaderFormats_ControlHeaderFormat1{
				ControlHeaderFormat1: &e2smrcv1.E2SmRcControlHeaderFormat1{
					UeId: ueID,
					RicStyleType: &e2smcommonies.RicStyleType{
						Value: rst,
					},
					RicControlActionId: &e2smrcv1.RicControlActionId{
						Value: rcaID,
					},
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlHeaderFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlHeaderFormat2() (*e2smrcv1.E2SmRcControlHeader, error) {

	ch := &e2smrcv1.E2SmRcControlHeader{
		RicControlHeaderFormats: &e2smrcv1.RicControlHeaderFormats{
			RicControlHeaderFormats: &e2smrcv1.RicControlHeaderFormats_ControlHeaderFormat2{
				ControlHeaderFormat2: &e2smrcv1.E2SmRcControlHeaderFormat2{},
			},
		},
	}

	return ch, nil
}

func CreateRicControlDecisionAccept() e2smrcv1.RicControlDecision {
	return e2smrcv1.RicControlDecision_RIC_CONTROL_DECISION_ACCEPT
}

func CreateRicControlDecisionReject() e2smrcv1.RicControlDecision {
	return e2smrcv1.RicControlDecision_RIC_CONTROL_DECISION_REJECT
}
