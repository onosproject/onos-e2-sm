// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreIndicationMessageFormat1(t *testing.T) {

	var plmnID = "ONF"

	cgi := &e2sm_rc_pre_ies.CellGlobalId{
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
	}
	earfcn := &e2sm_rc_pre_ies.Earfcn{
		Value: 253,
	}

	cellSize := e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO
	pci := &e2sm_rc_pre_ies.Pci{
		Value: 11,
	}

	e2SmIindicationMsg := &e2sm_rc_pre_ies.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_ies.E2SmRcPreIndicationMessageFormat1{
			Cgi:       cgi,
			DlEarfcn:  earfcn,
			CellSize:  cellSize,
			PciPool:   make([]*e2sm_rc_pre_ies.PciRange, 0),
			Pci:       pci,
			Neighbors: make([]*e2sm_rc_pre_ies.Nrt, 0),
		},
	}

	pciPool := &e2sm_rc_pre_ies.PciRange{
		LowerPci: &e2sm_rc_pre_ies.Pci{
			Value: 10,
		},
		UpperPci: &e2sm_rc_pre_ies.Pci{
			Value: 20,
		},
	}
	e2SmIindicationMsg.IndicationMessageFormat1.PciPool = append(e2SmIindicationMsg.IndicationMessageFormat1.PciPool, pciPool)

	neighbors := &e2sm_rc_pre_ies.Nrt{
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
	e2SmIindicationMsg.IndicationMessageFormat1.Neighbors = append(e2SmIindicationMsg.IndicationMessageFormat1.Neighbors, neighbors)

	xer, err := XerEncodeE2SmRcPreIndicationMessageFormat1(e2SmIindicationMsg)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationMessage-Format1 XER\n%s", string(xer))
}
