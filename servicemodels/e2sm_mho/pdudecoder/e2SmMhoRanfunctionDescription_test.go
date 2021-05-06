// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SmMhoRanfunctionDescription(t *testing.T) {
	e2smMhoRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-MHO-RANfunction-Description.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmMhoPdu, err := mhoctypes.XerDecodeE2SmMhoRanfunctionDescription(e2smMhoRanfunctionDescription)
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmMhoRanfunctionDescription(e2SmMhoPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "Oid", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 1, len(*ricReportList))
}

func Test_DecodeE2SmMhoRanfunctionDescriptionBytes(t *testing.T) {
	rfd, err := mhoctypes.PerDecodeE2SmMhoRanfunctionDescription(ranFuncDescBytesMHO)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	assert.Equal(t, "ONF", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "OpenNetworking", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "Oid", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, 3, int(rfd.GetRanFunctionName().GetRanFunctionInstance()))
	assert.Equal(t, 1, len(rfd.RicReportStyleList))
	assert.Equal(t, 1, len(rfd.RicEventTriggerStyleList))
}

var ranFuncDescBytesMHO = []byte{
	0x20, 0x20, 0x4f, 0x4e, 0x46, 0x00, 0x00, 0x02, 0x4f, 0x69, 0x64, 0x06, 0x80, 0x4f, 0x70, 0x65,
	0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x01, 0x03, 0x60, 0x00, 0x01,
	0x0d, 0x03, 0x80, 0x4f, 0x4e, 0x46, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x01, 0x2a, 0x00, 0x01, 0x0c,
	0x04, 0x00, 0x4f, 0x4e, 0x46, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x01, 0x15, 0x01, 0x38,
}
