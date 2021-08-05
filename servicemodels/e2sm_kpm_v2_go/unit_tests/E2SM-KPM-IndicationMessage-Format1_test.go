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

var refPerIndMsgF1 = "00000000  74 30 38 00 00 03 6f 6e  66 00 14 00 00 40 20 74  |t08...onf....@ t|" +
	"00000010  72 69 61 6c 01 3f ff e0  21 22 23 40 40 01 02 03  |rial.?..!\"#@@...|" +
	"00000020  00 0a 7c 0f 00 0f 00 01  70 00 00 fa 00 00 04 00  |..|.....p.......|" +
	"00000030  00 7a 00 01 c7 00 03 14  00 00 00 40 03 08 30 39  |.z.........@..09|" +
	"00000040  44 09 80 d9 0d 42 c1 47  ae 14 7b 00              |D....B.G..{.|"

func createIndicationMessageFormat1() (*e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat1, error) {

	var integer int64 = 12345
	var rl float64 = 6789.51
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

	incflg := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord:     &measRecord,
		IncompleteFlag: &incflg,
	}

	measData := &e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessageFormat1(subscriptionID, &granularity, &cellObjID, &measInfoList, measData)
	//if err := newE2SmKpmPdu.Validate(); err != nil {
	//	return nil, err
	//}
	return newE2SmKpmPdu.GetIndicationMessageFormats().GetIndicationMessageFormat1(), nil
}

func Test_perEncodingE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1, err := createIndicationMessageFormat1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*imf1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER - decoded\n%v", result)
}

func Test_perE2SmKpmIndicationMessageFormat1CompareBytes(t *testing.T) {

	imf1, err := createIndicationMessageFormat1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*imf1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsgF1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
