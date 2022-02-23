// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"gotest.tools/assert"
	"testing"
)

func TestMeasurementLabelFull(t *testing.T) {
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
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)
	assert.Assert(t, labelInfoItem != nil)
	t.Logf("Composed (full) MeasurementLabel is \n%v", labelInfoItem.GetMeasLabel())
}

func TestMeasurementLabelPartially(t *testing.T) {
	//plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	//var qfi int32 = 62
	var qci int32 = 15
	//var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	//var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	//var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	//var distX int32 = 123
	//var distY int32 = 456
	//var distZ int32 = 789
	//var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	//startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(nil, sst, sd, &fiveQI, nil, &qci, &qciMax, nil, &arpMax, nil,
		&bitrateRange, &layerMuMimo, nil, nil, nil, nil, nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, labelInfoItem != nil)
	t.Logf("Composed (partially) MeasurementLabel is \n%v", labelInfoItem.GetMeasLabel())
}

func TestE2SmKpmIndicationMessageFormat1(t *testing.T) {
	var integer int64 = 12345
	var rl float64 = 6789
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
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
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START
	var measurementName = "trial"

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName)
	assert.NilError(t, err)
	measInfoItem.SetLabelInfoList(&labelInfoList)

	measInfoList := e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemReal(rl))

	measDataItem, err := CreateMeasurementDataItem(&measRecord)
	assert.NilError(t, err)
	measDataItem.SetIncompleteFlag()

	measData := e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationMessageFormat1(subscriptionID, &measData)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
	newE2SmKpmPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID).SetMeasInfoList(&measInfoList)
	t.Logf("Composed IndicationMessage-Format1 is \n %v \n", newE2SmKpmPdu)
}

func TestE2SmKpmIndicationMessageFormat2(t *testing.T) {
	var integer int64 = 12345
	var rl float64 = 6789
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"
	var ueID = "SomeUE"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := CreateTestCondInfo(CreateTestCondTypeRSRP(), tce, CreateTestCondValueEnum(valEnum))
	assert.NilError(t, err)

	mci, err := CreateMatchingCondItemTestCondInfo(tci)
	assert.NilError(t, err)

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)

	mUEIDitem, err := CreateMatchingUEIDItem([]byte(ueID))
	assert.NilError(t, err)

	mUEIDlist := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	mUEIDlist.Value = append(mUEIDlist.Value, mUEIDitem)

	measCondUEIDItem, err := CreateMeasurementCondUEIDItem(measName, mcl)
	assert.NilError(t, err)
	measCondUEIDItem.SetMatchingUEUDlist(mUEIDlist)

	measCondUEIDList := e2sm_kpm_v2_go.MeasurementCondUeidList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondUeidItem, 0),
	}
	measCondUEIDList.Value = append(measCondUEIDList.Value, measCondUEIDItem)

	measRecord := e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemReal(rl))

	measDataItem, err := CreateMeasurementDataItem(&measRecord)
	assert.NilError(t, err)
	measDataItem.SetIncompleteFlag()

	measData := e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationMessageFormat2(subscriptionID, &measCondUEIDList, &measData)
	assert.NilError(t, err)
	newE2SmKpmPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID)
	assert.Assert(t, newE2SmKpmPdu != nil)
	t.Logf("Composed IndicationMessage-Format2 is \n %v \n", newE2SmKpmPdu)
}
