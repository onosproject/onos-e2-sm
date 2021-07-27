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

func createChoice2Msg() (*test_sm_ies.Choice2, error) {

	choice2 := test_sm_ies.Choice2{
		Choice2: &test_sm_ies.Choice2_Choice2A{
			Choice2A: -153,
		},
		//Choice2: &test_sm.Choice2_Choice2B{
		//	Choice2B: -153,
		//},
	}

	return &choice2, nil
}

func Test_xerEncodingChoice2(t *testing.T) {

	choice2, err := createChoice2Msg()
	assert.NilError(t, err, "Error creating Choice2 PDU")

	xer, err := xerEncodeChoice2(choice2)
	assert.NilError(t, err)
	t.Logf("Choice2 XER\n%s", string(xer))

	result, err := xerDecodeChoice2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice2 XER - decoded\n%v", result)
	assert.DeepEqual(t, choice2.GetChoice2A(), result.GetChoice2A())
	//assert.DeepEqual(t, choice2.GetChoice2B(), result.GetChoice2B())
}

func Test_perEncodingChoice2(t *testing.T) {

	choice2, err := createChoice2Msg()
	assert.NilError(t, err, "Error creating Choice2 PDU")

	per, err := perEncodeChoice2(choice2)
	assert.NilError(t, err)
	t.Logf("Choice2 PER\n%v", hex.Dump(per))

	result, err := perDecodeChoice2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice2 PER - decoded\n%v", result)
	assert.DeepEqual(t, choice2.GetChoice2A(), result.GetChoice2A())
	//assert.DeepEqual(t, choice2.GetChoice2B(), result.GetChoice2B())
}
