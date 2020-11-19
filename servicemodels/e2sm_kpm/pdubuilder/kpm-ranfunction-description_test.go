// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	//"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmRanfunctionDescriptionMsg(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 0
	var ricEventStyleType int32 = 0
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 0
	var ricReportStyleType int32 = 0
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 0
	var ricIndicationMessageFormatType int32 = 0
	newE2SmKpmPdu, err := CreateE2SmKpmRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid, ranFunctionDescription,
		ranFunctionInstance, ricEventStyleType, ricEventStyleName, ricEventFormatType, ricReportStyleType, ricReportStyleName,
		ricIndicationHeaderFormatType, ricIndicationMessageFormatType)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)

	//kpmctypes.PerEncodeE2SmKpmIndicationMessage(newE2SmKpmPdu)
}
