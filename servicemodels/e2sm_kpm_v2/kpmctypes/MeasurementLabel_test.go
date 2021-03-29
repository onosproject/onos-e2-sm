// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementLabel() *e2sm_kpm_v2.MeasurementLabel {
	return &e2sm_kpm_v2.MeasurementLabel{
		PlmnId: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x1, 0x2, 0x3},
		},
		SliceId: &e2sm_kpm_v2.Snssai{
			SD:  []byte{0x01, 0x02, 0x03},
			SSt: []byte{0x01},
		},
		//FiveQi: &e2sm_kpm_v2.FiveQi{
		//	Value: 23,
		//},
		//QFi: &e2sm_kpm_v2.Qfi{
		//	Value: 52,
		//},
		QCi: &e2sm_kpm_v2.Qci{
			Value: 24,
		},
		QCimax: &e2sm_kpm_v2.Qci{
			Value: 30,
		},
		//QCimin: &e2sm_kpm_v2.Qci{
		//	Value: 1,
		//},
		//ARpmax: &e2sm_kpm_v2.Arp{
		//	Value: 15,
		//},
		ARpmin: &e2sm_kpm_v2.Arp{
			Value: 1,
		},
		BitrateRange:     25,
		LayerMuMimo:      1,
		SUm:              e2sm_kpm_v2.SUM_SUM_TRUE,
		DistBinX:         123,
		DistBinY:         -1,
		DistBinZ:         789,
		PreLabelOverride: -1,
		StartEndInd:      e2sm_kpm_v2.StartEndInd_START_END_IND_END,
	}
}

func Test_xerEncodeMeasurementLabel(t *testing.T) {

	ml := createMeasurementLabel()

	xer, err := xerEncodeMeasurementLabel(ml)
	assert.NilError(t, err)
	//assert.Equal(t, 568, len(xer))
	assert.Equal(t, 401, len(xer)) // without FiveQI, QFi, QCimin, ARPmax, DistBinY and PreLabelOverride
	t.Logf("MeasurementLabel XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementLabel(t *testing.T) {

	ml := createMeasurementLabel()

	xer, err := xerEncodeMeasurementLabel(ml)
	assert.NilError(t, err)
	//assert.Equal(t, 568, len(xer))
	assert.Equal(t, 401, len(xer)) // without FiveQI, QFi, QCimin, ARPmax, DistBinY and PreLabelOverride
	t.Logf("MeasurementLabel XER\n%s", string(xer))

	result, err := xerDecodeMeasurementLabel(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementLabel XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementLabel(t *testing.T) {

	ml := createMeasurementLabel()

	per, err := perEncodeMeasurementLabel(ml)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	assert.Equal(t, 28, len(per)) // without FiveQI, QFi, QCimin, ARPmax, DistBinY and PreLabelOverride
	t.Logf("MeasurementLabel XER\n%s", string(per))
}

func Test_perDecodeMeasurementLabel(t *testing.T) {

	ml := createMeasurementLabel()

	per, err := perEncodeMeasurementLabel(ml)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	assert.Equal(t, 28, len(per)) // without FiveQI, QFi, QCimin, ARPmax, DistBinY and PreLabelOverride
	t.Logf("MeasurementLabel PER\n%s", string(per))

	result, err := perDecodeMeasurementLabel(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementLabel PER - decoded\n%v", result)
}
