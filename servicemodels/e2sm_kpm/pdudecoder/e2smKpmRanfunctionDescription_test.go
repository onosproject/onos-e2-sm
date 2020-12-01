// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
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
	assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "OID123", string(ranFunctionName.RanFunctionE2SmOid))
	assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 6, len(*ricReportList))
}
