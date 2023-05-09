// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: LicenseRef-Intel

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTestExtension1Msg() (*test_sm_ies.TestExtension1, error) {

	var ie1 int32 = 32
	ie2 := []byte{0xF0, 0xB9, 0x32}
	ext1 := []byte{0xF0, 0xB9, 0x32, 0x77, 0xFF}

	testExtension1 := test_sm_ies.TestExtension1{
		Item1: ie1,
		Item2: ie2,
		Ext1:  ext1,
	}

	return &testExtension1, nil
}

func Test_xerEncodingTestExtension1(t *testing.T) {

	testExtension1, err := createTestExtension1Msg()
	assert.NilError(t, err, "Error creating TestExtension1 PDU")

	xer, err := xerEncodeTestExtension1(testExtension1)
	assert.NilError(t, err)
	t.Logf("TestExtension1 XER\n%s", string(xer))

	result, err := xerDecodeTestExtension1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestExtension1 XER - decoded\n%v", result)
	assert.Equal(t, testExtension1.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testExtension1.GetItem2(), result.GetItem2())
	assert.DeepEqual(t, testExtension1.GetExt1(), result.GetExt1())
}

func Test_perEncodingTestExtension1(t *testing.T) {

	testExtension1, err := createTestExtension1Msg()
	assert.NilError(t, err, "Error creating TestExtension1 PDU")

	per, err := PerEncodeTestExtension1(testExtension1)
	assert.NilError(t, err)
	t.Logf("TestExtension1 PER\n%v", hex.Dump(per))

	//// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(testExtension1, "valueExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)
	//
	////Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestExtension1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestExtension1 PER - decoded\n%v", result)
	assert.Equal(t, testExtension1.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testExtension1.GetItem2(), result.GetItem2())
	assert.DeepEqual(t, testExtension1.GetExt1(), result.GetExt1())
}
