// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1

	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(controlMessagePriority)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: % x", per)

}
