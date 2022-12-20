// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

// var refPerMDIfull = "00000000  00 03 00 15 20 09 80 d0  16 33 33 33 33 33 33 40  |.... ....333333@|"
var refPerMDInoRealnoOptional = "00000000  00 02 00 15 40                                    |....@|"
var refPerMDInoReal = "00000000  40 02 00 15 40                                    |@...@|"

//func createMeasurementDataItemFull() (*e2smkpmv2.MeasurementDataItem, error) {
//
//	measRecord := &e2smkpmv2.MeasurementRecord{
//		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
//	}
//
//	item1 := &e2smkpmv2.MeasurementRecordItem{
//		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
//			Integer: 21,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item1)
//
//	item2 := &e2smkpmv2.MeasurementRecordItem{
//		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Real{
//			Real: 22.2,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item2)
//
//	item3 := &e2smkpmv2.MeasurementRecordItem{
//		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
//			NoValue: 0,
//		},
//	}
//	measRecord.Value = append(measRecord.Value, item3)
//
//	incf := e2smkpmv2.IncompleteFlag_INCOMPLETE_FLAG_TRUE
//	measDataItem := &e2smkpmv2.MeasurementDataItem{
//		MeasRecord:     measRecord,
//		IncompleteFlag: &incf,
//	}
//
//	//if err := measDataItem.Validate(); err != nil {
//	//	return nil, err
//	//}
//	return measDataItem, nil
//}

func createMeasurementDataItemNoReal() (*e2smkpmv2.MeasurementDataItem, error) {

	measRecord := &e2smkpmv2.MeasurementRecord{
		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
	}

	item1 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item3 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	incf := e2smkpmv2.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2smkpmv2.MeasurementDataItem{
		MeasRecord:     measRecord,
		IncompleteFlag: &incf,
	}

	//if err := measDataItem.Validate(); err != nil {
	//	return nil, err
	//}
	return measDataItem, nil
}

func createMeasurementDataItemNoRealNoOptional() (*e2smkpmv2.MeasurementDataItem, error) {

	measRecord := &e2smkpmv2.MeasurementRecord{
		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
	}

	item1 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item3 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	measDataItem := &e2smkpmv2.MeasurementDataItem{
		MeasRecord: measRecord,
	}

	//if err := measDataItem.Validate(); err != nil {
	//	return nil, err
	//}
	return measDataItem, nil
}

//func Test_perEncodeMeasurementDataItemFull(t *testing.T) {
//
//	mdi, err := createMeasurementDataItemFull()
//	assert.NilError(t, err)
//
//	aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mdi, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))
//
//	result := e2smkpmv2.MeasurementDataItem{}
//	err = aper.UnmarshalWithParams(per, &result, "valueExt")
//	assert.NilError(t, err)
//	assert.Assert(t, &result != nil)
//	t.Logf("MeasurementData-Item PER - decoded\n%v", &result)
//	assert.Equal(t, mdi.GetIncompleteFlag().Number(), result.GetIncompleteFlag().Number())
//	assert.Equal(t, mdi.GetMeasRecord().GetValue()[0].GetInteger(), result.GetMeasRecord().GetValue()[0].GetInteger())
//	assert.Equal(t, mdi.GetMeasRecord().GetValue()[1].GetReal(), result.GetMeasRecord().GetValue()[1].GetReal())
//	assert.Equal(t, mdi.GetMeasRecord().GetValue()[2].GetNoValue(), result.GetMeasRecord().GetValue()[2].GetNoValue())
//}
//
//func Test_perMeasurementDataItemCompareBytesFull(t *testing.T) {
//
//	mdi, err := createMeasurementDataItemFull()
//	assert.NilError(t, err)
//
//	aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mdi, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))
//
//	//Comparing with reference bytes
//	perRefBytes, err := hexlib.DumpToByte(refPerMDIfull)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, per, perRefBytes)
//}

func Test_perEncodeMeasurementDataItemNoReal(t *testing.T) {

	mdi, err := createMeasurementDataItemNoReal()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementDataItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementData-Item PER - decoded\n%v", &result)
	assert.Equal(t, mdi.GetIncompleteFlag().Number(), result.GetIncompleteFlag().Number())
	assert.Equal(t, mdi.GetMeasRecord().GetValue()[0].GetInteger(), result.GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, mdi.GetMeasRecord().GetValue()[1].GetNoValue(), result.GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perMeasurementDataItemCompareBytesNoReal(t *testing.T) {

	mdi, err := createMeasurementDataItemNoReal()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMDInoReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodeMeasurementDataItemNoRealNoOptional(t *testing.T) {

	mdi, err := createMeasurementDataItemNoRealNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementDataItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementData-Item PER - decoded\n%v", &result)
	assert.Equal(t, mdi.GetMeasRecord().GetValue()[0].GetInteger(), result.GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, mdi.GetMeasRecord().GetValue()[1].GetNoValue(), result.GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perMeasurementDataItemCompareBytesNoRealNoOptional(t *testing.T) {

	mdi, err := createMeasurementDataItemNoRealNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mdi, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementData-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMDInoRealnoOptional)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
