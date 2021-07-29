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

func createConstrainedChoice4Msg() (*test_sm_ies.ConstrainedChoice4, error) {

	constrainedChoice4 := test_sm_ies.ConstrainedChoice4{
		ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
			ConstrainedChoice4A: 32,
		},
	}

	return &constrainedChoice4, nil
}

func Test_xerEncodingConstrainedChoice4(t *testing.T) {

	constrainedChoice4, err := createConstrainedChoice4Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice4 PDU")

	xer, err := xerEncodeConstrainedChoice4(constrainedChoice4)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice4 XER\n%s", string(xer))

	result, err := xerDecodeConstrainedChoice4(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice4 XER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice4.GetConstrainedChoice4A(), result.GetConstrainedChoice4A())
}

func Test_perEncodingConstrainedChoice4(t *testing.T) {

	constrainedChoice4, err := createConstrainedChoice4Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice4 PDU")

	per, err := perEncodeConstrainedChoice4(constrainedChoice4)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice4 PER\n%v", hex.Dump(per))

	result, err := perDecodeConstrainedChoice4(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice4 PER - decoded\n%v", result)
	assert.Equal(t, constrainedChoice4.GetConstrainedChoice4A(), result.GetConstrainedChoice4A())
}
