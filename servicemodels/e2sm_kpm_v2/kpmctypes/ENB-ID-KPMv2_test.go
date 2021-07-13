// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createEnbIDMacro() *e2sm_kpm_v2.EnbId {

	return &e2sm_kpm_v2.EnbId{
		EnbId: &e2sm_kpm_v2.EnbId_MacroENbId{
			MacroENbId: &e2sm_kpm_v2.BitString{
				Value: []byte{0x00, 0x00, 0x30},
				Len:   20,
			},
		},
	}
}

func createEnbIDHome() *e2sm_kpm_v2.EnbId {

	return &e2sm_kpm_v2.EnbId{
		EnbId: &e2sm_kpm_v2.EnbId_HomeENbId{
			HomeENbId: &e2sm_kpm_v2.BitString{
				Value: []byte{0x12, 0x34, 0x00, 0x00},
				Len:   28,
			},
		},
	}
}

func Test_xerEncodeEnbID(t *testing.T) {

	enbID := createEnbIDMacro()

	xer, err := xerEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 99, len(xer))
	t.Logf("EnbID (Macro) XER\n%s", string(xer))

	enbID = createEnbIDHome()

	xer, err = xerEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 105, len(xer))
	t.Logf("EnbID (Home) XER\n%s", string(xer))
}

func Test_xerDecodeEnbID(t *testing.T) {

	enbID := createEnbIDMacro()

	xer, err := xerEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 99, len(xer))
	t.Logf("EnbID (Macro) XER\n%s", string(xer))

	result, err := xerDecodeEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbID (Macro) Value decoded\n%x", result.GetMacroENbId().Value)
	assert.DeepEqual(t, enbID.GetMacroENbId().Value, result.GetMacroENbId().Value)
	t.Logf("EnbID (Macro) XER - decoded\n%s", result)

	enbID = createEnbIDHome()

	xer, err = xerEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 105, len(xer))
	t.Logf("EnbID (Home) XER\n%s", string(xer))

	result, err = xerDecodeEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbID (Home) XER - decoded\n%s", result)
}

func Test_perEncodeEnbID(t *testing.T) {

	enbID := createEnbIDMacro()

	per, err := perEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbID (Macro) PER\n%v", hex.Dump(per))

	enbID = createEnbIDHome()

	per, err = perEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 5, len(per))
	t.Logf("EnbID (Home) PER\n%v", hex.Dump(per))
}

func Test_perDecodeEnbID(t *testing.T) {

	enbID := createEnbIDMacro()

	per, err := perEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbID (Macro) PER\n%v", hex.Dump(per))

	result, err := perDecodeEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbID (Macro) PER - decoded\n%s", result)

	enbID = createEnbIDHome()

	per, err = perEncodeEnbID(enbID)
	assert.NilError(t, err)
	//assert.Equal(t, 5, len(per))
	t.Logf("EnbID (Home) PER\n%v", hex.Dump(per))

	result, err = perDecodeEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbID (Home) PER - decoded\n%s", result)
}
