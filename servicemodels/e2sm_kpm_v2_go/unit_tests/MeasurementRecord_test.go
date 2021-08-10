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

var refPerMeasRecord = "00000000  03 00 15 20 03 80 01 0b  40                       |... ....@|"
var refPerMeasRecordNoReal = "00000000  02 00 15 40                                       |...@|"

func createMeasurementRecord() (*e2sm_kpm_v2_go.MeasurementRecord, error) {
	res := &e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	res.Value = append(res.Value, item1)

	//ToDo - bring back once handling of REAL types is implemented
	//item2 := &e2sm_kpm_v2_go.MeasurementRecordItem{
	//	MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Real{
	//		Real: 22,
	//	},
	//}
	//res.Value = append(res.Value, item2)

	item3 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	res.Value = append(res.Value, item3)

	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}
	return res, nil
}

func Test_perEncodingMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mr, "")
	assert.NilError(t, err)
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementRecord{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecord PER - decoded\n%v", &result)
	assert.Equal(t, 2, len(mr.GetValue()))
	assert.Equal(t, mr.GetValue()[0].GetInteger(), result.GetValue()[0].GetInteger())
	//assert.Equal(t, mr.GetValue()[1].GetReal(), result.GetValue()[1].GetReal())
	assert.Equal(t, mr.GetValue()[1].GetNoValue(),result.GetValue()[1].GetNoValue())
}

func Test_perMeasurementRecordCompareBytes(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mr, "")
	assert.NilError(t, err)
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasRecordNoReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
