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

func createEngnbID() *e2sm_kpm_v2.EngnbId {

	return &e2sm_kpm_v2.EngnbId{
		EngnbId: &e2sm_kpm_v2.EngnbId_GNbId{
			GNbId: &e2sm_kpm_v2.BitString{
				Value: []byte{0xd4, 0xbc, 0x0c},
				Len:   22,
			},
		},
	}
}

func Test_xerEncodeEngnbID(t *testing.T) {

	engnbID := createEngnbID()

	xer, err := xerEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 81, len(xer))
	t.Logf("ENGNBID XER\n%s", string(xer))
}

func Test_xerDecodeEngnbID(t *testing.T) {

	engnbID := createEngnbID()

	xer, err := xerEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 81, len(xer))
	t.Logf("ENGNBID XER\n%s", string(xer))

	result, err := xerDecodeEngnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ENGNBID XER - decoded\n%s", result)
}

func Test_perEncodeEngnbID(t *testing.T) {

	engnbID := createEngnbID()

	per, err := perEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("ENGNBID PER\n%v", hex.Dump(per))
}

func Test_perDecodeEngnbID(t *testing.T) {

	engnbID := createEngnbID()

	per, err := perEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("ENGNBID PER\n%v", hex.Dump(per))

	result, err := perDecodeEngnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ENGNBID PER - decoded\n%s", result)
}
