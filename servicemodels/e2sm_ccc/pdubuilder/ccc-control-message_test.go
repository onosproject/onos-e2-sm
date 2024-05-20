// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcControlMessage(t *testing.T) {
	controlMessageFormat := &e2smcccv1.ControlMessageFormat{}
	result, err := CreateE2SmCCcRIcControlMessage(controlMessageFormat)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateE2SmCCcControlMessageFormat1(t *testing.T) {
	var value []*e2smcccv1.ConfigurationStructureWrite
	for i := 0; i < 1; i++ {
		value = append(value, &e2smcccv1.ConfigurationStructureWrite{})
	}
	listOfConfigurationStructures := &e2smcccv1.ListOfConfigurationStructures{
		Value: value,
	}
	result, err := CreateE2SmCCcControlMessageFormat1(listOfConfigurationStructures)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateConfigurationStructureWrite(t *testing.T) {
	ranConfigurationStructureName := &e2smcccv1.RanConfigurationStructureName{}
	oldValuesOfAttributes := &e2smcccv1.ValuesOfAttributes{}
	newValuesOfAttributes := &e2smcccv1.ValuesOfAttributes{}
	result, err := CreateConfigurationStructureWrite(ranConfigurationStructureName, oldValuesOfAttributes, newValuesOfAttributes)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateE2SmCCcControlMessageFormat2(t *testing.T) {
	var value []*e2smcccv1.CellControlled
	for i := 0; i < 1; i++ {
		value = append(value, &e2smcccv1.CellControlled{})
	}
	listOfCellsReported := &e2smcccv1.ListOfCells{
		Value: value,
	}
	result, err := CreateE2SmCCcControlMessageFormat2(listOfCellsReported)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
