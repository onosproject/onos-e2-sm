// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}

	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationHeader(cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Header: \n%v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("PER decoded RC-PRE-IndicationHeader is \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}
