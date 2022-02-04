// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(&controlMessagePriority, cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Control Header: %s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("XER ControlHeader - decoded \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}

func TestE2SmRcPreControlHeaderExcludeOptionalField(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(nil, cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Control Header: %s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("XER ControlHeader - decoded \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}

func TestE2SmRcPreControlHeaderExcludeAllOptionalFields(t *testing.T) {
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Control Header: %s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("XER ControlHeader - decoded \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), resultPer.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}
