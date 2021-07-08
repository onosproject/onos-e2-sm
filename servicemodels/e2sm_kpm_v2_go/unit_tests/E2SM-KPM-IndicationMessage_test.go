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

var refPerIndMsg1 = "00000000  0e 80 30 38 00 00 03 6f  6e 66 00 14 00 00 40 20  |..08...onf....@ |\n" +
	"00000010  74 72 69 61 6c 01 3f ff  e0 21 22 23 40 40 01 02  |trial.?..!\"#@@..|\n" +
	"00000020  03 00 0a 7c 0f 00 0f 00  01 72 40 00 fa 00 00 04  |...|.....r@.....|\n" +
	"00000030  00 00 7a 00 01 c7 00 03  14 00 00 00 40 03 08 30  |..z.........@..0|\n" +
	"00000040  39 44 09 80 d0 16 33 33  33 33 33 33 00           |9D....333333.|"
var refPerIndMsg2 = "00000000  2d 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |-08...onf....@ t|\n" +
	"00000010  72 69 61 6c 00 00 48 21  02 00 c9 00 00 00 06 53  |rial..H!.......S|\n" +
	"00000020  6f 6d 65 55 45 00 00 40  03 08 30 39 44 09 80 d0  |omeUE..@..09D...|\n" +
	"00000030  16 33 33 33 33 33 33 00                           |.333333.|"

func createE2SMKPMIndicationMessageFormat1() (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	var rl float64 = 22.2
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

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	measInfoItem, _ := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)

	measInfoList := e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem, _ := pdubuilder.CreateMeasurementDataItem(&measRecord)
	measDataItem.IncompleteFlag = &incf

	measData := e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, &granularity, &cellObjID, &measInfoList, &measData)
	//if err := newE2SmKpmPdu.Validate(); err != nil {
	//	return nil, err
	//}
	return newE2SmKpmPdu, nil
}

func createE2SMKPMIndicationMessageFormat2() (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	var integer int64 = 12345
	var rl float64 = 22.2
	var cellObjID string = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"
	var ueID string = "SomeUE"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, _ := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))

	mci, _ := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)

	mUEIDitem, _ := pdubuilder.CreateMatchingUEIDItem([]byte(ueID))

	mUEIDlist := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	mUEIDlist.Value = append(mUEIDlist.Value, mUEIDitem)

	measCondUEIDItem, _ := pdubuilder.CreateMeasurementCondUEIDItem(measName, mcl, mUEIDlist)

	measCondUEIDList := e2sm_kpm_v2_go.MeasurementCondUeidList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondUeidItem, 0),
	}
	measCondUEIDList.Value = append(measCondUEIDList.Value, measCondUEIDItem)

	measRecord := e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measDataItem, _ := pdubuilder.CreateMeasurementDataItem(&measRecord)
	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem.IncompleteFlag = &incf

	measData := e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessageFormat2(subscriptionID, &granularity, &cellObjID, &measCondUEIDList, &measData)
	//if err := newE2SmKpmPdu.Validate(); err != nil {
	//	return nil, err
	//}
	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmIndicationMessage1(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*im, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER - decoded\n%v", result)
}

func Test_perE2SmKpmIndicationMessage1CompareBytes(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*im, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format1) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsg1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingE2SmKpmIndicationMessage2(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*im, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER - decoded\n%v", result)
}

func Test_perE2SmKpmIndicationMessage2CompareBytes(t *testing.T) {

	im, err := createE2SMKPMIndicationMessageFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*im, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessage (Format2) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsg2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
