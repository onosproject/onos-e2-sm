// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerCellObjectID0 = "00000000  00 00 00                                          |...|"
var refPerCellObjectID1 = "00000000  00 00 03 31 32 33                                 |...123|"
var refPerCellObjectID2 = "00000000  00 00 18 4f 70 65 6e 4e  65 74 77 6f 72 6b 69 6e  |...OpenNetworkin|\n" +
	"00000010  67 46 6f 75 6e 64 61 74  69 6f 6e                 |gFoundation|"

func createCellObjectID0() *e2sm_kpm_v2_go.CellObjectId {

	return &e2sm_kpm_v2_go.CellObjectId{
		Value: "",
	}
}

func createCellObjectID1() *e2sm_kpm_v2_go.CellObjectId {

	return &e2sm_kpm_v2_go.CellObjectId{
		Value: "123",
	}
}

func createCellObjectID2() *e2sm_kpm_v2_go.CellObjectId {

	return &e2sm_kpm_v2_go.CellObjectId{
		Value: "OpenNetworkingFoundation",
	}
}

func Test_perEncodingCellObjectID0(t *testing.T) {

	coID := createCellObjectID0()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellObjectId{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellObjectID PER - decoded\n%v", &result)
	assert.Equal(t, coID.GetValue(), result.GetValue())
}

func Test_perCellObjectID0CompareBytes(t *testing.T) {

	coID := createCellObjectID0()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellObjectID0)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingCellObjectID1(t *testing.T) {

	coID := createCellObjectID1()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellObjectId{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellObjectID PER - decoded\n%v", &result)
	assert.Equal(t, coID.GetValue(), result.GetValue())
}

func Test_perCellObjectID1CompareBytes(t *testing.T) {

	coID := createCellObjectID1()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellObjectID1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingCellObjectID2(t *testing.T) {

	coID := createCellObjectID2()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellObjectId{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellObjectID PER - decoded\n%v", &result)
	assert.Equal(t, coID.GetValue(), result.GetValue())
}

func Test_perCellObjectID2CompareBytes(t *testing.T) {

	coID := createCellObjectID2()

	per, err := aper.MarshalWithParams(coID, "")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellObjectID2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
