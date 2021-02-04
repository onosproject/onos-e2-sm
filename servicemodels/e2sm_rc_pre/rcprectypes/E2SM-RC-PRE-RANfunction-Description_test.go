// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func createCommonE2SmRcPreRanfunctionDescription() (*e2sm_rc_pre_ies.E2SmRcPreRanfunctionDescription, error) {

	var ranFunctionShortName = "ORAN-E2SM-RC-PRE"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
	if err != nil {
		return nil, err
	}

	return newE2SmRcPrePdu, nil
}

func Test_xerEncodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	xer, err := xerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1290, len(xer))
	t.Logf("E2SM-RC-PRE-RANfunction-Description XER\n%s", string(xer))
}

func Test_xerDecodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	xer, err := xerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1290, len(xer))
	t.Logf("E2SM-RC-PRE-RANfunction-Description XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	per, err := PerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 76, len(per))
	t.Logf("E2SM-RC-PRE-RANfunction-Description PER\n%x", string(per))
}

func Test_perDecodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	per, err := PerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 76, len(per))
	t.Logf("E2SM-RC-PRE-RANfunction-Description PER\n%s", string(per))

	result, err := PerDecodeE2SmRcPreRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
