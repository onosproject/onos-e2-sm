// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeEUTRACellIdentity(t *testing.T) {

	eUTRACellIdentity := &e2sm_rc_pre_ies.EutracellIdentity{
		Value: &e2sm_rc_pre_ies.BitString{
			Value: 0x9bcd4ab, //uint64
			Len:   28,        //uint32
		},
	}

	xer, err := xerEncodeEUTRACellIdentity(eUTRACellIdentity)
	assert.NilError(t, err)
	t.Logf("EUTRACellIdentity XER\n%s", string(xer))

	per, err := perEncodeEUTRACellIdentity(eUTRACellIdentity)
	assert.NilError(t, err)
	t.Logf("EUTRACellIdentity XER\n%v", per)

}
