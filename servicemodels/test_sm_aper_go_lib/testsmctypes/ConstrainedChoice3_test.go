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

func createConstrainedChoice3Msg() (*test_sm_ies.ConstrainedChoice3, error) {

	// constrainedChoice3 := pdubuilder.CreateConstrainedChoice3() //ToDo - fill in arguments here(if this function exists

	constrainedChoice3 := test_sm_ies.ConstrainedChoice3{
		ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A{
			ConstrainedChoice3A: 32,
		},
		//ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3B{
		//	ConstrainedChoice3B: 15,
		//},
		//ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C{
		//	ConstrainedChoice3C: 1,
		//},
		//ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D{
		//	ConstrainedChoice3D: 1,
		//},
	}

	//if err := constrainedChoice3.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating ConstrainedChoice3 %s", err.Error())
	//}
	return &constrainedChoice3, nil
}

func Test_xerEncodingConstrainedChoice3(t *testing.T) {

	constrainedChoice3, err := createConstrainedChoice3Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice3 PDU")

	xer, err := xerEncodeConstrainedChoice3(constrainedChoice3)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice3 XER\n%s", string(xer))

	result, err := xerDecodeConstrainedChoice3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice3 XER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice3.GetConstrainedChoice3A(), result.GetConstrainedChoice3A())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3B(), result.GetConstrainedChoice3B())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3C(), result.GetConstrainedChoice3C())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3D(), result.GetConstrainedChoice3D())
}

func Test_perEncodingConstrainedChoice3(t *testing.T) {

	constrainedChoice3, err := createConstrainedChoice3Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice3 PDU")

	per, err := perEncodeConstrainedChoice3(constrainedChoice3)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice3 PER\n%v", hex.Dump(per))

	result, err := perDecodeConstrainedChoice3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice3 PER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice3.GetConstrainedChoice3A(), result.GetConstrainedChoice3A())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3B(), result.GetConstrainedChoice3B())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3C(), result.GetConstrainedChoice3C())
	//assert.Equal(t, constrainedChoice3.GetConstrainedChoice3D(), result.GetConstrainedChoice3D())
}
