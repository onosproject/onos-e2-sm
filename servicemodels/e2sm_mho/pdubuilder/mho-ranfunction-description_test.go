// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	//mhoctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
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
	retsl := make([]*e2sm_mho.RicEventTriggerStyleList, 0)
	retsi := CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_mho.RicReportStyleList, 0)
	rrsi := CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmMhoPdu, err := CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription, retsl, rrsl)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)
	if newE2SmMhoPdu != nil {
		newE2SmMhoPdu.RanFunctionName.RanFunctionInstance = ranFunctionInstance
	}

	per, err := mhoctypes.PerEncodeE2SmMhoRanfunctionDescription(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))
}
