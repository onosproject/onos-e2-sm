// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMIndicationMessageFormat1() (*e2sm_kpm_v2.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	var rl float64 = 22.2
	var cellObjID string = "onf"
	var granularity uint32 = 21
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
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName string = "trial"

	labelInfoItem, _ := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd)
	labelInfoItem.MeasLabel.FiveQi = &e2sm_kpm_v2.FiveQi{
		Value: fiveQI,
	}
	labelInfoItem.MeasLabel.QFi = &e2sm_kpm_v2.Qfi{
		Value: qfi,
	}
	labelInfoItem.MeasLabel.QCi = &e2sm_kpm_v2.Qci{
		Value: qci,
	}
	labelInfoItem.MeasLabel.QCimin = &e2sm_kpm_v2.Qci{
		Value: qciMin,
	}
	labelInfoItem.MeasLabel.QCimax = &e2sm_kpm_v2.Qci{
		Value: qciMax,
	}
	labelInfoItem.MeasLabel.ARpmin = &e2sm_kpm_v2.Arp{
		Value: arpMin,
	}
	labelInfoItem.MeasLabel.ARpmax = &e2sm_kpm_v2.Arp{
		Value: arpMax,
	}
	labelInfoItem.MeasLabel.BitrateRange = bitrateRange
	labelInfoItem.MeasLabel.LayerMuMimo = layerMuMimo
	labelInfoItem.MeasLabel.DistBinX = distX
	labelInfoItem.MeasLabel.DistBinY = distY
	labelInfoItem.MeasLabel.DistBinZ = distZ
	labelInfoItem.MeasLabel.StartEndInd = startEndIndication
	labelInfoItem.MeasLabel.PreLabelOverride = e2sm_kpm_v2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	labelInfoItem.MeasLabel.SUm = e2sm_kpm_v2.SUM_SUM_TRUE

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	measInfoItem, _ := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, _ := pdubuilder.CreateMeasurementDataItem(&measRecord)
	measDataItem.IncompleteFlag = e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE

	measData := e2sm_kpm_v2.MeasurementData{
		Value: make([]*e2sm_kpm_v2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, cellObjID, &measInfoList, &measData)
	newE2SmKpmPdu.GetIndicationMessageFormat1().GranulPeriod.Value = granularity
	if err := newE2SmKpmPdu.Validate(); err != nil {
		return nil, err
	}
	return newE2SmKpmPdu, nil
}

func createE2SMKPMIndicationMessageFormat2() (*e2sm_kpm_v2.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	var rl float64 = 22.2
	var cellObjID string = "onf"
	var granularity uint32 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"
	var ueID string = "SomeUE"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, _ := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))

	mci, _ := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)

	mcl := &e2sm_kpm_v2.MatchingCondList{
		Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)

	mUEIDitem, _ := pdubuilder.CreateMatchingUEIDItem(ueID)

	mUEIDlist := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	mUEIDlist.Value = append(mUEIDlist.Value, mUEIDitem)

	measCondUEIDItem, _ := pdubuilder.CreateMeasurementCondUEIDItem(measName, mcl, mUEIDlist)

	measCondUEIDList := e2sm_kpm_v2.MeasurementCondUeidList{
		Value: make([]*e2sm_kpm_v2.MeasurementCondUeidItem, 0),
	}
	measCondUEIDList.Value = append(measCondUEIDList.Value, measCondUEIDItem)

	measRecord := e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, _ := pdubuilder.CreateMeasurementDataItem(&measRecord)
	measDataItem.IncompleteFlag = e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE

	measData := e2sm_kpm_v2.MeasurementData{
		Value: make([]*e2sm_kpm_v2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessageFormat2(subscriptionID, cellObjID, &measCondUEIDList, &measData)
	newE2SmKpmPdu.GetIndicationMessageFormat2().GranulPeriod.Value = granularity
	if err := newE2SmKpmPdu.Validate(); err != nil {
		return nil, err
	}
	return newE2SmKpmPdu, nil
}

func Test_xerEncodeE2SmKpmIndicationMessage(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	//assert.Equal(t, 2391, len(xer))
	t.Logf("E2SmKpmIndicationMessage (Format1) XER\n%s", string(xer))

	im, err = createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	xer, err = XerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	//assert.Equal(t, 1848, len(xer))
	t.Logf("E2SmKpmIndicationMessage (Format2) XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmIndicationMessage(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	//assert.Equal(t, 2391, len(xer))
	t.Logf("E2SmKpmIndicationMessage (Format1) XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmIndicationMessage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format1) XER - decoded\n%s", result)

	im, err = createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	xer, err = XerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	//assert.Equal(t, 1848, len(xer))
	t.Logf("E2SmKpmIndicationMessage (Format2) XER\n%s", string(xer))

	result, err = XerDecodeE2SmKpmIndicationMessage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format2) XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmIndicationMessage(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	assert.Equal(t, 77, len(per))
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%s", string(per))

	im, err = createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err = PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	assert.Equal(t, 56, len(per))
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmIndicationMessage(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	per, err := PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	assert.Equal(t, 77, len(per))
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%s", string(per))

	result, err := PerDecodeE2SmKpmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER - decoded\n%s", result)

	im, err = createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err = PerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	assert.Equal(t, 56, len(per))
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%s", string(per))

	result, err = PerDecodeE2SmKpmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER - decoded\n%s", result)
}
