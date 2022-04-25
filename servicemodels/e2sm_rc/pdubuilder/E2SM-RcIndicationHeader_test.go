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

func TestE2SmRcIndicationHeaderFormat1(t *testing.T) {

	msg, err := CreateE2SmRcIndicationHeaderFormat1()
	assert.NilError(t, err)
	msg.GetRicIndicationHeaderFormats().GetIndicationHeaderFormat1().SetRicEventTriggerConditionID(2)

	aper, err := encoder.PerEncodeE2SmRcIndicationHeader(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationHeader(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcIndicationHeaderFormat2(t *testing.T) {

	ueID, err := CreateUeIDNgEnbDu(123)
	assert.NilError(t, err)

	msg, err := CreateE2SmRcIndicationHeaderFormat2(ueID, 1, 1)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationHeader(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationHeader(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
