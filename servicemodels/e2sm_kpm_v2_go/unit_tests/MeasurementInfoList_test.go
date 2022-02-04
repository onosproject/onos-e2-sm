// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMIL = "00000000  00 00 50 00 14 01 3f ff  e0 01 02 03 43 c0 01 02  |..P...?.....C...|\n" +
	"00000010  03 00 17 7c 18 00 1e 00  01 70 00 00 18 00 00 00  |...|.....p......|\n" +
	"00000020  00 00 7a 00 01 c7 00 03  14 20                    |..z...... |"
var refPerMILnoOptional = "00000000  00 00 10 00 14                                    |.....|"

func createMeasurementInfoList() (*e2sm_kpm_v2_go.MeasurementInfoList, error) {

	labelInfoList := &e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	labelInfoItem := &e2sm_kpm_v2_go.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			SliceId: &e2sm_kpm_v2_go.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x0f},
			},
			FiveQi: &e2sm_kpm_v2_go.FiveQi{
				Value: 23,
			},
			QFi: &e2sm_kpm_v2_go.Qfi{
				Value: 62,
			},
			QCi: &e2sm_kpm_v2_go.Qci{
				Value: 24,
			},
			QCimax: &e2sm_kpm_v2_go.Qci{
				Value: 30,
			},
			QCimin: &e2sm_kpm_v2_go.Qci{
				Value: 1,
			},
			ARpmax: &e2sm_kpm_v2_go.Arp{
				Value: 15,
			},
			ARpmin: &e2sm_kpm_v2_go.Arp{
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

	list := &e2sm_kpm_v2_go.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2_go.MeasurementType{
			MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}

	res := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	res.Value = append(res.Value, list)

	//if err := labelInfoList.Validate(); err != nil {
	//	return nil, err
	//}
	return res, nil
}

func createMeasurementInfoListNoOptional() (*e2sm_kpm_v2_go.MeasurementInfoList, error) {

	list := &e2sm_kpm_v2_go.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2_go.MeasurementType{
			MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
					Value: 21,
				},
			},
		},
	}

	res := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	res.Value = append(res.Value, list)

	//if err := labelInfoList.Validate(); err != nil {
	//	return nil, err
	//}
	return res, nil
}

func Test_perEncodeMeasurementInfoList(t *testing.T) {

	mil, err := createMeasurementInfoList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mil, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementInfoList{}
	err = aper.UnmarshalWithParams(per, &result, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-List PER - decoded\n%v", &result)
	assert.DeepEqual(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mil.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, mil.GetValue()[0].GetMeasType().GetMeasId().GetValue(), result.GetValue()[0].GetMeasType().GetMeasId().GetValue())
}

func Test_perMeasurementInfoListCompareBytes(t *testing.T) {

	mil, err := createMeasurementInfoList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mil, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMIL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodeMeasurementInfoListNoOptional(t *testing.T) {

	mil, err := createMeasurementInfoListNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mil, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementInfoList{}
	err = aper.UnmarshalWithParams(per, &result, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-List PER - decoded\n%v", &result)
	assert.Equal(t, mil.GetValue()[0].GetMeasType().GetMeasId().GetValue(), result.GetValue()[0].GetMeasType().GetMeasId().GetValue())
}

func Test_perMeasurementInfoListCompareBytesNoOptional(t *testing.T) {

	mil, err := createMeasurementInfoListNoOptional()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mil, "", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMILnoOptional)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
