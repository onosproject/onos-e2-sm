// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementRecordItemInteger() *e2sm_kpm_v2.MeasurementRecordItem {
	return &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
}

func createMeasurementRecordItemReal() *e2sm_kpm_v2.MeasurementRecordItem {
	return &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Real{
			Real: 22,
		},
	}
}

func createMeasurementRecordItemNoValue() *e2sm_kpm_v2.MeasurementRecordItem {
	return &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func Test_xerEncodeMeasurementRecordItem(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	xer, err := xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (Integer) XER\n%s", string(xer))

	mri = createMeasurementRecordItemReal()

	xer, err = xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (Real) XER\n%s", string(xer))

	mri = createMeasurementRecordItemNoValue()

	xer, err = xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (NoValue) XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementRecordItem(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	xer, err := xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (Integer) XER\n%s", string(xer))

	result, err := xerDecodeMeasurementRecordItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)

	mri = createMeasurementRecordItemReal()

	xer, err = xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (Real) XER\n%s", string(xer))

	result, err = xerDecodeMeasurementRecordItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)

	mri = createMeasurementRecordItemNoValue()

	xer, err = xerEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementRecordItem (NoValue) XER\n%s", string(xer))

	result, err = xerDecodeMeasurementRecordItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeMeasurementRecordItem(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	per, err := perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (Integer) PER\n%s", string(per))

	mri = createMeasurementRecordItemReal()

	per, err = perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (Real) PER\n%s", string(per))

	mri = createMeasurementRecordItemNoValue()

	per, err = perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (NoValue) PER\n%s", string(per))
}

func Test_perDecodeMeasurementRecordItem(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	per, err := perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (Integer) PER\n%s", string(per))

	result, err := perDecodeMeasurementRecordItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)

	mri = createMeasurementRecordItemReal()

	per, err = perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (Real) PER\n%s", string(per))

	result, err = perDecodeMeasurementRecordItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)

	mri = createMeasurementRecordItemNoValue()

	per, err = perEncodeMeasurementRecordItem(mri)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementRecordItem (NoValue) PER\n%s", string(per))

	result, err = perDecodeMeasurementRecordItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
