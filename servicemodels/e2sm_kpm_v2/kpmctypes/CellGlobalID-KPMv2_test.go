// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createCellGlobalID() *e2sm_kpm_v2.CellGlobalId {

	return &e2sm_kpm_v2.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_kpm_v2.Eutracgi{
				EUtracellIdentity: &e2sm_kpm_v2.EutracellIdentity{
					Value: &e2sm_kpm_v2.BitString{
						Value: 0x9bcd4,
						Len:   28,
					},
				},
				PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
					Value: []byte("ONF"),
				},
			},
		},
	}
}

func Test_xerEncodeCellGlobalID(t *testing.T) {

	cellGlobalID := createCellGlobalID()

	xer, err := xerEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	//assert.Equal(t, 210, len(xer))
	t.Logf("CellGlobalID XER\n%s", string(xer))
}

func Test_xerDecodeCellGlobalID(t *testing.T) {

	cellGlobalID := createCellGlobalID()

	xer, err := xerEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	//assert.Equal(t, 210, len(xer))
	t.Logf("CellGlobalID XER\n%s", string(xer))

	result, err := xerDecodeCellGlobalID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellGlobalID XER - decoded\n%s", result)
}

func Test_perEncodeCellGlobalID(t *testing.T) {

	cellGlobalID := createCellGlobalID()

	per, err := perEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("CellGlobalID PER\n%s", string(per))
}

func Test_perDecodeCellGlobalID(t *testing.T) {

	cellGlobalID := createCellGlobalID()

	per, err := perEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("CellGlobalID PER\n%s", string(per))

	result, err := perDecodeCellGlobalID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellGlobalID PER - decoded\n%v", result)
}
