// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createChoice4Msg() (*test_sm_ies.Choice4, error) {

	// choice4 := pdubuilder.CreateChoice4() //ToDo - fill in arguments here(if this function exists

	choice4 := test_sm_ies.Choice4{
		Choice4: &test_sm_ies.Choice4_Choice4A{
			Choice4A: 32,
		},
	}

	//if err := choice4.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating Choice4 %s", err.Error())
	//}
	return &choice4, nil
}

func Test_xerEncodingChoice4(t *testing.T) {

	choice4, err := createChoice4Msg()
	assert.NilError(t, err, "Error creating Choice4 PDU")

	xer, err := xerEncodeChoice4(choice4)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("Choice4 XER\n%s", string(xer))

	result, err := xerDecodeChoice4(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice4 XER - decoded\n%v", result)
	assert.DeepEqual(t, choice4.GetChoice4A(), result.GetChoice4A())
}

func Test_perEncodingChoice4(t *testing.T) {

	choice4, err := createChoice4Msg()
	assert.NilError(t, err, "Error creating Choice4 PDU")

	per, err := perEncodeChoice4(choice4)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("Choice4 PER\n%v", hex.Dump(per))

	result, err := perDecodeChoice4(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice4 PER - decoded\n%v", result)
	assert.DeepEqual(t, choice4.GetChoice4A(), result.GetChoice4A())
}
