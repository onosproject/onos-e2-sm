// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoEventTriggerDefinitionFormat1() (*e2sm_mho.E2SmMhoEventTriggerDefinitionFormat1, error) {

	var rtPeriod int32 = 12

	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod)
	if err != nil {
		return nil, err
	}

	return newE2SmMhoPdu.GetEventDefinitionFormat1(), nil

}

func Test_XerEncodeE2SmMhoEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmMhoEventTriggerDefinitionFormat1, err := createE2SmMhoEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoEventTriggerDefinitionFormat1(E2SmMhoEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 176, len(xer))
	t.Logf("E2SM-MHO-EventTriggerDefinition-Format1 XER\n%s", string(xer))
}

func Test_XerDecodeE2SmMhoEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmMhoEventTriggerDefinitionFormat1, err := createE2SmMhoEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoEventTriggerDefinitionFormat1(E2SmMhoEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 176, len(xer))
	t.Logf("E2SM-MHO-EventTriggerDefinition-Format1 XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoEventTriggerDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_PerEncodeE2SmMhoEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmMhoEventTriggerDefinitionFormat1, err := createE2SmMhoEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := PerEncodeE2SmMhoEventTriggerDefinitionFormat1(E2SmMhoEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-MHO-EventTriggerDefinition-Format1 PER\n%s", string(per))
}

func Test_PerDecodeE2SmMhoEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmMhoEventTriggerDefinitionFormat1, err := createE2SmMhoEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := PerEncodeE2SmMhoEventTriggerDefinitionFormat1(E2SmMhoEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-MHO-EventTriggerDefinition-Format1 PER\n%s", string(per))

	result, err := PerDecodeE2SmMhoEventTriggerDefinitionFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
