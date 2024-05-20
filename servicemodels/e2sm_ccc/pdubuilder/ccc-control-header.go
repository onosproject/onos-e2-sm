// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcControlHeader(controlHeaderFormat *e2smcccv1.ControlHeaderFormat) (*e2smcccv1.E2SmCCcRIcControlHeader, error) {

	msg := &e2smcccv1.E2SmCCcRIcControlHeader{}
	msg.ControlHeaderFormat = controlHeaderFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcControlHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlHeaderFormat1(ricStyleType *e2smcommoniesv1.RicStyleType) (*e2smcccv1.E2SmCCcControlHeaderFormat1, error) {

	msg := &e2smcccv1.E2SmCCcControlHeaderFormat1{}
	msg.RicStyleType = ricStyleType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}
