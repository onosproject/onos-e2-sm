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

var refPerMRIinteger = "00000000  00 15                                             |..|"

// var refPerMRIreal = "0000000  20 03 80 01 0b                                    | ....|"
var refPerMRInovalue = "00000000  40                                                |@|"

func createMeasurementRecordItemInteger() *e2smkpmv2.MeasurementRecordItem {
	return &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
}

//func createMeasurementRecordItemReal() *e2smkpmv2.MeasurementRecordItem {
//	return &e2smkpmv2.MeasurementRecordItem{
//		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Real{
//			Real: 22,
//		},
//	}
//}

func createMeasurementRecordItemNoValue() *e2smkpmv2.MeasurementRecordItem {
	return &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func Test_perEncodingMeasurementRecordItemInteger(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementRecordItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecordItem (Integer) PER - decoded\n%v", &result)
	assert.Equal(t, mri.GetInteger(), result.GetInteger())
}

func Test_perMeasurementRecordItemIntegerCompareBytes(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMRIinteger)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

//func Test_perEncodingMeasurementRecordItemReal(t *testing.T) {
//
//	mri := createMeasurementRecordItemReal()
//
//	aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mri, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementRecordItem (Real) PER\n%v", hex.Dump(per))
//
//	result := e2smkpmv2.MeasurementRecordItem{}
//	err = aper.UnmarshalWithParams(per, &result, "valueExt")
//	assert.NilError(t, err)
//	assert.Assert(t, &result != nil)
//	t.Logf("MeasurementRecordItem (Real) PER - decoded\n%v", &result)
//	assert.Equal(t, mri.GetReal(), result.GetReal())
//}
//
//func Test_perMeasurementRecordItemRealCompareBytes(t *testing.T) {
//
//	mri := createMeasurementRecordItemReal()
//
//	aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mri, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementRecordItem (Real) PER\n%v", hex.Dump(per))
//
//	//Comparing with reference bytes
//	perRefBytes, err := hexlib.DumpToByte(refPerMRIreal)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, per, perRefBytes)
//}

func Test_perEncodingMeasurementRecordItemNull(t *testing.T) {

	mri := createMeasurementRecordItemNoValue()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (No value) PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementRecordItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecordItem (No value) PER - decoded\n%v", &result)
	assert.Equal(t, mri.GetNoValue(), result.GetNoValue())
}

func Test_perMeasurementRecordItemNullCompareBytes(t *testing.T) {

	mri := createMeasurementRecordItemNoValue()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (No value) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMRInovalue)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
