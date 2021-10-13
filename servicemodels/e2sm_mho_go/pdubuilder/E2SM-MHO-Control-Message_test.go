// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoControlMsg(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	servingCgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
						Len:   28, //uint32
					},
				},
			},
		},
	}

	ueID := &e2sm_mho_go.UeIdentity{
		Value: []byte("1234"),
	}

	targetCgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
						Len:   28, //uint32
					},
				},
			},
		},
	}
	newE2SmMhoPdu, err := CreateE2SmMhoControlMessage(servingCgi, ueID, targetCgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoControlMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationMessage: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoControlMessage(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-ControlMessage is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
