// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 324, len(xer))
	t.Logf("E2SM-KPM-EventTriggerDefinition XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 324, len(xer))
	t.Logf("E2SM-KPM-EventTriggerDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmEventTriggerDefinition(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, e2SmKpmEventTriggerDefinitionFormat1.GetPolicyTestList(), result.GetPolicyTestList(), "Encoded and decoded E2SM-KPM-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}

func Test_perEncodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-KPM-EventTriggerDefinition PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-KPM-EventTriggerDefinition PER\n%s", string(per))

	result, err := PerDecodeE2SmKpmEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, e2SmKpmEventTriggerDefinition, result, "Encoded and decoded E2SM-KPM-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}
