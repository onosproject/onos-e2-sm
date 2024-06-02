// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	encoder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/encoder"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIceventTriggerDefinition(t *testing.T) {
	eventTriggerDefinitionFormat := &e2smcccv1.EventTriggerDefinitionFormat{}
	result, err := CreateE2SmCCcRIceventTriggerDefinition(eventTriggerDefinitionFormat)

	encodedMsg, err := encoder.PerEncodeE2SmCCcRIceventTriggerDefinition(result)
	assert.NoError(t, err)
	assert.NotNil(t, encodedMsg)
	decodedMsg, err := encoder.PerDecodeE2SmCCcRIceventTriggerDefinition(encodedMsg)
	assert.NoError(t, err)
	assert.NotNil(t, decodedMsg)
	assert.Equal(t, result, decodedMsg)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateE2SmCCcEventTriggerDefinitionFormat1(t *testing.T) {
	var value []*e2smcccv1.RanconfigurationStructureForEventTrigger
	for i := 0; i < 1; i++ {
		value = append(value, &e2smcccv1.RanconfigurationStructureForEventTrigger{})
	}
	listOfNodeLevelConfigurationStructuresForEventTrigger := &e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger{
		Value: value,
	}
	result, err := CreateE2SmCCcEventTriggerDefinitionFormat1(listOfNodeLevelConfigurationStructuresForEventTrigger)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)

}

func TestCreateListOfRanconfigurationStructuresForEventTrigger(t *testing.T) {
	var value []*e2smcccv1.RanconfigurationStructureForEventTrigger
	for i := 0; i < 1; i++ {
		value = append(value, &e2smcccv1.RanconfigurationStructureForEventTrigger{})
	}
	result, err := CreateListOfRanconfigurationStructuresForEventTrigger(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateRanConfigurationStructureName(t *testing.T) {
	value := []byte("Test")
	result, err := CreateRanConfigurationStructureName(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateListOfAttributes(t *testing.T) {
	value := []*e2smcccv1.AttributeName{}
	result, err := CreateListOfAttributes(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateAttributeName(t *testing.T) {
	value := []byte("Test")
	result, err := CreateAttributeName(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateE2SmCCcEventTriggerDefinitionFormat3(t *testing.T) {
	period := int32(10)
	result, err := CreateE2SmCCcEventTriggerDefinitionFormat3(period)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1(t *testing.T) {
	e2SmCccEventTriggerDefinitionFormat1 := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1{}
	result, err := CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1(e2SmCccEventTriggerDefinitionFormat1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2(t *testing.T) {
	e2SmCccEventTriggerDefinitionFormat2 := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2{}
	result, err := CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2(e2SmCccEventTriggerDefinitionFormat2)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3(t *testing.T) {
	// Mock input parameter
	e2SmCccEventTriggerDefinitionFormat3 := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3{}
	result, err := CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3(e2SmCccEventTriggerDefinitionFormat3)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}
