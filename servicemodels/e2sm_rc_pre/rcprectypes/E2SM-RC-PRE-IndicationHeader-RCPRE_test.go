// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_XerEncodeE2SmRcPreIndicationHeader(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeader() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader XER\n%s", string(xer))

	per, err := PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader PER\n%v", hex.Dump(per))
}

func Test_XerDecodeE2SmRcPreIndicationHeader(t *testing.T) {
	E2SmRcPreRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-RC-PRE-Indication-Header-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	E2SmRcPrePdu, err := XerDecodeE2SmRcPreIndicationHeader(E2SmRcPreRanfunctionDescription)
	assert.NilError(t, err)
	assert.NilError(t, E2SmRcPrePdu.Validate())
	t.Logf("E2SM-RC-PRE-IndicationHeader (EUTRACGI) XER - decoded\n%v", E2SmRcPrePdu)
}

func Test_PerDecodeE2SmRcPreIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeader() newE2SmRcPrePdu is nil")

	per, err := PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-IndicationHeader is \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_PerDecodeE2SmRcPreIndicationHeadeNrCGI(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x9f, 0xf0},
		Len:   36, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeaderFormat1() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader (NRCGI) XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreIndicationHeader(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader (NRCGI) XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	per, err := PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader (NRCGI) PER\n%v", hex.Dump(per))

	resultPer, err := PerDecodeE2SmRcPreIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader (NRCGI) PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}
