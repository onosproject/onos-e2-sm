// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreEventTriggerDefinition(t *testing.T) {
	var rtPeriod int64 = 12

	E2SmRcPrePdu, err := CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, E2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded EventTriggerDefinition (Periodic): \n%v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreEventTriggerDefinition(per)
	assert.NilError(t, err)
	t.Logf("PER decoded EventTriggerDefinition (Periodic): \n%v", resultPer)
	assert.Equal(t, E2SmRcPrePdu.String(), resultPer.String())

	E2SmRcPrePduUponChange, err := CreateE2SmRcPreEventTriggerDefinitionUponChange()
	assert.NilError(t, err)
	assert.Assert(t, E2SmRcPrePduUponChange != nil)

	per1, err := encoder.PerEncodeE2SmRcPreEventTriggerDefinition(E2SmRcPrePduUponChange)
	assert.NilError(t, err)
	t.Logf("PER Encoded EventTriggerDefinition (Upon Change): \n%v", hex.Dump(per1))

	result1Per, err := encoder.PerDecodeE2SmRcPreEventTriggerDefinition(per1)
	assert.NilError(t, err)
	t.Logf("PER decoded EventTriggerDefinition (Upon Change): \n%v", result1Per)
	assert.Equal(t, E2SmRcPrePdu.String(), resultPer.String())
}
