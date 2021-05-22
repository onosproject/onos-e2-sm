// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementTypeID() *e2sm_kpm_v2.MeasurementTypeId {
	return &e2sm_kpm_v2.MeasurementTypeId{
		Value: 123,
	}
}

func Test_xerEncodeMeasurementTypeID(t *testing.T) {

	mtID := createMeasurementTypeID()

	xer, err := xerEncodeMeasurementTypeID(mtID)
	assert.NilError(t, err)
	//assert.Equal(t, 43, len(xer))
	t.Logf("MeasurementTypeID XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementTypeID(t *testing.T) {

	mtID := createMeasurementTypeID()

	xer, err := xerEncodeMeasurementTypeID(mtID)
	assert.NilError(t, err)
	//assert.Equal(t, 43, len(xer))
	t.Logf("MeasurementTypeID XER\n%s", string(xer))

	result, err := xerDecodeMeasurementTypeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementTypeID XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementTypeID(t *testing.T) {

	mtID := createMeasurementTypeID()

	per, err := perEncodeMeasurementTypeID(mtID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("MeasurementTypeID PER\n%s", string(per))
}

func Test_perDecodeMeasurementTypeID(t *testing.T) {

	mtID := createMeasurementTypeID()

	per, err := perEncodeMeasurementTypeID(mtID)
	assert.NilError(t, err)
	//assert.Equal(t, 3, len(per))
	t.Logf("MeasurementTypeID PER\n%s", string(per))

	result, err := perDecodeMeasurementTypeID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementTypeID PER - decoded\n%s", result)
}
