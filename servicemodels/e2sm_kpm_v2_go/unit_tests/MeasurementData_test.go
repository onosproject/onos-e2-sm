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

//var refPerMDfull = "00000000  00 00 40 03 00 15 20 09  80 d0 16 33 33 33 33 33  |..@... ....33333|\n" +
//	"00000010  33 40                                             |3@|"
var refPerMDnoReal = "0000000  00 00 40 02 00 15 40                              |..@...@|"
var refPerMDnoRealnoOptional = "00000000  00 00 00 02 00 15 40                              |......@|"

//func createMeasurementDataFull() (*e2sm_kpm_v2_go.MeasurementData, error) {
//
//	measRecord := &e2sm_kpm_v2_go.MeasurementRecord{
//		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
//	}
//
//	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
//		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
//			Integer: 21,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item1)
//
//	item2 := &e2sm_kpm_v2_go.MeasurementRecordItem{
//		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Real{
//			Real: 22.2,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item2)
//
//	item3 := &e2sm_kpm_v2_go.MeasurementRecordItem{
//		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
//			NoValue: 0,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item3)
//
//	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
//	measDataItem := &e2sm_kpm_v2_go.MeasurementDataItem{
//		MeasRecord:     measRecord,
//		IncompleteFlag: &incf,
//	}
//
//	measData := &e2sm_kpm_v2_go.MeasurementData{
//		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
//	}
//	measData.Value = append(measData.Value, measDataItem)
//
//	//if err := measData.Validate(); err != nil {
//	//	return nil, err
//	//}
//	return measData, nil
//}

func createMeasurementDataNoReal() (*e2sm_kpm_v2_go.MeasurementData, error) {

	measRecord := &e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

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

func createMeasurementDataNoRealNoOptional() (*e2sm_kpm_v2_go.MeasurementData, error) {

	measRecord := &e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item3 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	measDataItem := &e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord: measRecord,
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

//func Test_perEncodingMeasurementDataFull(t *testing.T) {
//
//	md, err := createMeasurementDataFull()
//	assert.NilError(t, err)
//
//	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
//	per, err := aper.Marshal(md)
//	assert.NilError(t, err)
//	t.Logf("MeasurementData PER\n%v", hex.Dump(per))
//
//	result := e2sm_kpm_v2_go.MeasurementData{}
//	err = aper.Unmarshal(per, &result)
//	assert.NilError(t, err)
//	assert.Assert(t, &result != nil)
//	t.Logf("MeasurementData PER - decoded\n%v", &result)
//	assert.Equal(t, md.GetValue()[0].GetIncompleteFlag().Number(), result.GetValue()[0].GetIncompleteFlag().Number())
//	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
//	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[1].GetReal(), result.GetValue()[0].GetMeasRecord().GetValue()[1].GetReal())
//	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[2].GetNoValue(), result.GetValue()[0].GetMeasRecord().GetValue()[2].GetNoValue())
//}
//
//func Test_perMeasurementDataFullCompareBytes(t *testing.T) {
//
//	md, err := createMeasurementDataFull()
//	assert.NilError(t, err)
//
//	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
//	per, err := aper.Marshal(md)
//	assert.NilError(t, err)
//	t.Logf("MeasurementData PER\n%v", hex.Dump(per))
//
//	//Comparing with reference bytes
//	perRefBytes, err := hexlib.DumpToByte(refPerMDfull)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, per, perRefBytes)
//}

func Test_perEncodingMeasurementDataNoReal(t *testing.T) {

	md, err := createMeasurementDataNoReal()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementData{}
	err = aper.Unmarshal(per, &result, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementData PER - decoded\n%v", &result)
	assert.Equal(t, md.GetValue()[0].GetIncompleteFlag().Number(), result.GetValue()[0].GetIncompleteFlag().Number())
	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perMeasurementDataNoRealCompareBytes(t *testing.T) {

	md, err := createMeasurementDataNoReal()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMDnoReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingMeasurementDataNoRealNoOptional(t *testing.T) {

	md, err := createMeasurementDataNoRealNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementData{}
	err = aper.Unmarshal(per, &result, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementData PER - decoded\n%v", &result)
	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, md.GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perMeasurementDataNoRealNoOptionalCompareBytes(t *testing.T) {

	md, err := createMeasurementDataNoRealNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.Marshal(md, e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMDnoRealnoOptional)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
