// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func createCommonE2SmRcPreRanfunctionDescription() (*e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription, error) {

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

	retsl := make([]*e2sm_rc_pre_v2.RicEventTriggerStyleList, 0)
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_rc_pre_v2.RicReportStyleList, 0)
	rrsi := pdubuilder.CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription, retsl, rrsl)
	if err != nil {
		return nil, err
	}

	if newE2SmRcPrePdu != nil {
		newE2SmRcPrePdu.RanFunctionName.RanFunctionInstance = ranFunctionInstance
	}

	return newE2SmRcPrePdu, nil
}

func Test_XerEncodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	xer, err := XerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1290, len(xer))
	t.Logf("E2SM-RC-PRE-RANfunction-Description XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	xer, err := XerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 1290, len(xer))
	t.Logf("E2SM-RC-PRE-RANfunction-Description XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_PerEncodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	per, err := PerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 74, len(per))
	t.Logf("E2SM-RC-PRE-RANfunction-Description PER\n%x", hex.Dump(per))
}

func Test_PerDecodeE2SmRcPreRanfunctionDescription(t *testing.T) {

	ranfunctionDescriptionMsg, err := createCommonE2SmRcPreRanfunctionDescription()
	assert.NilError(t, err, "Something went wrong when the E2SmRcPreRanfunctionDescription was created")

	per, err := PerEncodeE2SmRcPreRanfunctionDescription(ranfunctionDescriptionMsg)
	assert.NilError(t, err)
	assert.Equal(t, 74, len(per))
	t.Logf("E2SM-RC-PRE-RANfunction-Description PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
