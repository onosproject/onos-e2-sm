// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SmRcPreEventTriggerDefinitionFormat1() (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1, error) {

	var rtPeriod uint32 = 12

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod)
	if err != nil {
		return nil, err
	}

	return newE2SmRcPrePdu.GetEventDefinitionFormat1(), nil

}

func Test_XerEncodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	//assert.Equal(t, 182, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	//assert.Equal(t, 182, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreEventTriggerDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Decoded EventTriggerDefinition-Format1 from XER is \n%v", result)
}

func Test_PerEncodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 PER\n%v", hex.Dump(per))
}

func Test_PerDecodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := PerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreEventTriggerDefinitionFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Decoded EventTriggerDefinition-Format1 from PER is \n%v", result)
}
