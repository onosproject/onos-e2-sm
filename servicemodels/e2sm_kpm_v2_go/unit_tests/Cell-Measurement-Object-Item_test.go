// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerCellMeasObjItem1 = "00000000  00 00 03 31 32 33 40 4f  4e 46 d4 bc 09 00        |...123@ONF....|"
var refPerCellMeasObjItem2 = "00000000  00 00 18 4f 70 65 6e 4e  65 74 77 6f 72 6b 69 6e  |...OpenNetworkin|\n" +
	"00000010  67 46 6f 75 6e 64 61 74  69 6f 6e 00 4f 4e 46 d4  |gFoundation.ONF.|\n" +
	"00000020  bc 09 00 f0                                       |....|"

func createCellMeasurementObjectItem1() *e2smkpmv2.CellMeasurementObjectItem {

	return &e2smkpmv2.CellMeasurementObjectItem{
		CellObjectId: &e2smkpmv2.CellObjectId{
			Value: "123",
		},
		CellGlobalId: &e2smkpmv2.CellGlobalId{
			CellGlobalId: &e2smkpmv2.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2smkpmv2.Eutracgi{
					EUtracellIdentity: &e2smkpmv2.EutracellIdentity{
						Value: &asn1.BitString{
							Value: []byte{0xd4, 0xbc, 0x09, 0x00},
							Len:   28,
						},
					},
					PLmnIdentity: &e2smkpmv2.PlmnIdentity{
						Value: []byte("ONF"),
					},
				},
			},
		},
	}
}

func createCellMeasurementObjectItem2() *e2smkpmv2.CellMeasurementObjectItem {

	return &e2smkpmv2.CellMeasurementObjectItem{
		CellObjectId: &e2smkpmv2.CellObjectId{
			Value: "OpenNetworkingFoundation",
		},
		CellGlobalId: &e2smkpmv2.CellGlobalId{
			CellGlobalId: &e2smkpmv2.CellGlobalId_NrCgi{
				NrCgi: &e2smkpmv2.Nrcgi{
					NRcellIdentity: &e2smkpmv2.NrcellIdentity{
						Value: &asn1.BitString{
							Value: []byte{0xd4, 0xbc, 0x09, 0x00, 0xf0},
							Len:   36,
						},
					},
					PLmnIdentity: &e2smkpmv2.PlmnIdentity{
						Value: []byte("ONF"),
					},
				},
			},
		},
	}
}

func Test_perEncodingCellMeasurementObjectItem1(t *testing.T) {

	cmoi := createCellMeasurementObjectItem1()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cmoi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellMeasurementObjectItem PER\n%v", hex.Dump(per))

	result := e2smkpmv2.CellMeasurementObjectItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellMeasurementObjectItem PER - decoded\n%v", &result)
	assert.Equal(t, cmoi.GetCellObjectId().GetValue(), result.GetCellObjectId().GetValue())
	assert.DeepEqual(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue(), result.GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_perCellMeasurementObjectItem1CompareBytes(t *testing.T) {

	cmoi := createCellMeasurementObjectItem1()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cmoi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellMeasurementObjectItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellMeasObjItem1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingCellMeasurementObjectItem2(t *testing.T) {

	cmoi := createCellMeasurementObjectItem2()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cmoi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellMeasurementObjectItem PER\n%v", hex.Dump(per))

	result := e2smkpmv2.CellMeasurementObjectItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellMeasurementObjectItem PER - decoded\n%v", &result)
	assert.Equal(t, cmoi.GetCellObjectId().GetValue(), result.GetCellObjectId().GetValue())
	assert.DeepEqual(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue(), result.GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, cmoi.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_perCellMeasurementObjectItem2CompareBytes(t *testing.T) {

	cmoi := createCellMeasurementObjectItem2()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cmoi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellMeasurementObjectItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellMeasObjItem2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
