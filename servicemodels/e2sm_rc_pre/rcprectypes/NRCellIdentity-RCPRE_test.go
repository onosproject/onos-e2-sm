// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRCellIdentity(t *testing.T) {

	nrCellIdentity := &e2sm_rc_pre_v2.NrcellIdentity{
		Value: &e2sm_rc_pre_v2.BitString{
			Value: []byte{0xef, 0xab, 0xd4, 0xbc, 0x90},
			Len:   36, //uint32
		},
	}

	xer, err := xerEncodeNRCellIdentity(nrCellIdentity)
	assert.NilError(t, err)
	t.Logf("NRCellIdentity XER\n%s", string(xer))

	per, err := perEncodeNRCellIdentity(nrCellIdentity)
	assert.NilError(t, err)
	t.Logf("NRCellIdentity PER\n%v", hex.Dump(per))
}
