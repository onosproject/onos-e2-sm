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

func createTestConstrainedIntMsg() (*test_sm_ies.TestConstrainedInt, error) {

	testConstrainedInt := test_sm_ies.TestConstrainedInt{
		AttrCiA: 99,
		AttrCiB: 65534,
		AttrCiC: 11,
		AttrCiD: -424684151,
		AttrCiE: 15,
		AttrCiF: 10,
		//AttrCiG: 10,
	}

	return &testConstrainedInt, nil
}

func createTestConstrainedIntMsgCmpr() (*test_sm_ies.TestConstrainedInt, error) {

	testConstrainedInt := test_sm_ies.TestConstrainedInt{
		AttrCiA: 100,
		AttrCiB: 65534,
		AttrCiC: 2147483647,
		AttrCiD: -2147483647, //minimum value by default (defined as MIN in ASN1 syntax, MAX is not possible to estimate)
		AttrCiE: 20,
		AttrCiF: 10,
		//AttrCiG: 12,
	}

	return &testConstrainedInt, nil
}

func Test_xerEncodingTestConstrainedInt(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	xer, err := xerEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedInt(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt XER - decoded\n%v", result)
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())
}

func Test_perEncodingTestConstrainedInt(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsg()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	per, err := PerEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	perNew, err := aper.Marshal(testConstrainedInt, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt PER (with Go APER library)\n%v", hex.Dump(perNew))

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestConstrainedInt(perNew)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt PER - decoded\n%v", result)
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())
}

func Test_perEncodingTestConstrainedIntCmpr(t *testing.T) {

	testConstrainedInt, err := createTestConstrainedIntMsgCmpr()
	assert.NilError(t, err, "Error creating TestConstrainedInt PDU")

	per, err := PerEncodeTestConstrainedInt(testConstrainedInt)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	perNew, err := aper.Marshal(testConstrainedInt, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)
	t.Logf("TestConstrainedInt PER (with Go APER library)\n%v", hex.Dump(perNew))

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestConstrainedInt(perNew)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedInt PER - decoded\n%v", result)
	assert.Equal(t, testConstrainedInt.GetAttrCiA(), result.GetAttrCiA())
	assert.Equal(t, testConstrainedInt.GetAttrCiB(), result.GetAttrCiB())
	assert.Equal(t, testConstrainedInt.GetAttrCiC(), result.GetAttrCiC())
	assert.Equal(t, testConstrainedInt.GetAttrCiD(), result.GetAttrCiD())
	assert.Equal(t, testConstrainedInt.GetAttrCiE(), result.GetAttrCiE())
	assert.Equal(t, testConstrainedInt.GetAttrCiF(), result.GetAttrCiF())
}
