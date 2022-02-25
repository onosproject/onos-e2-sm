// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoControlMsg(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	servingCgi := &e2smv2.Cgi{
		Cgi: &e2smv2.Cgi_EUtraCgi{
			EUtraCgi: &e2smv2.EutraCgi{
				PLmnidentity: &e2smv2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2smv2.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
						Len:   28, //uint32
					},
				},
			},
		},
	}

	ueID, err := CreateUeIDGNb(1, []byte{0x01, 0x02, 0x03}, []byte{0xFF}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)

	targetCgi := &e2smv2.Cgi{
		Cgi: &e2smv2.Cgi_EUtraCgi{
			EUtraCgi: &e2smv2.EutraCgi{
				PLmnidentity: &e2smv2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2smv2.EutracellIdentity{
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
