// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIcControlMessage(controlMessageFormat *e2smcccv1.ControlMessageFormat) (*e2smcccv1.E2SmCCcRIcControlMessage, error) {

	msg := &e2smcccv1.E2SmCCcRIcControlMessage{}
	msg.ControlMessageFormat = controlMessageFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcControlMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlMessageFormat1(listOfConfigurationStructures *e2smcccv1.ListOfConfigurationStructures) (*e2smcccv1.E2SmCCcControlMessageFormat1, error) {

	msg := &e2smcccv1.E2SmCCcControlMessageFormat1{}
	msg.ListOfConfigurationStructures = listOfConfigurationStructures

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateConfigurationStructureWrite(ranConfigurationStructureName *e2smcccv1.RanConfigurationStructureName, oldValuesOfAttributes *e2smcccv1.ValuesOfAttributes, newValuesOfAttributes *e2smcccv1.ValuesOfAttributes) (*e2smcccv1.ConfigurationStructureWrite, error) {

	msg := &e2smcccv1.ConfigurationStructureWrite{}
	msg.RanConfigurationStructureName = ranConfigurationStructureName
	msg.OldValuesOfAttributes = oldValuesOfAttributes
	msg.NewValuesOfAttributes = newValuesOfAttributes

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateConfigurationStructureWrite() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlMessageFormat2(listOfCellsReported *e2smcccv1.ListOfCells) (*e2smcccv1.E2SmCCcControlMessageFormat2, error) {

	msg := &e2smcccv1.E2SmCCcControlMessageFormat2{}
	msg.ListOfCellsReported = listOfCellsReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlMessageFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}
