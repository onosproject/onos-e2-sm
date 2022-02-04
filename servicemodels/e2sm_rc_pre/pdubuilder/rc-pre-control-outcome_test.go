// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20

	newE2SmRcPrePdu, err := CreateE2SmRcPreControlOutcome(ranParameterID)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlOutcome(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Outcome: \n%s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreControlOutcome(xer)
	assert.NilError(t, err)
	t.Logf("XER Control Outcome - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.GetControlOutcomeFormat1().GetOutcomeElementList()[0].GetRanParameterId().GetValue(), result.GetControlOutcomeFormat1().GetOutcomeElementList()[0].GetRanParameterId().GetValue())

	per, err := rcprectypes.PerEncodeE2SmRcPreControlOutcome(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Outcome: \n%v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreControlOutcome(per)
	assert.NilError(t, err)
	t.Logf("XER Control Outcome - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.GetControlOutcomeFormat1().GetOutcomeElementList()[0].GetRanParameterId().GetValue(), resultPer.GetControlOutcomeFormat1().GetOutcomeElementList()[0].GetRanParameterId().GetValue())
}
