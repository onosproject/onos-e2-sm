// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoIndicationMsgF1(t *testing.T) {

	ueID := &e2sm_mho_go.UeIdentity{
		Value: []byte("1234"),
	}
	rsrp := &e2sm_mho_go.Rsrp{
		Value: 1234,
	}
	newE2SmMhoPdu, err := CreateE2SmMhoIndicationMsgFormat1(ueID, rsrp)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)
	t.Logf("E2SM-MHO-IndicationMessage is \n%v", newE2SmMhoPdu)

	per, err := encoder.PerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationMessage: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-IndicationMessage is \n%v", result)
	assert.DeepEqual(t, newE2SmMhoPdu.String(), result.String())
}

func TestE2SmMhoIndicationMsgF2(t *testing.T) {

	ueID := &e2sm_mho_go.UeIdentity{
		Value: []byte("1234"),
	}

	newE2SmMhoPdu, err := CreateE2SmMhoIndicationMsgFormat2(ueID, CreateRrcStatusConnected())
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationMessage: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-IndicationMessage is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
