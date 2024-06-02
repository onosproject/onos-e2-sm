// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcControlOutcome(controlOutcomeFormat *e2smcccv1.ControlOutcomeFormat) (*e2smcccv1.E2SmCCcRIcControlOutcome, error) {

	msg := &e2smcccv1.E2SmCCcRIcControlOutcome{}
	msg.ControlOutcomeFormat = controlOutcomeFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcControlOutcome() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateControlOutcomeFormatE2SmCccControlOutcomeFormat1(e2SmCccControlOutcomeFormat1 *e2smcccv1.E2SmCCcControlOutcomeFormat1) (*e2smcccv1.ControlOutcomeFormat, error) {

	item := &e2smcccv1.ControlOutcomeFormat{
		ControlOutcomeFormat: &e2smcccv1.ControlOutcomeFormat_E2SmCccControlOutcomeFormat1{
			E2SmCccControlOutcomeFormat1: e2SmCccControlOutcomeFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlOutcomeFormatE2SmCccControlOutcomeFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}

func CreateControlOutcomeFormatE2SmCccControlOutcomeFormat2(e2SmCccControlOutcomeFormat2 *e2smcccv1.E2SmCCcControlOutcomeFormat2) (*e2smcccv1.ControlOutcomeFormat, error) {

	item := &e2smcccv1.ControlOutcomeFormat{
		ControlOutcomeFormat: &e2smcccv1.ControlOutcomeFormat_E2SmCccControlOutcomeFormat2{
			E2SmCccControlOutcomeFormat2: e2SmCccControlOutcomeFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlOutcomeFormatE2SmCccControlOutcomeFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}
