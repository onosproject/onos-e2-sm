// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestE2SmKpmEventTriggerDefinition_NewE2SmKpmEventTriggerDefinition(t *testing.T) {

	e2SmKpmEventTriggerDefinition := NewE2SmKpmEventTriggerDefinition()

	assert.Assert(t, reflect.TypeOf(E2SmKpmEventTriggerDefinition{}) == reflect.TypeOf(*e2SmKpmEventTriggerDefinition), "E2SmKpmEventTriggerDefinition{} values are mismatched")
}

func TestE2SmKpmEventTriggerDefinition_SetE2SmKpmEventTriggerDefinition(t *testing.T) {

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
	e2SmKpmEventTriggerDefinition := NewE2SmKpmEventTriggerDefinition().SetE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinitionFormat1)

	assert.DeepEqual(t, e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition.RetrieveTriggerConditionIeItemByRtPeriod(rtPeriod1), triggerConditionIeItem1)
	assert.DeepEqual(t, e2SmKpmEventTriggerDefinition.E2SmKpmEventTriggerDefinition.RetrieveTriggerConditionIeItemByRtPeriod(rtPeriod2), triggerConditionIeItem2)
}

func TestE2SmKpmEventTriggerDefinition_GetE2SmKpmEventTriggerDefinition(t *testing.T) {

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
	e2SmKpmEventTriggerDefinition1 := NewE2SmKpmEventTriggerDefinition().SetE2SmKpmEventTriggerDefinition(e2SmKpmEventTriggerDefinitionFormat1)
	e2SmKpmEventTriggerDefinition2 := e2SmKpmEventTriggerDefinition1.GetE2SmKpmEventTriggerDefinition()

	assert.DeepEqual(t, e2SmKpmEventTriggerDefinition1.E2SmKpmEventTriggerDefinition, e2SmKpmEventTriggerDefinition2)

}
