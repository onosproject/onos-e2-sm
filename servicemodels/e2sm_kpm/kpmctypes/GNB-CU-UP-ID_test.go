// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGnbCuUpId(t *testing.T) {

	var gnbCuUpId int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpId,
	}

	xer, err := xerEncodeGnbCuUpId(c)
	assert.NilError(t, err)
	t.Logf("GNB-CU-UP-ID XER\n%s", string(xer))
}

func Test_decodeGnbCuUpId(t *testing.T) {

	var gnbCuUpId int64 = 153

	c := &e2sm_kpm_ies.GnbCuUpId{
		Value: gnbCuUpId,
	}

	gnbCuUpIdC := newGnbCuUpId(c)
	result, err := decodeGnbCuUpId(gnbCuUpIdC)
	assert.NilError(t, err, "Something weird should have happened. Try again")
	assert.Equal(t, gnbCuUpId, result.Value, "Something went wrong, comparison is incorrect")
}
