// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_XerEncodeE2SmRcPreIndicationHeader(t *testing.T) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	cellID := e2sm_rc_pre_ies.BitString{
		Value: 0x9bcd4ab, //uint64
		Len:   28,        //uint32
	}
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreIndicationHeader(plmnIDBytes, &cellID)
	assert.NilError(t, err, "Test_XerEncodeE2SmRcPreIndicationHeader)_ error is not nil")
	assert.Assert(t, newE2SmRcPrePdu != nil, "Test_XerEncodeE2SmRcPreIndicationHeader() newE2SmRcPrePdu is nil")

	xer, err := XerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader XER\n%s", string(xer))

	per, err := PerEncodeE2SmRcPreIndicationHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-IndicationHeader PER\n%v", per)
}

func Test_PerDecodeE2SmRcPreIndicationHeader(t *testing.T) {
	E2SmRcPreRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-RC-PRE-Indication-Header-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	E2SmRcPrePdu, err := XerDecodeE2SmRcPreIndicationHeader(E2SmRcPreRanfunctionDescription)
	assert.NilError(t, err)
	assert.NilError(t, E2SmRcPrePdu.Validate())

}
