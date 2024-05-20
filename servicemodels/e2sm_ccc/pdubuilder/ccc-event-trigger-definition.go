// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmCCcRIceventTriggerDefinition(eventTriggerDefinitionFormat *e2smcccv1.EventTriggerDefinitionFormat) (*e2smcccv1.E2SmCCcRIceventTriggerDefinition, error) {

	msg := &e2smcccv1.E2SmCCcRIceventTriggerDefinition{}
	msg.EventTriggerDefinitionFormat = eventTriggerDefinitionFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIceventTriggerDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat1(listOfNodeLevelConfigurationStructuresForEventTrigger *e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1{}
	msg.ListOfNodeLevelConfigurationStructuresForEventTrigger = listOfNodeLevelConfigurationStructuresForEventTrigger

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfRanconfigurationStructuresForEventTrigger(value []*e2smcccv1.RanconfigurationStructureForEventTrigger) (*e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger, error) {

	msg := &e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfRanconfigurationStructuresForEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanConfigurationStructureName(value []byte) (*e2smcccv1.RanConfigurationStructureName, error) {

	msg := &e2smcccv1.RanConfigurationStructureName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanConfigurationStructureName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfAttributes(value []*e2smcccv1.AttributeName) (*e2smcccv1.ListOfAttributes, error) {

	msg := &e2smcccv1.ListOfAttributes{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfAttributes() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAttributeName(value []byte) (*e2smcccv1.AttributeName, error) {

	msg := &e2smcccv1.AttributeName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAttributeName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat2(listOfCellLevelConfigurationStructuresForEventTrigger *e2smcccv1.ListOfCellLevelConfigurationStructuresForEventTrigger) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2{}
	msg.ListOfCellLevelConfigurationStructuresForEventTrigger = listOfCellLevelConfigurationStructuresForEventTrigger

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat3(period int32) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3{}
	msg.Period = period

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat3() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1(e2SmCccEventTriggerDefinitionFormat1 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat1{
			E2SmCccEventTriggerDefinitionFormat1: e2SmCccEventTriggerDefinitionFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}

func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2(e2SmCccEventTriggerDefinitionFormat2 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat2{
			E2SmCccEventTriggerDefinitionFormat2: e2SmCccEventTriggerDefinitionFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}

func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3(e2SmCccEventTriggerDefinitionFormat3 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat3{
			E2SmCccEventTriggerDefinitionFormat3: e2SmCccEventTriggerDefinitionFormat3,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3() error validating PDU %s", err.Error())
	}

	return item, nil
}
