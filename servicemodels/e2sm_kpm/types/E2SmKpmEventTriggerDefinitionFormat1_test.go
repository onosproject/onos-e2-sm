// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestE2SmKpmEventTriggerDefinitionFormat1_NewE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	e2SmKpmEventTriggerDefinitionFormat1 := NewE2SmKpmEventTriggerDefinitionFormat1()

	assert.Assert(t, reflect.TypeOf(E2SmKpmEventTriggerDefinitionFormat1{}) == reflect.TypeOf(*e2SmKpmEventTriggerDefinitionFormat1), "E2SmKpmEventTriggerDefinitionFormat1{} values are mismatched")
}

func TestE2SmKpmEventTriggerDefinitionFormat1_AddTriggerConditionIeItem(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)
	triggerConditionIeItem := NewTriggerConditionIeItem(rtPeriod)

	e2SmKpmEventTriggerDefinitionFormat1 := NewE2SmKpmEventTriggerDefinitionFormat1().AddTriggerConditionIeItem(triggerConditionIeItem)

	assert.Assert(t, e2SmKpmEventTriggerDefinitionFormat1.PolicyTestList != nil)
}

func TestE2SmKpmEventTriggerDefinitionFormat1_RetrieveTriggerConditionIeItemByRtPeriod(t *testing.T) {

	var rtPeriodIe1 int32 = 14
	rtPeriod1, err := NewRtPeriodIe(rtPeriodIe1)
	assert.NilError(t, err)
	triggerConditionIeItem1 := NewTriggerConditionIeItem(rtPeriod1)
	var rtPeriodIe2 int32 = 19
	rtPeriod2, err := NewRtPeriodIe(rtPeriodIe2)
	assert.NilError(t, err)
	triggerConditionIeItem2 := NewTriggerConditionIeItem(rtPeriod2)

	e2SmKpmEventTriggerDefinitionFormat1 := NewE2SmKpmEventTriggerDefinitionFormat1().
		AddTriggerConditionIeItem(triggerConditionIeItem1).AddTriggerConditionIeItem(triggerConditionIeItem2)

	assert.DeepEqual(t, e2SmKpmEventTriggerDefinitionFormat1.RetrieveTriggerConditionIeItemByRtPeriod(rtPeriod2), triggerConditionIeItem2)
}

func TestE2SmKpmEventTriggerDefinitionFormat1_GetE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	var rtPeriodIe1 int32 = 14
	rtPeriod1, err := NewRtPeriodIe(rtPeriodIe1)
	assert.NilError(t, err)
	triggerConditionIeItem1 := NewTriggerConditionIeItem(rtPeriod1)
	var rtPeriodIe2 int32 = 19
	rtPeriod2, err := NewRtPeriodIe(rtPeriodIe2)
	assert.NilError(t, err)
	triggerConditionIeItem2 := NewTriggerConditionIeItem(rtPeriod2)

	e2SmKpmEventTriggerDefinitionFormat1 := NewE2SmKpmEventTriggerDefinitionFormat1().
		AddTriggerConditionIeItem(triggerConditionIeItem1).AddTriggerConditionIeItem(triggerConditionIeItem2)
	e2SmKpmEventTriggerDefinitionFormat2 := e2SmKpmEventTriggerDefinitionFormat1.GetE2SmKpmEventTriggerDefinitionFormat1()

	assert.DeepEqual(t, e2SmKpmEventTriggerDefinitionFormat1, e2SmKpmEventTriggerDefinitionFormat2)

}
