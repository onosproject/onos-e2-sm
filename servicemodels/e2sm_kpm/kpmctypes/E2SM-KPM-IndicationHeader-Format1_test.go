// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmIndicationHeaderFormat1(t *testing.T) {

	var plmnID = "ONF"
	var gNbCuUpID int64 = 0
	var gNbDuID int64 = 0
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpID, gNbDuID, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err, "Test_xerEncodeE2SmKpmIndicationHeaderFormat1() error is not nil")
	assert.Assert(t, newE2SmKpmPdu != nil, "Test_xerEncodeE2SmKpmIndicationHeaderFormat1() newE2SmKpmPdu is nil")

	xer, err := xerEncodeE2SmKpmIndicationHeaderFormat1(newE2SmKpmPdu.GetIndicationHeaderFormat1())
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader-Format1 XER\n%s", string(xer))
}
