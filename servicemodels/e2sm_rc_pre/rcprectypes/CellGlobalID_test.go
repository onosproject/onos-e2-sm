// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeCellGlobalID(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellGlobalID := &e2sm_rc_pre_ies.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_ies.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_rc_pre_ies.Eutracgi{
				PLmnIdentity: &e2sm_rc_pre_ies.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_rc_pre_ies.EutracellIdentity{
					Value: &e2sm_rc_pre_ies.BitString{
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
