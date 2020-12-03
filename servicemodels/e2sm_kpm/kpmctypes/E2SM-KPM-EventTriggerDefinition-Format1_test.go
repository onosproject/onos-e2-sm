// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2SmKpmEventTriggerDefinitionFormat1() *e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1 {

	policyTestItem := &e2sm_kpm_ies.TriggerConditionIeItem{
		ReportPeriodIe: e2sm_kpm_ies.RtPeriodIe_RT_PERIOD_IE_MS10,
	}
	e2SmKpmEventTriggerDefinitionFormat1 := &e2sm_kpm_ies.E2SmKpmEventTriggerDefinitionFormat1{
		PolicyTestList: make([]*e2sm_kpm_ies.TriggerConditionIeItem, 0),
	}
	e2SmKpmEventTriggerDefinitionFormat1.PolicyTestList = append(e2SmKpmEventTriggerDefinitionFormat1.PolicyTestList, policyTestItem)

	return e2SmKpmEventTriggerDefinitionFormat1

}

func Test_xerEncodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	e2SmKpmEventTriggerDefinitionFormat1 := createE2SmKpmEventTriggerDefinitionFormat1()

	xer, err := xerEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 130, len(xer))
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	e2SmKpmEventTriggerDefinitionFormat1 := createE2SmKpmEventTriggerDefinitionFormat1()

	xer, err := xerEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 130, len(xer))
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmEventTriggerDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, e2SmKpmEventTriggerDefinitionFormat1.GetPolicyTestList(), result.GetPolicyTestList(), "Encoded and decoded E2SM-KPM-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}

//func Test_perEncodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {
//
//	e2SmKpmEventTriggerDefinitionFormat1 := createE2SmKpmEventTriggerDefinitionFormat1()
//
//	per, err := perEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
//	assert.NilError(t, err) //ToDo: obtaining "bad file descriptor" there.. ?
//	assert.Equal(t, 1, len(per))
//	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER\n%s", string(per))
//}

//func Test_perDecodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {
//
//	e2SmKpmEventTriggerDefinitionFormat1 := createE2SmKpmEventTriggerDefinitionFormat1()
//
//	per, err := perEncodeE2SmKpmEventTriggerDefinitionFormat1(e2SmKpmEventTriggerDefinitionFormat1)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per))
//	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER\n%s", string(per))
//
//	result, err := perDecodeE2SmKpmEventTriggerDefinitionFormat1(per)
//	assert.NilError(t, err)
//	assert.Equal(t, triggerConditionIeItem.GetReportPeriodIe(), result.GetReportPeriodIe(), "Encoded and decoded E2SM-KPM-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
//}
