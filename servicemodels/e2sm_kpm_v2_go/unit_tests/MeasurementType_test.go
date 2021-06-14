// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasurementType1 = "00000000  00 40 6f 6e 66                                    |.@onf|" //Name
var refPerMeasurementType2 = "00000000  40 00 7a                                          |@.z|"   //ID

func createMeasurementType1() (*e2sm_kpm_v2_go.MeasurementType, error) {
	return pdubuilder.CreateMeasurementTypeMeasName("onf")
}

func createMeasurementType2() (*e2sm_kpm_v2_go.MeasurementType, error) {
	return pdubuilder.CreateMeasurementTypeMeasID(123)
}

func Test_perEncodingMeasurementType1(t *testing.T) {

	mt1, err := createMeasurementType1()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mt1, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementType (Name) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementType (Name) - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementType1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingMeasurementType2(t *testing.T) {

	mt2, err := createMeasurementType2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mt2, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementType (ID) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementType (ID) - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasurementType2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
