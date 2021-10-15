// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1

	newE2SmMhoPdu, err := CreateE2SmMhoControlHeader(controlMessagePriority)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoControlHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationHeader: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoControlHeader(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-ControlHeader is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
