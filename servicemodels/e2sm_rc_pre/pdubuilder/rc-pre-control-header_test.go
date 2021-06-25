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

func TestE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   36,        //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(&controlMessagePriority, cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))
}

func TestE2SmRcPreControlHeaderExcludeOptionalField(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_v2.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   36,        //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(nil, cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))
}

func TestE2SmRcPreControlHeaderExcludeAllOptionalFields(t *testing.T) {
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader(nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: %v", hex.Dump(per))
}
