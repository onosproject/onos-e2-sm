// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/hex"
	kpmctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_XerDecodeE2SmKpmRanfunctionDescription(t *testing.T) {
	e2smKpmRanfunctionDescription, err := ioutil.ReadFile("../test/RiaB_RANfunctionDescription.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmKpmPdu, err := kpmctypes.XerDecodeE2SmKpmRanfunctionDescription(e2smKpmRanfunctionDescription)
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 1, len(*ricReportList))
}

func Test_PerDecodeE2SmKpmRanfunctionDescriptionBytes(t *testing.T) {
	t.Logf("PER bytes in HEX format for RanFunctionDescription are \n%v", hex.Dump(rfdE2apPdu))
	rfd, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(rfdE2apPdu)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
}

// These bytes are taken from E2SetupRequest PDU decoded in this PR:
// https://github.com/onosproject/onos-e2t/pull/300
var rfdE2apPdu = []byte{
	0x74, 0x18, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x4b, 0x50, 0x4d, 0x00,
	0x00, 0x18, 0x31, 0x2e, 0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33,
	0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x32, 0x2e, 0x32, 0x05, 0x00, 0x4b, 0x50, 0x4d,
	0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x00, 0x00, 0x00, 0x00, 0x58, 0x02, 0xf8, 0x10,
	0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x45, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x43,
	0x65, 0x6c, 0x6c, 0x46, 0x44, 0x44, 0x40, 0x02, 0xf8, 0x10, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x01,
	0x03, 0x80, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x31, 0x00, 0x01, 0x00, 0x06, 0x1e, 0x80,
	0x4f, 0x2d, 0x43, 0x55, 0x2d, 0x55, 0x50, 0x20, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x50, 0x43, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x06,
	0x00, 0x04, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61,
	0x62, 0x41, 0x74, 0x74, 0x2e, 0x73, 0x75, 0x6d, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52, 0x43,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2e, 0x73,
	0x75, 0x6d, 0x00, 0x00, 0x01, 0x42, 0xa0, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52,
	0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x73, 0x75, 0x6d, 0x00, 0x00, 0x02,
	0x41, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x65, 0x61, 0x6e, 0x00, 0x00,
	0x03, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x00, 0x00,
	0x04, 0x00, 0x01, 0x00, 0x01}
