// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreRanFunctionDescriptionMsg(t *testing.T) {
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
	retsl := make([]*e2smrcprev2.RicEventTriggerStyleList, 0)
	retsi := CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2smrcprev2.RicReportStyleList, 0)
	rrsi := CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmRcPrePdu, err := CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetRanFunctionName().SetRanFunctionInstance(ranFunctionInstance)
	newE2SmRcPrePdu.SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl)

	per, err := encoder.PerEncodeE2SmRcPreRanFunctionDescription(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRcPreRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Decoded RanFunctionDescription from PER is \n%v", result)
}

func TestE2SmRcPreRanFunctionDescriptionMsgExcludeOptionalFields(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	newE2SmRcPrePdu, err := CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetRanFunctionName().SetRanFunctionInstance(ranFunctionInstance)

	per, err := encoder.PerEncodeE2SmRcPreRanFunctionDescription(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRcPreRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Decoded RanFunctionDescription from PER is \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())
}

func Test_DecodeE2SmRcPreRanfunctionDescriptionBytes(t *testing.T) {
	t.Logf("Obtained bytes to decode are\n%v", hex.Dump(ranFuncDescBytesRadysis))
	rfd, err := encoder.PerDecodeE2SmRcPreRanFunctionDescription(ranFuncDescBytesRadysis)
	assert.NilError(t, err)

	assert.Equal(t, "ORAN-E2SM-RC-PRE", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "RC-PRE", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.100", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(0), rfd.GetRanFunctionName().GetRanFunctionInstance())
	assert.Equal(t, 1, len(rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()))
	assert.Equal(t, "PCI and NRT update for gNB", rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()[0].GetRicReportStyleName().GetValue())
	assert.Equal(t, int32(1), rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()[0].GetRicIndicationMessageFormatType().GetValue())
	assert.Equal(t, 1, len(rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()))
	assert.Equal(t, "RC-PRE-trigger", rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()[0].GetRicEventTriggerStyleName().GetValue())
	assert.Equal(t, int32(1), rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()[0].GetRicEventTriggerFormatType().GetValue())
}

var ranFuncDescBytesRadysis = []byte{
	0x00, 0xf0, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x2d, 0x50,
	0x52, 0x45, 0x00, 0x00, 0x1a, 0x31, 0x2e, 0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31,
	0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x32, 0x2e, 0x31, 0x30, 0x30,
	0x02, 0x80, 0x52, 0x43, 0x2d, 0x50, 0x52, 0x45, 0x60, 0x00, 0x01, 0x06, 0x80, 0x52, 0x43, 0x2d,
	0x50, 0x52, 0x45, 0x2d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x00, 0x01, 0x00, 0x01, 0x0c,
	0x80, 0x50, 0x43, 0x49, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x4e, 0x52, 0x54, 0x20, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x4e, 0x42, 0x00, 0x01, 0x00, 0x01}
