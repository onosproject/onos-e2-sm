// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createCommonE2SmKpmRanfunctionDescription() (*e2sm_kpm_ies.E2SmKpmRanfunctionDescription, error) {

	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 1
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
	if err != nil {
		return nil, err
	}

	return newE2SmKpmPdu, nil
}

func Test_xerEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmKpmRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmKpmRanfunctionDescription was created")

	xer, err := xerEncodeE2SmKpmRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1265, len(xer))
	t.Logf("E2SM-KPM-RANfunction-Description XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmKpmRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmKpmRanfunctionDescription was created")

	xer, err := xerEncodeE2SmKpmRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1265, len(xer))
	t.Logf("E2SM-KPM-RANfunction-Description XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeE2SmKpmRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmKpmRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmKpmRanfunctionDescription was created")

	per, err := perEncodeE2SmKpmRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(per))
	t.Logf("E2SM-KPM-RANfunction-Description PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmKpmRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmKpmRanfunctionDescription was created")

	per, err := perEncodeE2SmKpmRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(per))
	t.Logf("E2SM-KPM-RANfunction-Description PER\n%s", string(per))

	result, err := perDecodeE2SmKpmRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
