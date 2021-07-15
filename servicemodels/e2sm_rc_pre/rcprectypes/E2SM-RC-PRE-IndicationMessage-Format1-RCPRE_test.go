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

func Test_XerEncodeE2SmRcPreIndicationMessageFormat1(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	arfcn := &e2sm_rc_pre_v2.Arfcn{
		Arfcn: &e2sm_rc_pre_v2.Arfcn_EArfcn{
			EArfcn: &e2sm_rc_pre_v2.Earfcn{
				Value: 253,
			},
		},
	}

	cellSize := e2sm_rc_pre_v2.CellSize_CELL_SIZE_MACRO
	pci := &e2sm_rc_pre_v2.Pci{
		Value: 11,
	}

	e2SmIindicationMsg := &e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_v2.E2SmRcPreIndicationMessageFormat1{
			DlArfcn:   arfcn,
			CellSize:  cellSize,
			Pci:       pci,
			Neighbors: make([]*e2sm_rc_pre_v2.Nrt, 0),
		},
	}

	neighbors := &e2sm_rc_pre_v2.Nrt{
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
			Arfcn: &e2sm_rc_pre_v2.Arfcn_NrArfcn{
				NrArfcn: &e2sm_rc_pre_v2.Nrarfcn{
					Value: 253,
				},
			},
		},
	}
	e2SmIindicationMsg.IndicationMessageFormat1.Neighbors = append(e2SmIindicationMsg.IndicationMessageFormat1.Neighbors, neighbors)

	xer, err := XerEncodeE2SmRcPreIndicationMessageFormat1(e2SmIindicationMsg)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationMessage-Format1 XER\n%s", string(xer))
}
