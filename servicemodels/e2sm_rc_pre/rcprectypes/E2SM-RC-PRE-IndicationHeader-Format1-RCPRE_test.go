// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_EncodingE2SmRcPreIndicationHeaderFormat1EutraCGI(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeaderFormat1(newE2SmRcPrePdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (EutraCGI) XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreIndicationHeaderFormat1(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (EutraCGI) XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())

	per, err := PerEncodeE2SmRcPreIndicationHeaderFormat1(newE2SmRcPrePdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (EutraCGI) PER\n%v", hex.Dump(per))

	resultPer, err := PerDecodeE2SmRcPreIndicationHeaderFormat1(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (EutraCGI) PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_EncodingE2SmRcPreIndicationHeaderFormat1NrCGI(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x9f, 0xf0},
		Len:   36, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeaderFormat1(newE2SmRcPrePdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (NRCGI) XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreIndicationHeaderFormat1(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (NRCGI) XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	per, err := PerEncodeE2SmRcPreIndicationHeaderFormat1(newE2SmRcPrePdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (NRCGI) PER\n%v", hex.Dump(per))

	resultPer, err := PerDecodeE2SmRcPreIndicationHeaderFormat1(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader-Format1 (NRCGI) PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}
