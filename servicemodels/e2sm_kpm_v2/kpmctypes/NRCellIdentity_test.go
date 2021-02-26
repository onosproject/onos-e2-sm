// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createNrcellIdentity() *e2sm_kpm_v2.NrcellIdentity {
	return &e2sm_kpm_v2.NrcellIdentity{
		Value: &e2sm_kpm_v2.BitString{
			Value: 0x9bcd4,
			Len:   22,
		},
	}
}

func Test_xerEncodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("NRCellIdentity XER\n%s", string(xer))
}

func Test_xerDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("NRCellIdentity XER\n%s", string(xer))

	result, err := xerDecodeNrcellIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeNrcellidentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	per, err := perEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("NRCellIdentity PER\n%s", string(per))
}

func Test_perDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	per, err := perEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("NRCellIdentity PER\n%s", string(per))

	result, err := perDecodeNrcellIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
