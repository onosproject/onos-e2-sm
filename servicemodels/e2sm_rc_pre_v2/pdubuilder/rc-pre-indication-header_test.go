// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_v2/rcprev2ctypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_v2/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_rc_pre_v2.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationHeader(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprev2ctypes.XerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := rcprev2ctypes.PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: % x", per)

}
