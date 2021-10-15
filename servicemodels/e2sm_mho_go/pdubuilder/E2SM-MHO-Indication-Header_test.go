// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	cellID := asn1.BitString{
		Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
		Len:   28, //uint32
	}

	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	newE2SmMhoPdu, err := CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationHeader: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoIndicationHeader(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-IndicationHeader is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
