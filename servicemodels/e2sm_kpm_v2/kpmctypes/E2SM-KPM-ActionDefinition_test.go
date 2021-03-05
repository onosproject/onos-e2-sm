// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMActionDefinition() *e2sm_kpm_v2.E2SmKpmActionDefinition {

	var ricStyleType int32 = 12
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
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum int32 = 0
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var preLabel int32 = 0
	var startEndIndication int32 = 1
	var measurementName string = "trial"

	labelInfoItem, _ := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, fiveQI,
		qci, qciMax, qciMin, arpMax, arpMin, bitrateRange, layerMuMimo, sum,
		distX, distY, distZ, preLabel, startEndIndication)

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, _ := pdubuilder.CreateMeasurementType_MeasName(measurementName)
	measInfoItem, _ := pdubuilder.CreateMeasurementInfoItem(*measName, labelInfoList)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinition, _ := pdubuilder.CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmActionDefinition(ricStyleType, *actionDefinition)

	return newE2SmKpmPdu
}

func Test_xerEncodeE2SmKpmActionDefinition(t *testing.T) {

	actionDef := createE2SMKPMActionDefinition()

	xer, err := xerEncodeE2SmKpmActionDefinition(actionDef)
	assert.NilError(t, err)
	assert.Equal(t, 1884, len(xer))
	t.Logf("E2SmKpmActionDefinition XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmActionDefinition(t *testing.T) {

	actionDef := createE2SMKPMActionDefinition()

	xer, err := xerEncodeE2SmKpmActionDefinition(actionDef)
	assert.NilError(t, err)
	assert.Equal(t, 1884, len(xer))
	t.Logf("E2SmKpmActionDefinition XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmActionDefinition(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmActionDefinition XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmActionDefinition(t *testing.T) {

	actionDef := createE2SMKPMActionDefinition()

	per, err := perEncodeE2SmKpmActionDefinition(actionDef)
	assert.NilError(t, err)
	assert.Equal(t, 59, len(per))
	t.Logf("E2SmKpmActionDefinition PER\n%s", string(per))
}

//func Test_perDecodeE2SmKpmActionDefinition(t *testing.T) {
//
//	actionDef := createE2SMKPMActionDefinition()
//
//	per, err := perEncodeE2SmKpmActionDefinition(actionDef)
//	assert.NilError(t, err)
//	assert.Equal(t, 59, len(per))
//	t.Logf("E2SmKpmActionDefinition PER\n%s", string(per))
//
//	result, err := perDecodeE2SmKpmActionDefinition(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2SmKpmActionDefinition PER - decoded\n%s", result)
//}
