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

func createChoice3Msg() (*test_sm_ies.Choice3, error) {

	choice3 := test_sm_ies.Choice3{
		//Choice3: &test_sm_ies.Choice3_Choice3A{
		//	Choice3A: 32,
		//},
		Choice3: &test_sm_ies.Choice3_Choice3B{
			Choice3B: 32,
		},
		//Choice3: &test_sm_ies.Choice3_Choice3C{
		//	Choice3C: 32,
		//},
	}

	return &choice3, nil
}

func Test_xerEncodingChoice3(t *testing.T) {

	choice3, err := createChoice3Msg()
	assert.NilError(t, err, "Error creating Choice3 PDU")

	xer, err := xerEncodeChoice3(choice3)
	assert.NilError(t, err)
	t.Logf("Choice3 XER\n%s", string(xer))

	result, err := xerDecodeChoice3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice3 XER - decoded\n%v", result)
	//assert.DeepEqual(t, choice3.GetChoice3A(), result.GetChoice3A())
	assert.DeepEqual(t, choice3.GetChoice3B(), result.GetChoice3B())
	//assert.DeepEqual(t, choice3.GetChoice3C(), result.GetChoice3C())
}

func Test_perEncodingChoice3(t *testing.T) {

	choice3, err := createChoice3Msg()
	assert.NilError(t, err, "Error creating Choice3 PDU")

	per, err := perEncodeChoice3(choice3)
	assert.NilError(t, err)
	t.Logf("Choice3 PER\n%v", hex.Dump(per))

	result, err := perDecodeChoice3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Choice3 PER - decoded\n%v", result)
	//assert.DeepEqual(t, choice3.GetChoice3A(), result.GetChoice3A())
	assert.DeepEqual(t, choice3.GetChoice3B(), result.GetChoice3B())
	//assert.DeepEqual(t, choice3.GetChoice3C(), result.GetChoice3C())
}
