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

var refPerMeasurementTypeID = "00000000  00 00 7a                                          |..z|"

func createMeasurementTypeID() *e2sm_kpm_v2_go.MeasurementTypeId {
	return &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: 123,
	}
}

func Test_perEncodingMeasurementTypeID(t *testing.T) {

	mtID := createMeasurementTypeID()

	per, err := aper.MarshalWithParams(mtID, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementTypeID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementTypeId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementTypeID - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementTypeID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
