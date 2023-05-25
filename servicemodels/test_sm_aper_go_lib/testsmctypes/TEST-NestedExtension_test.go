// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTestNestedExtensionNoExt() *test_sm_ies.TestNestedExtension {

	return &test_sm_ies.TestNestedExtension{
		Item1: 21,
		Item2: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	}
}

func createTestNestedExtension() *test_sm_ies.TestNestedExtension {

	return &test_sm_ies.TestNestedExtension{
		Item1: 21,
		Item2: []byte{0xAA, 0xFF, 0xFF, 0xFF, 0xAA},
		Ch:    createAChoice2(),
	}
}

func Test_xerEncodingTestNestedExtension(t *testing.T) {

	testNestedExtension := createTestNestedExtension()

	xer, err := xerEncodeTestNestedExtension(testNestedExtension)
	assert.NilError(t, err)
	t.Logf("TEST-NestedExtension XER\n%s", string(xer))

	result, err := xerDecodeTestNestedExtension(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TEST-NestedExtension XER - decoded\n%v", result)
	assert.Equal(t, testNestedExtension.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetItem2(), result.GetItem2())
	assert.Equal(t, testNestedExtension.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetItem2(), result.GetItem2())
	assert.Equal(t, testNestedExtension.GetCh().GetCh2().GetItem()[0].GetItem1(), result.GetCh().GetCh2().GetItem()[0].GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetCh().GetCh2().GetItem()[0].GetItem2(), result.GetCh().GetCh2().GetItem()[0].GetItem2())
	assert.Equal(t, testNestedExtension.GetCh().GetCh2().GetItem()[1].GetItem1(), result.GetCh().GetCh2().GetItem()[1].GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetCh().GetCh2().GetItem()[1].GetItem2(), result.GetCh().GetCh2().GetItem()[1].GetItem2())
}

func Test_perEncodingTestNestedExtension(t *testing.T) {

	testNestedExtension := createTestNestedExtension()

	per, err := PerEncodeTestNestedExtension(testNestedExtension)
	assert.NilError(t, err)
	t.Logf("TEST-NestedExtension PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(testNestedExtension, "valueExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestNestedExtension(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TEST-NestedExtension PER - decoded\n%v", result)
	assert.Equal(t, testNestedExtension.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetItem2(), result.GetItem2())
	assert.Equal(t, testNestedExtension.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetItem2(), result.GetItem2())
	assert.Equal(t, testNestedExtension.GetCh().GetCh2().GetItem()[0].GetItem1(), result.GetCh().GetCh2().GetItem()[0].GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetCh().GetCh2().GetItem()[0].GetItem2(), result.GetCh().GetCh2().GetItem()[0].GetItem2())
	assert.Equal(t, testNestedExtension.GetCh().GetCh2().GetItem()[1].GetItem1(), result.GetCh().GetCh2().GetItem()[1].GetItem1())
	assert.DeepEqual(t, testNestedExtension.GetCh().GetCh2().GetItem()[1].GetItem2(), result.GetCh().GetCh2().GetItem()[1].GetItem2())
}
