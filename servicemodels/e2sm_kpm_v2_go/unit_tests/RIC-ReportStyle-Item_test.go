// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerRicReportStyleItem = "00000000  00 7b 03 80 53 6f 6d 65  4e 61 6d 65 80 02 01 c8  |.{..SomeName....|\n" +
	"00000010  00 00 41 60 53 6f 6d 65  4d 65 61 73 4e 61 6d 65  |..A`SomeMeasName|\n" +
	"00000020  00 03 14 00 01 00 01                              |.......|"

func createRicReportStyleItem() (*e2sm_kpm_v2_go.RicReportStyleItem, error) {

	measInfo := &e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	item := &e2sm_kpm_v2_go.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
			Value: 789,
		},
		MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
			Value: "SomeMeasName",
		},
	}

	measInfo.Value = append(measInfo.Value, item)

	res := e2sm_kpm_v2_go.RicReportStyleItem{
		RicReportStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: 123,
		},
		RicReportStyleName: &e2sm_kpm_v2_go.RicStyleName{
			Value: "SomeName",
		},
		RicActionFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: 456,
		},
		MeasInfoActionList: measInfo,
		RicIndicationHeaderFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: 1,
		},
		RicIndicationMessageFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: 1,
		},
	}

	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}
	return &res, nil
}

func Test_perEncodingRicReportStyleItem(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(*item, "valueExt")
	assert.NilError(t, err)
	t.Logf("RIC-ReportStyle-Item PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicReportStyleItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RIC-ReportStyle-Item - decoded\n%v", result)
}

func Test_perRicReportStyleItemCompareBytes(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(*item, "valueExt")
	assert.NilError(t, err)
	t.Logf("RIC-ReportStyle-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicReportStyleItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
