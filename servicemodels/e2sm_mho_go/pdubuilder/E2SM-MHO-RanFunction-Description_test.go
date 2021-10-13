// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoRanfunctionDescriptionMsg(t *testing.T) {
	var ranFunctionShortName = "ONF"
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
	retsl := make([]*e2sm_mho_go.RicEventTriggerStyleList, 0)
	retsi := CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_mho_go.RicReportStyleList, 0)
	rrsi := CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmMhoPdu, err := CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)
	newE2SmMhoPdu.GetRanFunctionName().SetRanFunctionInstance(ranFunctionInstance)
	newE2SmMhoPdu.SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl)

	per, err := encoder.PerEncodeE2SmMhoRanFunctionDescription(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoRanFunctionDescription(per)
	assert.NilError(t, err)
	t.Logf("Decoded RanFunctionDescription is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
