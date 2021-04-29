// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoControlMsg(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	servingCgi := &e2sm_mho.CellGlobalId{
		CellGlobalId: &e2sm_mho.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho.Eutracgi{
				PLmnIdentity: &e2sm_mho.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho.EutracellIdentity{
					Value: &e2sm_mho.BitString{
						Value: 0x9bcd4ab, //uint64
						Len:   28,        //uint32
					},
				},
			},
		},
	}

	ueId := &e2sm_mho.UeIdentity{
		Value: "1234",
	}

	targetCgi := &e2sm_mho.CellGlobalId{
		CellGlobalId: &e2sm_mho.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho.Eutracgi{
				PLmnIdentity: &e2sm_mho.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho.EutracellIdentity{
					Value: &e2sm_mho.BitString{
						Value: 0x9bcd4ab, //uint64
						Len:   28,        //uint32
					},
				},
			},
		},
	}
	newE2SmMhoPdu, err := CreateE2SmMhoControlMessage(servingCgi, ueId, targetCgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	xer, err := mhoctypes.XerEncodeE2SmMhoControlMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := mhoctypes.PerEncodeE2SmMhoControlMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Message: % x", per)
}
