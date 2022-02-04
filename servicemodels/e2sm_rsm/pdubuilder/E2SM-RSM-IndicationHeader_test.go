// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRsmIndicationHeaderNrCellID(t *testing.T) {

	plmnID := []byte{0x00, 0x01, 0x0F}
	nrCellID := &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x00, 0x10},
		Len:   36,
	}

	cgi, err := CreateNrCGI(plmnID, nrCellID)
	assert.NilError(t, err)

	ih := CreateE2SmRsmIndicationHeaderFormat1(cgi)
	t.Logf("Created E2SM-RSM-IndicationHeader is \n%v", ih)

	// APER validation
	per, err := encoder.PerEncodeE2SmRsmIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmIndicationHeader(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-IndicationHeader PER - decoded\n%v", result)
	assert.DeepEqual(t, ih.String(), result.String())
}

func Test_E2SmRsmIndicationHeaderEutraCellID(t *testing.T) {

	plmnID := []byte{0x00, 0x01, 0x0F}
	eutraCellID := &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x10},
		Len:   28,
	}

	cgi, err := CreateEutraCGI(plmnID, eutraCellID)
	assert.NilError(t, err)

	ih := CreateE2SmRsmIndicationHeaderFormat1(cgi)
	t.Logf("Created E2SM-RSM-IndicationHeader is \n%v", ih)

	// APER validation
	per, err := encoder.PerEncodeE2SmRsmIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmIndicationHeader(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-IndicationHeader PER - decoded\n%v", result)
	assert.DeepEqual(t, ih.String(), result.String())
}
