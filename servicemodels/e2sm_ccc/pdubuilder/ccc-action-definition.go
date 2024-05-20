// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcactionDefinition(ricStyleType *e2smcommoniesv1.RicStyleType, actionDefinitionFormat *e2smcccv1.ActionDefinitionFormat) (*e2smcccv1.E2SmCCcRIcactionDefinition, error) {

	msg := &e2smcccv1.E2SmCCcRIcactionDefinition{}
	msg.RicStyleType = ricStyleType
	msg.ActionDefinitionFormat = actionDefinitionFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcactionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfRanconfigurationStructuresForAdf(value []*e2smcccv1.RanconfigurationStructureForAdf) (*e2smcccv1.ListOfRanconfigurationStructuresForAdf, error) {

	msg := &e2smcccv1.ListOfRanconfigurationStructuresForAdf{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfRanconfigurationStructuresForAdf() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcActionDefinitionFormat1(listOfNodeLevelRanconfigurationStructuresForAdf *e2smcccv1.ListOfRanconfigurationStructuresForAdf) (*e2smcccv1.E2SmCCcActionDefinitionFormat1, error) {

	msg := &e2smcccv1.E2SmCCcActionDefinitionFormat1{}
	msg.ListOfNodeLevelRanconfigurationStructuresForAdf = listOfNodeLevelRanconfigurationStructuresForAdf

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcActionDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcActionDefinitionFormat2(listOfCellConfigurationsToBeReportedForAdf *e2smcccv1.ListOfCellConfigurationsToBeReportedForAdf) (*e2smcccv1.E2SmCCcActionDefinitionFormat2, error) {

	msg := &e2smcccv1.E2SmCCcActionDefinitionFormat2{}
	msg.ListOfCellConfigurationsToBeReportedForAdf = listOfCellConfigurationsToBeReportedForAdf

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcActionDefinitionFormat2() error validating PDU %s", err.Error())
	}
	return msg, nil
}
