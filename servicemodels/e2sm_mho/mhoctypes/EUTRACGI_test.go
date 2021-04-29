// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createEutracgi() *e2sm_mho.Eutracgi {

	return &e2sm_mho.Eutracgi{
		PLmnIdentity: &e2sm_mho.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		EUtracellIdentity: &e2sm_mho.EutracellIdentity{
			Value: &e2sm_mho.BitString{
				Value: 0x9bcd4,
				Len:   28,
			},
		},
	}
}

func Test_xerEncodeEutracgi(t *testing.T) {

	eutracgi := createEutracgi()

	xer, err := xerEncodeEutracgi(eutracgi)
	assert.NilError(t, err)
	assert.Equal(t, 153, len(xer))
	t.Logf("EUTRACGI XER\n%s", string(xer))
}

func Test_xerDecodeEutracgi(t *testing.T) {

	eutracgi := createEutracgi()

	xer, err := xerEncodeEutracgi(eutracgi)
	assert.NilError(t, err)
	assert.Equal(t, 153, len(xer))
	t.Logf("EUTRACGI XER\n%s", string(xer))

	result, err := xerDecodeEutracgi(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EUTRACGI XER - decoded\n%s", result)
}

func Test_perEncodeEutracgi(t *testing.T) {

	eutracgi := createEutracgi()

	per, err := perEncodeEutracgi(eutracgi)
	assert.NilError(t, err)
	assert.Equal(t, 8, len(per))
	t.Logf("EUTRACGI PER\n%s", string(per))
}

func Test_perDecodeEutracgi(t *testing.T) {

	eutracgi := createEutracgi()

	per, err := perEncodeEutracgi(eutracgi)
	assert.NilError(t, err)
	assert.Equal(t, 8, len(per))
	t.Logf("EUTRACGI PER\n%s", string(per))

	result, err := perDecodeEutracgi(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EUTRACGI PER - decoded\n%s", result)
}
