// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmMhoIndicationHeaderFormat1(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	cellID := e2sm_mho.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err, "Test_XerEncodeE2SmMhoIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmMhoPdu != nil, "Test_XerEncodeE2SmMhoIndicationHeaderFormat1() newE2SmMhoPdu is nil")

	xer, err := XerEncodeE2SmMhoIndicationHeaderFormat1(newE2SmMhoPdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-IndicationHeader-Format1 XER\n%s", string(xer))
}
