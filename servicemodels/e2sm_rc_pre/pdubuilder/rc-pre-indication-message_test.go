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

func TestE2SmRcPreIndicationMsg(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	dlArfcn := CreateEArfcn(253)
	var pci int32 = 11
	cellSize := e2sm_rc_pre_v2.CellSize_CELL_SIZE_MACRO

	cellID := e2sm_rc_pre_v2.BitString{
		Value: []byte{0xac, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}
	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	neighbor, err := CreateNrt(cgi, dlArfcn, cellSize, &e2sm_rc_pre_v2.Pci{
		Value: pci,
	})
	assert.NilError(t, err)

	neighbors := make([]*e2sm_rc_pre_v2.Nrt, 0)
	neighbors = append(neighbors, neighbor)
	neighbors = append(neighbors, neighbor)

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes, CreateEArfcn(253), cellSize, pci, neighbors)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Indication Message: \n%s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreIndicationMessage(xer)
	assert.NilError(t, err)
	t.Logf("XER decoded RC-PRE-IndicationMessage is \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())

	per, err := rcprectypes.PerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Message: \n%v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("PER decoded RC-PRE-IndicationMessage is \n%v", result)
	assert.DeepEqual(t, newE2SmRcPrePdu.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), resultPer.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, newE2SmRcPrePdu.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), resultPer.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}
