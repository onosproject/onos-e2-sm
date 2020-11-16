// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRCGI(t *testing.T) {

	var plmnIDnrcgi = "ONF"

	nrCGI := &e2sm_kpm_ies.Nrcgi{
		PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
			Value: []byte(plmnIDnrcgi),
		},
		NRcellIdentity: &e2sm_kpm_ies.NrcellIdentity{
			Value: &e2sm_kpm_ies.BitString{
				Value: 0x9bcd4, //uint64
				Len:   22,      //uint32
			},
		},
	}

	xer, err := xerEncodeNRCGI(nrCGI)
	assert.NilError(t, err)
	t.Logf("NRCGI XER\n%s", string(xer))
}
