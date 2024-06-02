// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateListOfCellsForRanfunctionDefinition(t *testing.T) {
	value := []*e2smcccv1.CellForRanfunctionDefinition{{}}
	result, err := CreateListOfCellsForRanfunctionDefinition(value)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)
}
