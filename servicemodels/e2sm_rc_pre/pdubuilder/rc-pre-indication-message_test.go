// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationMsg(t *testing.T) {
	var plmnID = "ONF"

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationMsg(plmnID)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Message: % x", per)

}
