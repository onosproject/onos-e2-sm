// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	var ranParameterValue int32 = 10
	e2SmRcPreControlOutcome, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID, ranParameterValue)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	assert.NilError(t, err)
	assert.Equal(t, 427, len(xer))
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	var ranParameterValue int32 = 10
	e2SmRcPreControlOutcome, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID, ranParameterValue)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	assert.NilError(t, err)
	assert.Equal(t, 427, len(xer))
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlOutcome(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%s", result)
}
