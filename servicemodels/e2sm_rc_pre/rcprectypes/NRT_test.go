// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRT(t *testing.T) {
	var plmnID = "ONF"
	cellSize := e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO
	neighbor := &e2sm_rc_pre_ies.Nrt{
		NrIndex: 1,
		Cgi: &e2sm_rc_pre_ies.CellGlobalId{
			CellGlobalId: &e2sm_rc_pre_ies.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_rc_pre_ies.Eutracgi{
					PLmnIdentity: &e2sm_rc_pre_ies.PlmnIdentity{
						Value: []byte(plmnID),
					},
					EUtracellIdentity: &e2sm_rc_pre_ies.EutracellIdentity{
						Value: &e2sm_rc_pre_ies.BitString{
							Value: 0x9bcd4ab, //uint64
							Len:   28,        //uint32
						},
					},
				},
			},
		},
		Pci: &e2sm_rc_pre_ies.Pci{
			Value: 11,
		},
		CellSize: cellSize,
		DlEarfcn: &e2sm_rc_pre_ies.Earfcn{
			Value: 253,
		},
	}

	xer, err := xerEncodeNRT(neighbor)
	assert.NilError(t, err)
	t.Logf("NRT XER\n%s", string(xer))
}
