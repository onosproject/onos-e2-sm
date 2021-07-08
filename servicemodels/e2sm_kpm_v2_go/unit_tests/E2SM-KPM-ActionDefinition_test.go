// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerADF1 string = "00000000  00 0c 00 00 03 6f 6e 66  00 00 40 20 74 72 69 61  |.....onf..@ tria|\n" +
	"00000010  6c 01 3f ff e0 21 22 23  40 40 01 02 03 00 0a 7c  |l.?..!\"#@@.....||\n" +
	"00000020  0f 00 0f 00 01 70 00 00  fa 00 00 04 00 00 7a 00  |.....p........z.|\n" +
	"00000030  01 c7 00 03 14 00 14 40  30 38                    |.......@08|"
var refPerADF2 string = "00000000  00 0c 20 06 53 6f 6d 65  55 45 00 00 03 6f 6e 66  |.. .SomeUE...onf|\n" +
	"00000010  00 00 40 20 74 72 69 61  6c 01 3f ff e0 21 22 23  |..@ trial.?..!\"#|\n" +
	"00000020  40 40 01 02 03 00 0a 7c  0f 00 0f 00 01 70 00 00  |@@.....|.....p..|\n" +
	"00000030  fa 00 00 04 00 00 7a 00  01 c7 00 03 14 00 14 40  |......z........@|\n" +
	"00000040  30 38                                             |08|"
var refPerADF3 string = "00000000  00 0c 40 00 03 6f 6e 66  00 00 00 40 74 72 69 61  |..@..onf...@tria|\n" +
	"00000010  6c 00 00 48 21 02 00 c9  00 14 40 30 38           |l..H!.....@08|"

func createE2SmKpmActionDefinitionFormat1() (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	var ricStyleType int32 = 12
	var cellObjID string = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"

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
	measInfoItem, err := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)
	if err != nil {
		return nil, err
	}
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
	var cellObjID string = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"

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
	measInfoItem, err := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)
	if err != nil {
		return nil, err
	}

	measInfoList := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := pdubuilder.CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	if err != nil{
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
	var cellObjID string = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"

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

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmActionDefinition{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER - decoded\n%v", result)
}

func Test_perE2SmKpmActionDefinitionF1CompareBytes(t *testing.T) {

	actionDefFormat1, err := createE2SmKpmActionDefinitionFormat1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format1) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerADF3)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingE2SmKpmActionDefinitionF2(t *testing.T) {

	actionDefFormat2, err := createE2SmKpmActionDefinitionFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat2, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format2) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmActionDefinition{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format2) PER - decoded\n%v", result)
}

func Test_perE2SmKpmActionDefinitionF2CompareBytes(t *testing.T) {

	actionDefFormat2, err := createE2SmKpmActionDefinitionFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat2, "valueExt")
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

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat3, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmActionDefinition{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER - decoded\n%v", result)
}

func Test_perE2SmKpmActionDefinitionF3CompareBytes(t *testing.T) {

	actionDefFormat3, err := createE2SmKpmActionDefinitionFormat3()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat3, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition (Format3) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerADF3)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
