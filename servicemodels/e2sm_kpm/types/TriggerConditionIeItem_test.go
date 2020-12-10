// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestTriggerConditionIeItem_NewTriggerConditionIeItem(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)
	triggerConditionIeItem := NewTriggerConditionIeItem(rtPeriod)

	assert.Assert(t, reflect.TypeOf(TriggerConditionIeItem{}) == reflect.TypeOf(*triggerConditionIeItem), "TriggerConditionIeItem{} values are mismatched")
	assert.DeepEqual(t, triggerConditionIeItem.ReportPeriodIe.RtPeriodIe, rtPeriodIe)
}

func TestTriggerConditionIeItem_GetRtPeriodIe(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)
	triggerConditionIeItem := NewTriggerConditionIeItem(rtPeriod)

	assert.DeepEqual(t, triggerConditionIeItem.GetRtPeriodIe(), rtPeriod)
}

func TestTriggerConditionIeItem_GetTriggerConditionIeItem(t *testing.T) {

	var rtPeriodIe int32 = 14
	rtPeriod, err := NewRtPeriodIe(rtPeriodIe)
	assert.NilError(t, err)

	triggerConditionIeItem1 := NewTriggerConditionIeItem(rtPeriod)
	triggerConditionIeItem2 := triggerConditionIeItem1.GetTriggerConditionIeItem()

	assert.DeepEqual(t, triggerConditionIeItem1, triggerConditionIeItem2)
}
