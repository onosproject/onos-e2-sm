// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

// var refPerIndMsgF1 = "00000000  74 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |t08...onf....@ t|\n" +
//
//	"00000010  72 69 61 6c 01 3f ff e0  21 22 23 40 40 01 02 03  |rial.?..!\"#@@...|\n" +
//	"00000020  00 0a 7c 0f 00 0f 00 01  72 40 00 fa 00 00 04 00  |..|.....r@......|\n" +
//	"00000030  00 7a 00 01 c7 00 03 14  00 00 00 40 03 08 30 39  |.z.........@..09|\n" +
//	"00000040  44 09 80 d9 0d 42 c1 47  ae 14 7b 00              |D....B.G..{.|"
var refPerIndMsgF1noReal = "00000000  74 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |t08...onf....@ t|\n" +
	"00000010  72 69 61 6c 01 3f ff e0  21 22 23 40 40 01 02 03  |rial.?..!\"#@@...|\n" +
	"00000020  00 0a 7c 0f 00 0f 00 01  72 40 00 fa 00 00 04 00  |..|.....r@......|\n" +
	"00000030  00 7a 00 01 c7 00 03 14  00 00 00 40 02 08 30 39  |.z.........@..09|\n" +
	"00000040  40                                                |@|"

func createIndicationMessageFormat1() (*e2smkpmv2.E2SmKpmIndicationMessageFormat1, error) {

	var integer int64 = 12345
	//var rl float64 = 6789.51
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

	incflg := e2smkpmv2.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2smkpmv2.MeasurementDataItem{
		MeasRecord:     &measRecord,
		IncompleteFlag: &incflg,
	}

	measData := &e2smkpmv2.MeasurementData{
		Value: make([]*e2smkpmv2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, measData)
	if err != nil {
		return nil, err
	}
	newE2SmKpmPdu.SetCellObjectID(cellObjID).SetGranularityPeriod(granularity).SetMeasInfoList(&measInfoList)
	if err = newE2SmKpmPdu.Validate(); err != nil {
		return nil, err
	}
	return newE2SmKpmPdu.GetIndicationMessageFormats().GetIndicationMessageFormat1(), nil
}

func Test_perEncodingE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1, err := createIndicationMessageFormat1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(imf1, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%v", hex.Dump(per))

	result := e2smkpmv2.E2SmKpmIndicationMessageFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER - decoded\n%v", &result)
	assert.Equal(t, imf1.GetGranulPeriod().GetValue(), result.GetGranulPeriod().GetValue())
	assert.Equal(t, imf1.GetSubscriptId().GetValue(), result.GetSubscriptId().GetValue())
	assert.Equal(t, imf1.GetCellObjId().GetValue(), result.GetCellObjId().GetValue())
	assert.DeepEqual(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, imf1.GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.Equal(t, imf1.GetMeasData().GetValue()[0].GetIncompleteFlag().Number(), result.GetMeasData().GetValue()[0].GetIncompleteFlag().Number())
	assert.Equal(t, imf1.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	//assert.Equal(t, imf1.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetReal(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetReal())
	assert.Equal(t, imf1.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perE2SmKpmIndicationMessageFormat1CompareBytes(t *testing.T) {

	imf1, err := createIndicationMessageFormat1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(imf1, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsgF1noReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
