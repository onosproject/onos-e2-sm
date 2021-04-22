// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_v2/rcprev2ctypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlMsg(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue int32 = 10
	var ranParameterID int32 = 1

	newE2SmRcPrePdu, err := CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameterValue)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprev2ctypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := rcprev2ctypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Message: % x", per)

}
