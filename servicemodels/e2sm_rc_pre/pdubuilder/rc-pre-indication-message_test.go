// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationMsg(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	var dlEarfcn int32 = 253
	var pci int32 = 11

	pciPool := &e2sm_rc_pre_ies.PciRange{
		LowerPci: &e2sm_rc_pre_ies.Pci{
			Value: 10,
		},
		UpperPci: &e2sm_rc_pre_ies.Pci{
			Value: 20,
		},
	}

	neighbors := &e2sm_rc_pre_ies.Nrt{
		NrIndex: 1,
		Cgi: &e2sm_rc_pre_ies.CellGlobalId{
			CellGlobalId: &e2sm_rc_pre_ies.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_rc_pre_ies.Eutracgi{
					PLmnIdentity: &e2sm_rc_pre_ies.PlmnIdentity{
						Value: plmnIDBytes,
					},
					EUtracellIdentity: &e2sm_rc_pre_ies.EutracellIdentity{
						Value: &e2sm_rc_pre_ies.BitString{
							Value: 0x9bcd4ac, //uint64
							Len:   28,        //uint32
						},
					},
				},
			},
		},
		Pci: &e2sm_rc_pre_ies.Pci{
			Value: 11,
		},
		CellSize: e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO,
		DlEarfcn: &e2sm_rc_pre_ies.Earfcn{
			Value: 253,
		},
	}

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationMsg(plmnIDBytes, &cellID, dlEarfcn, e2sm_rc_pre_ies.CellSize_CELL_SIZE_MACRO, pci, pciPool, neighbors)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Message: % x", per)

}
