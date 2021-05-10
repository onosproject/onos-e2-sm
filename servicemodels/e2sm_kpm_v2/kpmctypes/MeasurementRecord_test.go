// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementRecord() (*e2sm_kpm_v2.MeasurementRecord, error) {
	res := &e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	res.Value = append(res.Value, item1)

	item2 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Real{
			Real: 22,
		},
	}
	res.Value = append(res.Value, item2)

	item3 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	res.Value = append(res.Value, item3)

	if err := res.Validate(); err != nil {
		return nil, err
	}
	return res, nil
}

func Test_xerEncodeMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementRecord(mr)
	assert.NilError(t, err)
	assert.Equal(t, 140, len(xer))
	t.Logf("MeasurementRecord XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementRecord(mr)
	assert.NilError(t, err)
	assert.Equal(t, 140, len(xer))
	t.Logf("MeasurementRecord XER\n%s", string(xer))

	result, err := xerDecodeMeasurementRecord(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	assert.Equal(t, 3, len(result.GetValue()))
	t.Logf("MeasurementRecord XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementRecord(mr)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))
}

func Test_perDecodeMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementRecord(mr)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	result, err := perDecodeMeasurementRecord(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	assert.Equal(t, 3, len(result.GetValue()))
	t.Logf("MeasurementRecord PER - decoded\n%v", result)
}
