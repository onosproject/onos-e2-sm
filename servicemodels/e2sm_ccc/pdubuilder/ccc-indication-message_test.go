// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"
	encoder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/encoder"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcIndicationMessage(t *testing.T) {
	indicationMessageFormat := &e2smcccv1.IndicationMessageFormat{}
	result, err := CreateE2SmCCcRIcIndicationMessage(indicationMessageFormat)

	encodedMsg, err := encoder.PerEncodeE2SmCCcRIcIndicationMessage(result)
	assert.NoError(t, err)
	assert.NotNil(t, encodedMsg)
	decodedMsg, err := encoder.PerDecodeE2SmCCcRIcIndicationMessage(encodedMsg)
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

func TestCreateE2SmCCcIndicationMessageFormat1(t *testing.T) {
	// Mock input parameter
	listOfConfigurationStructuresReported := &e2smcccv1.ListOfConfigurationsReported{
		Value: []*e2smcccv1.ConfigurationStructure{{}},
	}
	result, err := CreateE2SmCCcIndicationMessageFormat1(listOfConfigurationStructuresReported)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateListOfConfigurationsReported(t *testing.T) {
	value := []*e2smcccv1.ConfigurationStructure{{}}
	result, err := CreateListOfConfigurationsReported(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}

func TestCreateE2SmCCcIndicationMessageFormat2(t *testing.T) {
	// Mock input parameter
	listOfCellsReported := &e2smcccv1.ListOfCellsReported{
		Value: []*e2smcccv1.CellReported{{}},
	}
	result, err := CreateE2SmCCcIndicationMessageFormat2(listOfCellsReported)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}
