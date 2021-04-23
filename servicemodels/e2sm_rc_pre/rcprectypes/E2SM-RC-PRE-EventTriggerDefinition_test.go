// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 312, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 312, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreEventTriggerDefinition(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_PerEncodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition PER\n%s", hex.Dump(per))
}

func Test_PerDecodeE2SmRcPreEventTriggerDefinition(t *testing.T) {

	E2SmRcPreEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPreEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	//10010c
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
