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

func Test_xerEncodeNRT(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellSize := e2sm_rc_pre_v2.CellSize_CELL_SIZE_MACRO
	neighbor := &e2sm_rc_pre_v2.Nrt{
		Cgi: &e2sm_rc_pre_v2.CellGlobalId{
			CellGlobalId: &e2sm_rc_pre_v2.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_rc_pre_v2.Eutracgi{
					PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
						Value: plmnIDBytes,
					},
					EUtracellIdentity: &e2sm_rc_pre_v2.EutracellIdentity{
						Value: &e2sm_rc_pre_v2.BitString{
							Value: []byte{0xba, 0x4d, 0xcb, 0x90},
							Len:   28, //uint32
						},
					},
				},
			},
		},
		Pci: &e2sm_rc_pre_v2.Pci{
			Value: 11,
		},
		CellSize: cellSize,
		DlArfcn: &e2sm_rc_pre_v2.Arfcn{
			Arfcn: &e2sm_rc_pre_v2.Arfcn_EArfcn{
				EArfcn: &e2sm_rc_pre_v2.Earfcn{
					Value: 253,
				},
			},
		},
	}

	xer, err := xerEncodeNRT(neighbor)
	assert.NilError(t, err)
	t.Logf("NRT XER\n%s", string(xer))
}
