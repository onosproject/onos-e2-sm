// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerAD3 string = "00000000  00 00 03 6f 6e 66 00 00  00 40 74 72 69 61 6c 00  |...onf...@trial.|\n" +
	"00000010  00 48 21 02 00 c9 00 14  40 30 38                 |.H!.....@08|"

func createActionDefinitionFormat3() (*e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat3, error) {

	var cellObjID string = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))
	if err != nil {
		return nil, err
	}

	mci, err := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)
	if err != nil {
		return nil, err
	}

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	if err != nil {
		return nil, err
	}
	measCondItem, err := pdubuilder.CreateMeasurementCondItem(measName, mcl)
	if err != nil {
		return nil, err
	}

	measCondList := e2sm_kpm_v2_go.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, err := pdubuilder.CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	if err != nil {
		return nil, err
	}
	//if err := actionDefinitionFormat3.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	//}
	return actionDefinitionFormat3, nil
}

func Test_perEncodingE2SmKpmActionDefinitionFormat3(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat3, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition-Format3 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat3{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-ActionDefinition-Format3 PER - decoded\n%v", result)
}

func Test_perE2SmKpmActionDefinitionFormat3CompareBytes(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*actionDefFormat3, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition-Format3 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerAD3)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
