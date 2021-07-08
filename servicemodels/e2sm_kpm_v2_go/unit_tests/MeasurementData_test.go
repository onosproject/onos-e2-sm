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

var refPerMD = "00000000  00 00 40 03 00 15 20 09  80 d0 16 33 33 33 33 33  |..@... ....33333|\n" +
	"00000010  33 40                                             |3@|"

func createMeasurementData() (*e2sm_kpm_v2_go.MeasurementData, error) {

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
		MeasRecord:     measRecord,
		IncompleteFlag: &incf,
	}

	measData := &e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	//if err := measData.Validate(); err != nil {
	//	return nil, err
	//}
	return measData, nil
}

func Test_perEncodingMeasurementData(t *testing.T) {

	md, err := createMeasurementData()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementData{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementData PER - decoded\n%v", result)
}

func Test_perMeasurementDataCompareBytes(t *testing.T) {

	md, err := createMeasurementData()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMD)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
