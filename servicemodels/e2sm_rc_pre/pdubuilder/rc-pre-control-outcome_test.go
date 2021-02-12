// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20
	var ranParameterValue int32 = 10

	newE2SmRcPrePdu, err := CreateE2SmRcPreControlOutcome(ranParameterID, ranParameterValue)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlOutcome(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Outcome: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlOutcome(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Outcome: % x", per)

}
