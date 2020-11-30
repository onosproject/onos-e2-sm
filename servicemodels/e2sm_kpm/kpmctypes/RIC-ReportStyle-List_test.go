// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createRicReportStyleItem() *e2sm_kpm_ies.RicReportStyleList {

	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56

	ricReportStyleItem := &e2sm_kpm_ies.RicReportStyleList{
		RicReportStyleType: &e2sm_kpm_ies.RicStyleType{
			Value: ricReportStyleType, //int32
		},
		RicReportStyleName: &e2sm_kpm_ies.RicStyleName{
			Value: ricReportStyleName, //string
		},
		RicIndicationHeaderFormatType: &e2sm_kpm_ies.RicFormatType{
			Value: ricIndicationHeaderFormatType, //int32
		},
		RicIndicationMessageFormatType: &e2sm_kpm_ies.RicFormatType{
			Value: ricIndicationMessageFormatType, //int32
		},
	}

	return ricReportStyleItem
}

func Test_xerEncodeRicReportStyleItem(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	xer, err := xerEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 308, len(xer))
	t.Logf("RIC-ReportStyle-List XER\n%s", string(xer))
}

func Test_xerDecodeRicReportStyleItem(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	xer, err := xerEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 308, len(xer))
	t.Logf("RIC-ReportStyle-List XER\n%s", string(xer))

	result, err := xerDecodeRicReportStyleItem(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricReportStyleItem.RicReportStyleType.Value, result.RicReportStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicReportStyleName.Value, result.RicReportStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationHeaderFormatType.Value, result.RicIndicationHeaderFormatType.Value, "Encoded and decoded RicIndicationHeaderFormatType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationMessageFormatType.Value, result.RicIndicationMessageFormatType.Value, "Encoded and decoded RicIndicationMessageFormatType values are not the same")
}

func Test_perEncodeRicReportStyleItem(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	per, err := perEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("RIC-ReportStyle-List PER\n%s", string(per))
}

func Test_perDecodeRicReportStyleItem(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	per, err := perEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("RIC-ReportStyle-List PER\n%s", string(per))

	result, err := perDecodeRicReportStyleItem(per)
	assert.NilError(t, err)
	assert.Equal(t, ricReportStyleItem.RicReportStyleType.Value, result.RicReportStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicReportStyleName.Value, result.RicReportStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationHeaderFormatType.Value, result.RicIndicationHeaderFormatType.Value, "Encoded and decoded RicIndicationHeaderFormatType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationMessageFormatType.Value, result.RicIndicationMessageFormatType.Value, "Encoded and decoded RicIndicationMessageFormatType values are not the same")
}
