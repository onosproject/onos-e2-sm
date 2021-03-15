// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createCellObjectID() *e2sm_kpm_v2.CellObjectId {

	return &e2sm_kpm_v2.CellObjectId{
		Value: "123",
	}
}

func Test_xerEncodeCellObjectID(t *testing.T) {

	objID := createCellObjectID()

	xer, err := xerEncodeCellObjectID(objID)
	assert.NilError(t, err)
	assert.Equal(t, 33, len(xer))
	t.Logf("CellObjectID XER\n%s", string(xer))
}

func Test_xerDecodeCellObjectID(t *testing.T) {

	objID := createCellObjectID()

	xer, err := xerEncodeCellObjectID(objID)
	assert.NilError(t, err)
	assert.Equal(t, 33, len(xer))
	t.Logf("CellObjectID XER\n%s", string(xer))

	result, err := xerDecodeCellObjectID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellObjectID XER - decoded\n%s", result)
}

func Test_perEncodeCellObjectID(t *testing.T) {

	objID := createCellObjectID()

	per, err := perEncodeCellObjectID(objID)
	assert.NilError(t, err)
	assert.Equal(t, 6, len(per))
	t.Logf("CellObjectID PER\n%s", string(per))
}

func Test_perDecodeCellObjectID(t *testing.T) {

	objID := createCellObjectID()

	per, err := perEncodeCellObjectID(objID)
	assert.NilError(t, err)
	assert.Equal(t, 6, len(per))
	t.Logf("CellObjectID PER\n%s", string(per))

	result, err := perDecodeCellObjectID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellObjectID PER - decoded\n%s", result)
}
