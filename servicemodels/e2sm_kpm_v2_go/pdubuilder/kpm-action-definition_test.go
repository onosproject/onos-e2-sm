// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmActionDefinitionFormat1(t *testing.T) {
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
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2smkpmv2.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2smkpmv2.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2smkpmv2.LabelInfoList{
		Value: make([]*e2smkpmv2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName)
	assert.NilError(t, err)
	measInfoItem.SetLabelInfoList(&labelInfoList)

	measInfoList := e2smkpmv2.MeasurementInfoList{
		Value: make([]*e2smkpmv2.MeasurementInfoItem, 0),
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

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2smkpmv2.LabelInfoList{
		Value: make([]*e2smkpmv2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName)
	assert.NilError(t, err)
	measInfoItem.SetLabelInfoList(&labelInfoList)

	measInfoList := &e2smkpmv2.MeasurementInfoList{
		Value: make([]*e2smkpmv2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	ueID := "SomeUE"
	actionDefinitionFormat2, err := CreateActionDefinitionFormat2([]byte(ueID), actionDefinitionFormat1)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat2(ricStyleType, actionDefinitionFormat2)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func TestE2SmKpmActionDefinitionFormat3(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	var valEnum int64 = 201
	tce := e2smkpmv2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := CreateTestCondInfo(CreateTestCondTypeRSRP(), tce, CreateTestCondValueEnum(valEnum))
	assert.NilError(t, err)

	mci, err := CreateMatchingCondItemTestCondInfo(tci)
	assert.NilError(t, err)

	mcl := &e2smkpmv2.MatchingCondList{
		Value: make([]*e2smkpmv2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measCondItem, err := CreateMeasurementCondItem(measName, mcl)
	assert.NilError(t, err)

	measCondList := e2smkpmv2.MeasurementCondList{
		Value: make([]*e2smkpmv2.MeasurementCondItem, 0),
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
	var granularity int64 = 1000
	var subscriptionID int64 = 3

	measName1, err := CreateMeasurementTypeMeasName("RRC.ConnEstabAtt.Sum")
	assert.NilError(t, err)
	measInfoItem1, err := CreateMeasurementInfoItem(measName1)
	assert.NilError(t, err)

	measName2, err := CreateMeasurementTypeMeasName("RRC.ConnEstabSucc.Sum")
	assert.NilError(t, err)
	measInfoItem2, err := CreateMeasurementInfoItem(measName2)
	assert.NilError(t, err)

	measName3, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.Sum")
	assert.NilError(t, err)
	measInfoItem3, err := CreateMeasurementInfoItem(measName3)
	assert.NilError(t, err)

	measName4, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.reconfigFail")
	assert.NilError(t, err)
	measInfoItem4, err := CreateMeasurementInfoItem(measName4)
	assert.NilError(t, err)

	measName5, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.HOFail")
	assert.NilError(t, err)
	measInfoItem5, err := CreateMeasurementInfoItem(measName5)
	assert.NilError(t, err)

	measName6, err := CreateMeasurementTypeMeasName("RRC.ConnReEstabAtt.Other")
	assert.NilError(t, err)
	measInfoItem6, err := CreateMeasurementInfoItem(measName6)
	assert.NilError(t, err)

	measName7, err := CreateMeasurementTypeMeasName("RRC.Conn.Avg")
	assert.NilError(t, err)
	measInfoItem7, err := CreateMeasurementInfoItem(measName7)
	assert.NilError(t, err)

	measName8, err := CreateMeasurementTypeMeasName("RRC.Conn.Max")
	assert.NilError(t, err)
	measInfoItem8, err := CreateMeasurementInfoItem(measName8)
	assert.NilError(t, err)

	measInfoList := e2smkpmv2.MeasurementInfoList{
		Value: make([]*e2smkpmv2.MeasurementInfoItem, 0),
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

	per, err := encoder.PerEncodeE2SmKpmActionDefinition(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("PER bytes are\n%v", hex.Dump(per))

	//comparing CGo bytes and these bytes:
	assert.DeepEqual(t, per, perCGo)

	t.Logf("%s", string(per))
}

var perCGo = []byte{
	0x00, 0x01, 0x00, 0x00, 0x0f, 0x31, 0x33, 0x38, 0x34, 0x32, 0x36, 0x30, 0x31, 0x34, 0x35, 0x35,
	0x30, 0x30, 0x30, 0x33, 0x00, 0x07, 0x00, 0x98, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0xa0, 0x52, 0x52,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2e,
	0x53, 0x75, 0x6d, 0x00, 0xa8, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0xf0, 0x52, 0x52, 0x43,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x00, 0xc0, 0x52, 0x52,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74,
	0x2e, 0x48, 0x4f, 0x46, 0x61, 0x69, 0x6c, 0x00, 0xb8, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x00, 0x58, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x2e, 0x41, 0x76, 0x67, 0x00,
	0x58, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x2e, 0x4d, 0x61, 0x78, 0x40, 0x03, 0xe7,
	0x00, 0x02}
