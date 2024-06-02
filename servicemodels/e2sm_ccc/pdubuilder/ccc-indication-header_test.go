// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (

	encoder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/encoder"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateE2SmCCcRIcIndicationHeader(t *testing.T) {
	indicationHeaderFormat := &e2smcccv1.IndicationHeaderFormat{}
	result, err := CreateE2SmCCcRIcIndicationHeader(indicationHeaderFormat)
	encodedMsg, err := encoder.PerEncodeE2SmCCcRIcIndicationHeader(result)
	assert.NoError(t, err)
	assert.NotNil(t, encodedMsg)
	decodedMsg, err := encoder.PerDecodeE2SmCCcRIcIndicationHeader(encodedMsg)
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

func TestCreateE2SmCCcIndicationHeaderFormat1(t *testing.T) {
	indicationReason := e2smcccv1.IndicationReason_INDICATION_REASON_UPON_SUBSCRIPTION
	eventTime := make([]byte, 8)
	result, err := CreateE2SmCCcIndicationHeaderFormat1(indicationReason, eventTime)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	err = result.Validate()
	assert.NoError(t, err)

}
