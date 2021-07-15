// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}

	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Indication Header: \n%s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreIndicationHeader(xer)
	assert.NilError(t, err)
	t.Logf("XER decoded RC-PRE-IndicationHeader is \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())

	per, err := rcprectypes.PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Header: \n%v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("PER decoded RC-PRE-IndicationHeader is \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), resultPer.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), resultPer.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())

}
