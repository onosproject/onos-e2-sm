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

var refPerMDI = "00000000  00 03 00 15 20 09 80 d0  16 33 33 33 33 33 33 40  |.... ....333333@|"

func createMeasurementDataItem() (*e2sm_kpm_v2_go.MeasurementDataItem, error) {

	measRecord := &e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item2 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Real{
			Real: 22.2,
		},
	}
	measRecord.Value = append(measRecord.Value, item2)

	item3 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord: measRecord,
		IncompleteFlag: &incf,
	}

	//if err := measDataItem.Validate(); err != nil {
	//	return nil, err
	//}
	return measDataItem, nil
}

func Test_perEncodeMeasurementDataItem(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementDataItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementData-Item PER - decoded\n%v", result)
}

func Test_perMeasurementDataItemCompareBytes(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMDI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
