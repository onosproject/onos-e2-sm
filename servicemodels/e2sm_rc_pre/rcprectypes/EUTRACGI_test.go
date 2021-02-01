// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeEUTRACGI(t *testing.T) {

	var plmnIDeutracgi = "ONF"

	eUTRACGI := &e2sm_rc_pre_ies.Eutracgi{
		PLmnIdentity: &e2sm_rc_pre_ies.PlmnIdentity{
			Value: []byte(plmnIDeutracgi),
		},
		EUtracellIdentity: &e2sm_rc_pre_ies.EutracellIdentity{
			Value: &e2sm_rc_pre_ies.BitString{
				Value: 0x9bcd4ab, //uint64
				Len:   28,        //uint32
			},
		},
	}

	xer, err := xerEncodeEUTRACGI(eUTRACGI)
	assert.NilError(t, err)
	t.Logf("EUTRACGI XER\n%s", string(xer))
}
