// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" // ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoIndicationMessageFormat2Msg() (*e2sm_mho.E2SmMhoIndicationMessageFormat2, error) {

	// e2SmMhoIndicationMessageFormat2 := pdubuilder.CreateE2SmMhoIndicationMessageFormat2() // ToDo - fill in arguments here(if this function exists

	e2SmMhoIndicationMessageFormat2 := e2sm_mho.E2SmMhoIndicationMessageFormat2{
		UeId:      &e2sm_mho.UeIdentity{Value: "1234"},
		RrcStatus: e2sm_mho.Rrcstatus_RRCSTATUS_IDLE,
	}

	if err := e2SmMhoIndicationMessageFormat2.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoIndicationMessageFormat2 %s", err.Error())
	}
	return &e2SmMhoIndicationMessageFormat2, nil
}

func Test_xerEncodingE2SmMhoIndicationMessageFormat2(t *testing.T) {

	e2SmMhoIndicationMessageFormat2, err := createE2SmMhoIndicationMessageFormat2Msg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessageFormat2 PDU")

	xer, err := xerEncodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2)
	assert.NilError(t, err)
	assert.Equal(t, 139, len(xer)) // ToDo - adjust length of the XER encoded message
	t.Logf("E2SmMhoIndicationMessageFormat2 XER\n%s", string(xer))

	result, err := xerDecodeE2SmMhoIndicationMessageFormat2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessageFormat2 XER - decoded\n%v", result)
	// ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoIndicationMessageFormat2.GetUeId(), result.GetUeId())
	//assert.Equal(t, e2SmMhoIndicationMessageFormat2.GetRrcStatus(), result.GetRrcStatus())

}

func Test_perEncodingE2SmMhoIndicationMessageFormat2(t *testing.T) {

	e2SmMhoIndicationMessageFormat2, err := createE2SmMhoIndicationMessageFormat2Msg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessageFormat2 PDU")

	per, err := perEncodeE2SmMhoIndicationMessageFormat2(e2SmMhoIndicationMessageFormat2)
	assert.NilError(t, err)
	assert.Equal(t, 7, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2SmMhoIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmMhoIndicationMessageFormat2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessageFormat2 PER - decoded\n%v", result)
	// ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoIndicationMessageFormat2.GetUeId(), result.GetUeId())
	//assert.Equal(t, e2SmMhoIndicationMessageFormat2.GetRrcStatus(), result.GetRrcStatus())

}
