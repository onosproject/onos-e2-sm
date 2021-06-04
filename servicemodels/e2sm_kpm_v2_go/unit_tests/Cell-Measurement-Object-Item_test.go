// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerCellMeasObjItem = "00000000  00 00 03 31 32 33 40 4f  4e 46 d4 bc 09 00        |...123@ONF....|"

func createCellMeasurementObjectItem() *e2sm_kpm_v2_go.CellMeasurementObjectItem {

	return &e2sm_kpm_v2_go.CellMeasurementObjectItem{
		CellObjectId: &e2sm_kpm_v2_go.CellObjectId{
			Value: "123",
		},
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId{
			CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_kpm_v2_go.Eutracgi{
					EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
						Value: &asn1.BitString{
							Value: 0x9bcd4,
							Len:   28,
						},
					},
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte("ONF"),
					},
				},
			},
		},
	}
}

func Test_perEncodingCellMeasurementObjectItem(t *testing.T) {

	cmoi := createCellMeasurementObjectItem()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*cmoi, "valueExt")
	assert.NilError(t, err)
	t.Logf("CellMeasurementObjectItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellMeasurementObjectItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("CellMeasurementObjectItem PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellMeasObjItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
