// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	"fmt"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/test_sm/pdubuilder"
	test_sm "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm" //ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createTestChoicesMsg() (*test_sm.TestChoices, error) {

	// testChoices := pdubuilder.CreateTestChoices() //ToDo - fill in arguments here(if this function exists

	testChoices := test_sm.TestChoices{
		OtherAttr: nil,
		Choice1:   nil,
		Choice2:   nil,
		Choice3:   nil,
		Choice4:   nil,
	}

	if err := testChoices.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestChoices %s", err.Error())
	}
	return &testChoices, nil
}

func Test_xerEncodingTestChoices(t *testing.T) {

	testChoices, err := createTestChoicesMsg()
	assert.NilError(t, err, "Error creating TestChoices PDU")

	xer, err := xerEncodeTestChoices(testChoices)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestChoices XER\n%s", string(xer))

	result, err := xerDecodeTestChoices(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestChoices XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testChoices.GetOtherAttr(), result.GetOtherAttr())
	assert.Equal(t, testChoices.GetChoice1(), result.GetChoice1())
	assert.Equal(t, testChoices.GetChoice2(), result.GetChoice2())
	assert.Equal(t, testChoices.GetChoice3(), result.GetChoice3())
	assert.Equal(t, testChoices.GetChoice4(), result.GetChoice4())

}

func Test_perEncodingTestChoices(t *testing.T) {

	testChoices, err := createTestChoicesMsg()
	assert.NilError(t, err, "Error creating TestChoices PDU")

	per, err := perEncodeTestChoices(testChoices)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestChoices PER\n%v", hex.Dump(per))

	result, err := perDecodeTestChoices(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestChoices PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testChoices.GetOtherAttr(), result.GetOtherAttr())
	assert.Equal(t, testChoices.GetChoice1(), result.GetChoice1())
	assert.Equal(t, testChoices.GetChoice2(), result.GetChoice2())
	assert.Equal(t, testChoices.GetChoice3(), result.GetChoice3())
	assert.Equal(t, testChoices.GetChoice4(), result.GetChoice4())

}
