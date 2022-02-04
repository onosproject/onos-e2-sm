// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createGlobalEnbID1() *e2sm_kpm_v2.GlobalEnbId {

	return &e2sm_kpm_v2.GlobalEnbId{
		PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		ENbId: &e2sm_kpm_v2.EnbId{
			EnbId: &e2sm_kpm_v2.EnbId_MacroENbId{
				MacroENbId: &e2sm_kpm_v2.BitString{
					Value: []byte{0xd4, 0xbc, 0x30},
					Len:   20,
				},
			},
		},
	}
}

func createGlobalEnbID2() *e2sm_kpm_v2.GlobalEnbId {

	return &e2sm_kpm_v2.GlobalEnbId{
		PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		ENbId: &e2sm_kpm_v2.EnbId{
			EnbId: &e2sm_kpm_v2.EnbId_HomeENbId{
				HomeENbId: &e2sm_kpm_v2.BitString{
					Value: []byte{0xd4, 0xbc, 0x09, 0x00},
					Len:   28,
				},
			},
		},
	}
}

func Test_xerEncodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID1()

	xer, err := xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 182, len(xer))
	t.Logf("GlobalEnbID (Macro) XER\n%s", string(xer))

	globalEnbID = createGlobalEnbID2()

	xer, err = xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 188, len(xer))
	t.Logf("GlobalEnbID (Home) XER\n%s", string(xer))
}

func Test_xerDecodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID1()

	xer, err := xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 182, len(xer))
	t.Logf("GlobalEnbID (Macro) XER\n%s", string(xer))

	result, err := xerDecodeGlobalEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID (Macro) XER - decoded\n%s", result)

	globalEnbID = createGlobalEnbID2()

	xer, err = xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 188, len(xer))
	t.Logf("GlobalEnbID (Home) XER\n%s", string(xer))

	result, err = xerDecodeGlobalEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID (Home) XER - decoded\n%s", result)
}

func Test_perEncodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID1()

	per, err := perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalEnbID (Macro) PER\n%v", hex.Dump(per))

	globalEnbID = createGlobalEnbID2()

	per, err = perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 9, len(per))
	t.Logf("GlobalEnbID (Home) PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID1()

	per, err := perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalEnbID (Macro) PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID (Macro) PER - decoded\n%s", result)

	globalEnbID = createGlobalEnbID2()

	per, err = perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 9, len(per))
	t.Logf("GlobalEnbID (Home) PER\n%v", hex.Dump(per))

	result, err = perDecodeGlobalEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID (Home) PER - decoded\n%s", result)
}
