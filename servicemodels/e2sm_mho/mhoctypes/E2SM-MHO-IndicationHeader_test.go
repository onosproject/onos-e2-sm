// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_XerEncodeE2SmMhoIndicationHeader(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_mho.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmMhoIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmMhoPdu != nil, "Test_XerEncodeE2SmMhoIndicationHeader() newE2SmMhoPdu is nil")

	xer, err := XerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader XER\n%s", string(xer))

	per, err := PerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader PER\n%v", hex.Dump(per))
}

func Test_XerDecodeE2SmMhoIndicationHeader(t *testing.T) {
	E2SmMhoRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-MHO-Indication-Header-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	E2SmMhoPdu, err := XerDecodeE2SmMhoIndicationHeader(E2SmMhoRanfunctionDescription)
	assert.NilError(t, err)
	assert.NilError(t, E2SmMhoPdu.Validate())
}

func Test_PerDecodeE2SmMhoIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_mho.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmMhoIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmMhoPdu != nil, "Test_XerEncodeE2SmMhoIndicationHeader() newE2SmMhoPdu is nil")

	per, err := PerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("Decoded RC-PRE-IndicationHeader is \n%v", result)
}
