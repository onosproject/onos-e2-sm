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

func createGnbDuID() *e2sm_kpm_v2.GnbDuId {

	return &e2sm_kpm_v2.GnbDuId{
		Value: 1234,
	}
}

func Test_xerEncodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	xer, err := xerEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 28, len(xer))
	t.Logf("GnbDuID XER\n%s", string(xer))
}

func Test_xerDecodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	xer, err := xerEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 28, len(xer))
	t.Logf("GnbDuID XER\n%s", string(xer))

	result, err := xerDecodeGnbDuID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbDuID XER - decoded\n%s", result)
}

func Test_perEncodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := perEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := perEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	result, err := perDecodeGnbDuID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbDuID PER - decoded\n%s", result)
}
