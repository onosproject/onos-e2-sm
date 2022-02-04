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

func createNrcellIdentity() *e2sm_kpm_v2.NrcellIdentity {
	return &e2sm_kpm_v2.NrcellIdentity{
		Value: &e2sm_kpm_v2.BitString{
			Value: []byte{0xd4, 0xbc, 0x09, 0x00, 0x00},
			Len:   36,
		},
	}
}

func Test_xerEncodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	//assert.Equal(t, 76, len(xer))
	t.Logf("NRCellIdentity XER\n%s", string(xer))
}

func Test_xerDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	//assert.Equal(t, 76, len(xer))
	t.Logf("NRCellIdentity XER\n%s", string(xer))

	result, err := xerDecodeNrcellIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NRCellIdentity XER - decoded\n%s", result)
}

func Test_perEncodeNrcellidentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	per, err := perEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	//assert.Equal(t, 5, len(per))
	t.Logf("NRCellIdentity PER\n%v", hex.Dump(per))
}

func Test_perDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	per, err := perEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	//assert.Equal(t, 5, len(per))
	t.Logf("NRCellIdentity PER\n%v", hex.Dump(per))

	result, err := perDecodeNrcellIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NRCellIdentity PER - decoded\n%v", result)
}
