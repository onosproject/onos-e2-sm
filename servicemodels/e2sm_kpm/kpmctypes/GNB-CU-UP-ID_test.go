// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGnbCuUpId(t *testing.T) {

	var gnbCuUpID int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpID,
	}

	xer, err := xerEncodeGnbCuUpID(c)
	assert.NilError(t, err)
	t.Logf("GNB-CU-UP-ID XER\n%s", string(xer))
}

func Test_decodeGnbCuUpId(t *testing.T) {

	var gnbCuUpID int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpID,
	}

	gnbCuUpIDC := newGnbCuUpID(c)
	result := decodeGnbCuUpID(gnbCuUpIDC)
	assert.Equal(t, gnbCuUpID, result.Value, "Something went wrong, comparison is incorrect")
}

func Test_perEncodeGnbCuUpId(t *testing.T) {

	var gnbCuUpID int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpID,
	}

	gnbCuUpIDPer, err := perEncodeGnbCuUpID(c)
	assert.NilError(t, err)
	assert.Assert(t, gnbCuUpIDPer != nil)
	assert.Equal(t, 2, len(gnbCuUpIDPer))
	t.Logf("GNB-CU-UP-ID PER\n%v", gnbCuUpIDPer)
}

func Test_perDecodeGnbCuUpId(t *testing.T) {

	var gnbCuUpID int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpID,
	}

	gnbCuUpIDPer, err := perEncodeGnbCuUpID(c)
	assert.NilError(t, err)
	assert.Assert(t, gnbCuUpIDPer != nil)
	t.Logf("GNB-CU-UP-ID PER\n%v", gnbCuUpIDPer)

	result, err := perDecodeGnbCuUpID(gnbCuUpIDPer)
	assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	t.Logf("GNB-CU-UP-ID PER decoded is \n%v", result)

	assert.Equal(t, gnbCuUpID, result.Value, "Encoded and decoded values are not the same")
}
