// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createCellMeasurementObjectItem() *e2sm_kpm_v2.CellMeasurementObjectItem {

	return &e2sm_kpm_v2.CellMeasurementObjectItem{
		CellObjectId: &e2sm_kpm_v2.CellObjectId{
			Value: "123",
		},
		CellGlobalId: &e2sm_kpm_v2.CellGlobalId{
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
		},
	}
}

func Test_xerEncodeCellMeasurementObjectItem(t *testing.T) {

	item := createCellMeasurementObjectItem()

	xer, err := xerEncodeCellMeasurementObjectItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 350, len(xer))
	t.Logf("CellMeasurementObjectItem XER\n%s", string(xer))
}

func Test_xerDecodeCellMeasurementObjectItem(t *testing.T) {

	item := createCellMeasurementObjectItem()

	xer, err := xerEncodeCellMeasurementObjectItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 350, len(xer))
	t.Logf("CellMeasurementObjectItem XER\n%s", string(xer))

	result, err := xerDecodeCellMeasurementObjectItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellMeasurementObjectItem XER - decoded\n%s", result)
}

func Test_perEncodeCellMeasurementObjectItem(t *testing.T) {

	item := createCellMeasurementObjectItem()

	per, err := perEncodeCellMeasurementObjectItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("CellMeasurementObjectItem PER\n%s", string(per))
}

func Test_perDecodeCellMeasurementObjectItem(t *testing.T) {

	item := createCellMeasurementObjectItem()

	per, err := perEncodeCellMeasurementObjectItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("CellMeasurementObjectItem PER\n%s", string(per))

	result, err := perDecodeCellMeasurementObjectItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellMeasurementObjectItem PER - decoded\n%s", result)
}
