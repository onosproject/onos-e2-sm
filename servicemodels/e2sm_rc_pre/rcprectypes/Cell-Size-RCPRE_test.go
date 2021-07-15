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

func createCellSize() e2sm_rc_pre_v2.CellSize {

	return e2sm_rc_pre_v2.CellSize_CELL_SIZE_MACRO
}

func Test_xerEncodeCellSize(t *testing.T) {

	cellSize := createCellSize()

	xer, err := xerEncodeCellSize(cellSize)
	assert.NilError(t, err)
	//assert.Equal(t, 32, len(xer))
	t.Logf("Cell-Size XER\n%s", string(xer))
}

func Test_xerDecodeCellSize(t *testing.T) {

	cellSize := createCellSize()

	xer, err := xerEncodeCellSize(cellSize)
	assert.NilError(t, err)
	//assert.Equal(t, 32, len(xer))
	t.Logf("Cell-Size XER\n%s", string(xer))

	result, err := xerDecodeCellSize(xer)
	assert.NilError(t, err)
	assert.Equal(t, cellSize, result, "Encoded and decoded CellSize values are not the same")
	t.Logf("Cell-Size XER - decoded\n%v", result)
}

func Test_perEncodeCellSize(t *testing.T) {

	cellSize := createCellSize()

	per, err := perEncodeCellSize(cellSize)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("Cell-Size PER\n%v", hex.Dump(per))
}

func Test_perDecodeCellSize(t *testing.T) {

	cellSize := createCellSize()

	per, err := perEncodeCellSize(cellSize)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("Cell-Size PER\n%v", hex.Dump(per))

	result, err := perDecodeCellSize(per)
	assert.NilError(t, err)
	assert.Equal(t, cellSize, result, "Encoded and decoded CellSize values are not the same")
	t.Logf("Cell-Size PER - decoded\n%v", result)
}
