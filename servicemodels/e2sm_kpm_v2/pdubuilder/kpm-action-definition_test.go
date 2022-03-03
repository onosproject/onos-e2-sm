// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	kpmv2ctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmActionDefinitionFormat1(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
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
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName = "trial"

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd)
	assert.NilError(t, err)
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

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName, &labelInfoList)
	assert.NilError(t, err)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinition, err := CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat1(ricStyleType, actionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func TestE2SmKpmActionDefinitionFormat2(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
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
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName = "trial"

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd)
	assert.NilError(t, err)
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

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName, &labelInfoList)
	assert.NilError(t, err)

	measInfoList := &e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	ueID := "SomeUE"
	actionDefinitionFormat2, err := CreateActionDefinitionFormat2(ueID, actionDefinitionFormat1)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat2(ricStyleType, actionDefinitionFormat2)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func TestE2SmKpmActionDefinitionFormat3(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity uint32 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := CreateTestCondInfo(CreateTestCondTypeRSRP(), tce, CreateTestCondValueEnum(valEnum))
	assert.NilError(t, err)

	mci, err := CreateMatchingCondItemTestCondInfo(tci)
	assert.NilError(t, err)

	mcl := &e2sm_kpm_v2.MatchingCondList{
		Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measCondItem, err := CreateMeasurementCondItem(measName, mcl)
	assert.NilError(t, err)

	measCondList := e2sm_kpm_v2.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, err := CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat3(ricStyleType, actionDefinitionFormat3)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func Test_Ray(t *testing.T) {
	var ricStyleType int32 = 1
	var cellObjID = "138426014550003"
	var granularity uint32 = 1000
	var subscriptionID int64 = 3

	measName1, err := CreateMeasurementTypeMeasName("RRC.ConnEstabAtt.Sum")
	assert.NilError(t, err)
	measInfoItem1, err := CreateMeasurementInfoItem(measName1, nil)
	assert.NilError(t, err)

	measName2, err := CreateMeasurementTypeMeasName("RRC.ConnEstabSucc.Sum")
	assert.NilError(t, err)
	measInfoItem2, err := CreateMeasurementInfoItem(measName2, nil)
	assert.NilError(t, err)

	measName3, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.Sum")
	assert.NilError(t, err)
	measInfoItem3, err := CreateMeasurementInfoItem(measName3, nil)
	assert.NilError(t, err)

	measName4, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.reconfigFail")
	assert.NilError(t, err)
	measInfoItem4, err := CreateMeasurementInfoItem(measName4, nil)
	assert.NilError(t, err)

	measName5, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.HOFail")
	assert.NilError(t, err)
	measInfoItem5, err := CreateMeasurementInfoItem(measName5, nil)
	assert.NilError(t, err)

	measName6, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.Other")
	assert.NilError(t, err)
	measInfoItem6, err := CreateMeasurementInfoItem(measName6, nil)
	assert.NilError(t, err)

	measName7, err := CreateMeasurementTypeMeasName("RRC.Conn.Avg")
	assert.NilError(t, err)
	measInfoItem7, err := CreateMeasurementInfoItem(measName7, nil)
	assert.NilError(t, err)

	measName8, err := CreateMeasurementTypeMeasName("RRC.Conn.Max")
	assert.NilError(t, err)
	measInfoItem8, err := CreateMeasurementInfoItem(measName8, nil)
	assert.NilError(t, err)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem1)
	measInfoList.Value = append(measInfoList.Value, measInfoItem2)
	measInfoList.Value = append(measInfoList.Value, measInfoItem3)
	measInfoList.Value = append(measInfoList.Value, measInfoItem4)
	measInfoList.Value = append(measInfoList.Value, measInfoItem5)
	measInfoList.Value = append(measInfoList.Value, measInfoItem6)
	measInfoList.Value = append(measInfoList.Value, measInfoItem7)
	measInfoList.Value = append(measInfoList.Value, measInfoItem8)

	actionDefinition, err := CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat1(ricStyleType, actionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
	t.Logf("E2SM-KPM-ActionDefinition-Format1 is\n%v", newE2SmKpmPdu)

	per, err := kpmv2ctypes.PerEncodeE2SmKpmActionDefinition(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("PER bytes are\n%v", hex.Dump(per))
}
