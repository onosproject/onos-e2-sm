//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcControlHeaderFormat1(t *testing.T) {

	ueID, err := CreateUeIDGnbCuUp(15)
	assert.NilError(t, err)

	msg, err := CreateE2SmRcControlHeaderFormat1(ueID, 12, 5)
	assert.NilError(t, err)
	msg.GetRicControlHeaderFormats().GetControlHeaderFormat1().SetRicControlDecision(CreateRicControlDecisionAccept())

	aper, err := encoder.PerEncodeE2SmRcControlHeader(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlHeader(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcControlHeaderFormat2(t *testing.T) {

	ueID, err := CreateUeIDGnbCuUp(15)
	assert.NilError(t, err)

	msg, err := CreateE2SmRcControlHeaderFormat2()
	assert.NilError(t, err)
	msg.GetRicControlHeaderFormats().GetControlHeaderFormat2().SetUeID(ueID).SetRicControlDecision(CreateRicControlDecisionReject())

	aper, err := encoder.PerEncodeE2SmRcControlHeader(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlHeader(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
