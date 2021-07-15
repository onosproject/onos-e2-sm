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

func Test_XerEncodeE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	e2SmRcPreControlOutcome, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	assert.NilError(t, err)
	//assert.Equal(t, 298, len(xer))
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	e2SmRcPreControlOutcome, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	assert.NilError(t, err)
	//assert.Equal(t, 298, len(xer))
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlOutcome(xer)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(result.GetControlOutcomeFormat1().GetOutcomeElementList()))
	oe0 := result.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, int32(20), oe0.GetRanParameterId().Value)
	t.Logf("E2SM-RC-PRE-ControlOutcome XER\n%v", result)
}

func Test_PerDecodeE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	e2SmRcPreControlOutcome, err := pdubuilder.CreateE2SmRcPreControlOutcome(ranParameterID)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmRcPreControlOutcome(e2SmRcPreControlOutcome)
	assert.NilError(t, err)
	assert.Equal(t, 6, len(per))
	t.Logf("E2SM-RC-PRE-ControlOutcome PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreControlOutcome(per)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(result.GetControlOutcomeFormat1().GetOutcomeElementList()))
	oe0 := result.GetControlOutcomeFormat1().GetOutcomeElementList()[0]
	assert.Equal(t, int32(20), oe0.GetRanParameterId().Value)
	t.Logf("E2SM-RC-PRE-ControlOutcome decoded PER is \n%v", result)
}
