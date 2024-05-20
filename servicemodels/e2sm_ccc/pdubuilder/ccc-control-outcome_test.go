// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcControlOutcome(t *testing.T) {
	controlOutcomeFormat := &e2smcccv1.ControlOutcomeFormat{}
	result, err := CreateE2SmCCcRIcControlOutcome(controlOutcomeFormat)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateControlOutcomeFormatE2SmCccControlOutcomeFormat1(t *testing.T) {
	e2SmCccControlOutcomeFormat1 := &e2smcccv1.E2SmCCcControlOutcomeFormat1{}
	result, err := CreateControlOutcomeFormatE2SmCccControlOutcomeFormat1(e2SmCccControlOutcomeFormat1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateControlOutcomeFormatE2SmCccControlOutcomeFormat2(t *testing.T) {
	e2SmCccControlOutcomeFormat2 := &e2smcccv1.E2SmCCcControlOutcomeFormat2{}
	result, err := CreateControlOutcomeFormatE2SmCccControlOutcomeFormat2(e2SmCccControlOutcomeFormat2)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
