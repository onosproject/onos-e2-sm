// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}

	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader(&controlMessagePriority, cgi)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	//assert.Equal(t, 491, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}

	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader(&controlMessagePriority, cgi)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	//assert.Equal(t, 491, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlHeader decoded XER is\n%v", result)
	assert.DeepEqual(t, e2SmRcPreControlHeader.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, e2SmRcPreControlHeader.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_PerDecodeE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xba, 0x4d, 0xcb, 0x90},
		Len:   28, //uint32
	}

	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader(&controlMessagePriority, cgi)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("E2SM-RC-PRE-ControlHeader PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlHeader decoded PER is\n%v", result)
	assert.DeepEqual(t, e2SmRcPreControlHeader.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, e2SmRcPreControlHeader.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetControlHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}
