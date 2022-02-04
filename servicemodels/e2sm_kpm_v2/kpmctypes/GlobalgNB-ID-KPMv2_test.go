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

func createGlobalgNbID() *e2sm_kpm_v2.GlobalgNbId {

	return &e2sm_kpm_v2.GlobalgNbId{
		PlmnId: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GnbId: &e2sm_kpm_v2.GnbIdChoice{
			GnbIdChoice: &e2sm_kpm_v2.GnbIdChoice_GnbId{
				GnbId: &e2sm_kpm_v2.BitString{
					Value: []byte{0xd4, 0xbc, 0x0c},
					Len:   22,
				},
			},
		},
	}
}

func Test_xerEncodeGlobalgNbID(t *testing.T) {

	globalGnbID := createGlobalgNbID()

	xer, err := xerEncodeGlobalgNbID(globalGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 160, len(xer))
	t.Logf("GlobalgNbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalgNbID(t *testing.T) {

	globalGnbID := createGlobalgNbID()

	xer, err := xerEncodeGlobalgNbID(globalGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 160, len(xer))
	t.Logf("GlobalgNbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalgNbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalgNbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalgNbID(t *testing.T) {

	globalGnbID := createGlobalgNbID()

	per, err := perEncodeGlobalgNbID(globalGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalgNbID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalgNbID(t *testing.T) {

	globalGnbID := createGlobalgNbID()

	per, err := perEncodeGlobalgNbID(globalGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("GlobalgNbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalgNbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalgNbID PER - decoded\n%v", result)
}
