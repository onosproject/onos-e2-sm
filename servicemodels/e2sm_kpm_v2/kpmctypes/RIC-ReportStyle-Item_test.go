// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createRicReportStyleItem() (*e2sm_kpm_v2.RicReportStyleItem, error) {

	measInfo := &e2sm_kpm_v2.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoActionItem, 0),
	}

	item := &e2sm_kpm_v2.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2.MeasurementTypeId{
			Value: 789,
		},
		MeasName: &e2sm_kpm_v2.MeasurementTypeName{
			Value: "SomeMeasName",
		},
	}

	measInfo.Value = append(measInfo.Value, item)

	res := e2sm_kpm_v2.RicReportStyleItem{
		RicReportStyleType: &e2sm_kpm_v2.RicStyleType{
			Value: 123,
		},
		RicReportStyleName: &e2sm_kpm_v2.RicStyleName{
			Value: "SomeName",
		},
		RicActionFormatType: &e2sm_kpm_v2.RicFormatType{
			Value: 456,
		},
		MeasInfoActionList: measInfo,
		RicIndicationHeaderFormatType: &e2sm_kpm_v2.RicFormatType{
			Value: 1,
		},
		RicIndicationMessageFormatType: &e2sm_kpm_v2.RicFormatType{
			Value: 1,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, err
	}
	return &res, nil
}

func Test_xerEncodeRicReportStyleItem(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicReportStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 572, len(xer))
	t.Logf("RicReportStyleItem XER\n%s", string(xer))
}

func Test_xerDecodeRicReportStyleItem(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicReportStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 572, len(xer))
	t.Logf("RicReportStyleItem XER\n%s", string(xer))

	result, err := xerDecodeRicReportStyleItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicReportStyleItem XER - decoded\n%s", result)
}

func Test_perEncodeRicReportStyleItem(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	per, err := perEncodeRicReportStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 39, len(per))
	t.Logf("RicReportStyleItem PER\n%s", string(per))
}

func Test_perDecodeRicReportStyleItem(t *testing.T) {

	item, err := createRicReportStyleItem()
	assert.NilError(t, err)

	per, err := perEncodeRicReportStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 39, len(per))
	t.Logf("RicReportStyleItem PER\n%s", string(per))

	result, err := perDecodeRicReportStyleItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicReportStyleItem PER - decoded\n%v", result)
}
