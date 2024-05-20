// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcIndicationMessage(indicationMessageFormat *e2smcccv1.IndicationMessageFormat) (*e2smcccv1.E2SmCCcRIcIndicationMessage, error) {

	msg := &e2smcccv1.E2SmCCcRIcIndicationMessage{}
	msg.IndicationMessageFormat = indicationMessageFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcIndicationMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationMessageFormat1(listOfConfigurationStructuresReported *e2smcccv1.ListOfConfigurationsReported) (*e2smcccv1.E2SmCCcIndicationMessageFormat1, error) {

	msg := &e2smcccv1.E2SmCCcIndicationMessageFormat1{}
	msg.ListOfConfigurationStructuresReported = listOfConfigurationStructuresReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfConfigurationsReported(value []*e2smcccv1.ConfigurationStructure) (*e2smcccv1.ListOfConfigurationsReported, error) {

	msg := &e2smcccv1.ListOfConfigurationsReported{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfConfigurationsReported() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationMessageFormat2(listOfCellsReported *e2smcccv1.ListOfCellsReported) (*e2smcccv1.E2SmCCcIndicationMessageFormat2, error) {

	msg := &e2smcccv1.E2SmCCcIndicationMessageFormat2{}
	msg.ListOfCellsReported = listOfCellsReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationMessageFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}
