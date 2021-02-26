// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoActionList() *e2sm_kpm_v2.MeasurementInfoActionList {

	res := &e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	item := &e2sm_kpm_v2.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2.MeasurementTypeId{
			Value: 21,
		},
		MeasName: &e2sm_kpm_v2.MeasurementTypeName{
			Value: "onf",
		},
	}
	res.Value = append(res.Value, item)

	return res
}

func Test_xerEncodeMeasurementInfoActionList(t *testing.T) {

	mial := createMeasurementInfoActionList()

	xer, err := xerEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoActionList XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoActionList(t *testing.T) {

	mial := createMeasurementInfoActionList()

	xer, err := xerEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoActionList XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoActionList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeMeasurementInfoActionList(t *testing.T) {

	mial := createMeasurementInfoActionList()

	per, err := perEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoActionList PER\n%s", string(per))
}

func Test_perDecodeMeasurementInfoActionList(t *testing.T) {

	mial := createMeasurementInfoActionList()

	per, err := perEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoActionList PER\n%s", string(per))

	result, err := perDecodeMeasurementInfoActionList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
