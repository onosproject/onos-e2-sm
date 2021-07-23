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

func createTestUnconstrainedIntMsg() (*test_sm.TestUnconstrainedInt, error) {

	// testUnconstrainedInt := pdubuilder.CreateTestUnconstrainedInt() //ToDo - fill in arguments here(if this function exists

	testUnconstrainedInt := test_sm.TestUnconstrainedInt{
		AttrUciA: nil,
		AttrUciB: nil,
	}

	if err := testUnconstrainedInt.Validate(); err != nil {
		return nil, fmt.Errorf("error validating TestUnconstrainedInt %s", err.Error())
	}
	return &testUnconstrainedInt, nil
}

func Test_xerEncodingTestUnconstrainedInt(t *testing.T) {

	testUnconstrainedInt, err := createTestUnconstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedInt PDU")

	xer, err := xerEncodeTestUnconstrainedInt(testUnconstrainedInt)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("TestUnconstrainedInt XER\n%s", string(xer))

	result, err := xerDecodeTestUnconstrainedInt(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedInt XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testUnconstrainedInt.GetAttrUciA(), result.GetAttrUciA())
	assert.Equal(t, testUnconstrainedInt.GetAttrUciB(), result.GetAttrUciB())

}

func Test_perEncodingTestUnconstrainedInt(t *testing.T) {

	testUnconstrainedInt, err := createTestUnconstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedInt PDU")

	per, err := perEncodeTestUnconstrainedInt(testUnconstrainedInt)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("TestUnconstrainedInt PER\n%v", hex.Dump(per))

	result, err := perDecodeTestUnconstrainedInt(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedInt PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, testUnconstrainedInt.GetAttrUciA(), result.GetAttrUciA())
	assert.Equal(t, testUnconstrainedInt.GetAttrUciB(), result.GetAttrUciB())

}
