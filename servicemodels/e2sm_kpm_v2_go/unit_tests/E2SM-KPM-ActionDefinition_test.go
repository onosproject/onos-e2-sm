// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerADF1 = "00000000  00 0c 00 00 03 6f 6e 66  00 00 40 20 74 72 69 61  |.....onf..@ tria|\n" +
	"00000010  6c 01 3f ff e0 21 22 23  40 40 01 02 03 00 0a 7c  |l.?..!\"#@@.....||\n" +
	"00000020  0f 00 0f 00 01 70 00 00  fa 00 00 04 00 00 7a 00  |.....p........z.|\n" +
	"00000030  01 c7 00 03 14 00 14 40  30 38                    |.......@08|"
var refPerADF2 = "00000000  00 0c 20 06 53 6f 6d 65  55 45 00 00 03 6f 6e 66  |.. .SomeUE...onf|\n" +
	"00000010  00 00 40 20 74 72 69 61  6c 01 3f ff e0 21 22 23  |..@ trial.?..!\"#|\n" +
	"00000020  40 40 01 02 03 00 0a 7c  0f 00 0f 00 01 70 00 00  |@@.....|.....p..|\n" +
	"00000030  fa 00 00 04 00 00 7a 00  01 c7 00 03 14 00 14 40  |......z........@|\n" +
	"00000040  30 38                                             |08|"
var refPerADF3 = "00000000  00 0c 40 00 03 6f 6e 66  00 00 00 40 74 72 69 61  |..@..onf...@tria|\n" +
	"00000010  6c 00 00 48 21 02 00 c9  00 14 40 30 38           |l..H!.....@08|"

func createE2SmKpmActionDefinitionFormat1() (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	if err != nil {
		return nil, err
	}
	measInfoItem, err := pdubuilder.CreateMeasurementInfoItem(measName)
	if err != nil {
		return nil, err
	}
	measInfoItem.SetLabelInfoList(&labelInfoList)

	measInfoList := e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := pdubuilder.CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	if err != nil {
		return nil, err
	}

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmActionDefinitionFormat1(ricStyleType, actionDefinitionFormat1)
	if err != nil {
		return nil, err
	}
	//if err := actionDefinitionFormat1.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat1 %s", err.Error())
	//}
	return newE2SmKpmPdu, nil
}

func createE2SmKpmActionDefinitionFormat2() (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	if err != nil {
		return nil, err
	}
	measInfoItem, err := pdubuilder.CreateMeasurementInfoItem(measName)
	if err != nil {
		return nil, err
	}
	measInfoItem.SetLabelInfoList(&labelInfoList)

	measInfoList := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := pdubuilder.CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	if err != nil {
		return nil, err
	}

	ueID := "SomeUE"
	actionDefinitionFormat2, err := pdubuilder.CreateActionDefinitionFormat2([]byte(ueID), actionDefinitionFormat1)
	if err != nil {
		return nil, err
	}

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmActionDefinitionFormat2(ricStyleType, actionDefinitionFormat2)
	if err != nil {
		return nil, err
	}
	//if err := actionDefinitionFormat2.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	//}
	return newE2SmKpmPdu, nil
}

func createE2SmKpmActionDefinitionFormat3() (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))
	if err != nil {
		return nil, err
	}

	mci, err := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)
	if err != nil {
		return nil, err
	}

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	if err != nil {
		return nil, err
	}
	measCondItem, err := pdubuilder.CreateMeasurementCondItem(measName, mcl)
	if err != nil {
		return nil, err
	}

	measCondList := e2sm_kpm_v2_go.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, err := pdubuilder.CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	if err != nil {
		return nil, err
	}

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmActionDefinitionFormat3(ricStyleType, actionDefinitionFormat3)
	if err != nil {
		return nil, err
	}
	//if err := actionDefinitionFormat3.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	//}
	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmActionDefinitionF1(t *testing.T) {

	actionDefFormat1, err := createE2SmKpmActionDefinitionFormat1()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat1)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmActionDefinition(per)
	//err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER - decoded\n%v", result)
	assert.Equal(t, actionDefFormat1.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetCellObjId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetCellObjId().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetGranulPeriod().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetGranulPeriod().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetSubscriptId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetSubscriptId().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.DeepEqual(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, actionDefFormat1.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
}

func Test_perE2SmKpmActionDefinitionF1CompareBytes(t *testing.T) {

	actionDefFormat1, err := createE2SmKpmActionDefinitionFormat1()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat1)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerADF1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingE2SmKpmActionDefinitionF2(t *testing.T) {

	actionDefFormat2, err := createE2SmKpmActionDefinitionFormat2()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat2)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format2) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmActionDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format2) PER - decoded\n%v", result)
	assert.Equal(t, actionDefFormat2.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetCellObjId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetCellObjId().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetGranulPeriod().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetGranulPeriod().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetSubscriptId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetSubscriptId().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.DeepEqual(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.DeepEqual(t, actionDefFormat2.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetUeId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat2().GetUeId().GetValue())
}

func Test_perE2SmKpmActionDefinitionF2CompareBytes(t *testing.T) {

	actionDefFormat2, err := createE2SmKpmActionDefinitionFormat2()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat2)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format2) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerADF2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingE2SmKpmActionDefinitionF3(t *testing.T) {

	actionDefFormat3, err := createE2SmKpmActionDefinitionFormat3()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat3)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmActionDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER - decoded\n%v", result)
	assert.Equal(t, actionDefFormat3.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetGranulPeriod().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetGranulPeriod().GetValue())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetSubscriptId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetSubscriptId().GetValue())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetCellObjId().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetCellObjId().GetValue())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueEnum(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueEnum())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetRSrp().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetRSrp().Number())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number())
	assert.Equal(t, actionDefFormat3.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetActionDefinitionFormats().GetActionDefinitionFormat3().GetMeasCondList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
}

func Test_perE2SmKpmActionDefinitionF3CompareBytes(t *testing.T) {

	actionDefFormat3, err := createE2SmKpmActionDefinitionFormat3()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(actionDefFormat3)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerADF3)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
