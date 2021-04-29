// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createNrcgi() *e2sm_mho.Nrcgi {
	return &e2sm_mho.Nrcgi{
		PLmnIdentity: &e2sm_mho.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		NRcellIdentity: &e2sm_mho.NrcellIdentity{
			Value: &e2sm_mho.BitString{
				Value: 0x9bcd4,
				Len:   36,
			},
		},
	}
}

func Test_xerEncodeNrcgi(t *testing.T) {

	nrcgi := createNrcgi()

	xer, err := xerEncodeNrcgi(nrcgi)
	assert.NilError(t, err)
	assert.Equal(t, 149, len(xer))
	t.Logf("NRCGI XER\n%s", string(xer))
}

func Test_xerDecodeNrcgi(t *testing.T) {

	nrcgi := createNrcgi()

	xer, err := xerEncodeNrcgi(nrcgi)
	assert.NilError(t, err)
	assert.Equal(t, 149, len(xer))
	t.Logf("NRCGI XER\n%s", string(xer))

	result, err := xerDecodeNrcgi(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NRCGI XER - decode\n%s", result)
}

func Test_perEncodeNrcgi(t *testing.T) {

	nrcgi := createNrcgi()

	per, err := perEncodeNrcgi(nrcgi)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("NRCGI PER\n%s", string(per))
}

func Test_perDecodeNrcgi(t *testing.T) {

	nrcgi := createNrcgi()

	per, err := perEncodeNrcgi(nrcgi)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("NRCGI PER\n%s", string(per))

	result, err := perDecodeNrcgi(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NRCGI PER - decoded\n%v", result)
}
