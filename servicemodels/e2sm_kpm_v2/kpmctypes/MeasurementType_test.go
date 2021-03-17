// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementType1() (*e2sm_kpm_v2.MeasurementType, error) {
	return pdubuilder.CreateMeasurementTypeMeasName("onf")
}

func createMeasurementType2() (*e2sm_kpm_v2.MeasurementType, error) {
	return pdubuilder.CreateMeasurementTypeMeasID(123)
}

func Test_xerEncodeMeasurementType(t *testing.T) {

	mt1, err := createMeasurementType1()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementType(mt1)
	assert.NilError(t, err)
	assert.Equal(t, 66, len(xer))
	t.Logf("MeasurementType (Name) XER\n%s", string(xer))

	mt2, err := createMeasurementType2()
	assert.NilError(t, err)

	xer, err = xerEncodeMeasurementType(mt2)
	assert.NilError(t, err)
	assert.Equal(t, 62, len(xer))
	t.Logf("MeasurementType (ID) XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementType(t *testing.T) {

	mt1, err := createMeasurementType1()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementType(mt1)
	assert.NilError(t, err)
	assert.Equal(t, 66, len(xer))
	t.Logf("MeasurementType (Name) XER\n%s", string(xer))

	result, err := xerDecodeMeasurementType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementType (Name) XER - decoded\n%s", result)

	mt2, err := createMeasurementType2()
	assert.NilError(t, err)

	xer, err = xerEncodeMeasurementType(mt2)
	assert.NilError(t, err)
	assert.Equal(t, 62, len(xer))
	t.Logf("MeasurementType (ID) XER\n%s", string(xer))

	result, err = xerDecodeMeasurementType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementType (ID) XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementType(t *testing.T) {

	mt1, err := createMeasurementType1()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementType(mt1)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("MeasurementType (Name)  PER\n%s", string(per))

	mt2, err := createMeasurementType2()
	assert.NilError(t, err)

	per, err = perEncodeMeasurementType(mt2)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("MeasurementType (ID)  PER\n%s", string(per))
}

func Test_perDecodeMeasurementType(t *testing.T) {

	mt1, err := createMeasurementType1()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementType(mt1)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("MeasurementType (Name) PER\n%s", string(per))

	result, err := perDecodeMeasurementType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementType (Name) PER - decoded\n%v", result)

	mt2, err := createMeasurementType2()
	assert.NilError(t, err)

	per, err = perEncodeMeasurementType(mt2)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("MeasurementType (ID) PER\n%s", string(per))

	result, err = perDecodeMeasurementType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementType (ID) PER - decoded\n%v", result)
}
