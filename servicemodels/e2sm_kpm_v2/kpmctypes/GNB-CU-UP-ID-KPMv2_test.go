// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createGnbCuUpID() *e2sm_kpm_v2.GnbCuUpId {

	return &e2sm_kpm_v2.GnbCuUpId{
		Value: 1234,
	}
}

func Test_xerEncodeGnbCuUpID(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	xer, err := xerEncodeGnbCuUpID(gnbCuUpID)
	assert.NilError(t, err)
	//assert.Equal(t, 34, len(xer))
	t.Logf("GnbCuUpID XER\n%s", string(xer))
}

func Test_xerDecodeGnbCuUpID(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	xer, err := xerEncodeGnbCuUpID(gnbCuUpID)
	assert.NilError(t, err)
	//assert.Equal(t, 34, len(xer))
	t.Logf("GnbCuUpID XER\n%s", string(xer))

	result, err := xerDecodeGnbCuUpID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbCuUpID XER - decoded\n%s", result)
}

func Test_perEncodeGnbCuUpID(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	per, err := perEncodeGnbCuUpID(gnbCuUpID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("GnbCuUpID PER\n%s", string(per))
}

func Test_perDecodeGnbCuUpID(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	per, err := perEncodeGnbCuUpID(gnbCuUpID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("GnbCuUpID PER\n%s", string(per))

	result, err := perDecodeGnbCuUpID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbCuUpID PER - decoded\n%s", result)
}
