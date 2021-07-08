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

var refPerMeasurementInfoActionList = "00000000  00 00 40 40 6f 6e 66 00  00 14                    |..@@onf...|"

func createMeasurementInfoActionList() (*e2sm_kpm_v2_go.MeasurementInfoActionList, error) {

	res := &e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	item := &e2sm_kpm_v2_go.MeasurementInfoActionItem{
		MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
			Value: 21,
		},
		MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
			Value: "onf",
		},
	}
	res.Value = append(res.Value, item)

	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}
	return res, nil
}

func Test_perEncodingMeasurementInfoActionList(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	per, err := aper.Marshal(mial)
	assert.NilError(t, err)
	t.Logf("MeasurementInfoActionList PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementInfoActionList{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfoActionList - decoded\n%v", result)
}

func Test_perMeasurementInfoActionListCompareBytes(t *testing.T) {

	mial, err := createMeasurementInfoActionList()
	assert.NilError(t, err)

	per, err := aper.Marshal(mial)
	assert.NilError(t, err)
	t.Logf("MeasurementInfoActionList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementInfoActionList)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
