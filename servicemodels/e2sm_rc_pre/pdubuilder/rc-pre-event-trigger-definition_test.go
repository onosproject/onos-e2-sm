// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreEventTriggerDefinition(t *testing.T) {
	var rtPeriod uint32 = 12

	E2SmRcPrePdu, err := CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, E2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER EventTriggerDefinition (Periodic): \n%s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreEventTriggerDefinition(xer)
	assert.NilError(t, err)
	t.Logf("XER decoded EventTriggerDefinition (Periodic): \n%v", result)
	assert.Equal(t, E2SmRcPrePdu.GetEventDefinitionFormat1().GetReportingPeriodMs(), result.GetEventDefinitionFormat1().GetReportingPeriodMs())

	per, err := rcprectypes.PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded EventTriggerDefinition (Periodic): \n%v", hex.Dump(per))

	resultPer, err := rcprectypes.PerDecodeE2SmRcPreEventTriggerDefinition(per)
	assert.NilError(t, err)
	t.Logf("PER decoded EventTriggerDefinition (Periodic): \n%v", result)
	assert.Equal(t, E2SmRcPrePdu.GetEventDefinitionFormat1().GetReportingPeriodMs(), resultPer.GetEventDefinitionFormat1().GetReportingPeriodMs())

	E2SmRcPrePduUponChange, err := CreateE2SmRcPreEventTriggerDefinitionUponChange()
	assert.NilError(t, err)
	assert.Assert(t, E2SmRcPrePduUponChange != nil)

	xer1, err := rcprectypes.XerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePduUponChange)
	assert.NilError(t, err)
	t.Logf("XER EventTriggerDefinition (Upon Change): \n%s", string(xer1))

	result1, err := rcprectypes.XerDecodeE2SmRcPreEventTriggerDefinition(xer1)
	assert.NilError(t, err)
	t.Logf("XER decoded EventTriggerDefinition (Upno Change): \n%v", result1)
	assert.Equal(t, int32(E2SmRcPrePduUponChange.GetEventDefinitionFormat1().GetTriggerType()), int32(e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE))

	per1, err := rcprectypes.PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePduUponChange)
	assert.NilError(t, err)
	t.Logf("PER Encoded EventTriggerDefinition (Upon Change): \n%v", hex.Dump(per1))

	result1Per, err := rcprectypes.PerDecodeE2SmRcPreEventTriggerDefinition(per1)
	assert.NilError(t, err)
	t.Logf("PER decoded EventTriggerDefinition (Upon Change): \n%v", result1Per)
	assert.Equal(t, int32(E2SmRcPrePduUponChange.GetEventDefinitionFormat1().GetTriggerType()), int32(e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE))
}
