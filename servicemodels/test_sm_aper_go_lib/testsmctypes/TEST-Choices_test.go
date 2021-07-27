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

func createTestChoicesMsg() (*test_sm_ies.TestChoices, error) {

	testChoices := test_sm_ies.TestChoices{
		OtherAttr: "nil",
		Choice1:   &test_sm_ies.Choice1{
			Choice1: &test_sm_ies.Choice1_Choice1A{
				Choice1A: 32,
			},
		},
		Choice2:   &test_sm_ies.Choice2{
			Choice2: &test_sm_ies.Choice2_Choice2B{
				Choice2B: -153,
			},
		},
		Choice3:   &test_sm_ies.Choice3{
			Choice3: &test_sm_ies.Choice3_Choice3C{
				Choice3C: 32,
			},
		},
		Choice4:   &test_sm_ies.Choice4{
			Choice4: &test_sm_ies.Choice4_Choice4A{
				Choice4A: 32,
			},
		},
	}

	return &testChoices, nil
}

func Test_xerEncodingTestChoices(t *testing.T) {

	testChoices, err := createTestChoicesMsg()
	assert.NilError(t, err, "Error creating TestChoices PDU")

	xer, err := xerEncodeTestChoices(testChoices)
	assert.NilError(t, err)
	t.Logf("TestChoices XER\n%s", string(xer))

	result, err := xerDecodeTestChoices(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestChoices XER - decoded\n%v", result)
	assert.Equal(t, testChoices.GetOtherAttr(), result.GetOtherAttr())
	assert.Equal(t, testChoices.GetChoice1().GetChoice1A(), result.GetChoice1().GetChoice1A())
	assert.Equal(t, testChoices.GetChoice2().GetChoice2B(), result.GetChoice2().GetChoice2B())
	assert.Equal(t, testChoices.GetChoice3().GetChoice3C(), result.GetChoice3().GetChoice3C())
	assert.Equal(t, testChoices.GetChoice4().GetChoice4A(), result.GetChoice4().GetChoice4A())
}

func Test_perEncodingTestChoices(t *testing.T) {

	testChoices, err := createTestChoicesMsg()
	assert.NilError(t, err, "Error creating TestChoices PDU")

	per, err := perEncodeTestChoices(testChoices)
	assert.NilError(t, err)
	t.Logf("TestChoices PER\n%v", hex.Dump(per))

	result, err := perDecodeTestChoices(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestChoices PER - decoded\n%v", result)
	assert.Equal(t, testChoices.GetOtherAttr(), result.GetOtherAttr())
	assert.Equal(t, testChoices.GetChoice1().GetChoice1A(), result.GetChoice1().GetChoice1A())
	assert.Equal(t, testChoices.GetChoice2().GetChoice2B(), result.GetChoice2().GetChoice2B())
	assert.Equal(t, testChoices.GetChoice3().GetChoice3C(), result.GetChoice3().GetChoice3C())
	assert.Equal(t, testChoices.GetChoice4().GetChoice4A(), result.GetChoice4().GetChoice4A())
}
