// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationHeader(t *testing.T) {
	var plmnID = "ONF"
	var gNbCuUpID int64 = 12
	var gNbDuID int64 = 13
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpID, gNbDuID, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)

	xer, err := kpmctypes.XerEncodeE2SmKpmIndicationHeader(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Header: %s", string(xer))

	per, err := kpmctypes.PerEncodeE2SmKpmIndicationHeader(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Header: % x", per)

}
