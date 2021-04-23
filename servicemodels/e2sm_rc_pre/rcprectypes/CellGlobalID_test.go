// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeCellGlobalID(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellGlobalID := &e2sm_rc_pre_v2.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_v2.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_rc_pre_v2.Eutracgi{
				PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_rc_pre_v2.EutracellIdentity{
					Value: &e2sm_rc_pre_v2.BitString{
						Value: 0x9bcd4ab, //uint64
						Len:   28,        //uint32
					},
				},
			},
			//TODO: Implement other types
		},
	}

	xer, err := xerEncodeCellGlobalID(cellGlobalID)
	assert.NilError(t, err)
	t.Logf("CellGlobalID XER\n%s", string(xer))
}
