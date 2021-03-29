// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementDataItem() (*e2sm_kpm_v2.MeasurementDataItem, error) {

	measRecord := &e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item2 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Real{
			Real: 22.2,
		},
	}
	measRecord.Value = append(measRecord.Value, item2)

	item3 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	measDataItem := &e2sm_kpm_v2.MeasurementDataItem{
		MeasRecord: measRecord,
		//IncompleteFlag: e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE,
		IncompleteFlag: -1,
	}

	if err := measDataItem.Validate(); err != nil {
		return nil, err
	}
	return measDataItem, nil
}

func Test_xerEncodeMeasurementDataItem(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementDataItem(mdi)
	assert.NilError(t, err)
	//assert.Equal(t, 262, len(xer))
	assert.Equal(t, 217, len(xer)) // without IncompleteFlag
	t.Logf("MeasurementDataItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementDataItem(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementDataItem(mdi)
	assert.NilError(t, err)
	//assert.Equal(t, 262, len(xer))
	assert.Equal(t, 217, len(xer)) // without IncompleteFlag
	t.Logf("MeasurementDataItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementDataItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//measRecord := result.GetValue()[0]
	//assert.Equal(t, 3, len(measRecord.GetValue()))
	t.Logf("MeasurementDataItem XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementDataItem(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementDataItem(mdi)
	assert.NilError(t, err)
	//assert.Equal(t, 17, len(per))
	assert.Equal(t, 17, len(per)) // without IncompleteFlag
	t.Logf("MeasurementDataItem PER\n%s", string(per))
}

func Test_perDecodeMeasurementDataItem(t *testing.T) {

	mdi, err := createMeasurementDataItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementDataItem(mdi)
	assert.NilError(t, err)
	//assert.Equal(t, 17, len(per))
	assert.Equal(t, 17, len(per)) // without IncompleteFlag
	t.Logf("MeasurementDataItem PER\n%s", string(per))

	result, err := perDecodeMeasurementDataItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MeasurementDataItem PER - decoded\n%v", result)
}
