// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcIndicationHeader(indicationHeaderFormat *e2smcccv1.IndicationHeaderFormat) (*e2smcccv1.E2SmCCcRIcIndicationHeader, error) {

	msg := &e2smcccv1.E2SmCCcRIcIndicationHeader{}
	msg.IndicationHeaderFormat = indicationHeaderFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcIndicationHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationHeaderFormat1(indicationReason e2smcccv1.IndicationReason, eventTime []byte) (*e2smcccv1.E2SmCCcIndicationHeaderFormat1, error) {

	msg := &e2smcccv1.E2SmCCcIndicationHeaderFormat1{}
	msg.IndicationReason = indicationReason
	msg.EventTime = eventTime

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}
