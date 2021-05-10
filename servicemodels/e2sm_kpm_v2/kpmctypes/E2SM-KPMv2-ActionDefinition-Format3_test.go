// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createActionDefinitionFormat3() (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat3, error) {

	var cellObjID string = "onf"
	var granularity uint32 = 21
	var subscriptionID int64 = 12345
	var measurementName string = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, _ := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))

	mci, _ := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)

	mcl := &e2sm_kpm_v2.MatchingCondList{
		Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	measCondItem, _ := pdubuilder.CreateMeasurementCondItem(measName, mcl)

	measCondList := e2sm_kpm_v2.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, _ := pdubuilder.CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	if err := actionDefinitionFormat3.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	}
	return actionDefinitionFormat3, nil
}

func Test_xerEncodeE2SmKpmActionDefinitionFormat3(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmActionDefinitionFormat3(actionDefFormat3)
	assert.NilError(t, err)
	assert.Equal(t, 819, len(xer))
	t.Logf("E2SmKpmActionDefinitionFormat3 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmActionDefinitionFormat3(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmActionDefinitionFormat3(actionDefFormat3)
	assert.NilError(t, err)
	assert.Equal(t, 819, len(xer))
	t.Logf("E2SmKpmActionDefinitionFormat3 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmActionDefinitionFormat3(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmActionDefinitionFormat3 XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmActionDefinitionFormat3(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmActionDefinitionFormat3(actionDefFormat3)
	assert.NilError(t, err)
	assert.Equal(t, 27, len(per))
	t.Logf("E2SmKpmActionDefinitionFormat3 PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmActionDefinitionFormat3(t *testing.T) {

	actionDefFormat3, err := createActionDefinitionFormat3()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmActionDefinitionFormat3(actionDefFormat3)
	assert.NilError(t, err)
	assert.Equal(t, 27, len(per))
	t.Logf("E2SmKpmActionDefinitionFormat3 PER\n%s", string(per))

	result, err := perDecodeE2SmKpmActionDefinitionFormat3(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmActionDefinitionFormat3 PER - decoded\n%s", result)
}
