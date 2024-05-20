// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	"testing"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcControlHeader(t *testing.T) {

	controlHeaderFormat := &e2smcccv1.ControlHeaderFormat{}
	result, err := CreateE2SmCCcRIcControlHeader(controlHeaderFormat)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateE2SmCCcControlHeaderFormat1(t *testing.T) {
	ricStyleType := &e2smcommoniesv1.RicStyleType{}
	result, err := CreateE2SmCCcControlHeaderFormat1(ricStyleType)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
