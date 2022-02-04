// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmMhoEventTriggerDefinition(t *testing.T) {

	E2SmMhoEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoEventTriggerDefinition(E2SmMhoEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 306, len(xer))
	t.Logf("E2SM-MHO-EventTriggerDefinition XER\n%s", string(xer))
}

func Test_XerDecodeE2SmMhoEventTriggerDefinition(t *testing.T) {

	E2SmMhoEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoEventTriggerDefinition(E2SmMhoEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 306, len(xer))
	t.Logf("E2SM-MHO-EventTriggerDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoEventTriggerDefinition(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_PerEncodeE2SmMhoEventTriggerDefinition(t *testing.T) {

	E2SmMhoEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmMhoEventTriggerDefinition(E2SmMhoEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-MHO-EventTriggerDefinition PER\n%s", hex.Dump(per))
}

func Test_PerDecodeE2SmMhoEventTriggerDefinition(t *testing.T) {

	E2SmMhoEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(12)
	assert.NilError(t, err)
	per, err := PerEncodeE2SmMhoEventTriggerDefinition(E2SmMhoEventTriggerDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	//10010c
	t.Logf("E2SM-MHO-EventTriggerDefinition PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
