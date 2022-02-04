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

func Test_xerEncodeEUTRACellIdentity(t *testing.T) {

	eUTRACellIdentity := &e2sm_rc_pre_v2.EutracellIdentity{
		Value: &e2sm_rc_pre_v2.BitString{
			Value: []byte{0xba, 0x4d, 0xcb, 0x90},
			Len:   28, //uint32
		},
	}

	xer, err := xerEncodeEUTRACellIdentity(eUTRACellIdentity)
	assert.NilError(t, err)
	t.Logf("EUTRACellIdentity XER\n%s", string(xer))

	per, err := perEncodeEUTRACellIdentity(eUTRACellIdentity)
	assert.NilError(t, err)
	t.Logf("EUTRACellIdentity PER\n%v", hex.Dump(per))

}
