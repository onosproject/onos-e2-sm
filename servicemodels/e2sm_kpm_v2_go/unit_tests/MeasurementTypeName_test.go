// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasurementTypeName = "00000000  01 00 31 32 33                                    |..123|"

func createMeasurementTypeName() *e2sm_kpm_v2_go.MeasurementTypeName {
	return &e2sm_kpm_v2_go.MeasurementTypeName{
		Value: "123",
	}
}

func Test_perEncodingMeasurementTypeName(t *testing.T) {

	mtn := createMeasurementTypeName()

	per, err := aper.MarshalWithParams(mtn, "sizeExt")
	assert.NilError(t, err)
	t.Logf("MeasurementTypeName PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementTypeName{}
	err = aper.UnmarshalWithParams(per, &result, "sizeExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementTypeName - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementTypeName)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
