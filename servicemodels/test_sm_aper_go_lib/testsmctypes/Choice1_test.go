// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createChoice1Msg() (*test_sm_ies.Choice1, error) {

	choice1 := test_sm_ies.Choice1{
		Choice1: &test_sm_ies.Choice1_Choice1A{
			Choice1A: 32,
		},
	}

	return &choice1, nil
}

func Test_xerEncodingChoice1(t *testing.T) {

	choice1, err := createChoice1Msg()
	assert.NilError(t, err, "Error creating Choice1 PDU")

	xer, err := xerEncodeChoice1(choice1)
	assert.NilError(t, err)
	t.Logf("Choice1 XER\n%s", string(xer))

	result, err := xerDecodeChoice1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice1 XER - decoded\n%v", result)
	assert.Equal(t, choice1.GetChoice1A(), result.GetChoice1A())
}

func Test_perEncodingChoice1(t *testing.T) {

	choice1, err := createChoice1Msg()
	assert.NilError(t, err, "Error creating Choice1 PDU")

	per, err := perEncodeChoice1(choice1)
	assert.NilError(t, err)
	t.Logf("Choice1 PER\n%v", hex.Dump(per))

	result, err := perDecodeChoice1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice1 PER - decoded\n%v", result)
	assert.Equal(t, choice1.GetChoice1A(), result.GetChoice1A())
}
