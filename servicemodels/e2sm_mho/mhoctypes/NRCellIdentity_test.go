// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createNrcellIdentity() *e2sm_mho.NrcellIdentity {
	return &e2sm_mho.NrcellIdentity{
		Value: &e2sm_mho.BitString{
			Value: 0x9bcd4,
			Len:   36,
		},
	}
}

func Test_xerEncodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 76, len(xer))
	t.Logf("NRCellIdentity XER\n%s", string(xer))
}

func Test_xerDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	xer, err := xerEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 76, len(xer))
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
	assert.Equal(t, 5, len(per))
	t.Logf("NRCellIdentity PER\n%s", string(per))
}

func Test_perDecodeNrcellIdentity(t *testing.T) {

	nrcID := createNrcellIdentity()

	per, err := perEncodeNrcellIdentity(nrcID)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("NRCellIdentity PER\n%s", string(per))

	result, err := perDecodeNrcellIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NRCellIdentity PER - decoded\n%v", result)
}
