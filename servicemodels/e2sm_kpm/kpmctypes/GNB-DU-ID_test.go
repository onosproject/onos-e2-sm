// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGnbDuId(t *testing.T) {

	var gnbDuId int64 = 153

	c := &e2sm_kpm_ies.GnbDuId{
		Value: gnbDuId,
	}

	xer, err := xerEncodeGnbDuId(c)
	assert.NilError(t, err)
	t.Logf("GNB-DU-ID XER\n%s", string(xer))
}

func Test_decodeGnbDuId(t *testing.T) {

	var gnbDuId int64 = 153

	c := &e2sm_kpm_ies.GnbDuId{
		Value: gnbDuId,
	}

	gnbDuIdC := newGnbDuId(c)
	result := decodeGnbDuId(gnbDuIdC)
	assert.Equal(t, gnbDuId, result.Value, "Something went wrong, comparison is incorrect")
}

func Test_perEncodeGnbDuId(t *testing.T) {

	var gnbDuId int64 = 153

	c := &e2sm_kpm_ies.GnbDuId{
		Value: gnbDuId,
	}

	gnbDuIdPer, err := perEncodeGnbDuId(c)
	assert.NilError(t, err)
	assert.Assert(t, gnbDuIdPer != nil)
	t.Logf("GNB-DU-ID PER\n%v", gnbDuIdPer)
}

func Test_perDecodeGnbDuId(t *testing.T) {

	var gnbDuId int64 = 153

	c := &e2sm_kpm_ies.GnbDuId{
		Value: gnbDuId,
	}

	gnbDuIdPer, err := perEncodeGnbDuId(c)
	assert.NilError(t, err)
	assert.Assert(t, gnbDuIdPer != nil)
	t.Logf("GNB-DU-ID PER\n%v", gnbDuIdPer)

	result, err := perDecodeGnbDuId(gnbDuIdPer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GNB-DU-ID PER decoded is \n%v", result)

	assert.Equal(t, gnbDuId, result.Value, "Encoded and decoded values are not the same")
}
