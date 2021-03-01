// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader(controlMessagePriority, plmnIDBytes, &cellID)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 491, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader(controlMessagePriority, plmnIDBytes, &cellID)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 491, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", result)

	//assert.Equal(t, e2SmRcPreControlHeader.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
