// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createEutracellIdentity() *e2sm_kpm_v2.EutracellIdentity {

	return &e2sm_kpm_v2.EutracellIdentity{
		Value: &e2sm_kpm_v2.BitString{
			Value: 0x9bcd4,
			Len:   28,
		},
	}
}

func Test_xerEncodeEutracellIdentity(t *testing.T) {

	eci := createEutracellIdentity()

	xer, err := xerEncodeEutracellIdentity(eci)
	assert.NilError(t, err)
	assert.Equal(t, 74, len(xer))
	t.Logf("EUTRACellIdentity XER\n%s", string(xer))
}

func Test_xerDecodeEutracellIdentity(t *testing.T) {

	eci := createEutracellIdentity()

	xer, err := xerEncodeEutracellIdentity(eci)
	assert.NilError(t, err)
	assert.Equal(t, 74, len(xer))
	t.Logf("EUTRACellIdentity XER\n%s", string(xer))

	result, err := xerDecodeEutracellIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EUTRACellIdentity XER - decoded\n%s", result)
}

func Test_perEncodeEutracellIdentity(t *testing.T) {

	eci := createEutracellIdentity()

	per, err := perEncodeEutracellIdentity(eci)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EUTRACellIdentity PER\n%s", string(per))
}

func Test_perDecodeEutracellIdentity(t *testing.T) {

	eci := createEutracellIdentity()

	per, err := perEncodeEutracellIdentity(eci)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EUTRACellIdentity PER\n%s", string(per))

	result, err := perDecodeEutracellIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EUTRACellIdentity PER - decoded\n%s", result)
}
