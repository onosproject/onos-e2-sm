// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmActionDefinition(t *testing.T) {

	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 97, len(xer))
	t.Logf("E2SM-KPM-ActionDefinition XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmActionDefinition(t *testing.T) {

	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(12)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 97, len(xer))
	t.Logf("E2SM-KPM-ActionDefinition XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmActionDefinition(xer)
	assert.NilError(t, err)
	assert.Equal(t, e2SmKpmActionDefinition.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}

func Test_perEncodeE2SmKpmActionDefinition(t *testing.T) {

	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(12)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-KPM-ActionDefinition PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmActionDefinition(t *testing.T) {

	e2SmKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinition(12)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmActionDefinition(e2SmKpmActionDefinition)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2SM-KPM-ActionDefinition PER\n%s", string(per))

	result, err := PerDecodeE2SmKpmActionDefinition(per)
	assert.NilError(t, err)
	assert.Equal(t, e2SmKpmActionDefinition.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
