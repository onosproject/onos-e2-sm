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

func createGlobalenGnbID() *e2sm_kpm_v2.GlobalenGnbId {

	return &e2sm_kpm_v2.GlobalenGnbId{
		PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GNbId: &e2sm_kpm_v2.EngnbId{
			EngnbId: &e2sm_kpm_v2.EngnbId_GNbId{
				GNbId: &e2sm_kpm_v2.BitString{
					Value: []byte{0xd4, 0xbc, 0x0C},
					Len:   22,
				},
			},
		},
	}
}

func Test_xerEncodeGlobalenGnbID(t *testing.T) {

	globalenGnbID := createGlobalenGnbID()

	xer, err := xerEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 176, len(xer))
	t.Logf("GlobalenGnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalenGnbID(t *testing.T) {

	globalenGnbID := createGlobalenGnbID()

	xer, err := xerEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 176, len(xer))
	t.Logf("GlobalenGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalenGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalenGnbID(t *testing.T) {

	globalenGnbID := createGlobalenGnbID()

	per, err := perEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalenGnbID(t *testing.T) {

	globalenGnbID := createGlobalenGnbID()

	per, err := perEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalenGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
}
