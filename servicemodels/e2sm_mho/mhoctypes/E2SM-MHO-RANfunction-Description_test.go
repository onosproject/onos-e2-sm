// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoRanfunctionDescriptionMsg() (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {

	var ranFunctionShortName = "ORAN-E2SM-MHO"
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

	retsl := make([]*e2sm_mho.RicEventTriggerStyleList, 0)
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_mho.RicReportStyleList, 0)
	rrsi := pdubuilder.CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription, retsl, rrsl)
	if err != nil {
		return nil, err
	}

	if newE2SmMhoPdu != nil {
		newE2SmMhoPdu.RanFunctionName.RanFunctionInstance = ranFunctionInstance
	}

	return newE2SmMhoPdu, nil

	//e2SmMhoRanfunctionDescription := e2sm_mho.E2SmMhoRanfunctionDescription{
	//	RanFunctionName:          nil,
	//	RicEventTriggerStyleList: make([]*e2sm_mho.RicEventTriggerStyleList, 0), //ToDo - Check if protobuf structure is implemented correctly (mainly naming)
	//	RicReportStyleList:       make([]*e2sm_mho.RicReportStyleList, 0),       //ToDo - Check if protobuf structure is implemented correctly (mainly naming)

	//}

	//if err := e2SmMhoRanfunctionDescription.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmMhoRanfunctionDescription %s", err.Error())
	//}
	//return &e2SmMhoRanfunctionDescription, nil
}

func Test_XerEncodingE2SmMhoRanfunctionDescription(t *testing.T) {

	e2SmMhoRanfunctionDescription, err := createE2SmMhoRanfunctionDescriptionMsg()
	assert.NilError(t, err, "Error creating E2SmMhoRanfunctionDescription PDU")

	xer, err := XerEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)
	assert.NilError(t, err)
	assert.Equal(t, 1275, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("E2SmMhoRanfunctionDescription XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoRanfunctionDescription XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoRanfunctionDescription.GetRanFunctionName(), result.GetRanFunctionName())

	assert.Equal(t, 1, len(result.GetRicEventTriggerStyleList())) //ToDo - adjust length of a list
	//assert.DeepEqual(t, e2SmMhoRanfunctionDescription.GetRicEventTriggerStyleList(), result.GetRicEventTriggerStyleList())
	assert.Equal(t, 1, len(result.GetRicReportStyleList())) //ToDo - adjust length of a list
	//assert.DeepEqual(t, e2SmMhoRanfunctionDescription.GetRicReportStyleList(), result.GetRicReportStyleList())

}

func Test_PerEncodingE2SmMhoRanfunctionDescription(t *testing.T) {

	e2SmMhoRanfunctionDescription, err := createE2SmMhoRanfunctionDescriptionMsg()
	assert.NilError(t, err, "Error creating E2SmMhoRanfunctionDescription PDU")

	per, err := PerEncodeE2SmMhoRanfunctionDescription(e2SmMhoRanfunctionDescription)
	assert.NilError(t, err)
	assert.Equal(t, 73, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2SmMhoRanfunctionDescription PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoRanfunctionDescription PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoRanfunctionDescription.GetRanFunctionName(), result.GetRanFunctionName())

	assert.Equal(t, 1, len(result.GetRicEventTriggerStyleList())) //ToDo - adjust length of a list
	//assert.DeepEqual(t, e2SmMhoRanfunctionDescription.GetRicEventTriggerStyleList(), result.GetRicEventTriggerStyleList())
	assert.Equal(t, 1, len(result.GetRicReportStyleList())) //ToDo - adjust length of a list
	//assert.DeepEqual(t, e2SmMhoRanfunctionDescription.GetRicReportStyleList(), result.GetRicReportStyleList())

}
