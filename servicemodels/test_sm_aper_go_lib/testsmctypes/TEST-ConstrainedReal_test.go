// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func createTestConstrainedRealMsg() (*test_sm_ies.TestConstrainedReal, error) {

	testConstrainedReal := test_sm_ies.TestConstrainedReal{
		AttrCrA: 99.97,
		AttrCrB: 11.20,
		AttrCrC: -153.0,
		AttrCrD: 20.0,
		AttrCrE: 10.0,
		AttrCrF: 10.0,
	}

	return &testConstrainedReal, nil
}

func Test_xerEncodingTestConstrainedReal(t *testing.T) {

	testConstrainedReal, err := createTestConstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestConstrainedReal PDU")

	xer, err := xerEncodeTestConstrainedReal(testConstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestConstrainedReal XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedReal(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedReal XER - decoded\n%v", result)
	assert.Equal(t, testConstrainedReal.GetAttrCrA(), result.GetAttrCrA())
	assert.Equal(t, testConstrainedReal.GetAttrCrB(), result.GetAttrCrB())
	assert.Equal(t, testConstrainedReal.GetAttrCrC(), result.GetAttrCrC())
	assert.Equal(t, testConstrainedReal.GetAttrCrD(), result.GetAttrCrD())
	assert.Equal(t, testConstrainedReal.GetAttrCrE(), result.GetAttrCrE())
	assert.Equal(t, testConstrainedReal.GetAttrCrF(), result.GetAttrCrF())
}

func Test_perEncodingTestConstrainedReal(t *testing.T) {

	testConstrainedReal, err := createTestConstrainedRealMsg()
	assert.NilError(t, err, "Error creating TestConstrainedReal PDU")

	per, err := perEncodeTestConstrainedReal(testConstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestConstrainedReal PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	perNew, err := aper.Marshal(testConstrainedReal, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)
	t.Logf("TestConstrainedReal PER (with Go APER library)\n%v", hex.Dump(perNew))

	//Comparing bytes against each other
	assert.DeepEqual(t, per, perNew)

	result, err := perDecodeTestConstrainedReal(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedReal PER - decoded\n%v", result)
	assert.Equal(t, testConstrainedReal.GetAttrCrA(), result.GetAttrCrA())
	assert.Equal(t, testConstrainedReal.GetAttrCrB(), result.GetAttrCrB())
	assert.Equal(t, testConstrainedReal.GetAttrCrC(), result.GetAttrCrC())
	assert.Equal(t, testConstrainedReal.GetAttrCrD(), result.GetAttrCrD())
	assert.Equal(t, testConstrainedReal.GetAttrCrE(), result.GetAttrCrE())
	assert.Equal(t, testConstrainedReal.GetAttrCrF(), result.GetAttrCrF())

	testConstrainedReal.AttrCrF = 12345.6789
	per1, err := perEncodeTestConstrainedReal(testConstrainedReal)
	assert.NilError(t, err)
	t.Logf("TestConstrainedReal PER\n%v", hex.Dump(per1))

	// Generating APER bytes with Go APER lib
	perNew1, err := aper.Marshal(testConstrainedReal, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)
	t.Logf("TestConstrainedReal PER (with Go APER library)\n%v", hex.Dump(perNew1))

	//Comparing bytes against each other
	assert.DeepEqual(t, per1, perNew1)

	result1, err := perDecodeTestConstrainedReal(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("TestConstrainedReal PER - decoded\n%v", result1)
	assert.Equal(t, testConstrainedReal.GetAttrCrA(), result1.GetAttrCrA())
	assert.Equal(t, testConstrainedReal.GetAttrCrB(), result1.GetAttrCrB())
	assert.Equal(t, testConstrainedReal.GetAttrCrC(), result1.GetAttrCrC())
	assert.Equal(t, testConstrainedReal.GetAttrCrD(), result1.GetAttrCrD())
	assert.Equal(t, testConstrainedReal.GetAttrCrE(), result1.GetAttrCrE())
	assert.Equal(t, testConstrainedReal.GetAttrCrF(), result1.GetAttrCrF())
}
