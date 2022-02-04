// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoIndicationHeader(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_mho.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}

	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmMhoPdu, err := CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	xer, err := mhoctypes.XerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := mhoctypes.PerEncodeE2SmMhoIndicationHeader(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: % x", per)

}
