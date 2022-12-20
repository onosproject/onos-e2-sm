// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

// var refPerIndMsg1 = "00000000  0e 80 30 38 00 00 03 6f  6e 66 00 14 00 00 40 20  |..08...onf....@ |\n" +
//
//	"00000010  74 72 69 61 6c 01 3f ff  e0 21 22 23 40 40 01 02  |trial.?..!\"#@@..|\n" +
//	"00000020  03 00 0a 7c 0f 00 0f 00  01 72 40 00 fa 00 00 04  |...|.....r@.....|\n" +
//	"00000030  00 00 7a 00 01 c7 00 03  14 00 00 00 40 03 08 30  |..z.........@..0|\n" +
//	"00000040  39 44 09 80 d0 16 33 33  33 33 33 33 00           |9D....333333.|"
//
// var refPerIndMsg2 = "00000000  2d 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |-08...onf....@ t|\n" +
//
//	"00000010  72 69 61 6c 00 00 48 21  02 00 c9 00 00 00 06 53  |rial..H!.......S|\n" +
//	"00000020  6f 6d 65 55 45 00 00 40  03 08 30 39 44 09 80 d0  |omeUE..@..09D...|\n" +
//	"00000030  16 33 33 33 33 33 33 00                           |.333333.|"
var refPerIndMsg1noReal = "00000000  0e 80 30 38 00 00 03 6f  6e 66 00 14 00 00 40 20  |..08...onf....@ |\n" +
	"00000010  74 72 69 61 6c 01 3f ff  e0 21 22 23 40 40 01 02  |trial.?..!\"#@@..|\n" +
	"00000020  03 00 0a 7c 0f 00 0f 00  01 72 40 00 fa 00 00 04  |...|.....r@.....|\n" +
	"00000030  00 00 7a 00 01 c7 00 03  14 00 00 00 40 02 08 30  |..z.........@..0|\n" +
	"00000040  39 40                                             |9@|"
var refPerIndMsg2noReal = "00000000  2d 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |-08...onf....@ t|\n" +
	"00000010  72 69 61 6c 00 00 48 21  02 00 c9 00 00 00 06 53  |rial..H!.......S|\n" +
	"00000020  6f 6d 65 55 45 00 00 40  02 08 30 39 40           |omeUE..@..09@|"

func createE2SMKPMIndicationMessageFormat1() (*e2smkpmv2.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	//var rl float64 = 22.2
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
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2smkpmv2.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2smkpmv2.StartEndInd_START_END_IND_START

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	labelInfoList := e2smkpmv2.LabelInfoList{
		Value: make([]*e2smkpmv2.LabelInfoItem, 0),
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

	measInfoList := e2smkpmv2.MeasurementInfoList{
		Value: make([]*e2smkpmv2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2smkpmv2.MeasurementRecord{
		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	//measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, err := pdubuilder.CreateMeasurementDataItem(&measRecord)
	if err != nil {
		return nil, err
	}
	measDataItem.SetIncompleteFlag()

	measData := e2smkpmv2.MeasurementData{
		Value: make([]*e2smkpmv2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, &measData)
	if err != nil {
		return nil, err
	}
	newE2SmKpmPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID).SetMeasInfoList(&measInfoList)
	if err = newE2SmKpmPdu.Validate(); err != nil {
		return nil, err
	}
	return newE2SmKpmPdu, nil
}

func createE2SMKPMIndicationMessageFormat2() (*e2smkpmv2.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	//var rl float64 = 22.2
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"
	var ueID = "SomeUE"

	var valEnum int64 = 201
	tce := e2smkpmv2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, _ := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))

	mci, _ := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)

	mcl := &e2smkpmv2.MatchingCondList{
		Value: make([]*e2smkpmv2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)

	mUEIDitem, _ := pdubuilder.CreateMatchingUEIDItem([]byte(ueID))

	mUEIDlist := &e2smkpmv2.MatchingUeidList{
		Value: make([]*e2smkpmv2.MatchingUeidItem, 0),
	}
	mUEIDlist.Value = append(mUEIDlist.Value, mUEIDitem)

	measCondUEIDItem, err := pdubuilder.CreateMeasurementCondUEIDItem(measName, mcl)
	if err != nil {
		return nil, err
	}
	measCondUEIDItem.SetMatchingUEUDlist(mUEIDlist)

	measCondUEIDList := e2smkpmv2.MeasurementCondUeidList{
		Value: make([]*e2smkpmv2.MeasurementCondUeidItem, 0),
	}
	measCondUEIDList.Value = append(measCondUEIDList.Value, measCondUEIDItem)

	measRecord := e2smkpmv2.MeasurementRecord{
		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	//measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, err := pdubuilder.CreateMeasurementDataItem(&measRecord)
	if err != nil {
		return nil, err
	}
	measDataItem.SetIncompleteFlag()

	measData := e2smkpmv2.MeasurementData{
		Value: make([]*e2smkpmv2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMessageFormat2(subscriptionID, &measCondUEIDList, &measData)
	if err != nil {
		return nil, err
	}
	newE2SmKpmPdu.SetCellObjectID(cellObjID).SetGranularityPeriod(granularity)
	if err := newE2SmKpmPdu.Validate(); err != nil {
		return nil, err
	}
	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmIndicationMessage1(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER - decoded\n%v", &result)
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetGranulPeriod().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetGranulPeriod().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetSubscriptId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetSubscriptId().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetCellObjId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetCellObjId().GetValue())
	assert.DeepEqual(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetIncompleteFlag().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetIncompleteFlag().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
	//assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[2].GetReal(), result.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[2].GetReal())
}

func Test_perE2SmKpmIndicationMessage1CompareBytes(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsg1noReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingE2SmKpmIndicationMessage2(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER - decoded\n%v", result)
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetGranulPeriod().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetGranulPeriod().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetSubscriptId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetSubscriptId().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetCellObjId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetCellObjId().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number())
	assert.DeepEqual(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasCondUeidList().GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetIncompleteFlag().Number(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetIncompleteFlag().Number())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
	//assert.Equal(t, im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[2].GetReal(), result.GetIndicationMessageFormats().GetIndicationMessageFormat2().GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[2].GetReal())
}

func Test_perE2SmKpmIndicationMessage2CompareBytes(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsg2noReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
