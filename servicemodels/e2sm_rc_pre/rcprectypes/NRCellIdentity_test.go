// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRCellIdentity(t *testing.T) {

	nrCellIdentity := &e2sm_rc_pre_ies.NrcellIdentity{
		Value: &e2sm_rc_pre_ies.BitString{
			Value: 0x9bcd4abef, //uint64
			Len:   36,          //uint32
		},
	}

	xer, err := xerEncodeNRCellIdentity(nrCellIdentity)
	assert.NilError(t, err)
	t.Logf("NRCellIdentity XER\n%s", string(xer))

	per, err := perEncodeNRCellIdentity(nrCellIdentity)
	assert.NilError(t, err)
	t.Logf("NRCellIdentity XER\n%v", per)

}
