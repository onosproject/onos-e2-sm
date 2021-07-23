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

func createTestConstrainedChoicesMsg() (*test_sm.TestConstrainedChoices, error) {

	// testConstrainedChoices := pdubuilder.CreateTestConstrainedChoices() //ToDo - fill in arguments here(if this function exists

	testConstrainedChoices := test_sm.TestConstrainedChoices{
		OtherCattr:         nil,
		ConstrainedChoice1: nil,
		ConstrainedChoice2: nil,
		ConstrainedChoice3: nil,
		ConstrainedChoice4: nil,
	}

	if err := testConstrainedChoices.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestConstrainedChoices %s", err.Error())
	}
	return &testConstrainedChoices, nil
}

func Test_xerEncodingTestConstrainedChoices(t *testing.T) {

	testConstrainedChoices, err := createTestConstrainedChoicesMsg()
	assert.NilError(t, err, "Error creating TestConstrainedChoices PDU")

	xer, err := xerEncodeTestConstrainedChoices(testConstrainedChoices)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestConstrainedChoices XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedChoices(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedChoices XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testConstrainedChoices.GetOtherCattr(), result.GetOtherCattr())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice1(), result.GetConstrainedChoice1())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice2(), result.GetConstrainedChoice2())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice3(), result.GetConstrainedChoice3())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice4(), result.GetConstrainedChoice4())

}

func Test_perEncodingTestConstrainedChoices(t *testing.T) {

	testConstrainedChoices, err := createTestConstrainedChoicesMsg()
	assert.NilError(t, err, "Error creating TestConstrainedChoices PDU")

	per, err := perEncodeTestConstrainedChoices(testConstrainedChoices)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestConstrainedChoices PER\n%v", hex.Dump(per))

	result, err := perDecodeTestConstrainedChoices(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedChoices PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testConstrainedChoices.GetOtherCattr(), result.GetOtherCattr())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice1(), result.GetConstrainedChoice1())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice2(), result.GetConstrainedChoice2())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice3(), result.GetConstrainedChoice3())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice4(), result.GetConstrainedChoice4())

}
