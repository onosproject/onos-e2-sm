// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_EncodeCellGlobalIDEutraCGI(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellGlobalID := &e2sm_rc_pre_v2.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_v2.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_rc_pre_v2.Eutracgi{
				PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_rc_pre_v2.EutracellIdentity{
					Value: &e2sm_rc_pre_v2.BitString{
						Value: 0x9bcd4ab, //uint64
						Len:   28,        //uint32
					},
				},
			},
		},
	}

	xer, err := xerEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (EutraCGI) XER\n%s", string(xer))

	result, err := xerDecodeCellGlobalID(xer)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (EutraCGI) XER - decoded\n%s", result)

	per, err := perEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (EutraCGI) PER\n%s", hex.Dump(xer))

	resultPer, err := perDecodeCellGlobalID(per)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (EutraCGI) PER - decoded\n%s", resultPer)
}

func Test_EncodeCellGlobalIDNrCGI(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellGlobalID := &e2sm_rc_pre_v2.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_v2.CellGlobalId_NrCgi{
			NrCgi: &e2sm_rc_pre_v2.Nrcgi{
				PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				NRcellIdentity: &e2sm_rc_pre_v2.NrcellIdentity{
					Value: &e2sm_rc_pre_v2.BitString{
						Value: 0x9bcd4ab, //uint64
						Len:   36,        //uint32
					},
				},
			},
		},
	}

	xer, err := xerEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (NrCGI) XER\n%s", string(xer))

	result, err := xerDecodeCellGlobalID(xer)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (NrCGI) XER - decoded\n%s", result)

	per, err := perEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (NrCGI) PER\n%s", hex.Dump(xer))

	resultPer, err := perDecodeCellGlobalID(per)
	assert.NilError(t, err)
	t.Logf("CellGlobalID (NrCGI) PER - decoded\n%s", resultPer)
}
