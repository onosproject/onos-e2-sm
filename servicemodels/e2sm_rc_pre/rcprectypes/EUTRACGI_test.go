// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeEUTRACGI(t *testing.T) {

	var plmnIDeutracgi = "ONF"

	eUTRACGI := &e2sm_rc_pre_v2.Eutracgi{
		PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
			Value: []byte(plmnIDeutracgi),
		},
		EUtracellIdentity: &e2sm_rc_pre_v2.EutracellIdentity{
			Value: &e2sm_rc_pre_v2.BitString{
				Value: 0x9bcd4ab, //uint64
				Len:   28,        //uint32
			},
		},
	}

	xer, err := xerEncodeEUTRACGI(eUTRACGI)
	assert.NilError(t, err)
	t.Logf("EUTRACGI XER\n%s", string(xer))
}
