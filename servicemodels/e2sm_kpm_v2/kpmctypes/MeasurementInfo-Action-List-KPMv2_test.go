// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoActionList() (*e2sm_kpm_v2.MeasurementInfoActionList, error) {

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

	if err := res.Validate(); err != nil {
		return nil, err
	}
	return res, nil
}

func Test_xerEncodeMeasurementInfoActionList(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	//assert.Equal(t, 191, len(xer))
	t.Logf("MeasurementInfoActionList XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoActionList(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	//assert.Equal(t, 191, len(xer))
	t.Logf("MeasurementInfoActionList XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoActionList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoActionList XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementInfoActionList(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("MeasurementInfoActionList PER\n%v", hex.Dump(per))
}

func Test_perDecodeMeasurementInfoActionList(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoActionList(mial)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("MeasurementInfoActionList PER\n%v", hex.Dump(per))

	result, err := perDecodeMeasurementInfoActionList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoActionList PER - decoded\n%v", result)
}
