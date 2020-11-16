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
	t.Logf("GnbIdChoice XER\n%s", string(xer))
}
