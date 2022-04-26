//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcControlOutcomeFormat1(rpl []*e2smrcv1.E2SmRcControlOutcomeFormat1Item) (*e2smrcv1.E2SmRcControlOutcome, error) {

	ch := &e2smrcv1.E2SmRcControlOutcome{
		RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats{
			RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat1{
				ControlOutcomeFormat1: &e2smrcv1.E2SmRcControlOutcomeFormat1{
					RanPList: rpl,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlOutcomeFormat1Item(rpID int64, rpv *e2smrcv1.RanparameterValue) (*e2smrcv1.E2SmRcControlOutcomeFormat1Item, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat1Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterValue: rpv,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}
