// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	kpmctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SmKpmRanfunctionDescription(t *testing.T) {
	e2smKpmRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-KPM-RANfunction-Description.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmKpmPdu, err := kpmctypes.XerDecodeE2SmKpmRanfunctionDescription(e2smKpmRanfunctionDescription)
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "oid123", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 1, len(*ricReportList))
}

func Test_DecodeE2SmKpmRanfunctionDescriptionBytes(t *testing.T) {
	rfd, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(ranFuncDescBytes)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	assert.Equal(t, "onf", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "someDescription", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "oid123", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, 21, int(rfd.GetRanFunctionName().GetRanFunctionInstance()))

	assert.Equal(t, 1, len(rfd.GetRicReportStyleList()))
	assert.Equal(t, 1, len(rfd.GetRicEventTriggerStyleList()))
}

var ranFuncDescBytes = []byte{
	0x74, 0x04, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x05, 0x6f, 0x69, 0x64, 0x31, 0x32, 0x33,
	0x07, 0x00, 0x73, 0x6f, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x01, 0x15,
	0x00, 0x00, 0x43, 0x00, 0x21, 0x22, 0x23, 0x00, 0xd4, 0xbc, 0x08, 0x80, 0x30, 0x39, 0x20, 0x1a, 0x85, 0x00, 0x00,
	0x00, 0x00, 0x03, 0x4f, 0x4e, 0x46, 0x00, 0x21, 0x22, 0x23, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x01, 0x0b, 0x01,
	0x00, 0x6f, 0x6e, 0x66, 0x01, 0x0f, 0x00, 0x01, 0x0b, 0x01, 0x00, 0x6f, 0x6e, 0x66, 0x01, 0x0f, 0x00, 0x00, 0x41,
	0xa0, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x00, 0x00, 0x17, 0x01,
	0x2f, 0x01, 0x18}
