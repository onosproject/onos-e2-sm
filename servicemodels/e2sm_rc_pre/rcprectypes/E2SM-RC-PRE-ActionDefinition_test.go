// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreActionDefinition(t *testing.T) {

	e2SmRcPreActionDefinition, err := pdubuilder.CreateE2SmRcPreActionDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreActionDefinition(e2SmRcPreActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 103, len(xer))
	t.Logf("E2SM-RC-PRE-ActionDefinition XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreActionDefinition(t *testing.T) {

	e2SmRcPreActionDefinition, err := pdubuilder.CreateE2SmRcPreActionDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreActionDefinition(e2SmRcPreActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 103, len(xer))
	t.Logf("E2SM-RC-PRE-ActionDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreActionDefinition(xer)
	assert.NilError(t, err)
	assert.Equal(t, e2SmRcPreActionDefinition.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}

func Test_PerEncodeE2SmRcPreActionDefinition(t *testing.T) {

	e2SmRcPreActionDefinition, err := pdubuilder.CreateE2SmRcPreActionDefinition(12)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmRcPreActionDefinition(e2SmRcPreActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-RC-PRE-ActionDefinition PER\n%s", string(per))
}

func Test_PerDecodeE2SmRcPreActionDefinition(t *testing.T) {

	e2SmRcPreActionDefinition, err := pdubuilder.CreateE2SmRcPreActionDefinition(12)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmRcPreActionDefinition(e2SmRcPreActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-RC-PRE-ActionDefinition PER\n%s", string(per))

	result, err := PerDecodeE2SmRcPreActionDefinition(per)
	assert.NilError(t, err)
	assert.Equal(t, e2SmRcPreActionDefinition.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
