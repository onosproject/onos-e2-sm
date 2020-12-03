// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createTriggerConditionIeItem() *e2sm_kpm_ies.TriggerConditionIeItem {

	return &e2sm_kpm_ies.TriggerConditionIeItem{
		ReportPeriodIe: e2sm_kpm_ies.RtPeriodIe_RT_PERIOD_IE_MS10,
	}

}

func Test_xerEncodeTriggerConditionIeItem(t *testing.T) {

	triggerConditionIeItem := createTriggerConditionIeItem()

	xer, err := xerEncodeTriggerConditionIeItem(triggerConditionIeItem)
	assert.NilError(t, err)
	assert.Equal(t, 104, len(xer))
	t.Logf("Trigger-ConditionIE-Item XER\n%s", string(xer))
}

func Test_xerDecodeTriggerConditionIeItem(t *testing.T) {

	triggerConditionIeItem := createTriggerConditionIeItem()

	xer, err := xerEncodeTriggerConditionIeItem(triggerConditionIeItem)
	assert.NilError(t, err)
	assert.Equal(t, 104, len(xer))
	t.Logf("Trigger-ConditionIE-Item XER\n%s", string(xer))

	result, err := xerDecodeTriggerConditionIeItem(xer)
	assert.NilError(t, err)
	assert.Equal(t, triggerConditionIeItem.GetReportPeriodIe(), result.GetReportPeriodIe(), "Encoded and decoded Trigger-ConditionIE-Item -- RtPeriodIe values are not the same")
}

func Test_perEncodeTriggerConditionIeItem(t *testing.T) {

	triggerConditionIeItem := createTriggerConditionIeItem()

	per, err := perEncodeTriggerConditionIeItem(triggerConditionIeItem)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("Trigger-ConditionIE-Item PER\n%s", string(per))
}

func Test_perDecodeTriggerConditionIeItem(t *testing.T) {

	triggerConditionIeItem := createTriggerConditionIeItem()

	per, err := perEncodeTriggerConditionIeItem(triggerConditionIeItem)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("Trigger-ConditionIE-Item PER\n%s", string(per))

	result, err := perDecodeTriggerConditionIeItem(per)
	assert.NilError(t, err)
	assert.Equal(t, triggerConditionIeItem.GetReportPeriodIe(), result.GetReportPeriodIe(), "Encoded and decoded Trigger-ConditionIE-Item -- RtPeriodIe values are not the same")
}
