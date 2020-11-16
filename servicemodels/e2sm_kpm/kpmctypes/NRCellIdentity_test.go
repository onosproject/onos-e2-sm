// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRCellIdentity(t *testing.T) {

	nrCellIdentity := &e2sm_kpm_ies.NrcellIdentity{
		Value: &e2sm_kpm_ies.BitString{
			Value: 0x9bcd4, //uint64
			Len:   22,      //uint32
		},
	}

	xer, err := xerEncodeNRCellIdentity(nrCellIdentity)
	assert.NilError(t, err)
	t.Logf("SNSSAI XER\n%s", string(xer))
}
