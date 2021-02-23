// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmActionDefinition(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID string = "onf"
	var granularity int32 = 21
	var subscriptionID int64 = 12345
	plmnID := []byte{0x21, 0x22, 0x23}
	var sst int32 = 14
	var sd int32 = 54
	var fiveQI int32 = 10
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum int32 = 69
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var preLabel int32 = 11
	var startEndIndication int32 = 1
	var measurementName string = "trial"

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, fiveQI,
		qci, qciMax, qciMin, arpMax, arpMin, bitrateRange, layerMuMimo, sum,
		distX, distY, distZ, preLabel, startEndIndication)
	assert.NilError(t, err)

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

	actionDefinition, err := CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinition(ricStyleType, *actionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
