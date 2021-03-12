// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationMessage(t *testing.T) {
	var integer int64 = 12345
	var rl float64 = 6789
	var cellObjID string = "onf"
	var granularity int32 = 21
	var subscriptionID int64 = 12345
	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 100
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName string = "trial"

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, fiveQI,
		qci, qciMax, qciMin, arpMax, arpMin, bitrateRange, layerMuMimo,
		distX, distY, distZ, startEndIndication)

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementType_MeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(*measName, labelInfoList)
	assert.NilError(t, err)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItem_Integer(integer))
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItem_NoValue())
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItem_Real(rl))

	measData := e2sm_kpm_v2.MeasurementData{
		Value: make([]*e2sm_kpm_v2.MeasurementRecord, 0),
	}
	measData.Value = append(measData.Value, &measRecord)

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationMessage(subscriptionID, cellObjID, granularity, measInfoList, measData)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
