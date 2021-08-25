// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
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

	//ToDo - embed APER validation
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

	//ToDo - embed APER validation
}
