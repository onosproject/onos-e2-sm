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

func createConstrainedChoice1Msg() (*test_sm_ies.ConstrainedChoice1, error) {

	// constrainedChoice1 := pdubuilder.CreateConstrainedChoice1() //ToDo - fill in arguments here(if this function exists

	constrainedChoice1 := test_sm_ies.ConstrainedChoice1{
		ConstrainedChoice1: &test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A{
			ConstrainedChoice1A: 32,
		},
	}

	//if err := constrainedChoice1.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating ConstrainedChoice1 %s", err.Error())
	//}
	return &constrainedChoice1, nil
}

func Test_xerEncodingConstrainedChoice1(t *testing.T) {

	constrainedChoice1, err := createConstrainedChoice1Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice1 PDU")

	xer, err := xerEncodeConstrainedChoice1(constrainedChoice1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("ConstrainedChoice1 XER\n%s", string(xer))

	result, err := xerDecodeConstrainedChoice1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice1 XER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice1.GetConstrainedChoice1A(), result.GetConstrainedChoice1A())
}

func Test_perEncodingConstrainedChoice1(t *testing.T) {

	constrainedChoice1, err := createConstrainedChoice1Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice1 PDU")

	per, err := perEncodeConstrainedChoice1(constrainedChoice1)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("ConstrainedChoice1 PER\n%v", hex.Dump(per))

	result, err := perDecodeConstrainedChoice1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice1 PER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice1.GetConstrainedChoice1A(), result.GetConstrainedChoice1A())
}
