// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_xerEncodeE2SmKpmIndicationHeader(t *testing.T) {

	var plmnID = "ONF"
	var gNbCuUpID int64 = 10
	var gNbDuID int64 = 20
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpID, gNbDuID, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err, "Test_xerEncodeE2SmKpmIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmKpmPdu != nil, "Test_xerEncodeE2SmKpmIndicationHeader() newE2SmKpmPdu is nil")

	xer, err := XerEncodeE2SmKpmIndicationHeader(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader XER\n%s", string(xer))

	per, err := PerEncodeE2SmKpmIndicationHeader(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader PER\n%v", per)
}

func Test_perDecodeE2SmKpmIndicationHeader(t *testing.T) {
	e2smKpmRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-KPM-Indication-Header-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmKpmPdu, err := XerDecodeE2SmKpmIndicationHeader(e2smKpmRanfunctionDescription)
	assert.NilError(t, err)
	assert.NilError(t, e2SmKpmPdu.Validate())

}
