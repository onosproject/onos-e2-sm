// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SmRcPreRanfunctionDescription(t *testing.T) {
	e2smRcPreRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-RC-PRE-RANfunction-Description.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmRcPrePdu, err := rcprectypes.XerDecodeE2SmRcPreRanfunctionDescription(e2smRcPreRanfunctionDescription)
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmRcPreRanfunctionDescription(e2SmRcPrePdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "OID123", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 1, len(*ricReportList))
}

func Test_DecodeE2SmRcPreRanfunctionDescriptionBytes(t *testing.T) {
	rfd, err := rcprectypes.PerDecodeE2SmRcPreRanfunctionDescription(ranFuncDescBytesRC)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	assert.Equal(t, "ORAN-E2SM-RC-PRE", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "RC PRE", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "OID123", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, 3, int(rfd.GetRanFunctionName().GetRanFunctionInstance()))
	assert.Equal(t, 1, len(rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()))
	assert.Equal(t, 1, len(rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()))
}

var ranFuncDescBytesRC = []byte{
	0x20, 0xf0, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x2d, 0x50, 0x52,
	0x45, 0x00, 0x00, 0x05, 0x4f, 0x49, 0x44, 0x31, 0x32, 0x33, 0x02, 0x80, 0x52, 0x43, 0x20, 0x50, 0x52,
	0x45, 0x01, 0x03, 0x60, 0x00, 0x01, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63,
	0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x01, 0x05, 0x00, 0x01, 0x01, 0x0c, 0x80, 0x50, 0x43, 0x49,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x4e, 0x52, 0x54, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x65, 0x4e, 0x42, 0x01, 0x01, 0x01, 0x01,
}
