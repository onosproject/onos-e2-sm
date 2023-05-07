// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: LicenseRef-Intel

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTestExtension2Msg() (*test_sm_ies.TestExtension2, error) {

	var ie1 int32 = 32
	ie2 := []byte{0xF0, 0xB9, 0x32}
	ext1 := test_sm_ies.EnumThree_ENUM_THREE_THREE
	ext2 := test_sm_ies.EnumThree_ENUM_THREE_THREE
	ext3 := test_sm_ies.EnumThree_ENUM_THREE_THREE

	testExtension2 := test_sm_ies.TestExtension2{
		Item1: ie1,
		Item2: ie2,
		Ext1:  &ext1,
		Ext2:  &ext2,
		Ext3:  &ext3,
	}

	return &testExtension2, nil
}

func Test_xerEncodingTestExtension2(t *testing.T) {

	testExtension2, err := createTestExtension2Msg()
	assert.NilError(t, err, "Error creating TestExtension2 PDU")

	xer, err := xerEncodeTestExtension2(testExtension2)
	assert.NilError(t, err)
	t.Logf("TestExtension2 XER\n%s", string(xer))

	result, err := xerDecodeTestExtension2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestExtension2 XER - decoded\n%v", result)
	assert.Equal(t, testExtension2.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testExtension2.GetItem2(), result.GetItem2())
	assert.Equal(t, testExtension2.GetExt1(), result.GetExt1())
	assert.Equal(t, testExtension2.GetExt2(), result.GetExt2())
}

func Test_perEncodingTestExtension2(t *testing.T) {

	testExtension2, err := createTestExtension2Msg()
	assert.NilError(t, err, "Error creating TestExtension2 PDU")

	per, err := PerEncodeTestExtension2(testExtension2)
	assert.NilError(t, err)
	t.Logf("TestExtension2 PER\n%v", hex.Dump(per))

	//// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(testExtension2, "valueExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)
	//
	////Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestExtension2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestExtension2 PER - decoded\n%v", result)
	assert.Equal(t, testExtension2.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, testExtension2.GetItem2(), result.GetItem2())
	assert.Equal(t, testExtension2.GetExt1(), result.GetExt1())
	assert.Equal(t, testExtension2.GetExt2(), result.GetExt2())
}
