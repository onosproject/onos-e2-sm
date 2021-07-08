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

var refPerMeasurementInfoActionItem = "00000000  40 40 6f 6e 66 00 00 14                           |@@onf...|"

func createMeasurementInfoActionItem() *e2sm_kpm_v2_go.MeasurementInfoActionItem {

	return &e2sm_kpm_v2_go.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
			Value: 21,
		},
		MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
			Value: "onf",
		},
	}
}

func Test_perEncodingMeasurementInfoActionItem(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	per, err := aper.MarshalWithParams(miai, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementInfoActionItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementInfoActionItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfoActionItem - decoded\n%v", result)
}

func Test_perMeasurementInfoActionItemCompareBytes(t *testing.T) {

	miai := createMeasurementInfoActionItem()

	per, err := aper.MarshalWithParams(miai, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementInfoActionItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementInfoActionItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
