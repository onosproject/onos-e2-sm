// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGnbIdChoice(t *testing.T) {
	
	c := &e2sm_kpm_ies.GnbIdChoice{
		GnbIdChoice: &e2sm_kpm_ies.GnbIdChoice_GnbId{
			GnbId: &e2sm_kpm_ies.BitString{
				Value:                0xABCDEF9,
				Len:                  28,
			},
		},
	}
	
	xer, err := xerEncodeGnbIdChoice(c)
	assert.NilError(t, err)
	t.Logf("GNB-ID-Choice XER\n%s", string(xer))
}
