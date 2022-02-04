// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createUeIdentityMsg(ueIdentityValue string) (*e2sm_mho.UeIdentity, error) {

	// ueIdentity := pdubuilder.CreateUeIdentity() //ToDo - fill in arguments here(if this function exists

	ueIdentity := e2sm_mho.UeIdentity{
		Value: ueIdentityValue,
	}

	if err := ueIdentity.Validate(); err != nil {
		return nil, fmt.Errorf("error validating UeIdentity %s", err.Error())
	}
	return &ueIdentity, nil
}

func Test_xerEncodingUeIdentity(t *testing.T) {

	ueIdentity, err := createUeIdentityMsg("1234")
	assert.NilError(t, err, "Error creating UeIdentity PDU")

	xer, err := xerEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	assert.Equal(t, 39, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("UeIdentity XER\n%s", string(xer))

	result, err := xerDecodeUeIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("UeIdentity XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, ueIdentity.GetValue(), result.GetValue())

}

func Test_perEncodingUeIdentity(t *testing.T) {

	ueIdentity, err := createUeIdentityMsg("1234")
	assert.NilError(t, err, "Error creating UeIdentity PDU")

	per, err := perEncodeUeIdentity(ueIdentity)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("UeIdentity PER\n%v", hex.Dump(per))

	result, err := perDecodeUeIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("UeIdentity PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, ueIdentity.GetValue(), result.GetValue())

}
