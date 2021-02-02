// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 330, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition XER\n%s", string(xer))
}

func Test_xerDecodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 330, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreEventTriggerDefinition(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, E2SmRcPreEventTriggerDefinitionFormat1.GetPolicyTestList(), result.GetPolicyTestList(), "Encoded and decoded E2SM-RC-PRE-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}

func Test_perEncodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition PER\n%s", string(per))
}

func Test_perDecodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition PER\n%s", string(per))

	result, err := PerDecodeE2SmRcPreEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, E2SmRcPreEventTriggerDefinition, result, "Encoded and decoded E2SM-RC-PRE-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}
