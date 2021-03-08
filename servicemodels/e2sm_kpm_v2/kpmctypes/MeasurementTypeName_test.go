// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementTypeName() *e2sm_kpm_v2.MeasurementTypeName {
	return &e2sm_kpm_v2.MeasurementTypeName{
		Value: "123",
	}
}

func Test_xerEncodeMeasurementTypeName(t *testing.T) {

	mtn := createMeasurementTypeName()

	xer, err := xerEncodeMeasurementTypeName(mtn)
	assert.NilError(t, err)
	assert.Equal(t, 47, len(xer))
	t.Logf("MeasurementTypeName XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementTypeName(t *testing.T) {

	mtn := createMeasurementTypeName()

	xer, err := xerEncodeMeasurementTypeName(mtn)
	assert.NilError(t, err)
	assert.Equal(t, 47, len(xer))
	t.Logf("MeasurementTypeName XER\n%s", string(xer))

	result, err := xerDecodeMeasurementTypeName(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementTypeName XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementTypeName(t *testing.T) {

	mtn := createMeasurementTypeName()

	per, err := perEncodeMeasurementTypeName(mtn)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("MeasurementTypeName PER\n%s", string(per))
}

func Test_perDecodeMeasurementTypeName(t *testing.T) {

	mtn := createMeasurementTypeName()

	per, err := perEncodeMeasurementTypeName(mtn)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("MeasurementTypeName PER\n%s", string(per))

	result, err := perDecodeMeasurementTypeName(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementTypeName PER - decoded\n%s", result)
}
