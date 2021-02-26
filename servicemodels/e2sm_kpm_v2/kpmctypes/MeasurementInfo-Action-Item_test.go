// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoActionItem() *e2sm_kpm_v2.MeasurementInfoActionItem {

	return &e2sm_kpm_v2.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2.MeasurementTypeId{
			Value: 21,
		},
		MeasName: &e2sm_kpm_v2.MeasurementTypeName{
			Value: "onf",
		},
	}
}

func Test_xerEncodeMeasurementInfoActionItem(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	xer, err := xerEncodeMeasurementInfoActionItem(miai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoActionItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoActionItem(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	xer, err := xerEncodeMeasurementInfoActionItem(miai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoActionItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoActionItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeMeasurementInfoActionItem(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	per, err := perEncodeMeasurementInfoActionItem(miai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoActionItem PER\n%s", string(per))
}

func Test_perDecodeMeasurementInfoActionItem(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	per, err := perEncodeMeasurementInfoActionItem(miai)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoActionItem PER\n%s", string(per))

	result, err := perDecodeMeasurementInfoActionItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
