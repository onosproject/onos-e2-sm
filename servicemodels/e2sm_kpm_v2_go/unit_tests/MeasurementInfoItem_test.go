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

var refPerMII = "00000000  50 00 14 01 3f ff e0 01  02 03 47 c0 01 02 03 00  |P...?.....G.....|\n" +
	"00000010  17 7c 18 00 1e 00 01 70  00 00 18 00 00 00 00 00  |.|.....p........|\n" +
	"00000020  7a 00 01 c7 00 03 14 00                           |z.......|"

var refPerMIInoOptional = "00000000  10 00 14                                          |...|"

func createMeasurementInfoItem() (*e2smkpmv2.MeasurementInfoItem, error) {

	labelInfoList := &e2smkpmv2.LabelInfoList{
		Value: make([]*e2smkpmv2.LabelInfoItem, 0),
	}

	var br int32 = 25
	var lmm int32 = 1
	sum := e2smkpmv2.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2smkpmv2.StartEndInd_START_END_IND_START

	labelInfoItem := &e2smkpmv2.LabelInfoItem{
		MeasLabel: &e2smkpmv2.MeasurementLabel{
			PlmnId: &e2smkpmv2.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			SliceId: &e2smkpmv2.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x1f},
			},
			FiveQi: &e2smkpmv2.FiveQi{
				Value: 23,
			},
			QFi: &e2smkpmv2.Qfi{
				Value: 62,
			},
			QCi: &e2smkpmv2.Qci{
				Value: 24,
			},
			QCimax: &e2smkpmv2.Qci{
				Value: 30,
			},
			QCimin: &e2smkpmv2.Qci{
				Value: 1,
			},
			ARpmax: &e2smkpmv2.Arp{
				Value: 15,
			},
			ARpmin: &e2smkpmv2.Arp{
				Value: 1,
			},
			BitrateRange:     &br,
			LayerMuMimo:      &lmm,
			SUm:              &sum,
			DistBinX:         &dbx,
			DistBinY:         &dby,
			DistBinZ:         &dbz,
			PreLabelOverride: &plo,
			StartEndInd:      &seind,
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	res := &e2smkpmv2.MeasurementInfoItem{
		MeasType: &e2smkpmv2.MeasurementType{
			MeasurementType: &e2smkpmv2.MeasurementType_MeasId{
				MeasId: &e2smkpmv2.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}
	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}

	return res, nil
}

func createMeasurementInfoItemNoOptional() (*e2smkpmv2.MeasurementInfoItem, error) {
	res := &e2smkpmv2.MeasurementInfoItem{
		MeasType: &e2smkpmv2.MeasurementType{
			MeasurementType: &e2smkpmv2.MeasurementType_MeasId{
				MeasId: &e2smkpmv2.MeasurementTypeId{
					Value: 21,
				},
			},
		},
	}
	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}

	return res, nil
}

func Test_perEncodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-Item PER - decoded\n%v", &result)
	assert.DeepEqual(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, mii.GetMeasType().GetMeasId().GetValue(), result.GetMeasType().GetMeasId().GetValue())
}

func Test_perMeasurementInfoItemCompareBytes(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMII)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodeMeasurementInfoItemNoOptional(t *testing.T) {

	mii, err := createMeasurementInfoItemNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-Item PER - decoded\n%v", &result)
	assert.Equal(t, mii.GetMeasType().GetMeasId().GetValue(), result.GetMeasType().GetMeasId().GetValue())
}

func Test_perMeasurementInfoItemNoOptionalCompareBytes(t *testing.T) {

	mii, err := createMeasurementInfoItemNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMIInoOptional)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
