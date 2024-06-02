// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	encoder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/encoder"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcactionDefinition(t *testing.T) {
	ricStyleType := &e2smcommoniesv1.RicStyleType{}
	actionDefinitionFormat := &e2smcccv1.ActionDefinitionFormat{}

	result, err := CreateE2SmCCcRIcactionDefinition(ricStyleType, actionDefinitionFormat)

	encodedMsg, err := encoder.PerEncodeE2SmCCcRIcactionDefinition(result)
	assert.NoError(t, err)
	assert.NotNil(t, encodedMsg)
	decodedMsg, err := encoder.PerDecodeE2SmCCcRIcactionDefinition(encodedMsg)
	assert.NoError(t, err)
	assert.NotNil(t, decodedMsg)
	assert.Equal(t, result, decodedMsg)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)

	result, err := CreateE2SmCCcRIcactionDefinition(ricStyleType, actionDefinitionFormat)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateListOfRanconfigurationStructuresForAdf(t *testing.T) {
	validItem := &e2smcccv1.RanconfigurationStructureForAdf{
		//TODO
	}
	var value []*e2smcccv1.RanconfigurationStructureForAdf
	for i := 0; i < 1; i++ {
		value = append(value, validItem)
	}
	result, err := CreateListOfRanconfigurationStructuresForAdf(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	for _, item := range result.Value {
		err = item.Validate()
		assert.NoError(t, err)
	}
}

func TestCreateE2SmCCcActionDefinitionFormat1(t *testing.T) {
	validItem := &e2smcccv1.RanconfigurationStructureForAdf{
		//TODO
	}
	var value []*e2smcccv1.RanconfigurationStructureForAdf
	for i := 0; i < 1; i++ {
		value = append(value, validItem)
	}
	listOfNodeLevelRanconfigurationStructuresForAdf := &e2smcccv1.ListOfRanconfigurationStructuresForAdf{
		Value: value,
	}
	result, err := CreateE2SmCCcActionDefinitionFormat1(listOfNodeLevelRanconfigurationStructuresForAdf)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateE2SmCCcActionDefinitionFormat2(t *testing.T) {
	var value []*e2smcccv1.CellConfigurationToBeReportedForAdf
	for i := 0; i < 1; i++ {
		value = append(value, &e2smcccv1.CellConfigurationToBeReportedForAdf{
			// TODO
		})
	}
	listOfCellConfigurationsToBeReportedForAdf := &e2smcccv1.ListOfCellConfigurationsToBeReportedForAdf{
		Value: value,
	}
	result, err := CreateE2SmCCcActionDefinitionFormat2(listOfCellConfigurationsToBeReportedForAdf)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)

}
