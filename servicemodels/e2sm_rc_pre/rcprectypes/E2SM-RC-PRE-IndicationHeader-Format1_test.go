// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreIndicationHeaderFormat1(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(plmnIDBytes, &cellID)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeaderFormat1(newE2SmRcPrePdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 XER\n%s", string(xer))
}
