// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

func CreateFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func createTestUnconstrainedRealMsg() (*test_sm_ies.TestUnconstrainedReal, error) {

	testUnconstrainedReal := test_sm_ies.TestUnconstrainedReal{
		AttrUcrA: 21.7,
		AttrUcrB: -653.43,
	}

	return &testUnconstrainedReal, nil
}

func Test_xerEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	xer, err := xerEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedReal XER\n%s", string(xer))

	result, err := xerDecodeTestUnconstrainedReal(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal XER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())
}

func Test_perEncodingTestUnconstrainedReal(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	per, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedReal PER\n%v", hex.Dump(per))

	result, err := perDecodeTestUnconstrainedReal(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestUnconstrainedReal PER - decoded\n%v", result)
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result.GetAttrUcrB())

	// treating special case when REAL is equal to 0
	testUnconstrainedReal.AttrUcrA = 100.0
	testUnconstrainedReal.AttrUcrB = 65.0
	per1, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestUnconstrainedReal PER\n%v", hex.Dump(per1))

	result1, err := perDecodeTestUnconstrainedReal(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestUnconstrainedReal PER - decoded\n%v", result1)
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrA(), result1.GetAttrUcrA())
	assert.Equal(t, testUnconstrainedReal.GetAttrUcrB(), result1.GetAttrUcrB())
}

func Test_perEncoding(t *testing.T) {

	testUnconstrainedReal, err := createTestUnconstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestUnconstrainedReal PDU")

	for i := 1; i < 154; i++ {

		testUnconstrainedReal.AttrUcrA = float64(i)
		testUnconstrainedReal.AttrUcrB = float64(i) + 0.5

		per, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
		assert.NilError(t, err)
		t.Logf("TestUnconstrainedReal PER\n%v", hex.Dump(per))

		result, err := perDecodeTestUnconstrainedReal(per)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		t.Logf("TestUnconstrainedReal PER - decoded\n%v", result)
	}
}

func TestGoAperLibUnconstrainedReal(t *testing.T) {

	var min int32 = -1000
	var max int32 = 1000

	for i := min; i < max; i++ {
		testUnconstrainedReal := &test_sm_ies.TestUnconstrainedReal{
			AttrUcrA: CreateFloat64(float64(min), float64(max)),
			AttrUcrB: CreateFloat64(float64(min), float64(max)),
		}

		t.Logf("Encoding following message:\n%v", testUnconstrainedReal)

		perRef, err := perEncodeTestUnconstrainedReal(testUnconstrainedReal)
		assert.NilError(t, err)
		t.Logf("TestUnconstrainedReal APER\n%v", hex.Dump(perRef))

		per, err := aper.Marshal(testUnconstrainedReal, nil, nil)
		assert.NilError(t, err)
		t.Logf("TestUnconstrainedReal with Go APER library produces following APER\n%v", hex.Dump(per))
		//assert.DeepEqual(t, perRef, per)

		// looks like we can perfectly decode bytes both ways (CGo bytes with Go APER, and Go APER bytes with CGo)
		resultRef, err := perDecodeTestUnconstrainedReal(per)
		assert.NilError(t, err)
		assert.Assert(t, resultRef != nil)
		t.Logf("TestUnconstrainedReal PER - decoded\n%v", resultRef)

		result := &test_sm_ies.TestUnconstrainedReal{}
		err = aper.Unmarshal(perRef, result, nil, nil)
		assert.NilError(t, err)
		t.Logf("TestUnconstrainedReal PER - decoded with Go APER library\n%v", resultRef)
		assert.Equal(t, resultRef.String(), result.String())
	}
}
